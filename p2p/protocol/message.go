package protocol

import (
	"encoding/binary"
	"errors"
)

// Head protocol head 8+16+8+8+64+64 = 168 bit len
// 1+2+1+1+8+8 = 21 byte
type Head struct {
	version uint8
	magic   uint16
	msgType uint8
	gz      uint8
	reqId   uint64
	len     uint64
}

func NewHead(reqId, len uint64, msgType uint8) Head {
	return Head{
		version: version,
		magic:   magic,
		msgType: msgType,
		gz:      NO,
		reqId:   reqId,
		len:     len,
	}
}

func (h *Head) MsgType() uint8 {
	return h.msgType
}

func (h *Head) ReqId() uint64 {
	return h.reqId
}

func (h *Head) Len() uint64 {
	return h.len
}

func (h *Head) Gz() uint8 {
	return h.gz
}

// Message Custom protocol msg type
type Message struct {
	head    Head
	PayLoad []byte
}

func NewMessage(head Head, payLoad []byte) *Message {
	return &Message{
		head:    head,
		PayLoad: payLoad,
	}
}

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
