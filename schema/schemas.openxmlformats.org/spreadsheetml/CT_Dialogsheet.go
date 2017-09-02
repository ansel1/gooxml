// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"log"
)

type CT_Dialogsheet struct {
	// Sheet Properties
	SheetPr *CT_SheetPr
	// Dialog Sheet Views
	SheetViews *CT_SheetViews
	// Dialog Sheet Format Properties
	SheetFormatPr *CT_SheetFormatPr
	// Sheet Protection
	SheetProtection *CT_SheetProtection
	// Custom Sheet Views
	CustomSheetViews *CT_CustomSheetViews
	// Print Options
	PrintOptions *CT_PrintOptions
	// Page Margins
	PageMargins *CT_PageMargins
	// Page Setup Settings
	PageSetup *CT_PageSetup
	// Header & Footer Settings
	HeaderFooter *CT_HeaderFooter
	// Drawing
	Drawing *CT_Drawing
	// Legacy Drawing
	LegacyDrawing *CT_LegacyDrawing
	// Legacy Drawing Header Footer
	LegacyDrawingHF *CT_LegacyDrawing
	DrawingHF       *CT_DrawingHF
	OleObjects      *CT_OleObjects
	Controls        *CT_Controls
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_Dialogsheet() *CT_Dialogsheet {
	ret := &CT_Dialogsheet{}
	return ret
}

func (m *CT_Dialogsheet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.SheetPr != nil {
		sesheetPr := xml.StartElement{Name: xml.Name{Local: "x:sheetPr"}}
		e.EncodeElement(m.SheetPr, sesheetPr)
	}
	if m.SheetViews != nil {
		sesheetViews := xml.StartElement{Name: xml.Name{Local: "x:sheetViews"}}
		e.EncodeElement(m.SheetViews, sesheetViews)
	}
	if m.SheetFormatPr != nil {
		sesheetFormatPr := xml.StartElement{Name: xml.Name{Local: "x:sheetFormatPr"}}
		e.EncodeElement(m.SheetFormatPr, sesheetFormatPr)
	}
	if m.SheetProtection != nil {
		sesheetProtection := xml.StartElement{Name: xml.Name{Local: "x:sheetProtection"}}
		e.EncodeElement(m.SheetProtection, sesheetProtection)
	}
	if m.CustomSheetViews != nil {
		secustomSheetViews := xml.StartElement{Name: xml.Name{Local: "x:customSheetViews"}}
		e.EncodeElement(m.CustomSheetViews, secustomSheetViews)
	}
	if m.PrintOptions != nil {
		seprintOptions := xml.StartElement{Name: xml.Name{Local: "x:printOptions"}}
		e.EncodeElement(m.PrintOptions, seprintOptions)
	}
	if m.PageMargins != nil {
		sepageMargins := xml.StartElement{Name: xml.Name{Local: "x:pageMargins"}}
		e.EncodeElement(m.PageMargins, sepageMargins)
	}
	if m.PageSetup != nil {
		sepageSetup := xml.StartElement{Name: xml.Name{Local: "x:pageSetup"}}
		e.EncodeElement(m.PageSetup, sepageSetup)
	}
	if m.HeaderFooter != nil {
		seheaderFooter := xml.StartElement{Name: xml.Name{Local: "x:headerFooter"}}
		e.EncodeElement(m.HeaderFooter, seheaderFooter)
	}
	if m.Drawing != nil {
		sedrawing := xml.StartElement{Name: xml.Name{Local: "x:drawing"}}
		e.EncodeElement(m.Drawing, sedrawing)
	}
	if m.LegacyDrawing != nil {
		selegacyDrawing := xml.StartElement{Name: xml.Name{Local: "x:legacyDrawing"}}
		e.EncodeElement(m.LegacyDrawing, selegacyDrawing)
	}
	if m.LegacyDrawingHF != nil {
		selegacyDrawingHF := xml.StartElement{Name: xml.Name{Local: "x:legacyDrawingHF"}}
		e.EncodeElement(m.LegacyDrawingHF, selegacyDrawingHF)
	}
	if m.DrawingHF != nil {
		sedrawingHF := xml.StartElement{Name: xml.Name{Local: "x:drawingHF"}}
		e.EncodeElement(m.DrawingHF, sedrawingHF)
	}
	if m.OleObjects != nil {
		seoleObjects := xml.StartElement{Name: xml.Name{Local: "x:oleObjects"}}
		e.EncodeElement(m.OleObjects, seoleObjects)
	}
	if m.Controls != nil {
		secontrols := xml.StartElement{Name: xml.Name{Local: "x:controls"}}
		e.EncodeElement(m.Controls, secontrols)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Dialogsheet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Dialogsheet:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "sheetPr":
				m.SheetPr = NewCT_SheetPr()
				if err := d.DecodeElement(m.SheetPr, &el); err != nil {
					return err
				}
			case "sheetViews":
				m.SheetViews = NewCT_SheetViews()
				if err := d.DecodeElement(m.SheetViews, &el); err != nil {
					return err
				}
			case "sheetFormatPr":
				m.SheetFormatPr = NewCT_SheetFormatPr()
				if err := d.DecodeElement(m.SheetFormatPr, &el); err != nil {
					return err
				}
			case "sheetProtection":
				m.SheetProtection = NewCT_SheetProtection()
				if err := d.DecodeElement(m.SheetProtection, &el); err != nil {
					return err
				}
			case "customSheetViews":
				m.CustomSheetViews = NewCT_CustomSheetViews()
				if err := d.DecodeElement(m.CustomSheetViews, &el); err != nil {
					return err
				}
			case "printOptions":
				m.PrintOptions = NewCT_PrintOptions()
				if err := d.DecodeElement(m.PrintOptions, &el); err != nil {
					return err
				}
			case "pageMargins":
				m.PageMargins = NewCT_PageMargins()
				if err := d.DecodeElement(m.PageMargins, &el); err != nil {
					return err
				}
			case "pageSetup":
				m.PageSetup = NewCT_PageSetup()
				if err := d.DecodeElement(m.PageSetup, &el); err != nil {
					return err
				}
			case "headerFooter":
				m.HeaderFooter = NewCT_HeaderFooter()
				if err := d.DecodeElement(m.HeaderFooter, &el); err != nil {
					return err
				}
			case "drawing":
				m.Drawing = NewCT_Drawing()
				if err := d.DecodeElement(m.Drawing, &el); err != nil {
					return err
				}
			case "legacyDrawing":
				m.LegacyDrawing = NewCT_LegacyDrawing()
				if err := d.DecodeElement(m.LegacyDrawing, &el); err != nil {
					return err
				}
			case "legacyDrawingHF":
				m.LegacyDrawingHF = NewCT_LegacyDrawing()
				if err := d.DecodeElement(m.LegacyDrawingHF, &el); err != nil {
					return err
				}
			case "drawingHF":
				m.DrawingHF = NewCT_DrawingHF()
				if err := d.DecodeElement(m.DrawingHF, &el); err != nil {
					return err
				}
			case "oleObjects":
				m.OleObjects = NewCT_OleObjects()
				if err := d.DecodeElement(m.OleObjects, &el); err != nil {
					return err
				}
			case "controls":
				m.Controls = NewCT_Controls()
				if err := d.DecodeElement(m.Controls, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_Dialogsheet %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Dialogsheet
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Dialogsheet and its children
func (m *CT_Dialogsheet) Validate() error {
	return m.ValidateWithPath("CT_Dialogsheet")
}

// ValidateWithPath validates the CT_Dialogsheet and its children, prefixing error messages with path
func (m *CT_Dialogsheet) ValidateWithPath(path string) error {
	if m.SheetPr != nil {
		if err := m.SheetPr.ValidateWithPath(path + "/SheetPr"); err != nil {
			return err
		}
	}
	if m.SheetViews != nil {
		if err := m.SheetViews.ValidateWithPath(path + "/SheetViews"); err != nil {
			return err
		}
	}
	if m.SheetFormatPr != nil {
		if err := m.SheetFormatPr.ValidateWithPath(path + "/SheetFormatPr"); err != nil {
			return err
		}
	}
	if m.SheetProtection != nil {
		if err := m.SheetProtection.ValidateWithPath(path + "/SheetProtection"); err != nil {
			return err
		}
	}
	if m.CustomSheetViews != nil {
		if err := m.CustomSheetViews.ValidateWithPath(path + "/CustomSheetViews"); err != nil {
			return err
		}
	}
	if m.PrintOptions != nil {
		if err := m.PrintOptions.ValidateWithPath(path + "/PrintOptions"); err != nil {
			return err
		}
	}
	if m.PageMargins != nil {
		if err := m.PageMargins.ValidateWithPath(path + "/PageMargins"); err != nil {
			return err
		}
	}
	if m.PageSetup != nil {
		if err := m.PageSetup.ValidateWithPath(path + "/PageSetup"); err != nil {
			return err
		}
	}
	if m.HeaderFooter != nil {
		if err := m.HeaderFooter.ValidateWithPath(path + "/HeaderFooter"); err != nil {
			return err
		}
	}
	if m.Drawing != nil {
		if err := m.Drawing.ValidateWithPath(path + "/Drawing"); err != nil {
			return err
		}
	}
	if m.LegacyDrawing != nil {
		if err := m.LegacyDrawing.ValidateWithPath(path + "/LegacyDrawing"); err != nil {
			return err
		}
	}
	if m.LegacyDrawingHF != nil {
		if err := m.LegacyDrawingHF.ValidateWithPath(path + "/LegacyDrawingHF"); err != nil {
			return err
		}
	}
	if m.DrawingHF != nil {
		if err := m.DrawingHF.ValidateWithPath(path + "/DrawingHF"); err != nil {
			return err
		}
	}
	if m.OleObjects != nil {
		if err := m.OleObjects.ValidateWithPath(path + "/OleObjects"); err != nil {
			return err
		}
	}
	if m.Controls != nil {
		if err := m.Controls.ValidateWithPath(path + "/Controls"); err != nil {
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
