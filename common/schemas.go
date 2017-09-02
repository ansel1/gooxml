// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package common

// Consts for content types used throughout the package
const (
	// Common
	OfficeDocumentType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument"
	StylesType         = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/styles"
	ThemeType          = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/theme"
	ThemeContentType   = "application/vnd.openxmlformats-officedocument.theme+xml"
	SettingsType       = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/settings"
	ImageType          = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/image"
	CommentsType       = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/comments"
	ThumbnailType      = "http://schemas.openxmlformats.org/package/2006/relationships/metadata/thumbnail"

	ExtendedPropertiesType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/extended-properties"
	CorePropertiesType     = "http://schemas.openxmlformats.org/package/2006/relationships/metadata/core-properties"

	// SML
	WorksheetType            = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/worksheet"
	SharedStringsContentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sharedStrings+xml"
	SharedStingsType         = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/sharedStrings"

	// WML
	HeaderType      = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/header"
	FooterType      = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/footer"
	NumberingType   = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/numbering"
	FontTableType   = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/fontTable"
	WebSettingsType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/webSettings"
	FootNotesType   = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/footnotes"
	EndNotesType    = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/endnotes"

	// PML
	SlideType                   = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/slide"
	SlideContentType            = "application/vnd.openxmlformats-officedocument.presentationml.slide+xml"
	SlideMasterRelationshipType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/slideMaster"
	SlideMasterContentType      = "application/vnd.openxmlformats-officedocument.presentationml.slideMaster+xml"
	SlideLayoutType             = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/slideLayout"
	SlideLayoutContentType      = "application/vnd.openxmlformats-officedocument.presentationml.slideLayout+xml"
)