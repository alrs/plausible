package generator

import (
	"os"
	"testing"
)

func TestAddVendor(t *testing.T) {
	p := prefix{
		uint8(10),
		uint8(20),
		uint8(30),
	}
	v := make(Manuf)
	vr := vendorRecord{
		"test",
		p,
	}
	v.addPrefix(vr)
	t.Log(v)
}

func TestRandomMAC(t *testing.T) {
	p := prefix{
		uint8(10),
		uint8(20),
		uint8(30),
	}
	v := make(Manuf)
	vr := vendorRecord{
		"test",
		p,
	}
	v.addPrefix(vr)
	t.Log(v.RandomMAC("test"))
}

func TestParseLine(t *testing.T) {
	validLine := "00:00:02        BbnWasIn        # BBN (was internal usage only, no longer used)"
	invalidLine := "this should fail"
	t.Log("Parsing valid line.")
	parsed, err := parseLine(validLine)
	t.Log(parsed)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Parsing invalid line.")
	_, err = parseLine(invalidLine)
	if err == nil {
		t.Fatalf("Invalid line \"%s\" should have failed to parse.", invalidLine)
	}
}

func TestLoadVendor(t *testing.T) {
	f, err := os.Open(ManufPath)
	if err != nil {
		t.Fatal(err)
	}
	manuf := make(Manuf)
	manuf.LoadRecords(f)
	t.Log(manuf["google"])
}

func TestCompanyList(t *testing.T) {
	manuf := make(Manuf)
	vendorLines := []string{
		"00:00:08        FuturePn        # Officially Xerox, but 0:0:0:0:0:0 is more common",
		"00:00:01        Superlan        # SuperLAN-2U",
	}
	for _, v := range vendorLines {
		vr, err := parseLine(v)
		if err != nil {
			t.Fatal(err)
		}
		manuf.addPrefix(vr)
	}
	if len(manuf.CompanyList()) != len(vendorLines) {
		t.Fatalf("Company list is %d, expected %d", len(manuf.CompanyList()), len(vendorLines))
	} else {
		t.Logf("Vendor map populated with %d entries, CompanyList() returned same.", len(manuf.CompanyList()))
	}
}
