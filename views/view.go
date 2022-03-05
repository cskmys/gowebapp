package views

import "html/template"

type View struct {
	Template *template.Template
}

func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml") // for now, we just hardcode this and later we will make it take all files in the "layouts" folder

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
	}
}
