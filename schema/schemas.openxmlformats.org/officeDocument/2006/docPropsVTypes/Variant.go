// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package docPropsVTypes

import (
	"encoding/xml"
	"log"
	"time"
)

type Variant struct {
	CT_Variant
}

func NewVariant() *Variant {
	ret := &Variant{}
	ret.CT_Variant = *NewCT_Variant()
	return ret
}

func (m *Variant) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.CT_Variant.MarshalXML(e, start)
}

func (m *Variant) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Variant = *NewCT_Variant()
lVariant:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "variant":
				m.Variant = NewVariant()
				if err := d.DecodeElement(m.Variant, &el); err != nil {
					return err
				}
			case "vector":
				m.Vector = NewVector()
				if err := d.DecodeElement(m.Vector, &el); err != nil {
					return err
				}
			case "array":
				m.Array = NewArray()
				if err := d.DecodeElement(m.Array, &el); err != nil {
					return err
				}
			case "blob":
				m.Blob = new(string)
				if err := d.DecodeElement(m.Blob, &el); err != nil {
					return err
				}
			case "oblob":
				m.Oblob = new(string)
				if err := d.DecodeElement(m.Oblob, &el); err != nil {
					return err
				}
			case "empty":
				m.Empty = NewEmpty()
				if err := d.DecodeElement(m.Empty, &el); err != nil {
					return err
				}
			case "null":
				m.Null = NewNull()
				if err := d.DecodeElement(m.Null, &el); err != nil {
					return err
				}
			case "i1":
				m.I1 = new(int8)
				if err := d.DecodeElement(m.I1, &el); err != nil {
					return err
				}
			case "i2":
				m.I2 = new(int16)
				if err := d.DecodeElement(m.I2, &el); err != nil {
					return err
				}
			case "i4":
				m.I4 = new(int32)
				if err := d.DecodeElement(m.I4, &el); err != nil {
					return err
				}
			case "i8":
				m.I8 = new(int64)
				if err := d.DecodeElement(m.I8, &el); err != nil {
					return err
				}
			case "int":
				m.Int = new(int32)
				if err := d.DecodeElement(m.Int, &el); err != nil {
					return err
				}
			case "ui1":
				m.Ui1 = new(uint8)
				if err := d.DecodeElement(m.Ui1, &el); err != nil {
					return err
				}
			case "ui2":
				m.Ui2 = new(uint16)
				if err := d.DecodeElement(m.Ui2, &el); err != nil {
					return err
				}
			case "ui4":
				m.Ui4 = new(uint32)
				if err := d.DecodeElement(m.Ui4, &el); err != nil {
					return err
				}
			case "ui8":
				m.Ui8 = new(uint64)
				if err := d.DecodeElement(m.Ui8, &el); err != nil {
					return err
				}
			case "uint":
				m.Uint = new(uint32)
				if err := d.DecodeElement(m.Uint, &el); err != nil {
					return err
				}
			case "r4":
				m.R4 = new(float32)
				if err := d.DecodeElement(m.R4, &el); err != nil {
					return err
				}
			case "r8":
				m.R8 = new(float64)
				if err := d.DecodeElement(m.R8, &el); err != nil {
					return err
				}
			case "decimal":
				m.Decimal = new(float64)
				if err := d.DecodeElement(m.Decimal, &el); err != nil {
					return err
				}
			case "lpstr":
				m.Lpstr = new(string)
				if err := d.DecodeElement(m.Lpstr, &el); err != nil {
					return err
				}
			case "lpwstr":
				m.Lpwstr = new(string)
				if err := d.DecodeElement(m.Lpwstr, &el); err != nil {
					return err
				}
			case "bstr":
				m.Bstr = new(string)
				if err := d.DecodeElement(m.Bstr, &el); err != nil {
					return err
				}
			case "date":
				m.Date = new(time.Time)
				if err := d.DecodeElement(m.Date, &el); err != nil {
					return err
				}
			case "filetime":
				m.Filetime = new(time.Time)
				if err := d.DecodeElement(m.Filetime, &el); err != nil {
					return err
				}
			case "bool":
				m.Bool = new(bool)
				if err := d.DecodeElement(m.Bool, &el); err != nil {
					return err
				}
			case "cy":
				m.Cy = new(string)
				if err := d.DecodeElement(m.Cy, &el); err != nil {
					return err
				}
			case "error":
				m.Error = new(string)
				if err := d.DecodeElement(m.Error, &el); err != nil {
					return err
				}
			case "stream":
				m.Stream = new(string)
				if err := d.DecodeElement(m.Stream, &el); err != nil {
					return err
				}
			case "ostream":
				m.Ostream = new(string)
				if err := d.DecodeElement(m.Ostream, &el); err != nil {
					return err
				}
			case "storage":
				m.Storage = new(string)
				if err := d.DecodeElement(m.Storage, &el); err != nil {
					return err
				}
			case "ostorage":
				m.Ostorage = new(string)
				if err := d.DecodeElement(m.Ostorage, &el); err != nil {
					return err
				}
			case "vstream":
				m.Vstream = NewVstream()
				if err := d.DecodeElement(m.Vstream, &el); err != nil {
					return err
				}
			case "clsid":
				m.Clsid = new(string)
				if err := d.DecodeElement(m.Clsid, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on Variant %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lVariant
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Variant and its children
func (m *Variant) Validate() error {
	return m.ValidateWithPath("Variant")
}

// ValidateWithPath validates the Variant and its children, prefixing error messages with path
func (m *Variant) ValidateWithPath(path string) error {
	if err := m.CT_Variant.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
