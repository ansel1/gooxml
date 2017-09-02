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

type CT_MeasureGroups struct {
	// Measure Group Count
	CountAttr *uint32
	// OLAP Measure Group
	MeasureGroup []*CT_MeasureGroup
}

func NewCT_MeasureGroups() *CT_MeasureGroups {
	ret := &CT_MeasureGroups{}
	return ret
}

func (m *CT_MeasureGroups) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	if m.MeasureGroup != nil {
		semeasureGroup := xml.StartElement{Name: xml.Name{Local: "x:measureGroup"}}
		e.EncodeElement(m.MeasureGroup, semeasureGroup)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_MeasureGroups) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
lCT_MeasureGroups:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "measureGroup":
				tmp := NewCT_MeasureGroup()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.MeasureGroup = append(m.MeasureGroup, tmp)
			default:
				log.Printf("skipping unsupported element on CT_MeasureGroups %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_MeasureGroups
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_MeasureGroups and its children
func (m *CT_MeasureGroups) Validate() error {
	return m.ValidateWithPath("CT_MeasureGroups")
}

// ValidateWithPath validates the CT_MeasureGroups and its children, prefixing error messages with path
func (m *CT_MeasureGroups) ValidateWithPath(path string) error {
	for i, v := range m.MeasureGroup {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/MeasureGroup[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
