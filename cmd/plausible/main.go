package main

import (
	"flag"
	"fmt"
	"github.com/alrs/plausible"
	"log"
	"os"
)

var commit string
var buildTime string

func printManufacturerList(m plausible.Manuf) {
	companies := m.CompanyList()
	for _, c := range companies {
		fmt.Println(c)
	}
	os.Exit(0)
}

func main() {

	listArg := flag.Bool("l", false, "List manufacturers.")
	manuArg := flag.String("m", "google", "Select manufacturer.")
	dbArg := flag.String("d", "/usr/share/wireshark/manuf", "Path of database file.")
	flag.Parse()

	manuf, err := plausible.NewManuf(*dbArg)
	if err != nil {
		log.Fatal(err)
	}

	if *listArg {
		printManufacturerList(manuf)
	}

	mac, err := manuf.RandomMAC(*manuArg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mac)
}
