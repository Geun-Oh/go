package rw

import (
	"io"
	"os"
)

func ReadAndWrite() error {
	fi, err := os.Open("./1.txt")
	defer fi.Close()
	if err != nil {
		return err
	}

	fo, err := os.Create("./2.txt")
	defer fo.Close()
	if err != nil {
		return err
	}

	buff := make([]byte, 1024)

	// 루프

	for {
		// 읽기
		cnt, err := fi.Read(buff)
		if err != nil && err != io.EOF {
			return err
		}

		if cnt == 0 {
			break
		}

		// 쓰기
		_, err = fo.Write(buff[:cnt])
		if err != nil {
			return err
		}
	}
	return nil
}
