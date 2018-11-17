package config

import (
	"testing"

	"github.com/cqu903/bobcat/config"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfit(t *testing.T) {
	config.Config()
	if config.Conf.Port == 0 {
		t.Fail()
	}
	assert.NotEqual(t, 0, config.Conf.Port, "config.Conf.Port is 0,the init has wrong")
}
