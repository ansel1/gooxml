// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package terms_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/purl.org/dc/terms"
)

func TestURIConstructor(t *testing.T) {
	v := terms.NewURI()
	if v == nil {
		t.Errorf("terms.NewURI must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.URI should validate: %s", err)
	}
}

func TestURIMarshalUnmarshal(t *testing.T) {
	v := terms.NewURI()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewURI()
	xml.Unmarshal(buf, v2)
}
