package asn1go

import (
	"testing"
)

func TestUnmarshalAsn1(t *testing.T) {
	var asn1Blob = []byte(`value1 ProfileElement ::= header : {
  major-version 2,
  minor-version 1,
  profileType "GSMA Generic eUICC Test Profile",
  iccid '89000123456789012341'H,
  eUICC-Mandatory-services {
    usim NULL,
    isim NULL,
    csim NULL,
    usim-test-algorithm NULL,
    ber-tlv NULL
  },
  eUICC-Mandatory-GFSTEList {
    { 2 23 143 1 2 1 },
    { 2 23 143 1 2 3 },
    { 2 23 143 1 2 4 },
    { 2 23 143 1 2 5 },
    { 2 23 143 1 2 7 },
    { 2 23 143 1 2 8 },
    { 2 23 143 1 2 9 },
    { 2 23 143 1 2 10 },
    { 2 23 143 1 2 11 }
  }
}

`)

	type ProfileElement struct {
		Header struct {
			MajorVersion int
			MinorVersion int
			ProfileType  string
			Iccid        string
		}
	}
	var profileElement ProfileElement

	err := Unmarshal(asn1Blob, &profileElement)
	if err != nil {
		t.Error("error:", err)
	}
	t.Logf("%+v", profileElement)

}
