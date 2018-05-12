package flag

import (
	"flag"
	"testing"
)

func TestFlagsrc_Read(t *testing.T) {
	dbhost := flag.String("database-host", "", "db host")
	dbpw := flag.String("database-password", "", "db pw")

	flag.Set("database-host", "localhost")
	flag.Set("database-password", "some-password")
	flag.Parse()

	source := NewSource()
	c, err := source.Read()
	if err != nil {
		t.Error(err)
	}

	var actual = c.Data

	actualDB := actual["database"].(map[string]interface{})
	if actualDB["host"] != *dbhost {
		t.Errorf("expected %v got %v", *dbhost, actualDB["host"])
	}

	if actualDB["password"] != *dbpw {
		t.Errorf("expected %v got %v", *dbpw, actualDB["password"])
	}
}
