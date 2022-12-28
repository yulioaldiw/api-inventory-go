package repository

import (
	"database/sql"
	"log"

	"api-inventory-go/src/modules/user/model"
)

type userRepositoryPostgres struct {
	db *sql.DB
}

func NewUserRepositoryPostgres(db *sql.DB) *userRepositoryPostgres {
	return &userRepositoryPostgres{db: db}
}

func IsExist(id string, r *userRepositoryPostgres) error {
	query := `SELECT id FROM users WHERE id=$1`

	statement, err := r.db.Prepare(query)
	if err != nil {
		log.Print(err)
		return err
	}

	defer statement.Close()

	var resultId string
	err = statement.QueryRow(id).Scan(&resultId)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepositoryPostgres) Create(user *model.User) error {
	query := `
		INSERT INTO users (id, fullname, nickname, username, nohp, email, password, image, office, role, createdAt, updatedAt)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(
		user.ID,
		user.Fullname,
		user.Nickname,
		user.Username,
		user.NoHP,
		user.Email,
		user.Password,
		user.Image,
		user.Office,
		user.Role,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepositoryPostgres) Update(id string, user *model.User) error {
	err := IsExist(id, r)
	if err != nil {
		return err
	} else {
		query := `
		UPDATE users
		SET fullname=$1, nickname=$2, username=$3, nohp=$4, email=$5, password=$6, image=$7, office=$8, role=$9, updatedAt=$10
		WHERE id=$11`

		statement, err := r.db.Prepare(query)
		if err != nil {
			return err
		}

		defer statement.Close()

		_, err = statement.Exec(
			user.Fullname,
			user.Nickname,
			user.Username,
			user.NoHP,
			user.Email,
			user.Password,
			user.Image,
			user.Office,
			user.Role,
			user.UpdatedAt,
			id,
		)
		if err != nil {
			return err
		}

		return nil
	}
}

func (r *userRepositoryPostgres) Delete(id string) error {
	err := IsExist(id, r)
	if err != nil {
		return err
	} else {
		query := `DELETE FROM users WHERE id=$1`

		statement, err := r.db.Prepare(query)
		if err != nil {
			return err
		}

		defer statement.Close()

		_, err = statement.Exec(id)
		if err != nil {
			return err
		}

		return nil
	}
}

func (r *userRepositoryPostgres) FindById(id string) (*model.User, error) {
	query := `SELECT * FROM users WHERE id=$1`
	var user model.User

	statement, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	err = statement.QueryRow(id).Scan(
		&user.ID,
		&user.Fullname,
		&user.Nickname,
		&user.Username,
		&user.NoHP,
		&user.Email,
		&user.Password,
		&user.Image,
		&user.Office,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepositoryPostgres) FindAll() (model.Users, error) {
	query := `SELECT * FROM users`
	var users model.Users

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user model.User

		err = rows.Scan(
			&user.ID,
			&user.Fullname,
			&user.Nickname,
			&user.Username,
			&user.NoHP,
			&user.Email,
			&user.Password,
			&user.Image,
			&user.Office,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
