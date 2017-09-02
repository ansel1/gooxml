// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"log"
)

type EG_TextBulletSize struct {
	BuSzTx  *CT_TextBulletSizeFollowText
	BuSzPct *CT_TextBulletSizePercent
	BuSzPts *CT_TextBulletSizePoint
}

func NewEG_TextBulletSize() *EG_TextBulletSize {
	ret := &EG_TextBulletSize{}
	return ret
}

func (m *EG_TextBulletSize) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.BuSzTx != nil {
		sebuSzTx := xml.StartElement{Name: xml.Name{Local: "a:buSzTx"}}
		e.EncodeElement(m.BuSzTx, sebuSzTx)
	}
	if m.BuSzPct != nil {
		sebuSzPct := xml.StartElement{Name: xml.Name{Local: "a:buSzPct"}}
		e.EncodeElement(m.BuSzPct, sebuSzPct)
	}
	if m.BuSzPts != nil {
		sebuSzPts := xml.StartElement{Name: xml.Name{Local: "a:buSzPts"}}
		e.EncodeElement(m.BuSzPts, sebuSzPts)
	}
	return nil
}

func (m *EG_TextBulletSize) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_TextBulletSize:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "buSzTx":
				m.BuSzTx = NewCT_TextBulletSizeFollowText()
				if err := d.DecodeElement(m.BuSzTx, &el); err != nil {
					return err
				}
			case "buSzPct":
				m.BuSzPct = NewCT_TextBulletSizePercent()
				if err := d.DecodeElement(m.BuSzPct, &el); err != nil {
					return err
				}
			case "buSzPts":
				m.BuSzPts = NewCT_TextBulletSizePoint()
				if err := d.DecodeElement(m.BuSzPts, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on EG_TextBulletSize %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_TextBulletSize
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_TextBulletSize and its children
func (m *EG_TextBulletSize) Validate() error {
	return m.ValidateWithPath("EG_TextBulletSize")
}

// ValidateWithPath validates the EG_TextBulletSize and its children, prefixing error messages with path
func (m *EG_TextBulletSize) ValidateWithPath(path string) error {
	if m.BuSzTx != nil {
		if err := m.BuSzTx.ValidateWithPath(path + "/BuSzTx"); err != nil {
			return err
		}
	}
	if m.BuSzPct != nil {
		if err := m.BuSzPct.ValidateWithPath(path + "/BuSzPct"); err != nil {
			return err
		}
	}
	if m.BuSzPts != nil {
		if err := m.BuSzPts.ValidateWithPath(path + "/BuSzPts"); err != nil {
			return err
		}
	}
	return nil
}
