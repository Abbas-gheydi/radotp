// Code generated by radius-dict-gen. DO NOT EDIT.

package vendors

import (
	"net"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_Fortinet_VendorID = 12356
)

func _Fortinet_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_Fortinet_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _Fortinet_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Fortinet_VendorID {
			continue
		}
		for len(vsa) >= 3 {
			vsaTyp, vsaLen := vsa[0], vsa[1]
			if int(vsaLen) > len(vsa) || vsaLen < 3 {
				break
			}
			if vsaTyp == typ {
				values = append(values, vsa[2:int(vsaLen)])
			}
			vsa = vsa[int(vsaLen):]
		}
	}
	return
}

func _Fortinet_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Fortinet_VendorID {
			continue
		}
		for len(vsa) >= 3 {
			vsaTyp, vsaLen := vsa[0], vsa[1]
			if int(vsaLen) > len(vsa) || vsaLen < 3 {
				break
			}
			if vsaTyp == typ {
				return vsa[2:int(vsaLen)], true
			}
			vsa = vsa[int(vsaLen):]
		}
	}
	return
}

func _Fortinet_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _Fortinet_VendorID {
			i++
			continue
		}
		for j := 0; len(vsa[j:]) >= 3; {
			vsaTyp, vsaLen := vsa[0], vsa[1]
			if int(vsaLen) > len(vsa[j:]) || vsaLen < 3 {
				i++
				break
			}
			if vsaTyp == typ {
				vsa = append(vsa[:j], vsa[j+int(vsaLen):]...)
			}
			j += int(vsaLen)
		}
		if len(vsa) > 0 {
			copy(avp.Attribute[4:], vsa)
			i++
		} else {
			p.Attributes = append(p.Attributes[:i], p.Attributes[i+i:]...)
		}
	}
	return _Fortinet_AddVendor(p, typ, attr)
}

func _Fortinet_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _Fortinet_VendorID {
			i++
			continue
		}
		offset := 0
		for len(vsa[offset:]) >= 3 {
			vsaTyp, vsaLen := vsa[offset], vsa[offset+1]
			if int(vsaLen) > len(vsa) || vsaLen < 3 {
				continue vsaLoop
			}
			if vsaTyp == typ {
				copy(vsa[offset:], vsa[offset+int(vsaLen):])
				vsa = vsa[:len(vsa)-int(vsaLen)]
			} else {
				offset += int(vsaLen)
			}
		}
		if offset == 0 {
			p.Attributes = append(p.Attributes[:i], p.Attributes[i+1:]...)
		} else {
			i++
		}
	}
	return
}

func FortinetGroupName_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Fortinet_AddVendor(p, 1, a)
}

func FortinetGroupName_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Fortinet_AddVendor(p, 1, a)
}

func FortinetGroupName_Get(p *radius.Packet) (value []byte) {
	value, _ = FortinetGroupName_Lookup(p)
	return
}

func FortinetGroupName_GetString(p *radius.Packet) (value string) {
	value, _ = FortinetGroupName_LookupString(p)
	return
}

func FortinetGroupName_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Fortinet_GetsVendor(p, 1) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func FortinetGroupName_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Fortinet_GetsVendor(p, 1) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func FortinetGroupName_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Fortinet_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func FortinetGroupName_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Fortinet_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func FortinetGroupName_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Fortinet_SetVendor(p, 1, a)
}

func FortinetGroupName_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Fortinet_SetVendor(p, 1, a)
}

func FortinetGroupName_Del(p *radius.Packet) {
	_Fortinet_DelVendor(p, 1)
}

func FortinetClientIPAddress_Add(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _Fortinet_AddVendor(p, 2, a)
}

func FortinetClientIPAddress_Get(p *radius.Packet) (value net.IP) {
	value, _ = FortinetClientIPAddress_Lookup(p)
	return
}

