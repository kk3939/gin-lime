package userController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kk3939/gin-lime/controllers"
	"github.com/kk3939/gin-lime/entity"
	"github.com/kk3939/gin-lime/models/userModel"
)

func GetUsers(c *gin.Context) {
	users, err := userModel.GetUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": controllers.FmtErrMsg(err),
		})
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := userModel.GetUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": controllers.FmtErrMsg(err),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	user := entity.User{}
	c.BindJSON(&user)
	err := userModel.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": controllers.FmtErrMsg(err),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	if userGot, err := userModel.GetUser(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": controllers.FmtErrMsg(err),
		})
		return
	} else {
		user := entity.User{}
		c.BindJSON(&user)
		user.Id = userGot.Id
		if err := userModel.UpdateUser(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": controllers.FmtErrMsg(err),
			})
			return
		}
		c.JSON(http.StatusOK, user)
	}

}
func DeleteUser(c *gin.Context) {
	if userGot, err := userModel.GetUser(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": controllers.FmtErrMsg(err),
		})
		return
	} else {
		user := entity.User{}
		user.Id = userGot.Id
		if err := userModel.DeleteUser(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": controllers.FmtErrMsg(err),
			})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}
