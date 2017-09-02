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

type CT_ConditionalFormatting struct {
	// PivotTable Conditional Formatting
	PivotAttr *bool
	// Sequence of Refernces
	SqrefAttr *ST_Sqref
	// Conditional Formatting Rule
	CfRule []*CT_CfRule
	ExtLst *CT_ExtensionList
}

func NewCT_ConditionalFormatting() *CT_ConditionalFormatting {
	ret := &CT_ConditionalFormatting{}
	return ret
}

func (m *CT_ConditionalFormatting) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.PivotAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "pivot"},
			Value: fmt.Sprintf("%v", *m.PivotAttr)})
	}
	if m.SqrefAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sqref"},
			Value: fmt.Sprintf("%v", *m.SqrefAttr)})
	}
	e.EncodeToken(start)
	secfRule := xml.StartElement{Name: xml.Name{Local: "x:cfRule"}}
	e.EncodeElement(m.CfRule, secfRule)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ConditionalFormatting) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "pivot" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PivotAttr = &parsed
		}
		if attr.Name.Local == "sqref" {
			parsed, err := ParseSliceST_Sqref(attr.Value)
			if err != nil {
				return err
			}
			m.SqrefAttr = &parsed
		}
	}
lCT_ConditionalFormatting:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cfRule":
				tmp := NewCT_CfRule()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CfRule = append(m.CfRule, tmp)
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_ConditionalFormatting %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ConditionalFormatting
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ConditionalFormatting and its children
func (m *CT_ConditionalFormatting) Validate() error {
	return m.ValidateWithPath("CT_ConditionalFormatting")
}

// ValidateWithPath validates the CT_ConditionalFormatting and its children, prefixing error messages with path
func (m *CT_ConditionalFormatting) ValidateWithPath(path string) error {
	for i, v := range m.CfRule {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CfRule[%d]", path, i)); err != nil {
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