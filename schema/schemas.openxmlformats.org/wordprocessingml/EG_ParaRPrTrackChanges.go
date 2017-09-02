// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"log"
)

type EG_ParaRPrTrackChanges struct {
	// Inserted Paragraph
	Ins *CT_TrackChange
	// Deleted Paragraph
	Del *CT_TrackChange
	// Move Source Paragraph
	MoveFrom *CT_TrackChange
	// Move Destination Paragraph
	MoveTo *CT_TrackChange
}

func NewEG_ParaRPrTrackChanges() *EG_ParaRPrTrackChanges {
	ret := &EG_ParaRPrTrackChanges{}
	return ret
}

func (m *EG_ParaRPrTrackChanges) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Ins != nil {
		seins := xml.StartElement{Name: xml.Name{Local: "w:ins"}}
		e.EncodeElement(m.Ins, seins)
	}
	if m.Del != nil {
		sedel := xml.StartElement{Name: xml.Name{Local: "w:del"}}
		e.EncodeElement(m.Del, sedel)
	}
	if m.MoveFrom != nil {
		semoveFrom := xml.StartElement{Name: xml.Name{Local: "w:moveFrom"}}
		e.EncodeElement(m.MoveFrom, semoveFrom)
	}
	if m.MoveTo != nil {
		semoveTo := xml.StartElement{Name: xml.Name{Local: "w:moveTo"}}
		e.EncodeElement(m.MoveTo, semoveTo)
	}
	return nil
}

func (m *EG_ParaRPrTrackChanges) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_ParaRPrTrackChanges:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "ins":
				m.Ins = NewCT_TrackChange()
				if err := d.DecodeElement(m.Ins, &el); err != nil {
					return err
				}
			case "del":
				m.Del = NewCT_TrackChange()
				if err := d.DecodeElement(m.Del, &el); err != nil {
					return err
				}
			case "moveFrom":
				m.MoveFrom = NewCT_TrackChange()
				if err := d.DecodeElement(m.MoveFrom, &el); err != nil {
					return err
				}
			case "moveTo":
				m.MoveTo = NewCT_TrackChange()
				if err := d.DecodeElement(m.MoveTo, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on EG_ParaRPrTrackChanges %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_ParaRPrTrackChanges
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_ParaRPrTrackChanges and its children
func (m *EG_ParaRPrTrackChanges) Validate() error {
	return m.ValidateWithPath("EG_ParaRPrTrackChanges")
}

// ValidateWithPath validates the EG_ParaRPrTrackChanges and its children, prefixing error messages with path
func (m *EG_ParaRPrTrackChanges) ValidateWithPath(path string) error {
	if m.Ins != nil {
		if err := m.Ins.ValidateWithPath(path + "/Ins"); err != nil {
			return err
		}
	}
	if m.Del != nil {
		if err := m.Del.ValidateWithPath(path + "/Del"); err != nil {
			return err
		}
	}
	if m.MoveFrom != nil {
		if err := m.MoveFrom.ValidateWithPath(path + "/MoveFrom"); err != nil {
			return err
		}
	}
	if m.MoveTo != nil {
		if err := m.MoveTo.ValidateWithPath(path + "/MoveTo"); err != nil {
			return err
		}
	}
	return nil
}
