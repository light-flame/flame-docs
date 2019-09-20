package main

import (
	"io/ioutil"
	"regexp"
	"strings"
)

func readFile(url string) string {
	dat, err := ioutil.ReadFile(url)
	if err != nil {
		panic(err)
	}
	return string(dat)
}

func writeFile(url string, data string) {
	ioutil.WriteFile(url, []byte(data), 0644)
}

func main() {
	fileMd := readFile("README-CORE.md")

	re := regexp.MustCompile("\\$\\{(.*?)\\}")
	for {
		match := re.FindStringSubmatch(fileMd)
		if match == nil {
			break
		}
		path := strings.TrimLeft(match[0], "${")
		path = strings.TrimRight(path, "}")
		fileMd = strings.Replace(fileMd, match[0], readFile(path), -1)
	}

	writeFile("../flame-core/README.md", fileMd)
}
