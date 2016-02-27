package main

import (
	"bufio"
	"math/rand"
	"strings"
	"time"
)

var (
	uaPath        = "ua.txt"
	useragentList []string
)

// GetRandomUA get random user agent string
func GetRandomUA() string {
	if len(useragentList) == 0 {
		// inFile, err := os.Open(uaPath)
		// if err != nil {
		// 	log.Fatalln(err)
		// }
		// defer inFile.Close()
		//

		inFile := strings.NewReader(useragents)
		scanner := bufio.NewScanner(inFile)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			useragentList = append(useragentList, scanner.Text())
		}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ua := useragentList[r.Intn(len(useragentList))]

	return ua
}
