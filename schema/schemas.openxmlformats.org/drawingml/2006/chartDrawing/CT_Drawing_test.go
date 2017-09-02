// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chartDrawing_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/chartDrawing"
)

func TestCT_DrawingConstructor(t *testing.T) {
	v := chartDrawing.NewCT_Drawing()
	if v == nil {
		t.Errorf("chartDrawing.NewCT_Drawing must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chartDrawing.CT_Drawing should validate: %s", err)
	}
}

func TestCT_DrawingMarshalUnmarshal(t *testing.T) {
	v := chartDrawing.NewCT_Drawing()
	buf, _ := xml.Marshal(v)
	v2 := chartDrawing.NewCT_Drawing()
	xml.Unmarshal(buf, v2)
}
