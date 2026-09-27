package main

import "github.com/speedyhoon/jay/genjay/testdata/imports/ext"

type Section struct {
	Id       ext.ID `j:"-"`
	Name     string
	Color    ext.C16
	Sectors  ext.Sectors
	Sections [5]ext.Sections
	Project  *Project
	Order    ext.ID
}

type Project struct {
	Id   ext.ID `j:"-"`
	Name string
}
