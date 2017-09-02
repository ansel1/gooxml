// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package docPropsVTypes_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes"
)

func TestCT_VectorConstructor(t *testing.T) {
	v := docPropsVTypes.NewCT_Vector()
	if v == nil {
		t.Errorf("docPropsVTypes.NewCT_Vector must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed docPropsVTypes.CT_Vector should validate: %s", err)
	}
}

func TestCT_VectorMarshalUnmarshal(t *testing.T) {
	v := docPropsVTypes.NewCT_Vector()
	buf, _ := xml.Marshal(v)
	v2 := docPropsVTypes.NewCT_Vector()
	xml.Unmarshal(buf, v2)
}
