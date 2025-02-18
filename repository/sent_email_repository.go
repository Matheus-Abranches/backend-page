package repository

import (
	"database/sql"
	"fmt"
	"my-project/model"
)

type SentEmailRepository struct {
	connection *sql.DB
}

func NewSentEmailRepository(connection *sql.DB) SentEmailRepository {
	return SentEmailRepository{
		connection: connection,
	}
}

func (sr *SentEmailRepository) GetSentEmails() ([]model.SentEmail, error) {

	query := "SELECT date, version, user_id FROM SentEmail"
	rows, err := sr.connection.Query(query)

	if err != nil {
		fmt.Println("Erro ao procurar emails", err)
		return []model.SentEmail{}, err
	}

	var sentEmailList []model.SentEmail
	var sentEmailObj model.SentEmail

	var emailsList []model.SentEmail

	for rows.Next() {
		var emailObj model.SentEmail

		err = rows.Scan(
			&emailObj.Date,
			&emailObj.Version,
			&emailObj.UserID,
		)

		if err != nil {
			fmt.Println("Erro ao buscar emails enviados:", err)
			return []model.SentEmail{}, err
		}

		emailsList = append(sentEmailList, sentEmailObj)
	}

	rows.Close()
	return emailsList, nil
}

func (sr *SentEmailRepository) CreateSentEmail(email model.SentEmail) error {
	query := "INSERT INTO SentEmail (date, version, product, user_id) VALUES ($1, $2, $3, $4) RETURNING id"

	var id string
	err := sr.connection.QueryRow(query, email.Date, email.Version, email.UserID).Scan(&id)
	if err != nil {
		fmt.Println("Erro ao inserir email enviado:", err)
		return err
	}
	fmt.Println("Email inserido com sucesso, ID gerado:", id)
	return nil
}

func (sr *SentEmailRepository) UpdateSentEmail(id string, email model.SentEmail) error {
	query := "UPDATE SentEmail SET date = $1, version = $2, user_id = $3 WHERE id = $4 RETURNING id"

	var updatedID string
	err := sr.connection.QueryRow(query, email.Date, email.Version, email.UserID, id).Scan(&updatedID)
	if err != nil {
		fmt.Println("Erro ao atualizar email enviado:", err)
		return err
	}

	fmt.Println("Email atualizado com sucesso, ID:", updatedID)
	return nil
}

func (sr *SentEmailRepository) DeleteSentEmail(id string) error {
	checkQuery := "SELECT id FROM SentEmail WHERE id = $1"
	var existingID string
	err := sr.connection.QueryRow(checkQuery, id).Scan(&existingID)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Email não encontrado para deletar:", id)
			return nil
		}
		fmt.Println("Erro ao verificar email enviado:", err)
		return err
	}

	deleteQuery := "DELETE FROM SentEmail WHERE id = $1"
	_, err = sr.connection.Exec(deleteQuery, id)
	if err != nil {
		fmt.Println("Erro ao deletar email enviado:", err)
		return err
	}

	fmt.Println("Email deletado com sucesso, ID:", existingID)
	return nil
}
