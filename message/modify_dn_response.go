package message

import "fmt"

//
//        ModifyDNResponse ::= [APPLICATION 13] LDAPResult
func readModifyDNResponse(bytes *Bytes) (ret ModifyDNResponse, err error) {
	var res LDAPResult
	res, err = readTaggedLDAPResult(bytes, classApplication, TagModifyDNResponse)
	if err != nil {
		err = LdapError{fmt.Sprintf("readModifyDNResponse:\n%s", err.Error())}
		return
	}
	ret = ModifyDNResponse{
		LDAPResult: res,
	}
	return
}

//
//        ModifyDNResponse ::= [APPLICATION 13] LDAPResult
func (m ModifyDNResponse) write(bytes *Bytes) int {
	return m.LDAPResult.writeTagged(bytes, classApplication, TagModifyDNResponse)
}

//
//        ModifyDNResponse ::= [APPLICATION 13] LDAPResult
func (m ModifyDNResponse) size() int {
	return m.LDAPResult.sizeTagged(TagModifyDNResponse)
}
