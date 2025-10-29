package controllers

type LogoutController struct {
	BaseController
}

func (R *LogoutController) Get() {
	R.DestroySession()
	R.Data["json"] = map[string]interface{}{"success": true}
	R.ServeJSON()
}
