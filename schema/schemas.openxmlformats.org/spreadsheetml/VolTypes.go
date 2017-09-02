// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"log"
)

type VolTypes struct {
	CT_VolTypes
}

func NewVolTypes() *VolTypes {
	ret := &VolTypes{}
	ret.CT_VolTypes = *NewCT_VolTypes()
	return ret
}

func (m *VolTypes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:x"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xdr"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "x:volTypes"
	return m.CT_VolTypes.MarshalXML(e, start)
}

func (m *VolTypes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_VolTypes = *NewCT_VolTypes()
lVolTypes:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "volType":
				tmp := NewCT_VolType()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.VolType = append(m.VolType, tmp)
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on VolTypes %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lVolTypes
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the VolTypes and its children
func (m *VolTypes) Validate() error {
	return m.ValidateWithPath("VolTypes")
}

// ValidateWithPath validates the VolTypes and its children, prefixing error messages with path
func (m *VolTypes) ValidateWithPath(path string) error {
	if err := m.CT_VolTypes.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
