// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/presentationml"
)

func TestEG_ShowTypeConstructor(t *testing.T) {
	v := presentationml.NewEG_ShowType()
	if v == nil {
		t.Errorf("presentationml.NewEG_ShowType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed presentationml.EG_ShowType should validate: %s", err)
	}
}

func TestEG_ShowTypeMarshalUnmarshal(t *testing.T) {
	v := presentationml.NewEG_ShowType()
	buf, _ := xml.Marshal(v)
	v2 := presentationml.NewEG_ShowType()
	xml.Unmarshal(buf, v2)
}
