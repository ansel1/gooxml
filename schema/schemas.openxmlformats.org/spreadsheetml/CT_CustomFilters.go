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

type CT_CustomFilters struct {
	// And
	AndAttr *bool
	// Custom Filter Criteria
	CustomFilter []*CT_CustomFilter
}

func NewCT_CustomFilters() *CT_CustomFilters {
	ret := &CT_CustomFilters{}
	return ret
}

func (m *CT_CustomFilters) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.AndAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "and"},
			Value: fmt.Sprintf("%v", *m.AndAttr)})
	}
	e.EncodeToken(start)
	secustomFilter := xml.StartElement{Name: xml.Name{Local: "x:customFilter"}}
	e.EncodeElement(m.CustomFilter, secustomFilter)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CustomFilters) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "and" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AndAttr = &parsed
		}
	}
lCT_CustomFilters:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "customFilter":
				tmp := NewCT_CustomFilter()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CustomFilter = append(m.CustomFilter, tmp)
			default:
				log.Printf("skipping unsupported element on CT_CustomFilters %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CustomFilters
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CustomFilters and its children
func (m *CT_CustomFilters) Validate() error {
	return m.ValidateWithPath("CT_CustomFilters")
}

// ValidateWithPath validates the CT_CustomFilters and its children, prefixing error messages with path
func (m *CT_CustomFilters) ValidateWithPath(path string) error {
	for i, v := range m.CustomFilter {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CustomFilter[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
