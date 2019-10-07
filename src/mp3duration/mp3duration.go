package mp3duration

import (
	"errors"
	"io"
	"os"
)

var (
	versions = []string{"2.5", "x", "2", "1"}
	layers   = []string{"x", "3", "2", "1"}
	bitRates = map[string][]int{
		"V1Lx": []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"V1L1": []int{0, 32, 64, 96, 128, 160, 192, 224, 256, 288, 320, 352, 384, 416, 448},
		"V1L2": []int{0, 32, 48, 56, 64, 80, 96, 112, 128, 160, 192, 224, 256, 320, 384},
		"V1L3": []int{0, 32, 40, 48, 56, 64, 80, 96, 112, 128, 160, 192, 224, 256, 320},
		"V2Lx": []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"V2L1": []int{0, 32, 48, 56, 64, 80, 96, 112, 128, 144, 160, 176, 192, 224, 256},
		"V2L2": []int{0, 8, 16, 24, 32, 40, 48, 56, 64, 80, 96, 112, 128, 144, 160},
		"V2L3": []int{0, 8, 16, 24, 32, 40, 48, 56, 64, 80, 96, 112, 128, 144, 160},
		"VxLx": []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"VxL1": []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"VxL2": []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"VxL3": []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	sampleRates = map[string][]int{
		"x":   []int{0, 0, 0},
		"1":   []int{44100, 48000, 32000},
		"2":   []int{22050, 24000, 16000},
		"2.5": []int{11025, 12000, 8000},
	}
	samples = map[string]map[string]int{
		"x": {
			"x": 0,
			"1": 0,
			"2": 0,
			"3": 0,
		},
		"1": { //MPEGv1,     Layers 1,2,3
			"x": 0,
			"1": 384,
			"2": 1152,
			"3": 1152,
		},
		"2": { //MPEGv2/2.5, Layers 1,2,3
			"x": 0,
			"1": 384,
			"2": 1152,
			"3": 576,
		},
	}
)

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

// Calculate returns the duration of an mp3 file
func Calculate(filename string) (duration float64, err error) {
	var f *os.File
	f, err = os.Open("demo - vbr.mp3")
	if err != nil {
		return
	}
	defer f.Close()

	stats, statsErr := f.Stat()
	if statsErr != nil {
		return 0, statsErr
	}
	size := stats.Size()

	buffer := make([]byte, 100)
	var bytesRead int
	bytesRead, err = f.Read(buffer)
	if err != nil {
		return
	}
	if bytesRead < 100 {
		err = errors.New("Corrupt file")
		return
	}
	offset := int64(skipID3(buffer))

	buffer = make([]byte, 10)
	for offset < size {
		bytesRead, e := f.ReadAt(buffer, offset)
		if e != nil && e != io.EOF {
			err = e
			return
		}
		if bytesRead < 10 {
			return
		}
		offset = offset + int64(bytesRead)

		duration = duration + 1.0
	}

	return
}
