package timebased

import (
	"crypto/rand"
	"encoding/binary"
	"github.com/wayn3h0/go-errors"
	"github.com/wayn3h0/go-uuid"
	"net"
	"sync"
	"time"
)

const (
	// Intervals bewteen 1/1/1970 and 15/10/1582 (Julain days of 1 Jan 1970 - Julain days of 15 Oct 1582) * 100-Nanoseconds Per Day
	intervals = (2440587 - 2299160) * 86400 * 10000000
)

var (
	lastGenerated time.Time  // last generated time
	clockSequence uint16     // clock sequence for same tick
	nodeID        []byte     // node id (MAC Address)
	locker        sync.Mutex // global lock
)

// New returns a new time-based uuid.
func New() (uuid.UUID, error) {
	// Get and release a global lock
	locker.Lock()
	defer locker.Unlock()

	instance := make(uuid.UUID, 16)

	// get timestamp
	now := time.Now().UTC()
	timestamp := uint64(now.UnixNano()/100) + intervals // get timestamp
	if !now.After(lastGenerated) {
		clockSequence++ // last generated time known, then just increment clock sequence
	} else {
		b := make([]byte, 2)
		n, err := rand.Read(b)
		if err != nil {
			return nil, errors.Wrap(err, "uuid/timebased: generating pseudorandom numbers for clock sequence failed")
		}
		if n != len(b) {
			return nil, errors.New("uuid/timebased: generating pseudorandom numbers for clock sequence failed")
		}
		clockSequence = uint16(int(b[0])<<8 | int(b[1])) // set to a random value (network byte order)
	}

	lastGenerated = now // remember the last generated time

	timeLow := uint32(timestamp & 0xffffffff)
	timeMiddle := uint16((timestamp >> 32) & 0xffff)
	timeHigh := uint16((timestamp >> 48) & 0xfff)

	// network byte order(BigEndian)
	binary.BigEndian.PutUint32(instance[0:], timeLow)
	binary.BigEndian.PutUint16(instance[4:], timeMiddle)
	binary.BigEndian.PutUint16(instance[6:], timeHigh)
	binary.BigEndian.PutUint16(instance[8:], clockSequence)

	// get node id(mac address)
	if nodeID == nil {
		interfaces, err := net.Interfaces()
		if err != nil {
			return nil, err
		}

		for _, i := range interfaces {
			if len(i.HardwareAddr) >= 6 {
				nodeID = make([]byte, 6)
				copy(nodeID, i.HardwareAddr)
				break
			}
		}

		if nodeID == nil {
			nodeID = make([]byte, 6)
			n, err := rand.Read(nodeID)
			if err != nil {
				return nil, errors.Wrap(err, "uuid/timebased: generating pseudorandom numbers for node id failed")
			}
			if n != len(instance) {
				return nil, errors.New("uuid/timebased: generating pseudorandom numbers for node id failed")
			}
		}
	}

	copy(instance[10:], nodeID)

	uuid.SetVersion(instance, uuid.VersionTimeBased)
	uuid.SetLayout(instance, uuid.LayoutRFC4122)

	return instance, nil
}
