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

type CT_PivotAreaReferences struct {
	// Pivot Filter Count
	CountAttr *uint32
	// Reference
	Reference []*CT_PivotAreaReference
}

func NewCT_PivotAreaReferences() *CT_PivotAreaReferences {
	ret := &CT_PivotAreaReferences{}
	return ret
}

func (m *CT_PivotAreaReferences) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	sereference := xml.StartElement{Name: xml.Name{Local: "x:reference"}}
	e.EncodeElement(m.Reference, sereference)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PivotAreaReferences) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
lCT_PivotAreaReferences:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "reference":
				tmp := NewCT_PivotAreaReference()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Reference = append(m.Reference, tmp)
			default:
				log.Printf("skipping unsupported element on CT_PivotAreaReferences %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PivotAreaReferences
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PivotAreaReferences and its children
func (m *CT_PivotAreaReferences) Validate() error {
	return m.ValidateWithPath("CT_PivotAreaReferences")
}

// ValidateWithPath validates the CT_PivotAreaReferences and its children, prefixing error messages with path
func (m *CT_PivotAreaReferences) ValidateWithPath(path string) error {
	for i, v := range m.Reference {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Reference[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
