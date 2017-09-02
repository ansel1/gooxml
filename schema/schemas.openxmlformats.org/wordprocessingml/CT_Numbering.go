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
)

type CT_Numbering struct {
	// Picture Numbering Symbol Definition
	NumPicBullet []*CT_NumPicBullet
	// Abstract Numbering Definition
	AbstractNum []*CT_AbstractNum
	// Numbering Definition Instance
	Num []*CT_Num
	// Last Reviewed Abstract Numbering Definition
	NumIdMacAtCleanup *CT_DecimalNumber
}

func NewCT_Numbering() *CT_Numbering {
	ret := &CT_Numbering{}
	return ret
}

func (m *CT_Numbering) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.NumPicBullet != nil {
		senumPicBullet := xml.StartElement{Name: xml.Name{Local: "w:numPicBullet"}}
		e.EncodeElement(m.NumPicBullet, senumPicBullet)
	}
	if m.AbstractNum != nil {
		seabstractNum := xml.StartElement{Name: xml.Name{Local: "w:abstractNum"}}
		e.EncodeElement(m.AbstractNum, seabstractNum)
	}
	if m.Num != nil {
		senum := xml.StartElement{Name: xml.Name{Local: "w:num"}}
		e.EncodeElement(m.Num, senum)
	}
	if m.NumIdMacAtCleanup != nil {
		senumIdMacAtCleanup := xml.StartElement{Name: xml.Name{Local: "w:numIdMacAtCleanup"}}
		e.EncodeElement(m.NumIdMacAtCleanup, senumIdMacAtCleanup)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Numbering) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Numbering:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "numPicBullet":
				tmp := NewCT_NumPicBullet()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.NumPicBullet = append(m.NumPicBullet, tmp)
			case "abstractNum":
				tmp := NewCT_AbstractNum()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AbstractNum = append(m.AbstractNum, tmp)
			case "num":
				tmp := NewCT_Num()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Num = append(m.Num, tmp)
			case "numIdMacAtCleanup":
				m.NumIdMacAtCleanup = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.NumIdMacAtCleanup, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_Numbering %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Numbering
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Numbering and its children
func (m *CT_Numbering) Validate() error {
	return m.ValidateWithPath("CT_Numbering")
}

// ValidateWithPath validates the CT_Numbering and its children, prefixing error messages with path
func (m *CT_Numbering) ValidateWithPath(path string) error {
	for i, v := range m.NumPicBullet {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/NumPicBullet[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.AbstractNum {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AbstractNum[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Num {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Num[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.NumIdMacAtCleanup != nil {
		if err := m.NumIdMacAtCleanup.ValidateWithPath(path + "/NumIdMacAtCleanup"); err != nil {
			return err
		}
	}
	return nil
}