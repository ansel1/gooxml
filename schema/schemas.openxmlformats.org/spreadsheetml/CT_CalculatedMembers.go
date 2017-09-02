// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_CalculatedMembers struct {
	// Calculated Members Count
	CountAttr *uint32
	// Calculated Member
	CalculatedMember []*CT_CalculatedMember
}

func NewCT_CalculatedMembers() *CT_CalculatedMembers {
	ret := &CT_CalculatedMembers{}
	return ret
}

func (m *CT_CalculatedMembers) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	secalculatedMember := xml.StartElement{Name: xml.Name{Local: "x:calculatedMember"}}
	e.EncodeElement(m.CalculatedMember, secalculatedMember)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CalculatedMembers) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
		}
	}
lCT_CalculatedMembers:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "calculatedMember":
				tmp := NewCT_CalculatedMember()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CalculatedMember = append(m.CalculatedMember, tmp)
			default:
				log.Printf("skipping unsupported element on CT_CalculatedMembers %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CalculatedMembers
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CalculatedMembers and its children
func (m *CT_CalculatedMembers) Validate() error {
	return m.ValidateWithPath("CT_CalculatedMembers")
}

// ValidateWithPath validates the CT_CalculatedMembers and its children, prefixing error messages with path
func (m *CT_CalculatedMembers) ValidateWithPath(path string) error {
	for i, v := range m.CalculatedMember {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CalculatedMember[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
