package controller

import(
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/SaKu2110/test/model"
	"github.com/SaKu2110/test/service"
)

type IsController struct {
	DB	*gorm.DB
}

func (ctrl *IsController)SignInHandler(context *gin.Context){
	var request model.LoginRequest
	var user model.User
	var token string

	err := context.BindJSON(&request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error": err,
			"token": "",
		})
		return
	}
 
	err = service.CheckRequestValue(request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error": err,
			"token": "",
		})
		return
	}

	ctrl.DB..Table("users").Find(&user, "id=?", request.ID)
	if user.ID == "" {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error": "No users matched user id",
			"token": "",
		})
		return
	}
	if user.PASSWORD != request.PASSWORD {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error": "password is incorrect",
			"token": "",
		})
		return
	}

	token, err = CreateUserToken(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error": err,
			"token": "",
		})
		return 
	}

	context.JSON(http.StatusOK, gin.H{
		"status": true,
		"error": "",
		"token": token,
	})
}