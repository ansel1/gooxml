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

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/math"
)

type CT_Row struct {
	// Revision Identifier for Table Row Glyph Formatting
	RsidRPrAttr *string
	// Revision Identifier for Table Row
	RsidRAttr *string
	// Revision Identifier for Table Row Deletion
	RsidDelAttr *string
	// Revision Identifier for Table Row Properties
	RsidTrAttr *string
	// Table-Level Property Exceptions
	TblPrEx *CT_TblPrEx
	// Table Row Properties
	TrPr                  *CT_TrPr
	EG_ContentCellContent []*EG_ContentCellContent
}

func NewCT_Row() *CT_Row {
	ret := &CT_Row{}
	return ret
}

func (m *CT_Row) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.RsidRPrAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidRPr"},
			Value: fmt.Sprintf("%v", *m.RsidRPrAttr)})
	}
	if m.RsidRAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidR"},
			Value: fmt.Sprintf("%v", *m.RsidRAttr)})
	}
	if m.RsidDelAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidDel"},
			Value: fmt.Sprintf("%v", *m.RsidDelAttr)})
	}
	if m.RsidTrAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidTr"},
			Value: fmt.Sprintf("%v", *m.RsidTrAttr)})
	}
	e.EncodeToken(start)
	if m.TblPrEx != nil {
		setblPrEx := xml.StartElement{Name: xml.Name{Local: "w:tblPrEx"}}
		e.EncodeElement(m.TblPrEx, setblPrEx)
	}
	if m.TrPr != nil {
		setrPr := xml.StartElement{Name: xml.Name{Local: "w:trPr"}}
		e.EncodeElement(m.TrPr, setrPr)
	}
	if m.EG_ContentCellContent != nil {
		for _, c := range m.EG_ContentCellContent {
			c.MarshalXML(e, start)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Row) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "rsidRPr" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RsidRPrAttr = &parsed
		}
		if attr.Name.Local == "rsidR" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RsidRAttr = &parsed
		}
		if attr.Name.Local == "rsidDel" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RsidDelAttr = &parsed
		}
		if attr.Name.Local == "rsidTr" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RsidTrAttr = &parsed
		}
	}
