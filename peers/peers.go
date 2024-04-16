package peers

import (
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
)

type Peer struct {
	IP   net.IP
	Port uint16
}

// Unmarshal parses peer IP addresses and ports from a buffer
func Unmarshal(peersBin []byte) ([]Peer, error) {
	const peerSize = 6
	numPeers := len(peersBin) / peerSize
	if len(peersBin)%peerSize != 0 {
		err := fmt.Errorf("Received malformed peers")
		return nil, err
	}
	peers := make([]Peer, numPeers)
	for i := 0; i < numPeers; i++ {
		offset := i * peerSize
		peers[i].IP = net.IP(peersBin[offset : offset+4])                              //4 For IP
		peers[i].Port = binary.BigEndian.Uint16([]byte(peersBin[offset+4 : offset+6])) //2 For the Port , as it gets the last two integers that signify the Port
	}
	return peers, nil
}

func (p Peer) String() string {
	return net.JoinHostPort(p.IP.String(), strconv.Itoa(int(p.Port)))
}
