package main

import (
	"bufio"
	"fmt"
	"net/http"
	"regexp"
)

const apnic_ip_url = "http://ftp.apnic.net/stats/apnic/delegated-apnic-latesttt"

func filter(regex *regexp.Regexp, line string) (output string) {
	if regex.MatchString(line) {
		output = line
	} else {
		output = ""
	}
	return
}

func main() {
	regex, err := regexp.Compile("ipv4")
	if err != nil {
		panic(err)
	}

	response, err := http.Get(apnic_ip_url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		filter(regex, scanner.Text())
	}

	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s\n", string(contents))
	fmt.Print("hoge")
}
