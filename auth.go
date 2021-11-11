package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

//author: gurbanli

type ProductController struct {
}

	func(pc *ProductController) GetMyUsername (c *gin.Context){
	sess := sessions.Default(c)
	c.JSON(http.StatusOK, gin.H{
	"username": sess.Get("username"),
})
}

	func (pc *ProductController) GetProducts(c *gin.Context){
	sess := sessions.Default(c)
	if sess.Get("is_authenticated") == true{
	c.JSON(http.StatusOK, gin.H{
	"products":"Audi A6",
})
} else {
	c.JSON(http.StatusUnauthorized, gin.H{
	"message":"You must login to view this product",
})
}
}
