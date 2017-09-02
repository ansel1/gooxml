// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_TableCellProperties struct {
	MarLAttr         *ST_Coordinate32
	MarRAttr         *ST_Coordinate32
	MarTAttr         *ST_Coordinate32
	MarBAttr         *ST_Coordinate32
	VertAttr         ST_TextVerticalType
	AnchorAttr       ST_TextAnchoringType
	AnchorCtrAttr    *bool
	HorzOverflowAttr ST_TextHorzOverflowType
	LnL              *CT_LineProperties
	LnR              *CT_LineProperties
	LnT              *CT_LineProperties
	LnB              *CT_LineProperties
	LnTlToBr         *CT_LineProperties
	LnBlToTr         *CT_LineProperties
	Cell3D           *CT_Cell3D
	NoFill           *CT_NoFillProperties
	SolidFill        *CT_SolidColorFillProperties
	GradFill         *CT_GradientFillProperties
	BlipFill         *CT_BlipFillProperties
	PattFill         *CT_PatternFillProperties
	GrpFill          *CT_GroupFillProperties
	Headers          *CT_Headers
	ExtLst           *CT_OfficeArtExtensionList
}

func NewCT_TableCellProperties() *CT_TableCellProperties {
	ret := &CT_TableCellProperties{}
	return ret
}

