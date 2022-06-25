package config

import (
	"fmt"
	"testing"
)

func TestLoadFromJson(t *testing.T) {
	cfgFilePath := "../../etc/test.json"
	cfg := DefaultConfig()
	err := cfg.LoadFromJson(cfgFilePath)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(cfg)
}
