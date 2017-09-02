// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

func TestEG_TextUnderlineLineConstructor(t *testing.T) {
	v := drawingml.NewEG_TextUnderlineLine()
	if v == nil {
		t.Errorf("drawingml.NewEG_TextUnderlineLine must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed drawingml.EG_TextUnderlineLine should validate: %s", err)
	}
}

func TestEG_TextUnderlineLineMarshalUnmarshal(t *testing.T) {
	v := drawingml.NewEG_TextUnderlineLine()
	buf, _ := xml.Marshal(v)
	v2 := drawingml.NewEG_TextUnderlineLine()
	xml.Unmarshal(buf, v2)
}
