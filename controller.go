package controller

import (
	"Ders16_MYSQL/4.dto"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)




type AuthController struct {

}

func (ac *AuthController) Login(c *gin.Context){
	var loginRequest dto.LoginRequest
	err := c.ShouldBindJSON(&loginRequest)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	if loginRequest.Username == "admin" && loginRequest.Password == "12345"{
		sess := sessions.Default(c)
		sess.Set("username", loginRequest.Username)
		sess.Set("is_authenticated", true)
		if loginRequest.Username == "admin"{
			sess.Set("is_admin", true)
		}
		err := sess.Save()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Can't set session",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome !",
		})
	}else{
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Username or password is incorrect !",
		})
	}
}

func (ac *AuthController) Register(c *gin.Context){

}
