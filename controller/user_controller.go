package controller

import (
	"my-project/model"
	"my-project/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUseCase usecase.UserUseCase
}

func (u *UserController) GetUsers(ctx *gin.Context) {
	user, err := u.UserUseCase.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno no servidor"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (u *UserController) CreateUser(ctx *gin.Context) {
	var newUser model.User

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Corpo da requisição inválido"})
		return
	}
	err := u.UserUseCase.CreateUser(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao criar o usuário"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Usuário criado com sucesso"})
}

func (u *UserController) UpdateUser(ctx *gin.Context) {
	var updateUser model.User
	idParam := ctx.Param("id")

	idUint, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := ctx.ShouldBindJSON(&updateUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Corpo da requisição inválido"})
		return
	}

	err = u.UserUseCase.UpdateUser(uint(idUint), updateUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao atualizar usuário"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Usuário atualizado com sucesso"})
}

func (u UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	err := u.UserUseCase.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao excluir o registro do usuário"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Registro do usuário enviado excluído"})
}
