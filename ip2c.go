package main

import (
	//"errors"
	"bufio"
	"flag"
	"fmt"
	"github.com/nranchev/go-libGeoIP"
	"os"
	"strings"
)

func usage() {
	fmt.Fprint(os.Stderr, "usage: ip2c [-db db-file-path] ipaddress\n")

	for _, v := range []string{"db", "help"} {
		f := flag.Lookup(v)
		fmt.Fprintf(os.Stderr, "  -%s\t: %s\n", v, f.Usage)
	}
	os.Exit(1)
}

type Config struct {
	dbPath string
}

func Configure() *Config {
	geoIpDbPath := os.Getenv("GEOIP_DB_PATH")

	if geoIpDbPath == "" {
		currentDir, _ := os.Getwd()
		geoIpDbPath = strings.Join([]string{currentDir, "GeoIP.dat"}, "/")
	}

	dbPath := flag.String("db", geoIpDbPath, "File path of 'GeoIP.dat'")
	help := flag.Bool("help", false, "Show this message")

	flag.Usage = usage
	flag.Parse()

	if *help {
		flag.Usage()
	}

	return &Config{*dbPath}
}

func detectCountry(geoIp *libgeo.GeoIP, ipAddress string) {
	loc := geoIp.GetLocationByIP(ipAddress)
	if loc == nil {
		fmt.Println("")
		return
	}

	fmt.Println(loc.CountryName)
}

func readFromStdin(geoIp *libgeo.GeoIP) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		detectCountry(geoIp, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func readFromArg(geoIp *libgeo.GeoIP) {

	for _, ipAddress := range flag.Args() {
		detectCountry(geoIp, ipAddress)
	}
}

func main() {
	config := Configure()

	geoIp, err := libgeo.Load(config.dbPath)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	if flag.NArg() == 0 {
		readFromStdin(geoIp)
	} else {
		readFromArg(geoIp)
	}
}
