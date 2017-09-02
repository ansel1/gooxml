// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"log"
)

type CT_Scaling struct {
	LogBase     *CT_LogBase
	Orientation *CT_Orientation
	Max         *CT_Double
	Min         *CT_Double
	ExtLst      *CT_ExtensionList
}

func NewCT_Scaling() *CT_Scaling {
	ret := &CT_Scaling{}
	return ret
}

func (m *CT_Scaling) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.LogBase != nil {
		selogBase := xml.StartElement{Name: xml.Name{Local: "logBase"}}
		e.EncodeElement(m.LogBase, selogBase)
	}
	if m.Orientation != nil {
		seorientation := xml.StartElement{Name: xml.Name{Local: "orientation"}}
		e.EncodeElement(m.Orientation, seorientation)
	}
	if m.Max != nil {
		semax := xml.StartElement{Name: xml.Name{Local: "max"}}
		e.EncodeElement(m.Max, semax)
	}
	if m.Min != nil {
		semin := xml.StartElement{Name: xml.Name{Local: "min"}}
		e.EncodeElement(m.Min, semin)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Scaling) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Scaling:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "logBase":
				m.LogBase = NewCT_LogBase()
				if err := d.DecodeElement(m.LogBase, &el); err != nil {
					return err
				}
			case "orientation":
				m.Orientation = NewCT_Orientation()
				if err := d.DecodeElement(m.Orientation, &el); err != nil {
					return err
				}
			case "max":
				m.Max = NewCT_Double()
				if err := d.DecodeElement(m.Max, &el); err != nil {
					return err
				}
			case "min":
				m.Min = NewCT_Double()
				if err := d.DecodeElement(m.Min, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_Scaling %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Scaling
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Scaling and its children
func (m *CT_Scaling) Validate() error {
	return m.ValidateWithPath("CT_Scaling")
}

// ValidateWithPath validates the CT_Scaling and its children, prefixing error messages with path
func (m *CT_Scaling) ValidateWithPath(path string) error {
	if m.LogBase != nil {
		if err := m.LogBase.ValidateWithPath(path + "/LogBase"); err != nil {
			return err
		}
	}
	if m.Orientation != nil {
		if err := m.Orientation.ValidateWithPath(path + "/Orientation"); err != nil {
			return err
		}
	}
	if m.Max != nil {
		if err := m.Max.ValidateWithPath(path + "/Max"); err != nil {
			return err
		}
	}
	if m.Min != nil {
		if err := m.Min.ValidateWithPath(path + "/Min"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}