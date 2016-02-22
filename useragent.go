package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	path       = "ua.txt"
	useragents []string
)

// GetRandomUA get random user agent string
func GetRandomUA() string {
	ua := "Foo Bar"

	if len(useragents) == 0 {
		inFile, err := os.Open(path)
		if err != nil {
			log.Fatalln(err)
		}
		defer inFile.Close()

		scanner := bufio.NewScanner(inFile)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			useragents = append(useragents, scanner.Text())
		}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ua = useragents[r.Intn(len(useragents))]

	return ua
}
