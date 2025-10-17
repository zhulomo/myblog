package controllers

type LogoutController struct {
	BaseController
}

func (R *LogoutController) Get() {
	R.DestroySession()
	R.Redirect("/login", 302)
}
