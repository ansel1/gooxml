// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"
)

func TestCT_GroupShapeChoiceConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewCT_GroupShapeChoice()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewCT_GroupShapeChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.CT_GroupShapeChoice should validate: %s", err)
	}
}

func TestCT_GroupShapeChoiceMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewCT_GroupShapeChoice()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewCT_GroupShapeChoice()
	xml.Unmarshal(buf, v2)
}
