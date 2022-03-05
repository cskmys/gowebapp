package views

import (
	"html/template"
	"path/filepath"
)

type View struct {
	Template *template.Template
	Layout   string
}

func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...) // append takes a variadic parameter as the 2nd argument,
	// so we put "..." after "layoutFiles()" to unpack it's return which is a string slice and pass it as a variadic parameter

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
	layoutFiles, err := filepath.Glob("views/layouts/*.gohtml")
	if err != nil {
		panic(err)
	}
	return layoutFiles
}
