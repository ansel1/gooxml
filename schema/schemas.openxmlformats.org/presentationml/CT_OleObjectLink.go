// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_OleObjectLink struct {
	// Update Linked Embedded Objects Automatically
	UpdateAutomaticAttr *bool
	ExtLst              *CT_ExtensionList
}

func NewCT_OleObjectLink() *CT_OleObjectLink {
	ret := &CT_OleObjectLink{}
	return ret
}

func (m *CT_OleObjectLink) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.UpdateAutomaticAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "updateAutomatic"},
			Value: fmt.Sprintf("%v", *m.UpdateAutomaticAttr)})
	}
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_OleObjectLink) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "updateAutomatic" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UpdateAutomaticAttr = &parsed
		}
	}
lCT_OleObjectLink:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_OleObjectLink %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_OleObjectLink
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_OleObjectLink and its children
func (m *CT_OleObjectLink) Validate() error {
	return m.ValidateWithPath("CT_OleObjectLink")
}

// ValidateWithPath validates the CT_OleObjectLink and its children, prefixing error messages with path
func (m *CT_OleObjectLink) ValidateWithPath(path string) error {
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}