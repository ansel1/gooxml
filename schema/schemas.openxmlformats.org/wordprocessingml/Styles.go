// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"log"
)

type Styles struct {
	CT_Styles
}

func NewStyles() *Styles {
	ret := &Styles{}
	ret.CT_Styles = *NewCT_Styles()
	return ret
}

func (m *Styles) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:m"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/math"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:ma"}, Value: "http://schemas.openxmlformats.org/schemaLibrary/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:pic"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/picture"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:w"}, Value: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:wp"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "w:styles"
	return m.CT_Styles.MarshalXML(e, start)
}

func (m *Styles) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Styles = *NewCT_Styles()
lStyles:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "docDefaults":
				m.DocDefaults = NewCT_DocDefaults()
				if err := d.DecodeElement(m.DocDefaults, &el); err != nil {
					return err
				}
			case "latentStyles":
				m.LatentStyles = NewCT_LatentStyles()
				if err := d.DecodeElement(m.LatentStyles, &el); err != nil {
					return err
				}
			case "style":
				tmp := NewCT_Style()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Style = append(m.Style, tmp)
			default:
				log.Printf("skipping unsupported element on Styles %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lStyles
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Styles and its children
func (m *Styles) Validate() error {
	return m.ValidateWithPath("Styles")
}

// ValidateWithPath validates the Styles and its children, prefixing error messages with path
func (m *Styles) ValidateWithPath(path string) error {
	if err := m.CT_Styles.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}