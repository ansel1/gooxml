// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/diagram"
)

func TestStyleDefHdrLstConstructor(t *testing.T) {
	v := diagram.NewStyleDefHdrLst()
	if v == nil {
		t.Errorf("diagram.NewStyleDefHdrLst must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.StyleDefHdrLst should validate: %s", err)
	}
}

func TestStyleDefHdrLstMarshalUnmarshal(t *testing.T) {
	v := diagram.NewStyleDefHdrLst()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewStyleDefHdrLst()
	xml.Unmarshal(buf, v2)
}
