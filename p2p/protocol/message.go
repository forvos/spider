package protocol

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
