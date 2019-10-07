package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func skipID3(buffer []byte) int {
	var id3v2Flags, z0, z1, z2, z3 byte
	var tagSize, footerSize int

	//http://id3.org/d3v2.3.0
	if buffer[0] == 0x49 && buffer[1] == 0x44 && buffer[2] == 0x33 { //'ID3'
		id3v2Flags = buffer[5]
		if (id3v2Flags & 0x10) != 0 {
			footerSize = 10
		} else {
			footerSize = 0
		}

		//ID3 size encoding is crazy (7 bits in each of 4 bytes)
		z0 = buffer[6]
		z1 = buffer[7]
		z2 = buffer[8]
		z3 = buffer[9]
		if ((z0 & 0x80) == 0) && ((z1 & 0x80) == 0) && ((z2 & 0x80) == 0) && ((z3 & 0x80) == 0) {
			tagSize = (((int)(z0 & 0x7f)) * 2097152) +
				(((int)(z1 & 0x7f)) * 16384) +
				(((int)(z2 & 0x7f)) * 128) +
				((int)(z3 & 0x7f))
			return 10 + tagSize + footerSize
		}
	}
	return 0
}

func mp3duration(filename string) (duration float64, err error) {
	var f *os.File
	f, err = os.Open("demo - vbr.mp3")
	if err != nil {
		return
	}
	defer f.Close()

	buffer := make([]byte, 100)
	bytesRead, err := f.Read(buffer)
	if err != nil {
		return
	}
	if bytesRead < 100 {
		err = errors.New("Corrupt file")
		return
	}

	offset := skipID3(buffer)

	fmt.Printf("%d\n", offset)

	return
}

func main() {
	duration, err := mp3duration("demo - cbr.mp3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Duration %v\n", duration)
}
