package generator

import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"strings"
	"time"
	"unicode"
)

type prefix [3]uint8
type Manuf map[string][]prefix
type vendorRecord struct {
	company string
	prefix  prefix
}

const ManufPath = "/usr/share/wireshark/manuf"

var UnparseableLineError = errors.New("Unparseable vendor line.")
var NoSuchCompanyError = errors.New("No such company.")

// addPrefix takes a single prefix record and appends the prefix value
// to the company key in the vendor map.
func (m Manuf) addPrefix(vr vendorRecord) {
	cleanCompanyName := strings.TrimSpace(strings.ToLower(vr.company))
	m[cleanCompanyName] = append(m[cleanCompanyName], vr.prefix)
}

// CompanyList returns a list of every company key in the vendor map.
func (m Manuf) CompanyList() []string {
	companies := []string{}
	for k, _ := range m {
		companies = append(companies, k)
	}
	sort.Strings(companies)
	return companies
}

// LoadRecords reads a flat file database of the Wireshark manuf
// format and loads it into the vendor map.
func (m Manuf) LoadRecords(r io.Reader) (err error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		record, err := parseLine(line)
		if err != nil {
			switch err {
			case UnparseableLineError:
				break
			default:
				return err
			}
		}
		m.addPrefix(record)
	}
	return nil
}

// RandomMAC provides a randomly-generated MAC address from
// a randomly chosen portion of a company's assigned MAC
// address space.
func (m Manuf) RandomMAC(company string) (string, error) {
	companyRecord, ok := m[company]
	if !ok {
		return "", NoSuchCompanyError
	}
	rand.Seed(time.Now().UTC().UnixNano())
	length := len(companyRecord)
	p := companyRecord[rand.Intn(length)]
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", p[0], p[1], p[2],
		rand.Intn(255), rand.Intn(255), rand.Intn(255)), nil
}

// parseLine reads an individual line from a Wireshark "manuf" format
// database and returns a structured vendorRecord if it is able to parse
// the line.
func parseLine(line string) (vr vendorRecord, err error) {
	companyLeftColumn := 9
	prefixLength := 8
	octetPosition := []int{
		0,
		3,
		6,
	}

	sepPosition := []int{
		2,
		5,
	}

	if len(line) < prefixLength {
		return vr, UnparseableLineError
	}

	for _, pos := range sepPosition {
		if line[pos] != ':' {
			return vr, UnparseableLineError
		}
	}

	wsPosition := 8
	if !unicode.IsSpace(rune(line[wsPosition])) {
		return vr, UnparseableLineError
	}

	var prefix prefix
	var company string

	for k, v := range octetPosition {
		octet, err := hex.DecodeString(line[v : v+2])
		prefix[k] = uint8(octet[0])
		if err != nil {
			return vr, UnparseableLineError
		}
	}

	company = string(line[companyLeftColumn:len(line)])
	company = strings.Split(company, "#")[0]
	company = strings.TrimSpace(company)

	vr = vendorRecord{
		company,
		prefix,
	}
	return vr, nil
}
