// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_Tables struct {
	// Count of Tables
	CountAttr *uint32
	// No Value
	M []*CT_TableMissing
	// Character Value
	S []*CT_XStringElement
	// Index
	X []*CT_Index
}

func NewCT_Tables() *CT_Tables {
	ret := &CT_Tables{}
	return ret
}

func (m *CT_Tables) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	if m.M != nil {
		sem := xml.StartElement{Name: xml.Name{Local: "x:m"}}
		e.EncodeElement(m.M, sem)
	}
	if m.S != nil {
		ses := xml.StartElement{Name: xml.Name{Local: "x:s"}}
		e.EncodeElement(m.S, ses)
	}
	if m.X != nil {
		sex := xml.StartElement{Name: xml.Name{Local: "x:x"}}
		e.EncodeElement(m.X, sex)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Tables) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
		}
	}
lCT_Tables:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "m":
				tmp := NewCT_TableMissing()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.M = append(m.M, tmp)
			case "s":
				tmp := NewCT_XStringElement()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.S = append(m.S, tmp)
			case "x":
				tmp := NewCT_Index()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.X = append(m.X, tmp)
			default:
				log.Printf("skipping unsupported element on CT_Tables %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Tables
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Tables and its children
func (m *CT_Tables) Validate() error {
	return m.ValidateWithPath("CT_Tables")
}

// ValidateWithPath validates the CT_Tables and its children, prefixing error messages with path
func (m *CT_Tables) ValidateWithPath(path string) error {
	for i, v := range m.M {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/M[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.S {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/S[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.X {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/X[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}