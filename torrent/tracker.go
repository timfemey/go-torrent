package torrent

import (
	"net/url"
	"strconv"
)

func (t *Torrentfile) buildTrackerURL(peerID [20]byte, port uint16) (string, error) {
	base, err := url.Parse(t.Announce)
	if err != nil {
		return "", err
	}
	// info_hash: Identifies the file we’re trying to download.
	// It’s the infohash we calculated earlier from the bencoded info dict. The tracker will use this to figure out which peers to show us.

	//peer_id: A 20 byte name to identify ourselves to trackers and peers. We’ll just generate 20 random bytes for this.
	//Real BitTorrent clients have IDs like -TR2940-k8hj0wgej6ch which identify the client software and version—in this case, TR2940 stands for Transmission client 2.94.
	params := url.Values{
		"info_hash":  []string{string(t.InfoHash[:])},
		"peer_id":    []string{string(peerID[:])},
		"port":       []string{strconv.Itoa(int(port))},
		"uploaded":   []string{"0"},
		"downloaded": []string{"0"},
		"compact":    []string{"1"},
		"left":       []string{strconv.Itoa(t.Length)},
	}
	base.RawQuery = params.Encode()
	return base.String(), nil
}
