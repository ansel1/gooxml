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

type MapInfo struct {
	CT_MapInfo
}

func NewMapInfo() *MapInfo {
	ret := &MapInfo{}
	ret.CT_MapInfo = *NewCT_MapInfo()
	return ret
}

func (m *MapInfo) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:x"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xdr"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "x:MapInfo"
	return m.CT_MapInfo.MarshalXML(e, start)
}

func (m *MapInfo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_MapInfo = *NewCT_MapInfo()
	for _, attr := range start.Attr {
		if attr.Name.Local == "SelectionNamespaces" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SelectionNamespacesAttr = parsed
		}
	}
lMapInfo:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "Schema":
				tmp := NewCT_Schema()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Schema = append(m.Schema, tmp)
			case "Map":
				tmp := NewCT_Map()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Map = append(m.Map, tmp)
			default:
				log.Printf("skipping unsupported element on MapInfo %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lMapInfo
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the MapInfo and its children
func (m *MapInfo) Validate() error {
	return m.ValidateWithPath("MapInfo")
}

// ValidateWithPath validates the MapInfo and its children, prefixing error messages with path
func (m *MapInfo) ValidateWithPath(path string) error {
	if err := m.CT_MapInfo.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
