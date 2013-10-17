package main

import (
	"flag"
	"fmt"
	"github.com/nranchev/go-libGeoIP"
	"os"
	"strings"
)

var (
	dbPath    string
	ipAddress string
)

func init() {
	maxmindDbPath := os.Getenv("MAXMIND_DB_PATH")

	if maxmindDbPath == "" {
		currentDir, _ := os.Getwd()
		maxmindDbPath = strings.Join([]string{currentDir, "GeoIP.dat"}, "/")
	}

	flag.StringVar(&dbPath, "db", maxmindDbPath, "Maxmind db file path")
	flag.Parse()
}

func main() {
	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "usage: ip2c [--db db-file-path] ipaddress")
		os.Exit(1)
	}

	ipAddress := flag.Arg(0)

	geoIp, err := libgeo.Load(dbPath)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	loc := geoIp.GetLocationByIP(ipAddress)
	if loc == nil {
		fmt.Println("")
		return
	}

	fmt.Println(loc.CountryName)
}
