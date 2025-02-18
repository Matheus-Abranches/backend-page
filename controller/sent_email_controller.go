package controller

import (
	"my-project/model"
	"my-project/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SentEmailController struct {
	SentEmailUseCase usecase.SentEmailUseCase
}

func (s *SentEmailController) GetSentEmails(ctx *gin.Context) {
	sentEmail, err := s.SentEmailUseCase.GetSentEmails()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno no servidor"})
		return
	}
	ctx.JSON(http.StatusOK, sentEmail)
}

func (s *SentEmailController) CreateSentEmail(ctx *gin.Context) {
	var newEmail model.SentEmail

	if err := ctx.ShouldBindJSON(&newEmail); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Corpo da requisição inválido"})
		return
	}

	err := s.SentEmailUseCase.CreateSentEmail(newEmail)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao criar o registro de e-mail enviado"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Registro de e-mail enviado criado"})
}

func (s *SentEmailController) UpdateSentEmail(ctx *gin.Context) {
	var updatedEmail model.SentEmail
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&updatedEmail); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Corpo da requisição inválido"})
		return
	}

	err := s.SentEmailUseCase.UpdateSentEmail(id, updatedEmail)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao atualizar o registro de e-mail enviado"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Registro de e-mail enviado atualizado"})
}

func (s *SentEmailController) DeleteSentEmail(ctx *gin.Context) {
	id := ctx.Param("id")

	err := s.SentEmailUseCase.DeleteSentEmail(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao excluir o registro de e-mail enviado"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Registro de e-mail enviado excluído"})
}
