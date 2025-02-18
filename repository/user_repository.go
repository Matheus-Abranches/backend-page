package repository

import (
	"database/sql"
	"fmt"
	"my-project/model"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return UserRepository{
		connection: connection,
	}
}

func (ur *UserRepository) GetUsers() ([]model.User, error) {

	query := "SELECT id, email, full_name FROM users"
	rows, err := ur.connection.Query(query)

	if err != nil {
		fmt.Println("Erro ao procurar usuários", err)
		return []model.User{}, err
	}

	var usersList []model.User
	var userObj model.User

	for rows.Next() {
		err = rows.Scan(
			&userObj.ID,
			&userObj.FullName,
			&userObj.Email,
		)

		if err != nil {
			fmt.Println("Erro ao procurar usuários", err)
			return []model.User{}, err
		}

		usersList = append(usersList, userObj)
	}
	rows.Close()
	return usersList, nil
}

func (ur *UserRepository) CreateUser(user model.User) error {
	query := "INSERT INTO users (full_name, email, phone, agreed_to_terms, agreed_to_receives_news, created_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"

	var id string
	err := ur.connection.QueryRow(query, user.FullName, user.Email, user.Phone, user.AgreeTerms, user.AgreeReceivesNews, user.CreatAt, user.DeleteAt).Scan(&id)
	if err != nil {
		fmt.Println("Erro ao inserir usuário:", err)
		return err
	}

	fmt.Println("Usuário inserido com sucesso, ID gerado:", id)
	return nil
}

func (ur *UserRepository) UpdateUser(id uint, user model.User) error {
	query := "UPDATE users SET full_name = $1, email = $2, phone = $3, agreed_to_terms = $4, agreed_to_receive_news = $5, created_at = $6, deleted_at = $7 WHERE id = $8 RETURNING id"

	var updatedID uint
	err := ur.connection.QueryRow(query, user.FullName, user.Email, user.Phone, user.AgreeTerms, user.AgreeReceivesNews, user.CreatAt, user.DeleteAt, id).Scan(&updatedID)
	if err != nil {
		fmt.Println("Erro ao atualizar usuário:", err)
		return err
	}

	fmt.Println("Usuário atualizado com sucesso, ID:", updatedID)
	return nil
}

func (ur *UserRepository) DeleteUser(id string) error {
	checkQuery := "SELECT id FROM users WHERE id = $1"
	var existingID string
	err := ur.connection.QueryRow(checkQuery, id).Scan(&existingID)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Usuário não encontrado para deletar:", id)
			return nil
		}
		fmt.Println("Erro ao verificar usuário:", err)
		return err
	}

	deleteQuery := "DELETE FROM users WHERE id = $1"
	_, err = ur.connection.Exec(deleteQuery, id)
	if err != nil {
		fmt.Println("Erro ao deletar usuário:", err)
		return err
	}

	fmt.Println("Usuário deletado com sucesso, ID:", existingID)
	return nil
}
