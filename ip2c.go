package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/nranchev/go-libGeoIP"
	"os"
	"strings"
)

type Config struct {
	dbPath string
}

func Configure() (*Config, error) {
	if flag.NArg() == 0 {
		return nil, errors.New("usage: ip2c [--db db-file-path] ipaddress")
	}

	geoIpDbPath := os.Getenv("GEOIP_DB_PATH")

	if geoIpDbPath == "" {
		currentDir, _ := os.Getwd()
		geoIpDbPath = strings.Join([]string{currentDir, "GeoIP.dat"}, "/")
	}

	dbPath := flag.String("db", geoIpDbPath, "GeoIP.dat file path")
	flag.Parse()

	return &Config{*dbPath}, nil
}

func detectCountry(ipAddress string) string {
	return ""
}

func readFromArg() {

}

func readFromStdin() {

}

func main() {
	config, err := Configure()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if flag.NArg() == 0 {
		readFromStdin()
	} else {
		readFromArg()
	}

	ipAddress := flag.Arg(0)

	geoIp, err := libgeo.Load(config.dbPath)
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
