package views

import (
	"fmt"
	"html/template"
	"path/filepath"
)

var (
	LayoutDir     string = "views/layouts/"
	LayoutFileExt string = "gohtml"
)

type View struct {
	Template *template.Template
	Layout   string
}

func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

func layoutFiles() []string {
	layoutGlob := fmt.Sprintf("%s/*.%s", LayoutDir, LayoutFileExt)
	layoutFiles, err := filepath.Glob(layoutGlob)
	if err != nil {
		panic(err)
	}
	return layoutFiles
}