lCT_Row:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tblPrEx":
				m.TblPrEx = NewCT_TblPrEx()
				if err := d.DecodeElement(m.TblPrEx, &el); err != nil {
					return err
				}
			case "trPr":
				m.TrPr = NewCT_TrPr()
				if err := d.DecodeElement(m.TrPr, &el); err != nil {
					return err
				}
			case "tc":
				tmpcontentcellcontent := NewEG_ContentCellContent()
				tmp := NewCT_Tc()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				tmpcontentcellcontent.Tc = append(tmpcontentcellcontent.Tc, tmp)
				m.EG_ContentCellContent = append(m.EG_ContentCellContent, tmpcontentcellcontent)
			case "customXml":
				tmpcontentcellcontent := NewEG_ContentCellContent()
				tmpcontentcellcontent.CustomXml = NewCT_CustomXmlCell()
				if err := d.DecodeElement(tmpcontentcellcontent.CustomXml, &el); err != nil {
					return err
				}
				m.EG_ContentCellContent = append(m.EG_ContentCellContent, tmpcontentcellcontent)
			case "sdt":
				tmpcontentcellcontent := NewEG_ContentCellContent()
				tmpcontentcellcontent.Sdt = NewCT_SdtCell()
				if err := d.DecodeElement(tmpcontentcellcontent.Sdt, &el); err != nil {
					return err
				}
				m.EG_ContentCellContent = append(m.EG_ContentCellContent, tmpcontentcellcontent)
			case "proofErr":
				tmpcontentcellcontent := NewEG_ContentCellContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.ProofErr = NewCT_ProofErr()
				if err := d.DecodeElement(tmprunlevelelts.ProofErr, &el); err != nil {
					return err
				}
				m.EG_ContentCellContent = append(m.EG_ContentCellContent, tmpcontentcellcontent)
				tmpcontentcellcontent.EG_RunLevelElts = append(tmpcontentcellcontent.EG_RunLevelElts, tmprunlevelelts)
			case "permStart":
				tmpcontentcellcontent := NewEG_ContentCellContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.PermStart = NewCT_PermStart()
				if err := d.DecodeElement(tmprunlevelelts.PermStart, &el); err != nil {
					return err
				}
				m.EG_ContentCellContent = append(m.EG_ContentCellContent, tmpcontentcellcontent)
				tmpcontentcellcontent.EG_RunLevelElts = append(tmpcontentcellcontent.EG_RunLevelElts, tmprunlevelelts)
			case "permEnd":
				tmpcontentcellcontent := NewEG_ContentCellContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.PermEnd = NewCT_Perm()
				if err := d.DecodeElement(tmprunlevelelts.PermEnd, &el); err != nil {
					return err
				}
				m.EG_ContentCellContent = append(m.EG_ContentCellContent, tmpcontentcellcontent)
				tmpcontentcellcontent.EG_RunLevelElts = append(tmpcontentcellcontent.EG_RunLevelElts, tmprunlevelelts)
			case "ins":
				tmpcontentcellcontent := NewEG_ContentCellContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.Ins = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.Ins, &el); err != nil {
					return err
				}
				m.EG_ContentCellContent = append(m.EG_ContentCellContent, tmpcontentcellcontent)
				tmpcontentcellcontent.EG_RunLevelElts = append(tmpcontentcellcontent.EG_RunLevelElts, tmprunlevelelts)
			case "del":
				tmpcontentcellcontent := NewEG_ContentCellContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.Del = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.Del, &el); err != nil {
					return err
				}
				m.EG_ContentCellContent = append(m.EG_ContentCellContent, tmpcontentcellcontent)
				tmpcontentcellcontent.EG_RunLevelElts = append(tmpcontentcellcontent.EG_RunLevelElts, tmprunlevelelts)
			case "moveFrom":
				tmpcontentcellcontent := NewEG_ContentCellContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.MoveFrom = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.MoveFrom, &el); err != nil {
					return err
				}
				m.EG_ContentCellContent = append(m.EG_ContentCellContent, tmpcontentcellcontent)
				tmpcontentcellcontent.EG_RunLevelElts = append(tmpcontentcellcontent.EG_RunLevelElts, tmprunlevelelts)
			case "moveTo":
				tmpcontentcellcontent := NewEG_ContentCellContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.MoveTo = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.MoveTo, &el); err != nil {
					return err
				}
				m.EG_ContentCellContent = append(m.EG_ContentCellContent, tmpcontentcellcontent)
				tmpcontentcellcontent.EG_RunLevelElts = append(tmpcontentcellcontent.EG_RunLevelElts, tmprunlevelelts)
			case "bookmarkStart":
				tmpcontentcellcontent := NewEG_ContentCellContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.BookmarkStart = NewCT_Bookmark()
				if err := d.DecodeElement(tmprangemarkupelements.BookmarkStart, &el); err != nil {
					return err
				}
				m.EG_ContentCellContent = append(m.EG_ContentCellContent, tmpcontentcellcontent)
				tmpcontentcellcontent.EG_RunLevelElts = append(tmpcontentcellcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "bookmarkEnd":
				tmpcontentcellcontent := NewEG_ContentCellContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.BookmarkEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.BookmarkEnd, &el); err != nil {
					return err
				}
				m.EG_ContentCellContent = append(m.EG_ContentCellContent, tmpcontentcellcontent)
				tmpcontentcellcontent.EG_RunLevelElts = append(tmpcontentcellcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "moveFromRangeStart":
				tmpcontentcellcontent := NewEG_ContentCellContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveFromRangeStart = NewCT_MoveBookmark()
				if err := d.DecodeElement(tmprangemarkupelements.MoveFromRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentCellContent = append(m.EG_ContentCellContent, tmpcontentcellcontent)
				tmpcontentcellcontent.EG_RunLevelElts = append(tmpcontentcellcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "moveFromRangeEnd":
				tmpcontentcellcontent := NewEG_ContentCellContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveFromRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.MoveFromRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentCellContent = append(m.EG_ContentCellContent, tmpcontentcellcontent)
				tmpcontentcellcontent.EG_RunLevelElts = append(tmpcontentcellcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "moveToRangeStart":
				tmpcontentcellcontent := NewEG_ContentCellContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveToRangeStart = NewCT_MoveBookmark()
				if err := d.DecodeElement(tmprangemarkupelements.MoveToRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentCellContent = append(m.EG_ContentCellContent, tmpcontentcellcontent)
				tmpcontentcellcontent.EG_RunLevelElts = append(tmpcontentcellcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "moveToRangeEnd":
				tmpcontentcellcontent := NewEG_ContentCellContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveToRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.MoveToRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentCellContent = append(m.EG_ContentCellContent, tmpcontentcellcontent)
				tmpcontentcellcontent.EG_RunLevelElts = append(tmpcontentcellcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "commentRangeStart":
				tmpcontentcellcontent := NewEG_ContentCellContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CommentRangeStart = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.CommentRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentCellContent = append(m.EG_ContentCellContent, tmpcontentcellcontent)
				tmpcontentcellcontent.EG_RunLevelElts = append(tmpcontentcellcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "commentRangeEnd":
				tmpcontentcellcontent := NewEG_ContentCellContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CommentRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.CommentRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentCellContent = append(m.EG_ContentCellContent, tmpcontentcellcontent)
				tmpcontentcellcontent.EG_RunLevelElts = append(tmpcontentcellcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlInsRangeStart":
				tmpcontentcellcontent := NewEG_ContentCellContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlInsRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlInsRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentCellContent = append(m.EG_ContentCellContent, tmpcontentcellcontent)
				tmpcontentcellcontent.EG_RunLevelElts = append(tmpcontentcellcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlInsRangeEnd":
				tmpcontentcellcontent := NewEG_ContentCellContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlInsRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlInsRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentCellContent = append(m.EG_ContentCellContent, tmpcontentcellcontent)
				tmpcontentcellcontent.EG_RunLevelElts = append(tmpcontentcellcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlDelRangeStart":
				tmpcontentcellcontent := NewEG_ContentCellContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlDelRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlDelRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentCellContent = append(m.EG_ContentCellContent, tmpcontentcellcontent)
				tmpcontentcellcontent.EG_RunLevelElts = append(tmpcontentcellcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlDelRangeEnd":
				tmpcontentcellcontent := NewEG_ContentCellContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlDelRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlDelRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentCellContent = append(m.EG_ContentCellContent, tmpcontentcellcontent)
				tmpcontentcellcontent.EG_RunLevelElts = append(tmpcontentcellcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlMoveFromRangeStart":
				tmpcontentcellcontent := NewEG_ContentCellContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveFromRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveFromRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentCellContent = append(m.EG_ContentCellContent, tmpcontentcellcontent)
				tmpcontentcellcontent.EG_RunLevelElts = append(tmpcontentcellcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlMoveFromRangeEnd":
				tmpcontentcellcontent := NewEG_ContentCellContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveFromRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveFromRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentCellContent = append(m.EG_ContentCellContent, tmpcontentcellcontent)
				tmpcontentcellcontent.EG_RunLevelElts = append(tmpcontentcellcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlMoveToRangeStart":
				tmpcontentcellcontent := NewEG_ContentCellContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveToRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveToRangeStart, &el); err != nil {
					return err
				}
				m.EG_ContentCellContent = append(m.EG_ContentCellContent, tmpcontentcellcontent)
				tmpcontentcellcontent.EG_RunLevelElts = append(tmpcontentcellcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "customXmlMoveToRangeEnd":
				tmpcontentcellcontent := NewEG_ContentCellContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveToRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveToRangeEnd, &el); err != nil {
					return err
				}
				m.EG_ContentCellContent = append(m.EG_ContentCellContent, tmpcontentcellcontent)
				tmpcontentcellcontent.EG_RunLevelElts = append(tmpcontentcellcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case "oMathPara":
				tmpcontentcellcontent := NewEG_ContentCellContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmpmathcontent := NewEG_MathContent()
				tmpmathcontent.OMathPara = math.NewOMathPara()
				if err := d.DecodeElement(tmpmathcontent.OMathPara, &el); err != nil {
					return err
				}
				m.EG_ContentCellContent = append(m.EG_ContentCellContent, tmpcontentcellcontent)
				tmpcontentcellcontent.EG_RunLevelElts = append(tmpcontentcellcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_MathContent = append(tmprunlevelelts.EG_MathContent, tmpmathcontent)
			case "oMath":
				tmpcontentcellcontent := NewEG_ContentCellContent()
				tmprunlevelelts := NewEG_RunLevelElts()
				tmpmathcontent := NewEG_MathContent()
				tmpmathcontent.OMath = math.NewOMath()
				if err := d.DecodeElement(tmpmathcontent.OMath, &el); err != nil {
					return err
				}
				m.EG_ContentCellContent = append(m.EG_ContentCellContent, tmpcontentcellcontent)
				tmpcontentcellcontent.EG_RunLevelElts = append(tmpcontentcellcontent.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_MathContent = append(tmprunlevelelts.EG_MathContent, tmpmathcontent)
			default:
				log.Printf("skipping unsupported element on CT_Row %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Row
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Row and its children
func (m *CT_Row) Validate() error {
	return m.ValidateWithPath("CT_Row")
}

// ValidateWithPath validates the CT_Row and its children, prefixing error messages with path
func (m *CT_Row) ValidateWithPath(path string) error {
	if m.TblPrEx != nil {
		if err := m.TblPrEx.ValidateWithPath(path + "/TblPrEx"); err != nil {
			return err
		}
	}
	if m.TrPr != nil {
		if err := m.TrPr.ValidateWithPath(path + "/TrPr"); err != nil {
			return err
		}
	}
	for i, v := range m.EG_ContentCellContent {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_ContentCellContent[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}