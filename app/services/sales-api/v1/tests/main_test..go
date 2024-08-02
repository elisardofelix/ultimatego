package tests

import (
	"fmt"
	"testing"

	"github.com/elisardofelix/ultimatego/business/data/dbtest"
	"github.com/elisardofelix/ultimatego/foundation/docker"
)

var c *docker.Container

func TestMain(m *testing.M) {
	var err error
	c, err = dbtest.StartDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dbtest.StopDB(c)

	m.Run()
}
