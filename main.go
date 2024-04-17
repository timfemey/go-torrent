package main

import (
	"log"
	"os"

	"github.com/timfemey/go-torrent/torrent"
)

func main() {
	inPath := os.Args[1]
	outPath := os.Args[2]

	tf, err := torrent.Open(inPath)
	if err != nil {
		log.Fatal(err)
	}

	err = tf.DownloadToFile(outPath)
	if err != nil {
		log.Fatal(err)
	}
}
