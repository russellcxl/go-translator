package config

import (
	"fmt"
	"testing"
)

func Test_LoadConfig(t *testing.T) {
	c, e := LoadConfig("config.json")
	if e != nil {
		t.Error(e)
	}
	fmt.Println(c)
}
