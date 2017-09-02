// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"log"
)

type EG_ShowType struct {
	// Presenter Slide Show Mode
	Present *CT_Empty
	// Browse Slide Show Mode
	Browse *CT_ShowInfoBrowse
	// Kiosk Slide Show Mode
	Kiosk *CT_ShowInfoKiosk
}

func NewEG_ShowType() *EG_ShowType {
	ret := &EG_ShowType{}
	return ret
}

func (m *EG_ShowType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Present != nil {
		sepresent := xml.StartElement{Name: xml.Name{Local: "p:present"}}
		e.EncodeElement(m.Present, sepresent)
	}
	if m.Browse != nil {
		sebrowse := xml.StartElement{Name: xml.Name{Local: "p:browse"}}
		e.EncodeElement(m.Browse, sebrowse)
	}
	if m.Kiosk != nil {
		sekiosk := xml.StartElement{Name: xml.Name{Local: "p:kiosk"}}
		e.EncodeElement(m.Kiosk, sekiosk)
	}
	return nil
}

func (m *EG_ShowType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_ShowType:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "present":
				m.Present = NewCT_Empty()
				if err := d.DecodeElement(m.Present, &el); err != nil {
					return err
				}
			case "browse":
				m.Browse = NewCT_ShowInfoBrowse()
				if err := d.DecodeElement(m.Browse, &el); err != nil {
					return err
				}
			case "kiosk":
				m.Kiosk = NewCT_ShowInfoKiosk()
				if err := d.DecodeElement(m.Kiosk, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on EG_ShowType %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_ShowType
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_ShowType and its children
func (m *EG_ShowType) Validate() error {
	return m.ValidateWithPath("EG_ShowType")
}

// ValidateWithPath validates the EG_ShowType and its children, prefixing error messages with path
func (m *EG_ShowType) ValidateWithPath(path string) error {
	if m.Present != nil {
		if err := m.Present.ValidateWithPath(path + "/Present"); err != nil {
			return err
		}
	}
	if m.Browse != nil {
		if err := m.Browse.ValidateWithPath(path + "/Browse"); err != nil {
			return err
		}
	}
	if m.Kiosk != nil {
		if err := m.Kiosk.ValidateWithPath(path + "/Kiosk"); err != nil {
			return err
		}
	}
	return nil
}
