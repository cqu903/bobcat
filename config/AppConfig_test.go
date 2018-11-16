package config

import (
	"testing"

	"github.com/cqu903/bobcat/config"
)

func TestLoadConfit(t *testing.T) {
	config.Config()
	if config.Conf.Port == 0 {
		t.Fail()
	}
}
