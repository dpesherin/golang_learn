package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"partner.portal/internal/dto"
	"partner.portal/internal/models"
)

type JWTTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func RegController(c *gin.Context) {
	User := dto.User{}
	err := c.BindJSON(&User)
	var r string
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	r, err = models.CreateUser(User)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, r)
	return
}
func LoginController(c *gin.Context) {
	var access string
	var refresh string
	var err error
	var r dto.User

	User := dto.User{}
	err = c.BindJSON(&User)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	r, err = models.GetUserByLogin(User.Login)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	access, err = models.GenJWT(r, 1)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	refresh, err = models.GenJWT(r, 24)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	jwt := JWTTokens{access, refresh}

	c.JSON(http.StatusOK, jwt)
}
