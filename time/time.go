package time

import (
	"crypto/rand"
	"encoding/binary"
	"github.com/landjur/go-uuid/utility"
	"net"
	"sync"
	"time"
)

const (
	// Intervals bewteen 1/1/1970 and 15/10/1582 (Julain days of 1 Jan 1970 - Julain days of 15 Oct 1582) * 100-Nanoseconds Per Day
	intervals = (2440587 - 2299160) * 86400 * 10000000
)

var (
	lastGenerated time.Time // last generated time
	clockSequence uint16    // clock sequence for same tick
	nodeID        []byte    // node id (MAC Address)

	locker sync.Mutex // global lock
)

// New returns a new uuid time based.
func New() ([]byte, error) {
	// Get and release a global lock
	locker.Lock()
	defer locker.Unlock()

	id := make([]byte, 16)

	// get timestamp
	now := time.Now().UTC()
	timestamp := uint64(now.UnixNano()/100) + intervals // get timestamp
	if !now.After(lastGenerated) {
		clockSequence++ // last generated time known, then just increment clock sequence
	} else {
		b := make([]byte, 2)
		n, err := rand.Read(b)
		if n != len(b) || err != nil {
			return nil, err
		}
		clockSequence = uint16(int(b[0])<<8 | int(b[1])) // set to a random value (network byte order)
	}

	lastGenerated = now // remember the last generated time

	timeLow := uint32(timestamp & 0xffffffff)
	timeMiddle := uint16((timestamp >> 32) & 0xffff)
	timeHigh := uint16((timestamp >> 48) & 0xfff)

	// network byte order(BigEndian)
	binary.BigEndian.PutUint32(id[0:], timeLow)
	binary.BigEndian.PutUint16(id[4:], timeMiddle)
	binary.BigEndian.PutUint16(id[6:], timeHigh)
	binary.BigEndian.PutUint16(id[8:], clockSequence)

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
			if n != len(id) || err != nil {
				return nil, err
			}
		}
	}

	copy(id[10:], nodeID)

	utility.SetVersion(id, 1) // Version 1
	utility.SetLayout(id, 2)  // RFC4122 Layout

	return id, nil
}
