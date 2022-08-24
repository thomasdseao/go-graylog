package graylog

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"log"
	"net"
	"strconv"
)

type Transport string

type Config struct {
	Hostname  string
	Port      int
	Transport Transport
	ErrorLog  bool
}

const (
	UDP Transport = "udp"
	TCP Transport = "tcp"
)

type Message struct {
	Version      string                 `json:"version"`
	Host         string                 `json:"host"`
	ShortMessage string                 `json:"short_message"`
	FullMessage  string                 `json:"full_message,omitempty"`
	Timestamp    int64                  `json:"timestamp,omitempty"`
	Level        uint                   `json:"level,omitempty"`
	Extra        map[string]interface{} `json:"-,"`
}

type Gelf struct {
	Config Config
}

func NewGelf(config Config) *Gelf {
	gelf := &Gelf{
		Config{
			Hostname:  config.Hostname,
			Port:      config.Port,
			Transport: config.Transport,
			ErrorLog:  config.ErrorLog,
		},
	}

	return gelf
}

func (gelf *Gelf) compress(byte []byte) bytes.Buffer {
	var buffer bytes.Buffer
	comp := zlib.NewWriter(&buffer)

	comp.Write(byte)
	comp.Close()

	return buffer
}

func (gelf *Gelf) Send(message []byte) (bool, error) {
	messageStruct := Message{}
	err := json.Unmarshal(message, &messageStruct)
	if err != nil {
		if gelf.Config.ErrorLog {
			log.Printf("Unable to encode the message : %s", err)
		}
		return false, err
	}

	if err != nil {
		if gelf.Config.ErrorLog {
			log.Printf("Unable to encode the message : %s", err)
		}
		return false, err
	}
	compressed := gelf.compress(message)

	var addr = gelf.Config.Hostname + ":" + strconv.Itoa(gelf.Config.Port)
	var conn net.Conn

	if gelf.Config.Transport == TCP {
		tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
		if err != nil {
			if gelf.Config.ErrorLog {
				log.Printf("Unable to resolve TCP address : %s", err)
			}
			return false, err
		}
		conn, err = net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			if gelf.Config.ErrorLog {
				log.Printf("TCP Transport error : %s", err)
			}
			return false, err
		}
	} else {
		udpAddr, err := net.ResolveUDPAddr("udp", addr)
		if err != nil {
			if gelf.Config.ErrorLog {
				log.Printf("Unable to resolve UDP address : %s", err)
			}
			return false, err
		}
		conn, err = net.DialUDP("udp", nil, udpAddr)
		if err != nil {
			if gelf.Config.ErrorLog {
				log.Printf("UDP Transport error : %s", err)
			}
			return false, err
		}
	}

	_, err = conn.Write(compressed.Bytes())
	if err != nil {
		if gelf.Config.ErrorLog {
			log.Printf("Error during sending message : %s", err)
		}
		return false, err
	}
	conn.Close()

	return true, nil
}
