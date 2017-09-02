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
)

type CT_TLBuildParagraph struct {
	// Build Types
	BuildAttr ST_TLParaBuildType
	// Build Level
	BldLvlAttr *uint32
	// Animate Background
	AnimBgAttr *bool
	// Auto Update Animation Background
	AutoUpdateAnimBgAttr *bool
	// Reverse
	RevAttr *bool
	// Auto Advance Time
	AdvAutoAttr *ST_TLTime
	// Template effects
	TmplLst      *CT_TLTemplateList
	SpidAttr     *uint32
	GrpIdAttr    *uint32
	UiExpandAttr *bool
}

func NewCT_TLBuildParagraph() *CT_TLBuildParagraph {
	ret := &CT_TLBuildParagraph{}
	return ret
}

func (m *CT_TLBuildParagraph) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.BuildAttr != ST_TLParaBuildTypeUnset {
		attr, err := m.BuildAttr.MarshalXMLAttr(xml.Name{Local: "build"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.BldLvlAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "bldLvl"},
			Value: fmt.Sprintf("%v", *m.BldLvlAttr)})
	}
	if m.AnimBgAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "animBg"},
			Value: fmt.Sprintf("%v", *m.AnimBgAttr)})
	}
	if m.AutoUpdateAnimBgAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoUpdateAnimBg"},
			Value: fmt.Sprintf("%v", *m.AutoUpdateAnimBgAttr)})
	}
	if m.RevAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rev"},
			Value: fmt.Sprintf("%v", *m.RevAttr)})
	}
	if m.AdvAutoAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "advAuto"},
			Value: fmt.Sprintf("%v", *m.AdvAutoAttr)})
	}
	if m.SpidAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "spid"},
			Value: fmt.Sprintf("%v", *m.SpidAttr)})
	}
	if m.GrpIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "grpId"},
			Value: fmt.Sprintf("%v", *m.GrpIdAttr)})
	}
	if m.UiExpandAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uiExpand"},
			Value: fmt.Sprintf("%v", *m.UiExpandAttr)})
	}
	e.EncodeToken(start)
	if m.TmplLst != nil {
		setmplLst := xml.StartElement{Name: xml.Name{Local: "p:tmplLst"}}
		e.EncodeElement(m.TmplLst, setmplLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLBuildParagraph) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "build" {
			m.BuildAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "bldLvl" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.BldLvlAttr = &pt
		}
		if attr.Name.Local == "animBg" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AnimBgAttr = &parsed
		}
		if attr.Name.Local == "autoUpdateAnimBg" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AutoUpdateAnimBgAttr = &parsed
		}
		if attr.Name.Local == "rev" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RevAttr = &parsed
		}
		if attr.Name.Local == "advAuto" {
			parsed, err := ParseUnionST_TLTime(attr.Value)
			if err != nil {
				return err
			}
			m.AdvAutoAttr = &parsed
		}
		if attr.Name.Local == "spid" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.SpidAttr = &pt
		}
		if attr.Name.Local == "grpId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.GrpIdAttr = &pt
		}
		if attr.Name.Local == "uiExpand" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UiExpandAttr = &parsed
		}
	}
lCT_TLBuildParagraph:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tmplLst":
				m.TmplLst = NewCT_TLTemplateList()
				if err := d.DecodeElement(m.TmplLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_TLBuildParagraph %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLBuildParagraph
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLBuildParagraph and its children
func (m *CT_TLBuildParagraph) Validate() error {
	return m.ValidateWithPath("CT_TLBuildParagraph")
}

// ValidateWithPath validates the CT_TLBuildParagraph and its children, prefixing error messages with path
func (m *CT_TLBuildParagraph) ValidateWithPath(path string) error {
	if err := m.BuildAttr.ValidateWithPath(path + "/BuildAttr"); err != nil {
		return err
	}
	if m.AdvAutoAttr != nil {
		if err := m.AdvAutoAttr.ValidateWithPath(path + "/AdvAutoAttr"); err != nil {
			return err
		}
	}
	if m.TmplLst != nil {
		if err := m.TmplLst.ValidateWithPath(path + "/TmplLst"); err != nil {
			return err
		}
	}
	return nil
}