func FortinetClientIPAddress_Gets(p *radius.Packet) (values []net.IP, err error) {
	var i net.IP
	for _, attr := range _Fortinet_GetsVendor(p, 2) {
		i, err = radius.IPAddr(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func FortinetClientIPAddress_Lookup(p *radius.Packet) (value net.IP, err error) {
	a, ok := _Fortinet_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value, err = radius.IPAddr(a)
	return
}

func FortinetClientIPAddress_Set(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _Fortinet_SetVendor(p, 2, a)
}

func FortinetClientIPAddress_Del(p *radius.Packet) {
	_Fortinet_DelVendor(p, 2)
}

func FortinetVdomName_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Fortinet_AddVendor(p, 3, a)
}

func FortinetVdomName_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Fortinet_AddVendor(p, 3, a)
}

func FortinetVdomName_Get(p *radius.Packet) (value []byte) {
	value, _ = FortinetVdomName_Lookup(p)
	return
}

func FortinetVdomName_GetString(p *radius.Packet) (value string) {
	value, _ = FortinetVdomName_LookupString(p)
	return
}

func FortinetVdomName_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Fortinet_GetsVendor(p, 3) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func FortinetVdomName_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Fortinet_GetsVendor(p, 3) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func FortinetVdomName_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Fortinet_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func FortinetVdomName_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Fortinet_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func FortinetVdomName_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Fortinet_SetVendor(p, 3, a)
}

func FortinetVdomName_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Fortinet_SetVendor(p, 3, a)
}

func FortinetVdomName_Del(p *radius.Packet) {
	_Fortinet_DelVendor(p, 3)
}

func FortinetClientIPv6Address_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Fortinet_AddVendor(p, 4, a)
}

func FortinetClientIPv6Address_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Fortinet_AddVendor(p, 4, a)
}

func FortinetClientIPv6Address_Get(p *radius.Packet) (value []byte) {
	value, _ = FortinetClientIPv6Address_Lookup(p)
	return
}

func FortinetClientIPv6Address_GetString(p *radius.Packet) (value string) {
	value, _ = FortinetClientIPv6Address_LookupString(p)
	return
}

func FortinetClientIPv6Address_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Fortinet_GetsVendor(p, 4) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func FortinetClientIPv6Address_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Fortinet_GetsVendor(p, 4) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func FortinetClientIPv6Address_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Fortinet_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func FortinetClientIPv6Address_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Fortinet_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func FortinetClientIPv6Address_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Fortinet_SetVendor(p, 4, a)
}

func FortinetClientIPv6Address_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Fortinet_SetVendor(p, 4, a)
}

func FortinetClientIPv6Address_Del(p *radius.Packet) {
	_Fortinet_DelVendor(p, 4)
}

func FortinetInterfaceName_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Fortinet_AddVendor(p, 5, a)
}

func FortinetInterfaceName_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Fortinet_AddVendor(p, 5, a)
}

func FortinetInterfaceName_Get(p *radius.Packet) (value []byte) {
	value, _ = FortinetInterfaceName_Lookup(p)
	return
}

func FortinetInterfaceName_GetString(p *radius.Packet) (value string) {
	value, _ = FortinetInterfaceName_LookupString(p)
	return
}

func FortinetInterfaceName_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Fortinet_GetsVendor(p, 5) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func FortinetInterfaceName_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Fortinet_GetsVendor(p, 5) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func FortinetInterfaceName_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Fortinet_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func FortinetInterfaceName_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Fortinet_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func FortinetInterfaceName_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Fortinet_SetVendor(p, 5, a)
}

func FortinetInterfaceName_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Fortinet_SetVendor(p, 5, a)
}

func FortinetInterfaceName_Del(p *radius.Packet) {
	_Fortinet_DelVendor(p, 5)
}

func FortinetAccessProfile_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Fortinet_AddVendor(p, 6, a)
}

func FortinetAccessProfile_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Fortinet_AddVendor(p, 6, a)
}

func FortinetAccessProfile_Get(p *radius.Packet) (value []byte) {
	value, _ = FortinetAccessProfile_Lookup(p)
	return
}

func FortinetAccessProfile_GetString(p *radius.Packet) (value string) {
	value, _ = FortinetAccessProfile_LookupString(p)
	return
}

func FortinetAccessProfile_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Fortinet_GetsVendor(p, 6) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func FortinetAccessProfile_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Fortinet_GetsVendor(p, 6) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func FortinetAccessProfile_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Fortinet_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func FortinetAccessProfile_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Fortinet_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func FortinetAccessProfile_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Fortinet_SetVendor(p, 6, a)
}

func FortinetAccessProfile_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Fortinet_SetVendor(p, 6, a)
}

func FortinetAccessProfile_Del(p *radius.Packet) {
	_Fortinet_DelVendor(p, 6)
}
