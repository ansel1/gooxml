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

type CT_ConditionalFormats struct {
	// Conditional Format Count
	CountAttr *uint32
	// Conditional Formatting
	ConditionalFormat []*CT_ConditionalFormat
}

func NewCT_ConditionalFormats() *CT_ConditionalFormats {
	ret := &CT_ConditionalFormats{}
	return ret
}

func (m *CT_ConditionalFormats) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	seconditionalFormat := xml.StartElement{Name: xml.Name{Local: "x:conditionalFormat"}}
	e.EncodeElement(m.ConditionalFormat, seconditionalFormat)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ConditionalFormats) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
lCT_ConditionalFormats:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "conditionalFormat":
				tmp := NewCT_ConditionalFormat()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ConditionalFormat = append(m.ConditionalFormat, tmp)
			default:
				log.Printf("skipping unsupported element on CT_ConditionalFormats %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ConditionalFormats
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ConditionalFormats and its children
func (m *CT_ConditionalFormats) Validate() error {
	return m.ValidateWithPath("CT_ConditionalFormats")
}

// ValidateWithPath validates the CT_ConditionalFormats and its children, prefixing error messages with path
func (m *CT_ConditionalFormats) ValidateWithPath(path string) error {
	for i, v := range m.ConditionalFormat {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ConditionalFormat[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}