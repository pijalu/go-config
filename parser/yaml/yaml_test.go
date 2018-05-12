package yaml

import "testing"

func testParseGeneral(t *testing.T, data interface{}) {
	p := Parser{}
	m, err := p.parseYAML(data)
	if err != nil {
		t.Fatalf("Error parsing yaml: %v", err)
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

func TestParseYAML(t *testing.T) {
	// Test parsing as a string
	testParseGeneral(t, `
key:
    a: "1"
    b: "2"
`)
}

func TestParseYAMLasBytes(t *testing.T) {
	// Test parsing as an array of bytes
	testParseGeneral(t, []byte(`
key:
    a: "1"
    b: "2"
`))
}

func TestParseWithWrongData(t *testing.T) {
	p := Parser{}
	_, err := p.parseYAML(42)
	if err == nil {
		t.Fatal("Expected error parsing but go none !")
	}
}
