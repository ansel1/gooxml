// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_Table struct {
	TblPr   *CT_TableProperties
	TblGrid *CT_TableGrid
	Tr      []*CT_TableRow
}

func NewCT_Table() *CT_Table {
	ret := &CT_Table{}
	ret.TblGrid = NewCT_TableGrid()
	return ret
}

func (m *CT_Table) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.TblPr != nil {
		setblPr := xml.StartElement{Name: xml.Name{Local: "a:tblPr"}}
		e.EncodeElement(m.TblPr, setblPr)
	}
	setblGrid := xml.StartElement{Name: xml.Name{Local: "a:tblGrid"}}
	e.EncodeElement(m.TblGrid, setblGrid)
	if m.Tr != nil {
		setr := xml.StartElement{Name: xml.Name{Local: "a:tr"}}
		e.EncodeElement(m.Tr, setr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Table) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TblGrid = NewCT_TableGrid()
lCT_Table:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tblPr":
				m.TblPr = NewCT_TableProperties()
				if err := d.DecodeElement(m.TblPr, &el); err != nil {
					return err
				}
			case "tblGrid":
				if err := d.DecodeElement(m.TblGrid, &el); err != nil {
					return err
				}
			case "tr":
				tmp := NewCT_TableRow()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Tr = append(m.Tr, tmp)
			default:
				log.Printf("skipping unsupported element on CT_Table %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Table
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Table and its children
func (m *CT_Table) Validate() error {
	return m.ValidateWithPath("CT_Table")
}

// ValidateWithPath validates the CT_Table and its children, prefixing error messages with path
func (m *CT_Table) ValidateWithPath(path string) error {
	if m.TblPr != nil {
		if err := m.TblPr.ValidateWithPath(path + "/TblPr"); err != nil {
			return err
		}
	}
	if err := m.TblGrid.ValidateWithPath(path + "/TblGrid"); err != nil {
		return err
	}
	for i, v := range m.Tr {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Tr[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}