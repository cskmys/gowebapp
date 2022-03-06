package controllers

import (
	"gowebapp/views"
	"net/http"
)

type Users struct {
	NewView *views.View
}

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"), // hardcoding just for now
	}
}

func (u *Users) New(w http.ResponseWriter, _ *http.Request) {
	u.NewView.Render(w, nil)
}
