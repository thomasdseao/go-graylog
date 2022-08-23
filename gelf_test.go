package graylog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewGelf(t *testing.T) {
	gelf := NewGelf(Config{
		"graylog1.example.com",
		2202,
		UDP,
		true,
	})

	assert.Equal(t, gelf.Config.Port, 2202)
	assert.Equal(t, gelf.Config.Hostname, "graylog1.example.com")
	assert.Equal(t, gelf.Config.Transport, UDP)
	assert.Equal(t, gelf.Config.ErrorLog, true)
}

func Test_Gelf_ResolveUDPAddr_Error(t *testing.T) {
	gelf := NewGelf(Config{
		"graylog112.example.com",
		22000,
		UDP,
		true,
	})

	send, err := gelf.Send(Message{
		Version:      "1.1",
		Host:         "example.com",
		ShortMessage: "This is the short message",
	})

	assert.NotEqual(t, err, nil)
	assert.Equal(t, send, false)
}

func Test_Gelf_ResolveTCPAddr_Error(t *testing.T) {
	gelf := NewGelf(Config{
		"graylog112.example.com",
		22001,
		TCP,
		true,
	})

	send, err := gelf.Send(Message{
		Version:      "1.1",
		Host:         "example.com",
		ShortMessage: "This is the short message",
	})

	assert.NotEqual(t, err, nil)
	assert.Equal(t, send, false)
}
