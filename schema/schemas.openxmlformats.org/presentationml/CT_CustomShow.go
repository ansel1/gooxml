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

type CT_CustomShow struct {
	// Custom Show Name
	NameAttr string
	// Custom Show ID
	IdAttr uint32
	// List of Presentation Slides
	SldLst *CT_SlideRelationshipList
	ExtLst *CT_ExtensionList
}

func NewCT_CustomShow() *CT_CustomShow {
	ret := &CT_CustomShow{}
	ret.SldLst = NewCT_SlideRelationshipList()
	return ret
}

func (m *CT_CustomShow) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	e.EncodeToken(start)
	sesldLst := xml.StartElement{Name: xml.Name{Local: "p:sldLst"}}
	e.EncodeElement(m.SldLst, sesldLst)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CustomShow) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.SldLst = NewCT_SlideRelationshipList()
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
		}
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.IdAttr = uint32(parsed)
		}
	}
lCT_CustomShow:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "sldLst":
				if err := d.DecodeElement(m.SldLst, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_CustomShow %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CustomShow
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CustomShow and its children
func (m *CT_CustomShow) Validate() error {
	return m.ValidateWithPath("CT_CustomShow")
}

// ValidateWithPath validates the CT_CustomShow and its children, prefixing error messages with path
func (m *CT_CustomShow) ValidateWithPath(path string) error {
	if err := m.SldLst.ValidateWithPath(path + "/SldLst"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
