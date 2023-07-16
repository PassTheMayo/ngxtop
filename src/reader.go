package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"regexp"
)

var ipAddressRegEx = regexp.MustCompile(`\d{1,3}(\.\d{1,3}){3}`)

func ReaderGoroutine() {
	f, err := os.Open(LogPath)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.Seek(0, 2); err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')

		if err != nil {
			if errors.Is(err, io.EOF) {
				continue
			}

			panic(err)
		}

		matches := ipAddressRegEx.FindAllStringSubmatch(line, 1)

		if len(matches) < 1 {
			continue
		}

		AddRequest(matches[0][0])
	}
}
