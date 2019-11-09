package problem015

import (
	"io"
	"math/rand"
	"time"
)

func Run(streamReader io.Reader) (byte, error) {
	buffer := make([]byte, 64, 64)
	var selection byte
	totalReadSize := 0
	rand.Seed(time.Now().UnixNano())
	for readSize, err := streamReader.Read(buffer); readSize > 0; {
		for i := 0; i < readSize; i++ {
			totalReadSize++
			if rand.Intn(totalReadSize) == totalReadSize-1 {
				selection = buffer[i]
			}
		}
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return 0, err
			}
		}
		readSize, err = streamReader.Read(buffer)
	}
	return selection, nil
}
