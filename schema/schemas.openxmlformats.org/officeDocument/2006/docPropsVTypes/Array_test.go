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

func TestArrayConstructor(t *testing.T) {
	v := docPropsVTypes.NewArray()
	if v == nil {
		t.Errorf("docPropsVTypes.NewArray must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed docPropsVTypes.Array should validate: %s", err)
	}
}

func TestArrayMarshalUnmarshal(t *testing.T) {
	v := docPropsVTypes.NewArray()
	buf, _ := xml.Marshal(v)
	v2 := docPropsVTypes.NewArray()
	xml.Unmarshal(buf, v2)
}
