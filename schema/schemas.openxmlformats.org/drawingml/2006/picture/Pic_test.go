// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package picture_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/picture"
)

func TestPicConstructor(t *testing.T) {
	v := picture.NewPic()
	if v == nil {
		t.Errorf("picture.NewPic must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed picture.Pic should validate: %s", err)
	}
}

func TestPicMarshalUnmarshal(t *testing.T) {
	v := picture.NewPic()
	buf, _ := xml.Marshal(v)
	v2 := picture.NewPic()
	xml.Unmarshal(buf, v2)
}
