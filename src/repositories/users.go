package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

func (repo users) Create(user models.User) error {
	query := `insert into users(name, email, password) values($1, $2, $3)`

	_, err := repo.db.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (repo users) Search(query string) ([]models.User, error) {
	query = fmt.Sprintf("%%%s%%", query)

	rows, err := repo.db.Query("select id, name, email, created_at from users where name ilike $1", query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repo users) GetByID(ID uint64) (models.User, error) {
	rows, err := repo.db.Query(
		"select id, name, email, created_at from users where id = $1", ID,
	)

	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repo users) Update(ID uint64, user models.User) error {
	query := `update users set name=$1, email=$2 where id=$3;`

	_, err := repo.db.Exec(query, user.Name, user.Email, ID)
	if err != nil {
		return err
	}

	return nil
}

func (repo users) Delete(ID uint64) error {
	query := `delete from users where id=$1;`

	_, err := repo.db.Exec(query, ID)
	if err != nil {
		return err
	}

	return nil
}

func (repo users) GetByEmail(email string) (models.User, error) {
	rows, err := repo.db.Query(
		"select id, password from users where email = $1", email,
	)

	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Password,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}
