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
	"strconv"
)

type CT_AbstractNum struct {
	// Abstract Numbering Definition ID
	AbstractNumIdAttr int64
	// Abstract Numbering Definition Identifier
	Nsid *CT_LongHexNumber
	// Abstract Numbering Definition Type
	MultiLevelType *CT_MultiLevelType
	// Numbering Template Code
	Tmpl *CT_LongHexNumber
	// Abstract Numbering Definition Name
	Name *CT_String
	// Numbering Style Definition
	StyleLink *CT_String
	// Numbering Style Reference
	NumStyleLink *CT_String
	// Numbering Level Definition
	Lvl []*CT_Lvl
}

func NewCT_AbstractNum() *CT_AbstractNum {
	ret := &CT_AbstractNum{}
	return ret
}

func (m *CT_AbstractNum) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:abstractNumId"},
		Value: fmt.Sprintf("%v", m.AbstractNumIdAttr)})
	e.EncodeToken(start)
	if m.Nsid != nil {
		sensid := xml.StartElement{Name: xml.Name{Local: "w:nsid"}}
		e.EncodeElement(m.Nsid, sensid)
	}
	if m.MultiLevelType != nil {
		semultiLevelType := xml.StartElement{Name: xml.Name{Local: "w:multiLevelType"}}
		e.EncodeElement(m.MultiLevelType, semultiLevelType)
	}
	if m.Tmpl != nil {
		setmpl := xml.StartElement{Name: xml.Name{Local: "w:tmpl"}}
		e.EncodeElement(m.Tmpl, setmpl)
	}
	if m.Name != nil {
		sename := xml.StartElement{Name: xml.Name{Local: "w:name"}}
		e.EncodeElement(m.Name, sename)
	}
	if m.StyleLink != nil {
		sestyleLink := xml.StartElement{Name: xml.Name{Local: "w:styleLink"}}
		e.EncodeElement(m.StyleLink, sestyleLink)
	}
	if m.NumStyleLink != nil {
		senumStyleLink := xml.StartElement{Name: xml.Name{Local: "w:numStyleLink"}}
		e.EncodeElement(m.NumStyleLink, senumStyleLink)
	}
	if m.Lvl != nil {
		selvl := xml.StartElement{Name: xml.Name{Local: "w:lvl"}}
		e.EncodeElement(m.Lvl, selvl)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_AbstractNum) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "abstractNumId" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.AbstractNumIdAttr = parsed
		}
	}
lCT_AbstractNum:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "nsid":
				m.Nsid = NewCT_LongHexNumber()
				if err := d.DecodeElement(m.Nsid, &el); err != nil {
					return err
				}
			case "multiLevelType":
				m.MultiLevelType = NewCT_MultiLevelType()
				if err := d.DecodeElement(m.MultiLevelType, &el); err != nil {
					return err
				}
			case "tmpl":
				m.Tmpl = NewCT_LongHexNumber()
				if err := d.DecodeElement(m.Tmpl, &el); err != nil {
					return err
				}
			case "name":
				m.Name = NewCT_String()
				if err := d.DecodeElement(m.Name, &el); err != nil {
					return err
				}
			case "styleLink":
				m.StyleLink = NewCT_String()
				if err := d.DecodeElement(m.StyleLink, &el); err != nil {
					return err
				}
			case "numStyleLink":
				m.NumStyleLink = NewCT_String()
				if err := d.DecodeElement(m.NumStyleLink, &el); err != nil {
					return err
				}
			case "lvl":
				tmp := NewCT_Lvl()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Lvl = append(m.Lvl, tmp)
			default:
				log.Printf("skipping unsupported element on CT_AbstractNum %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_AbstractNum
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_AbstractNum and its children
func (m *CT_AbstractNum) Validate() error {
	return m.ValidateWithPath("CT_AbstractNum")
}

// ValidateWithPath validates the CT_AbstractNum and its children, prefixing error messages with path
func (m *CT_AbstractNum) ValidateWithPath(path string) error {
	if m.Nsid != nil {
		if err := m.Nsid.ValidateWithPath(path + "/Nsid"); err != nil {
			return err
		}
	}
	if m.MultiLevelType != nil {
		if err := m.MultiLevelType.ValidateWithPath(path + "/MultiLevelType"); err != nil {
			return err
		}
	}
	if m.Tmpl != nil {
		if err := m.Tmpl.ValidateWithPath(path + "/Tmpl"); err != nil {
			return err
		}
	}
	if m.Name != nil {
		if err := m.Name.ValidateWithPath(path + "/Name"); err != nil {
			return err
		}
	}
	if m.StyleLink != nil {
		if err := m.StyleLink.ValidateWithPath(path + "/StyleLink"); err != nil {
			return err
		}
	}
	if m.NumStyleLink != nil {
		if err := m.NumStyleLink.ValidateWithPath(path + "/NumStyleLink"); err != nil {
			return err
		}
	}
	for i, v := range m.Lvl {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Lvl[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
