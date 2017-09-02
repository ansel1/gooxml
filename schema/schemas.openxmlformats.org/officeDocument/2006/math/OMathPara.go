// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"
	"log"
)

type OMathPara struct {
	CT_OMathPara
}

func NewOMathPara() *OMathPara {
	ret := &OMathPara{}
	ret.CT_OMathPara = *NewCT_OMathPara()
	return ret
}

func (m *OMathPara) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/math"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:m"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/math"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:w"}, Value: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "m:oMathPara"
	return m.CT_OMathPara.MarshalXML(e, start)
}

func (m *OMathPara) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_OMathPara = *NewCT_OMathPara()
lOMathPara:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "oMathParaPr":
				m.OMathParaPr = NewCT_OMathParaPr()
				if err := d.DecodeElement(m.OMathParaPr, &el); err != nil {
					return err
				}
			case "oMath":
				tmp := NewCT_OMath()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.OMath = append(m.OMath, tmp)
			default:
				log.Printf("skipping unsupported element on OMathPara %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lOMathPara
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the OMathPara and its children
func (m *OMathPara) Validate() error {
	return m.ValidateWithPath("OMathPara")
}

// ValidateWithPath validates the OMathPara and its children, prefixing error messages with path
func (m *OMathPara) ValidateWithPath(path string) error {
	if err := m.CT_OMathPara.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}