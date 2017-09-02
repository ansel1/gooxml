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

func TestCT_GraphicalObjectFrameConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewCT_GraphicalObjectFrame()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewCT_GraphicalObjectFrame must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.CT_GraphicalObjectFrame should validate: %s", err)
	}
}

func TestCT_GraphicalObjectFrameMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewCT_GraphicalObjectFrame()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewCT_GraphicalObjectFrame()
	xml.Unmarshal(buf, v2)
}
