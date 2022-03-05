package views

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	LayoutDir     = "views/layouts/"
	LayoutFileExt = "gohtml"
)

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error { // "interface{}" type is used to pass argument of type "any"
	return v.Template.ExecuteTemplate(w, v.Layout, data)
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
