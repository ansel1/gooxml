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

type CT_PCDSDTCEntries struct {
	// Tuple Count
	CountAttr *uint32
	// No Value
	M []*CT_Missing
	// Numeric Value
	N []*CT_Number
	// Error Value
	E []*CT_Error
	// Character Value
	S []*CT_String
}

func NewCT_PCDSDTCEntries() *CT_PCDSDTCEntries {
	ret := &CT_PCDSDTCEntries{}
	return ret
}

func (m *CT_PCDSDTCEntries) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	if m.M != nil {
		sem := xml.StartElement{Name: xml.Name{Local: "x:m"}}
		e.EncodeElement(m.M, sem)
	}
	if m.N != nil {
		sen := xml.StartElement{Name: xml.Name{Local: "x:n"}}
		e.EncodeElement(m.N, sen)
	}
	if m.E != nil {
		see := xml.StartElement{Name: xml.Name{Local: "x:e"}}
		e.EncodeElement(m.E, see)
	}
	if m.S != nil {
		ses := xml.StartElement{Name: xml.Name{Local: "x:s"}}
		e.EncodeElement(m.S, ses)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PCDSDTCEntries) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
lCT_PCDSDTCEntries:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "m":
				tmp := NewCT_Missing()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.M = append(m.M, tmp)
			case "n":
				tmp := NewCT_Number()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.N = append(m.N, tmp)
			case "e":
				tmp := NewCT_Error()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.E = append(m.E, tmp)
			case "s":
				tmp := NewCT_String()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.S = append(m.S, tmp)
			default:
				log.Printf("skipping unsupported element on CT_PCDSDTCEntries %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PCDSDTCEntries
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PCDSDTCEntries and its children
func (m *CT_PCDSDTCEntries) Validate() error {
	return m.ValidateWithPath("CT_PCDSDTCEntries")
}

// ValidateWithPath validates the CT_PCDSDTCEntries and its children, prefixing error messages with path
func (m *CT_PCDSDTCEntries) ValidateWithPath(path string) error {
	for i, v := range m.M {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/M[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.N {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/N[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.E {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/E[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.S {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/S[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}