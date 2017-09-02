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

type CT_Dimensions struct {
	// OLAP Dimensions Count
	CountAttr *uint32
	// OLAP Dimension
	Dimension []*CT_PivotDimension
}

func NewCT_Dimensions() *CT_Dimensions {
	ret := &CT_Dimensions{}
	return ret
}

func (m *CT_Dimensions) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	if m.Dimension != nil {
		sedimension := xml.StartElement{Name: xml.Name{Local: "x:dimension"}}
		e.EncodeElement(m.Dimension, sedimension)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Dimensions) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
lCT_Dimensions:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "dimension":
				tmp := NewCT_PivotDimension()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Dimension = append(m.Dimension, tmp)
			default:
				log.Printf("skipping unsupported element on CT_Dimensions %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Dimensions
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Dimensions and its children
func (m *CT_Dimensions) Validate() error {
	return m.ValidateWithPath("CT_Dimensions")
}

// ValidateWithPath validates the CT_Dimensions and its children, prefixing error messages with path
func (m *CT_Dimensions) ValidateWithPath(path string) error {
	for i, v := range m.Dimension {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Dimension[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}