func (m *CT_TableCellProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.MarLAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "marL"},
			Value: fmt.Sprintf("%v", *m.MarLAttr)})
	}
	if m.MarRAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "marR"},
			Value: fmt.Sprintf("%v", *m.MarRAttr)})
	}
	if m.MarTAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "marT"},
			Value: fmt.Sprintf("%v", *m.MarTAttr)})
	}
	if m.MarBAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "marB"},
			Value: fmt.Sprintf("%v", *m.MarBAttr)})
	}
	if m.VertAttr != ST_TextVerticalTypeUnset {
		attr, err := m.VertAttr.MarshalXMLAttr(xml.Name{Local: "vert"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.AnchorAttr != ST_TextAnchoringTypeUnset {
		attr, err := m.AnchorAttr.MarshalXMLAttr(xml.Name{Local: "anchor"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.AnchorCtrAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "anchorCtr"},
			Value: fmt.Sprintf("%v", *m.AnchorCtrAttr)})
	}
	if m.HorzOverflowAttr != ST_TextHorzOverflowTypeUnset {
		attr, err := m.HorzOverflowAttr.MarshalXMLAttr(xml.Name{Local: "horzOverflow"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.LnL != nil {
		selnL := xml.StartElement{Name: xml.Name{Local: "a:lnL"}}
		e.EncodeElement(m.LnL, selnL)
	}
	if m.LnR != nil {
		selnR := xml.StartElement{Name: xml.Name{Local: "a:lnR"}}
		e.EncodeElement(m.LnR, selnR)
	}
	if m.LnT != nil {
		selnT := xml.StartElement{Name: xml.Name{Local: "a:lnT"}}
		e.EncodeElement(m.LnT, selnT)
	}
	if m.LnB != nil {
		selnB := xml.StartElement{Name: xml.Name{Local: "a:lnB"}}
		e.EncodeElement(m.LnB, selnB)
	}
	if m.LnTlToBr != nil {
		selnTlToBr := xml.StartElement{Name: xml.Name{Local: "a:lnTlToBr"}}
		e.EncodeElement(m.LnTlToBr, selnTlToBr)
	}
	if m.LnBlToTr != nil {
		selnBlToTr := xml.StartElement{Name: xml.Name{Local: "a:lnBlToTr"}}
		e.EncodeElement(m.LnBlToTr, selnBlToTr)
	}
	if m.Cell3D != nil {
		secell3D := xml.StartElement{Name: xml.Name{Local: "a:cell3D"}}
		e.EncodeElement(m.Cell3D, secell3D)
	}
	if m.NoFill != nil {
		senoFill := xml.StartElement{Name: xml.Name{Local: "a:noFill"}}
		e.EncodeElement(m.NoFill, senoFill)
	}
	if m.SolidFill != nil {
		sesolidFill := xml.StartElement{Name: xml.Name{Local: "a:solidFill"}}
		e.EncodeElement(m.SolidFill, sesolidFill)
	}
	if m.GradFill != nil {
		segradFill := xml.StartElement{Name: xml.Name{Local: "a:gradFill"}}
		e.EncodeElement(m.GradFill, segradFill)
	}
	if m.BlipFill != nil {
		seblipFill := xml.StartElement{Name: xml.Name{Local: "a:blipFill"}}
		e.EncodeElement(m.BlipFill, seblipFill)
	}
	if m.PattFill != nil {
		sepattFill := xml.StartElement{Name: xml.Name{Local: "a:pattFill"}}
		e.EncodeElement(m.PattFill, sepattFill)
	}
	if m.GrpFill != nil {
		segrpFill := xml.StartElement{Name: xml.Name{Local: "a:grpFill"}}
		e.EncodeElement(m.GrpFill, segrpFill)
	}
	if m.Headers != nil {
		seheaders := xml.StartElement{Name: xml.Name{Local: "a:headers"}}
		e.EncodeElement(m.Headers, seheaders)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TableCellProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "marL" {
			parsed, err := ParseUnionST_Coordinate32(attr.Value)
			if err != nil {
				return err
			}
			m.MarLAttr = &parsed
		}
		if attr.Name.Local == "marR" {
			parsed, err := ParseUnionST_Coordinate32(attr.Value)
			if err != nil {
				return err
			}
			m.MarRAttr = &parsed
		}
		if attr.Name.Local == "marT" {
			parsed, err := ParseUnionST_Coordinate32(attr.Value)
			if err != nil {
				return err
			}
			m.MarTAttr = &parsed
		}
		if attr.Name.Local == "marB" {
			parsed, err := ParseUnionST_Coordinate32(attr.Value)
			if err != nil {
				return err
			}
			m.MarBAttr = &parsed
		}
		if attr.Name.Local == "vert" {
			m.VertAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "anchor" {
			m.AnchorAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "anchorCtr" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AnchorCtrAttr = &parsed
		}
		if attr.Name.Local == "horzOverflow" {
			m.HorzOverflowAttr.UnmarshalXMLAttr(attr)
		}
	}
lCT_TableCellProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "lnL":
				m.LnL = NewCT_LineProperties()
				if err := d.DecodeElement(m.LnL, &el); err != nil {
					return err
				}
			case "lnR":
				m.LnR = NewCT_LineProperties()
				if err := d.DecodeElement(m.LnR, &el); err != nil {
					return err
				}
			case "lnT":
				m.LnT = NewCT_LineProperties()
				if err := d.DecodeElement(m.LnT, &el); err != nil {
					return err
				}
			case "lnB":
				m.LnB = NewCT_LineProperties()
				if err := d.DecodeElement(m.LnB, &el); err != nil {
					return err
				}
			case "lnTlToBr":
				m.LnTlToBr = NewCT_LineProperties()
				if err := d.DecodeElement(m.LnTlToBr, &el); err != nil {
					return err
				}
			case "lnBlToTr":
				m.LnBlToTr = NewCT_LineProperties()
				if err := d.DecodeElement(m.LnBlToTr, &el); err != nil {
					return err
				}
			case "cell3D":
				m.Cell3D = NewCT_Cell3D()
				if err := d.DecodeElement(m.Cell3D, &el); err != nil {
					return err
				}
			case "noFill":
				m.NoFill = NewCT_NoFillProperties()
				if err := d.DecodeElement(m.NoFill, &el); err != nil {
					return err
				}
			case "solidFill":
				m.SolidFill = NewCT_SolidColorFillProperties()
				if err := d.DecodeElement(m.SolidFill, &el); err != nil {
					return err
				}
			case "gradFill":
				m.GradFill = NewCT_GradientFillProperties()
				if err := d.DecodeElement(m.GradFill, &el); err != nil {
					return err
				}
			case "blipFill":
				m.BlipFill = NewCT_BlipFillProperties()
				if err := d.DecodeElement(m.BlipFill, &el); err != nil {
					return err
				}
			case "pattFill":
				m.PattFill = NewCT_PatternFillProperties()
				if err := d.DecodeElement(m.PattFill, &el); err != nil {
					return err
				}
			case "grpFill":
				m.GrpFill = NewCT_GroupFillProperties()
				if err := d.DecodeElement(m.GrpFill, &el); err != nil {
					return err
				}
			case "headers":
				m.Headers = NewCT_Headers()
				if err := d.DecodeElement(m.Headers, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_TableCellProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TableCellProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TableCellProperties and its children
func (m *CT_TableCellProperties) Validate() error {
	return m.ValidateWithPath("CT_TableCellProperties")
}

// ValidateWithPath validates the CT_TableCellProperties and its children, prefixing error messages with path
func (m *CT_TableCellProperties) ValidateWithPath(path string) error {
	if m.MarLAttr != nil {
		if err := m.MarLAttr.ValidateWithPath(path + "/MarLAttr"); err != nil {
			return err
		}
	}
	if m.MarRAttr != nil {
		if err := m.MarRAttr.ValidateWithPath(path + "/MarRAttr"); err != nil {
			return err
		}
	}
	if m.MarTAttr != nil {
		if err := m.MarTAttr.ValidateWithPath(path + "/MarTAttr"); err != nil {
			return err
		}
	}
	if m.MarBAttr != nil {
		if err := m.MarBAttr.ValidateWithPath(path + "/MarBAttr"); err != nil {
			return err
		}
	}
	if err := m.VertAttr.ValidateWithPath(path + "/VertAttr"); err != nil {
		return err
	}
	if err := m.AnchorAttr.ValidateWithPath(path + "/AnchorAttr"); err != nil {
		return err
	}
	if err := m.HorzOverflowAttr.ValidateWithPath(path + "/HorzOverflowAttr"); err != nil {
		return err
	}
	if m.LnL != nil {
		if err := m.LnL.ValidateWithPath(path + "/LnL"); err != nil {
			return err
		}
	}
	if m.LnR != nil {
		if err := m.LnR.ValidateWithPath(path + "/LnR"); err != nil {
			return err
		}
	}
	if m.LnT != nil {
		if err := m.LnT.ValidateWithPath(path + "/LnT"); err != nil {
			return err
		}
	}
	if m.LnB != nil {
		if err := m.LnB.ValidateWithPath(path + "/LnB"); err != nil {
			return err
		}
	}
	if m.LnTlToBr != nil {
		if err := m.LnTlToBr.ValidateWithPath(path + "/LnTlToBr"); err != nil {
			return err
		}
	}
	if m.LnBlToTr != nil {
		if err := m.LnBlToTr.ValidateWithPath(path + "/LnBlToTr"); err != nil {
			return err
		}
	}
	if m.Cell3D != nil {
		if err := m.Cell3D.ValidateWithPath(path + "/Cell3D"); err != nil {
			return err
		}
	}
	if m.NoFill != nil {
		if err := m.NoFill.ValidateWithPath(path + "/NoFill"); err != nil {
			return err
		}
	}
	if m.SolidFill != nil {
		if err := m.SolidFill.ValidateWithPath(path + "/SolidFill"); err != nil {
			return err
		}
	}
	if m.GradFill != nil {
		if err := m.GradFill.ValidateWithPath(path + "/GradFill"); err != nil {
			return err
		}
	}
	if m.BlipFill != nil {
		if err := m.BlipFill.ValidateWithPath(path + "/BlipFill"); err != nil {
			return err
		}
	}
	if m.PattFill != nil {
		if err := m.PattFill.ValidateWithPath(path + "/PattFill"); err != nil {
			return err
		}
	}
	if m.GrpFill != nil {
		if err := m.GrpFill.ValidateWithPath(path + "/GrpFill"); err != nil {
			return err
		}
	}
	if m.Headers != nil {
		if err := m.Headers.ValidateWithPath(path + "/Headers"); err != nil {
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
