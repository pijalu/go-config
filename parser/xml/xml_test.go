package xml

import "testing"

const XML = `
<?xml version=1.0?>
<key>
 <a>1</a>
 <b>2</b>
</key>
`

func testParseGeneral(t *testing.T, data interface{}) {
	p := Parser{}
	m, err := p.parseXML(data)
	if err != nil {
		t.Fatalf("Error parsing XML: %v", err)
	}

	// Check that data is loaded correctly (rest should be covered by noop)
	actualKeyA := m["key"].(map[string]interface{})["a"]
	actualKeyB := m["key"].(map[string]interface{})["b"]

	if actualKeyA != "1" {
		t.Fatalf("key.a expected %v but is %v (from %v)", "1", actualKeyA, m)
	}

	if actualKeyB != "2" {
		t.Fatalf("key.b expected %v but is %v (from %v)", "2", actualKeyB, m)
	}
}

func TestParseXML(t *testing.T) {
	// Test parsing as a string
	testParseGeneral(t, XML)
}

func TestParseXMLasBytes(t *testing.T) {
	// Test parsing as an array of bytes
	testParseGeneral(t, []byte(XML))
}

func TestParseWithWrongData(t *testing.T) {
	p := Parser{}
	_, err := p.parseXML(42)
	if err == nil {
		t.Fatal("Expected error parsing but go none !")
	}
}
