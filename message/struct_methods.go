package message

func NewLDAPMessageWithProtocolOp(po ProtocolResponse) *LDAPMessage {
	m := NewLDAPMessage()
	m.protocolOp = po
	controls := po.controls()
	if controls != nil {
		m.controls = &controls
	}
	return m
}
