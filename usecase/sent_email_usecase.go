package usecase

import (
	"my-project/model"
	"my-project/repository"
)

type SentEmailUseCase struct {
	repository repository.SentEmailRepository
}

func NewSentEmailUseCase(repo repository.SentEmailRepository) SentEmailUseCase {
	return SentEmailUseCase{
		repository: repo,
	}
}

func (su *SentEmailUseCase) GetSentEmails() ([]model.SentEmail, error) {
	return su.repository.GetSentEmails()
}

func (su *SentEmailUseCase) CreateSentEmail(email model.SentEmail) error {
	return su.repository.CreateSentEmail(email)
}

func (su *SentEmailUseCase) UpdateSentEmail(id string, email model.SentEmail) error {
	return su.repository.UpdateSentEmail(id, email)
}

func (su *SentEmailUseCase) DeleteSentEmail(id string) error {
	return su.repository.DeleteSentEmail(id)
}
