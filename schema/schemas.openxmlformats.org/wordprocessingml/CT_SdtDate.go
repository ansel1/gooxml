// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
	"log"
	"time"
)

type CT_SdtDate struct {
	// Last Known Date in XML Schema DateTime Format
	FullDateAttr *time.Time
	// Date Display Mask
	DateFormat *CT_String
	// Date Picker Language ID
	Lid *CT_Lang
	// Custom XML Data Date Storage Format
	StoreMappedDataAs *CT_SdtDateMappingType
	// Date Picker Calendar Type
	Calendar *CT_CalendarType
}

func NewCT_SdtDate() *CT_SdtDate {
	ret := &CT_SdtDate{}
	return ret
}

func (m *CT_SdtDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.FullDateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:fullDate"},
			Value: fmt.Sprintf("%v", *m.FullDateAttr)})
	}
	e.EncodeToken(start)
	if m.DateFormat != nil {
		sedateFormat := xml.StartElement{Name: xml.Name{Local: "w:dateFormat"}}
		e.EncodeElement(m.DateFormat, sedateFormat)
	}
	if m.Lid != nil {
		selid := xml.StartElement{Name: xml.Name{Local: "w:lid"}}
		e.EncodeElement(m.Lid, selid)
	}
	if m.StoreMappedDataAs != nil {
		sestoreMappedDataAs := xml.StartElement{Name: xml.Name{Local: "w:storeMappedDataAs"}}
		e.EncodeElement(m.StoreMappedDataAs, sestoreMappedDataAs)
	}
	if m.Calendar != nil {
		secalendar := xml.StartElement{Name: xml.Name{Local: "w:calendar"}}
		e.EncodeElement(m.Calendar, secalendar)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SdtDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "fullDate" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.FullDateAttr = &parsed
		}
	}
lCT_SdtDate:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "dateFormat":
				m.DateFormat = NewCT_String()
				if err := d.DecodeElement(m.DateFormat, &el); err != nil {
					return err
				}
			case "lid":
				m.Lid = NewCT_Lang()
				if err := d.DecodeElement(m.Lid, &el); err != nil {
					return err
				}
			case "storeMappedDataAs":
				m.StoreMappedDataAs = NewCT_SdtDateMappingType()
				if err := d.DecodeElement(m.StoreMappedDataAs, &el); err != nil {
					return err
				}
			case "calendar":
				m.Calendar = NewCT_CalendarType()
				if err := d.DecodeElement(m.Calendar, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_SdtDate %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SdtDate
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SdtDate and its children
func (m *CT_SdtDate) Validate() error {
	return m.ValidateWithPath("CT_SdtDate")
}

// ValidateWithPath validates the CT_SdtDate and its children, prefixing error messages with path
func (m *CT_SdtDate) ValidateWithPath(path string) error {
	if m.DateFormat != nil {
		if err := m.DateFormat.ValidateWithPath(path + "/DateFormat"); err != nil {
			return err
		}
	}
	if m.Lid != nil {
		if err := m.Lid.ValidateWithPath(path + "/Lid"); err != nil {
			return err
		}
	}
	if m.StoreMappedDataAs != nil {
		if err := m.StoreMappedDataAs.ValidateWithPath(path + "/StoreMappedDataAs"); err != nil {
			return err
		}
	}
	if m.Calendar != nil {
		if err := m.Calendar.ValidateWithPath(path + "/Calendar"); err != nil {
			return err
		}
	}
	return nil
}