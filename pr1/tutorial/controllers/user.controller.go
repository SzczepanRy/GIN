package contro

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tutorial.com/api/tutorial/models"
	"tutorial.com/api/tutorial/services"
)

type UserController struct {
	UserService services.UserService
}

func New(userservice services.UserService) UserController {

	return UserController{
		UserService: userservice,
	}

}

func (uc *UserController) createUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uc.UserService.CreateUser(&user)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) getUser(ctx *gin.Context) {

	username := ctx.Param("name")

	user, err := uc.UserService.GetUser(&username)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) getAll(ctx *gin.Context) {
	users, err := uc.UserService.GetAll()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) updateUser(ctx *gin.Context) {

	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uc.UserService.UpdateUser(&user)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) deleteUser(ctx *gin.Context) {

	username := ctx.Param("name")

	err := uc.UserService.DeleteUser(&username)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

}

func (uc *UserController) RedgisterUserRoutes(rg *gin.RouterGroup) {

	userroute := rg.Group("/user")

	userroute.POST("/create", uc.createUser)
	userroute.PATCH("/update", uc.updateUser)
	userroute.GET("/getall", uc.getAll)
	userroute.GET("/get/:name", uc.getUser)
	userroute.DELETE("/delete/:name", uc.deleteUser)

}
