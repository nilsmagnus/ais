package ais

import (
	"github.com/adrianmo/go-nmea"
	"testing"
)

func TestDecodePackage5(t *testing.T) {

	line := "\\s:2573345,c:1692279193*06\\!BSVDM,2,1,4,B,53nai482E4Shh@td0010E8hDp00000000000001?0`<665Uf:61RDj1PDSDh,0*48"
	packet, err := nmea.Parse(line)

	m, ok := packet.(nmea.VDMVDO)

	if !ok {
		t.Fatalf("Expected VDMVDO packet")
	}

	p, err := Decode(m.Payload)
	if err != nil {
		t.Fatalf("Error decoding package: %s", err.Error())
	}
	if p == nil {
		t.Fatalf("Expected non-nil package")
	}

	if p.GetMessageId() != 5 {
		t.Fatalf("Expected message id 5, got %d", p.GetMessageId())
	}

	sd, ok := p.(ShipStaticAndVoyageRelatedData)
	if !ok {
		t.Fatalf("Expected ShipStaticAndVoyageRelatedData")
	}

	if sd.RepeatIndicator != 0 {
		t.Fatalf("Expected repeat indicator 0, got %d", sd.RepeatIndicator)
	}

	if sd.UserID != 258634000 {
		t.Fatalf("Expected user id 258634000, got %d", sd.UserID)
	}

	if sd.AisVersion != 2 {
		t.Fatalf("Expected ais version 0, got %d", sd.AisVersion)
	}

	if sd.ImoNumber != 9769532 {
		t.Fatalf("Expected imo number 9769532, got %d", sd.ImoNumber)
	}

	if sd.CallSign != "LDOK" {
		t.Fatalf("Expected call sign LDOK, got %s", sd.CallSign)
	}

	if sd.Name != "PERLEN" {
		t.Fatalf("Expected call sign LDOK, got %s", sd.CallSign)
	}

	if sd.ShipType == 0 {
		t.Fatalf("shiptype should not be 0")
	}

}
