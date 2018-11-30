package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConf(t *testing.T) {
	conf := struct {
		Test string `yaml:"test"`
	}{}

	confPath = "configtest.yml"
	err := Conf(&conf)
	assert.Nil(t, err)

	expected := "test"
	assert.Equal(t, expected, conf.Test)
}
