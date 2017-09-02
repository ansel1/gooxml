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

type CT_Border struct {
	// Diagonal Up
	DiagonalUpAttr *bool
	// Diagonal Down
	DiagonalDownAttr *bool
	// Outline
	OutlineAttr *bool
	// Leading Edge Border
	Start *CT_BorderPr
	// Trailing Edge Border
	End *CT_BorderPr
	// Leading Edge Border
	Left *CT_BorderPr
	// Trailing Edge Border
	Right *CT_BorderPr
	// Top Border
	Top *CT_BorderPr
	// Bottom Border
	Bottom *CT_BorderPr
	// Diagonal
	Diagonal *CT_BorderPr
	// Vertical Inner Border
	Vertical *CT_BorderPr
	// Horizontal Inner Borders
	Horizontal *CT_BorderPr
}

func NewCT_Border() *CT_Border {
	ret := &CT_Border{}
	return ret
}

func (m *CT_Border) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.DiagonalUpAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "diagonalUp"},
			Value: fmt.Sprintf("%v", *m.DiagonalUpAttr)})
	}
	if m.DiagonalDownAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "diagonalDown"},
			Value: fmt.Sprintf("%v", *m.DiagonalDownAttr)})
	}
	if m.OutlineAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "outline"},
			Value: fmt.Sprintf("%v", *m.OutlineAttr)})
	}
	e.EncodeToken(start)
	if m.Start != nil {
		sestart := xml.StartElement{Name: xml.Name{Local: "x:start"}}
		e.EncodeElement(m.Start, sestart)
	}
	if m.End != nil {
		seend := xml.StartElement{Name: xml.Name{Local: "x:end"}}
		e.EncodeElement(m.End, seend)
	}
	if m.Left != nil {
		seleft := xml.StartElement{Name: xml.Name{Local: "x:left"}}
		e.EncodeElement(m.Left, seleft)
	}
	if m.Right != nil {
		seright := xml.StartElement{Name: xml.Name{Local: "x:right"}}
		e.EncodeElement(m.Right, seright)
	}
	if m.Top != nil {
		setop := xml.StartElement{Name: xml.Name{Local: "x:top"}}
		e.EncodeElement(m.Top, setop)
	}
	if m.Bottom != nil {
		sebottom := xml.StartElement{Name: xml.Name{Local: "x:bottom"}}
		e.EncodeElement(m.Bottom, sebottom)
	}
	if m.Diagonal != nil {
		sediagonal := xml.StartElement{Name: xml.Name{Local: "x:diagonal"}}
		e.EncodeElement(m.Diagonal, sediagonal)
	}
	if m.Vertical != nil {
		severtical := xml.StartElement{Name: xml.Name{Local: "x:vertical"}}
		e.EncodeElement(m.Vertical, severtical)
	}
	if m.Horizontal != nil {
		sehorizontal := xml.StartElement{Name: xml.Name{Local: "x:horizontal"}}
		e.EncodeElement(m.Horizontal, sehorizontal)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Border) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "diagonalUp" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DiagonalUpAttr = &parsed
		}
		if attr.Name.Local == "diagonalDown" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DiagonalDownAttr = &parsed
		}
		if attr.Name.Local == "outline" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.OutlineAttr = &parsed
		}
	}
lCT_Border:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "start":
				m.Start = NewCT_BorderPr()
				if err := d.DecodeElement(m.Start, &el); err != nil {
					return err
				}
			case "end":
				m.End = NewCT_BorderPr()
				if err := d.DecodeElement(m.End, &el); err != nil {
					return err
				}
			case "left":
				m.Left = NewCT_BorderPr()
				if err := d.DecodeElement(m.Left, &el); err != nil {
					return err
				}
			case "right":
				m.Right = NewCT_BorderPr()
				if err := d.DecodeElement(m.Right, &el); err != nil {
					return err
				}
			case "top":
				m.Top = NewCT_BorderPr()
				if err := d.DecodeElement(m.Top, &el); err != nil {
					return err
				}
			case "bottom":
				m.Bottom = NewCT_BorderPr()
				if err := d.DecodeElement(m.Bottom, &el); err != nil {
					return err
				}
			case "diagonal":
				m.Diagonal = NewCT_BorderPr()
				if err := d.DecodeElement(m.Diagonal, &el); err != nil {
					return err
				}
			case "vertical":
				m.Vertical = NewCT_BorderPr()
				if err := d.DecodeElement(m.Vertical, &el); err != nil {
					return err
				}
			case "horizontal":
				m.Horizontal = NewCT_BorderPr()
				if err := d.DecodeElement(m.Horizontal, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_Border %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Border
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Border and its children
func (m *CT_Border) Validate() error {
	return m.ValidateWithPath("CT_Border")
}

// ValidateWithPath validates the CT_Border and its children, prefixing error messages with path
func (m *CT_Border) ValidateWithPath(path string) error {
	if m.Start != nil {
		if err := m.Start.ValidateWithPath(path + "/Start"); err != nil {
			return err
		}
	}
	if m.End != nil {
		if err := m.End.ValidateWithPath(path + "/End"); err != nil {
			return err
		}
	}
	if m.Left != nil {
		if err := m.Left.ValidateWithPath(path + "/Left"); err != nil {
			return err
		}
	}
	if m.Right != nil {
		if err := m.Right.ValidateWithPath(path + "/Right"); err != nil {
			return err
		}
	}
	if m.Top != nil {
		if err := m.Top.ValidateWithPath(path + "/Top"); err != nil {
			return err
		}
	}
	if m.Bottom != nil {
		if err := m.Bottom.ValidateWithPath(path + "/Bottom"); err != nil {
			return err
		}
	}
	if m.Diagonal != nil {
		if err := m.Diagonal.ValidateWithPath(path + "/Diagonal"); err != nil {
			return err
		}
	}
	if m.Vertical != nil {
		if err := m.Vertical.ValidateWithPath(path + "/Vertical"); err != nil {
			return err
		}
	}
	if m.Horizontal != nil {
		if err := m.Horizontal.ValidateWithPath(path + "/Horizontal"); err != nil {
			return err
		}
	}
	return nil
}
