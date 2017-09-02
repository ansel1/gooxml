// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_ViewProperties struct {
	// Last View
	LastViewAttr ST_ViewType
	// Show Comments
	ShowCommentsAttr *bool
	// Normal View Properties
	NormalViewPr *CT_NormalViewProperties
	// Slide View Properties
	SlideViewPr *CT_SlideViewProperties
	// Outline View Properties
	OutlineViewPr *CT_OutlineViewProperties
	// Notes Text View Properties
	NotesTextViewPr *CT_NotesTextViewProperties
	// Slide Sorter View Properties
	SorterViewPr *CT_SlideSorterViewProperties
	// Notes View Properties
	NotesViewPr *CT_NotesViewProperties
	// Grid Spacing
	GridSpacing *drawingml.CT_PositiveSize2D
	ExtLst      *CT_ExtensionList
}

func NewCT_ViewProperties() *CT_ViewProperties {
	ret := &CT_ViewProperties{}
	return ret
}

func (m *CT_ViewProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.LastViewAttr != ST_ViewTypeUnset {
		attr, err := m.LastViewAttr.MarshalXMLAttr(xml.Name{Local: "lastView"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ShowCommentsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showComments"},
			Value: fmt.Sprintf("%v", *m.ShowCommentsAttr)})
	}
	e.EncodeToken(start)
	if m.NormalViewPr != nil {
		senormalViewPr := xml.StartElement{Name: xml.Name{Local: "p:normalViewPr"}}
		e.EncodeElement(m.NormalViewPr, senormalViewPr)
	}
	if m.SlideViewPr != nil {
		seslideViewPr := xml.StartElement{Name: xml.Name{Local: "p:slideViewPr"}}
		e.EncodeElement(m.SlideViewPr, seslideViewPr)
	}
	if m.OutlineViewPr != nil {
		seoutlineViewPr := xml.StartElement{Name: xml.Name{Local: "p:outlineViewPr"}}
		e.EncodeElement(m.OutlineViewPr, seoutlineViewPr)
	}
	if m.NotesTextViewPr != nil {
		senotesTextViewPr := xml.StartElement{Name: xml.Name{Local: "p:notesTextViewPr"}}
		e.EncodeElement(m.NotesTextViewPr, senotesTextViewPr)
	}
	if m.SorterViewPr != nil {
		sesorterViewPr := xml.StartElement{Name: xml.Name{Local: "p:sorterViewPr"}}
		e.EncodeElement(m.SorterViewPr, sesorterViewPr)
	}
	if m.NotesViewPr != nil {
		senotesViewPr := xml.StartElement{Name: xml.Name{Local: "p:notesViewPr"}}
		e.EncodeElement(m.NotesViewPr, senotesViewPr)
	}
	if m.GridSpacing != nil {
		segridSpacing := xml.StartElement{Name: xml.Name{Local: "p:gridSpacing"}}
		e.EncodeElement(m.GridSpacing, segridSpacing)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ViewProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "lastView" {
			m.LastViewAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "showComments" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowCommentsAttr = &parsed
		}
	}
lCT_ViewProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "normalViewPr":
				m.NormalViewPr = NewCT_NormalViewProperties()
				if err := d.DecodeElement(m.NormalViewPr, &el); err != nil {
					return err
				}
			case "slideViewPr":
				m.SlideViewPr = NewCT_SlideViewProperties()
				if err := d.DecodeElement(m.SlideViewPr, &el); err != nil {
					return err
				}
			case "outlineViewPr":
				m.OutlineViewPr = NewCT_OutlineViewProperties()
				if err := d.DecodeElement(m.OutlineViewPr, &el); err != nil {
					return err
				}
			case "notesTextViewPr":
				m.NotesTextViewPr = NewCT_NotesTextViewProperties()
				if err := d.DecodeElement(m.NotesTextViewPr, &el); err != nil {
					return err
				}
			case "sorterViewPr":
				m.SorterViewPr = NewCT_SlideSorterViewProperties()
				if err := d.DecodeElement(m.SorterViewPr, &el); err != nil {
					return err
				}
			case "notesViewPr":
				m.NotesViewPr = NewCT_NotesViewProperties()
				if err := d.DecodeElement(m.NotesViewPr, &el); err != nil {
					return err
				}
			case "gridSpacing":
				m.GridSpacing = drawingml.NewCT_PositiveSize2D()
				if err := d.DecodeElement(m.GridSpacing, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_ViewProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ViewProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ViewProperties and its children
func (m *CT_ViewProperties) Validate() error {
	return m.ValidateWithPath("CT_ViewProperties")
}

// ValidateWithPath validates the CT_ViewProperties and its children, prefixing error messages with path
func (m *CT_ViewProperties) ValidateWithPath(path string) error {
	if err := m.LastViewAttr.ValidateWithPath(path + "/LastViewAttr"); err != nil {
		return err
	}
	if m.NormalViewPr != nil {
		if err := m.NormalViewPr.ValidateWithPath(path + "/NormalViewPr"); err != nil {
			return err
		}
	}
	if m.SlideViewPr != nil {
		if err := m.SlideViewPr.ValidateWithPath(path + "/SlideViewPr"); err != nil {
			return err
		}
	}
	if m.OutlineViewPr != nil {
		if err := m.OutlineViewPr.ValidateWithPath(path + "/OutlineViewPr"); err != nil {
			return err
		}
	}
	if m.NotesTextViewPr != nil {
		if err := m.NotesTextViewPr.ValidateWithPath(path + "/NotesTextViewPr"); err != nil {
			return err
		}
	}
	if m.SorterViewPr != nil {
		if err := m.SorterViewPr.ValidateWithPath(path + "/SorterViewPr"); err != nil {
			return err
		}
	}
	if m.NotesViewPr != nil {
		if err := m.NotesViewPr.ValidateWithPath(path + "/NotesViewPr"); err != nil {
			return err
		}
	}
	if m.GridSpacing != nil {
		if err := m.GridSpacing.ValidateWithPath(path + "/GridSpacing"); err != nil {
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
