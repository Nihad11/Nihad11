package lib

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//author: gurbanli

type AdminController struct{

}

func (ac *AdminController) GetAdminPage(c *gin.Context){
	sess := sessions.Default(c)
	if sess.Get("is_admin") == true{
		c.String(200, "Admin Page")
	}else{
		c.String(401, "You are not admin !!!!")
	}
}