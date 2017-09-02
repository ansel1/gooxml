// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"log"
)

type Theme struct {
	CT_OfficeStyleSheet
}

func NewTheme() *Theme {
	ret := &Theme{}
	ret.CT_OfficeStyleSheet = *NewCT_OfficeStyleSheet()
	return ret
}

func (m *Theme) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "a:theme"
	return m.CT_OfficeStyleSheet.MarshalXML(e, start)
}

func (m *Theme) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_OfficeStyleSheet = *NewCT_OfficeStyleSheet()
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
		}
	}
lTheme:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "themeElements":
				if err := d.DecodeElement(m.ThemeElements, &el); err != nil {
					return err
				}
			case "objectDefaults":
				m.ObjectDefaults = NewCT_ObjectStyleDefaults()
				if err := d.DecodeElement(m.ObjectDefaults, &el); err != nil {
					return err
				}
			case "extraClrSchemeLst":
				m.ExtraClrSchemeLst = NewCT_ColorSchemeList()
				if err := d.DecodeElement(m.ExtraClrSchemeLst, &el); err != nil {
					return err
				}
			case "custClrLst":
				m.CustClrLst = NewCT_CustomColorList()
				if err := d.DecodeElement(m.CustClrLst, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on Theme %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lTheme
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Theme and its children
func (m *Theme) Validate() error {
	return m.ValidateWithPath("Theme")
}

// ValidateWithPath validates the Theme and its children, prefixing error messages with path
func (m *Theme) ValidateWithPath(path string) error {
	if err := m.CT_OfficeStyleSheet.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
