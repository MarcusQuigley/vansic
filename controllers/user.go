package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIdPattern *regexp.Regexp
}

func (uc userController) ServeHttp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from User Controller"))
}

func newUserController() *userController {
	return &userController{
		userIdPattern: regexp.MustCompile(`^/users/(\d+)/?`), //looking for path thats /user/some number
	}
}
