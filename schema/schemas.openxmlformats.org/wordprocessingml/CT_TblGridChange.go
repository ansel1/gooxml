// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_TblGridChange struct {
	// Annotation Identifier
	IdAttr  int64
	TblGrid *CT_TblGridBase
}

func NewCT_TblGridChange() *CT_TblGridChange {
	ret := &CT_TblGridChange{}
	ret.TblGrid = NewCT_TblGridBase()
	return ret
}

func (m *CT_TblGridChange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	e.EncodeToken(start)
	setblGrid := xml.StartElement{Name: xml.Name{Local: "w:tblGrid"}}
	e.EncodeElement(m.TblGrid, setblGrid)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TblGridChange) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TblGrid = NewCT_TblGridBase()
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
		}
	}
lCT_TblGridChange:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tblGrid":
				if err := d.DecodeElement(m.TblGrid, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_TblGridChange %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TblGridChange
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TblGridChange and its children
func (m *CT_TblGridChange) Validate() error {
	return m.ValidateWithPath("CT_TblGridChange")
}

// ValidateWithPath validates the CT_TblGridChange and its children, prefixing error messages with path
func (m *CT_TblGridChange) ValidateWithPath(path string) error {
	if err := m.TblGrid.ValidateWithPath(path + "/TblGrid"); err != nil {
		return err
	}
	return nil
}
