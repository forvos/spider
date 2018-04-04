package protocol

import (
	"encoding/binary"
	"errors"
)

const version = 0x01
const magic = 0x594C

const (
	NO  = 0x00
	ZIP = 0x01
)

type Protocol struct{}

func (p *Protocol) DeCode(data []byte) (*Message, error) {
	head := Head{}
	if len(data) < 21 {
		return nil, errors.New("length too small")
	}

	b := binary.BigEndian
	head.version = data[0]
	if uint(head.version) != version {
		return nil, errors.New("protocol version error")
	}

	head.magic = b.Uint16(data[1:3])
	if head.magic != magic {
		return nil, errors.New("magic error")
	}

	head.msgType = data[3]
	head.gz = data[4]
	head.reqId = b.Uint64(data[5:13])
	head.len = b.Uint64(data[13:21])

	if uint64(len(data)) < 21+head.len {
		return nil, errors.New("length too small")
	}

	return &Message{
		head:    head,
		PayLoad: data[21 : 21+head.len],
	}, nil
}

func (p *Protocol) EnCode(msg *Message) ([]byte, error) {
	var (
		b    = binary.BigEndian
		data = make([]byte, 21)
	)
	data[0] = version
	b.PutUint16(data[1:3], magic)
	data[3] = msg.head.msgType
	data[4] = msg.head.gz
	b.PutUint64(data[5:13], msg.head.reqId)
	b.PutUint64(data[13:], msg.head.len)
	data = append(data, msg.PayLoad...)
	return data, nil
}
