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

type CT_DashStopList struct {
	Ds []*CT_DashStop
}

func NewCT_DashStopList() *CT_DashStopList {
	ret := &CT_DashStopList{}
	return ret
}

func (m *CT_DashStopList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Ds != nil {
		seds := xml.StartElement{Name: xml.Name{Local: "a:ds"}}
		e.EncodeElement(m.Ds, seds)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DashStopList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_DashStopList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "ds":
				tmp := NewCT_DashStop()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Ds = append(m.Ds, tmp)
			default:
				log.Printf("skipping unsupported element on CT_DashStopList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DashStopList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DashStopList and its children
func (m *CT_DashStopList) Validate() error {
	return m.ValidateWithPath("CT_DashStopList")
}

// ValidateWithPath validates the CT_DashStopList and its children, prefixing error messages with path
func (m *CT_DashStopList) ValidateWithPath(path string) error {
	for i, v := range m.Ds {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Ds[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}