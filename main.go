package main

import (
	"flag"
	"fmt"
	"github.com/alrs/plausible/generator"
	"log"
	"os"
)

func printManufacturerList(m *generator.Manuf) {
	companies := m.CompanyList()
	for _, c := range companies {
		fmt.Println(c)
	}
	os.Exit(0)
}

func main() {

	listArg := flag.Bool("l", false, "List manufacturers.")
	manuArg := flag.String("m", "google", "Select manufacturer.")
	dbArg := flag.String("d", generator.ManufPath, "Path of database file.")

	flag.Parse()

	manuf, err := generator.NewManuf(*dbArg)
	if err != nil {
		log.Fatal(err)
	}

	if *listArg {
		printManufacturerList(&manuf)
	}

	mac, err := manuf.RandomMAC(*manuArg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mac)
}
