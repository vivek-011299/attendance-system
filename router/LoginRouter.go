package router

import "my-project/RestHandler"

func LoginButton() {
	Router.HandleFunc("/login", RestHandler.LoginRest)
	Router.HandleFunc("/signup",RestHandler.Signup)
}