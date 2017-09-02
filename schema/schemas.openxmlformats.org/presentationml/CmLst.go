// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"log"
)

type CmLst struct {
	CT_CommentList
}

func NewCmLst() *CmLst {
	ret := &CmLst{}
	ret.CT_CommentList = *NewCT_CommentList()
	return ret
}

func (m *CmLst) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/presentationml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:p"}, Value: "http://schemas.openxmlformats.org/presentationml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "p:cmLst"
	return m.CT_CommentList.MarshalXML(e, start)
}

func (m *CmLst) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_CommentList = *NewCT_CommentList()
lCmLst:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cm":
				tmp := NewCT_Comment()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Cm = append(m.Cm, tmp)
			default:
				log.Printf("skipping unsupported element on CmLst %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCmLst
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CmLst and its children
func (m *CmLst) Validate() error {
	return m.ValidateWithPath("CmLst")
}

// ValidateWithPath validates the CmLst and its children, prefixing error messages with path
func (m *CmLst) ValidateWithPath(path string) error {
	if err := m.CT_CommentList.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}