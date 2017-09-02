// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"log"
	"strconv"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type Presentation struct {
	CT_Presentation
}

func NewPresentation() *Presentation {
	ret := &Presentation{}
	ret.CT_Presentation = *NewCT_Presentation()
	return ret
}

func (m *Presentation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/presentationml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:p"}, Value: "http://schemas.openxmlformats.org/presentationml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "p:presentation"
	return m.CT_Presentation.MarshalXML(e, start)
}

func (m *Presentation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Presentation = *NewCT_Presentation()
	for _, attr := range start.Attr {
		if attr.Name.Local == "serverZoom" {
			parsed, err := ParseUnionST_Percentage(attr.Value)
			if err != nil {
				return err
			}
			m.ServerZoomAttr = &parsed
		}
		if attr.Name.Local == "firstSlideNum" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.FirstSlideNumAttr = &pt
		}
		if attr.Name.Local == "showSpecialPlsOnTitleSld" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowSpecialPlsOnTitleSldAttr = &parsed
		}
		if attr.Name.Local == "rtl" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RtlAttr = &parsed
		}
		if attr.Name.Local == "removePersonalInfoOnSave" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RemovePersonalInfoOnSaveAttr = &parsed
		}
		if attr.Name.Local == "compatMode" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CompatModeAttr = &parsed
		}
		if attr.Name.Local == "strictFirstAndLastChars" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.StrictFirstAndLastCharsAttr = &parsed
		}
		if attr.Name.Local == "embedTrueTypeFonts" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.EmbedTrueTypeFontsAttr = &parsed
		}
		if attr.Name.Local == "saveSubsetFonts" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SaveSubsetFontsAttr = &parsed
		}
		if attr.Name.Local == "autoCompressPictures" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AutoCompressPicturesAttr = &parsed
		}
		if attr.Name.Local == "bookmarkIdSeed" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.BookmarkIdSeedAttr = &pt
		}
		if attr.Name.Local == "conformance" {
			m.ConformanceAttr.UnmarshalXMLAttr(attr)
		}
	}
lPresentation:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "sldMasterIdLst":
				m.SldMasterIdLst = NewCT_SlideMasterIdList()
				if err := d.DecodeElement(m.SldMasterIdLst, &el); err != nil {
					return err
				}
			case "notesMasterIdLst":
				m.NotesMasterIdLst = NewCT_NotesMasterIdList()
				if err := d.DecodeElement(m.NotesMasterIdLst, &el); err != nil {
					return err
				}
			case "handoutMasterIdLst":
				m.HandoutMasterIdLst = NewCT_HandoutMasterIdList()
				if err := d.DecodeElement(m.HandoutMasterIdLst, &el); err != nil {
					return err
				}
			case "sldIdLst":
				m.SldIdLst = NewCT_SlideIdList()
				if err := d.DecodeElement(m.SldIdLst, &el); err != nil {
					return err
				}
			case "sldSz":
				m.SldSz = NewCT_SlideSize()
				if err := d.DecodeElement(m.SldSz, &el); err != nil {
					return err
				}
			case "notesSz":
				if err := d.DecodeElement(m.NotesSz, &el); err != nil {
					return err
				}
			case "smartTags":
				m.SmartTags = NewCT_SmartTags()
				if err := d.DecodeElement(m.SmartTags, &el); err != nil {
					return err
				}
			case "embeddedFontLst":
				m.EmbeddedFontLst = NewCT_EmbeddedFontList()
				if err := d.DecodeElement(m.EmbeddedFontLst, &el); err != nil {
					return err
				}
			case "custShowLst":
				m.CustShowLst = NewCT_CustomShowList()
				if err := d.DecodeElement(m.CustShowLst, &el); err != nil {
					return err
				}
			case "photoAlbum":
				m.PhotoAlbum = NewCT_PhotoAlbum()
				if err := d.DecodeElement(m.PhotoAlbum, &el); err != nil {
					return err
				}
			case "custDataLst":
				m.CustDataLst = NewCT_CustomerDataList()
				if err := d.DecodeElement(m.CustDataLst, &el); err != nil {
					return err
				}
			case "kinsoku":
				m.Kinsoku = NewCT_Kinsoku()
				if err := d.DecodeElement(m.Kinsoku, &el); err != nil {
					return err
				}
			case "defaultTextStyle":
				m.DefaultTextStyle = drawingml.NewCT_TextListStyle()
				if err := d.DecodeElement(m.DefaultTextStyle, &el); err != nil {
					return err
				}
			case "modifyVerifier":
				m.ModifyVerifier = NewCT_ModifyVerifier()
				if err := d.DecodeElement(m.ModifyVerifier, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on Presentation %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lPresentation
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Presentation and its children
func (m *Presentation) Validate() error {
	return m.ValidateWithPath("Presentation")
}

// ValidateWithPath validates the Presentation and its children, prefixing error messages with path
func (m *Presentation) ValidateWithPath(path string) error {
	if err := m.CT_Presentation.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}