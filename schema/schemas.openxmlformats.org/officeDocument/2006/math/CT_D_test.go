// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/math"
)

func TestCT_DConstructor(t *testing.T) {
	v := math.NewCT_D()
	if v == nil {
		t.Errorf("math.NewCT_D must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_D should validate: %s", err)
	}
}

func TestCT_DMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_D()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_D()
	xml.Unmarshal(buf, v2)
}
