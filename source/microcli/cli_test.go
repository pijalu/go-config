package microcli

import (
	"fmt"
	"testing"

	"github.com/micro/cli"
	"github.com/pijalu/go-config/source"
)

func TestClisrc_Read(t *testing.T) {
	var src source.Source
	app := cli.NewApp()
	app.Name = "testapp"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "db-host"},
	}
	app.Action = func(c *cli.Context) {
		src = NewSource(c)
	}
	app.Run([]string{"run", "-db-host", "localhost"})

	c, err := src.Read()
	if err != nil {
		t.Error(err)
	}

	actual := c.Data
	actualDB := actual["db"].(map[string]interface{})
	actualValue := actualDB["host"].(fmt.Stringer).String()

	if actualValue != "localhost" {
		t.Errorf("expected localhost, got %s", actualDB["name"])
	}
}
