package repository

import (
	"database/sql"
	"log"

	"api-inventory-go/src/modules/good/model"
)

type goodRepositoryPostgres struct {
	db *sql.DB
}

func NewGoodRepositoryPostgres(db *sql.DB) *goodRepositoryPostgres {
	return &goodRepositoryPostgres{db: db}
}

func IsExist(id string, r *goodRepositoryPostgres) error {
	query := `SELECT id FROM goods WHERE id=$1`

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

func (r *goodRepositoryPostgres) Create(good *model.Good) error {
	query := `
	INSERT INTO goods (id, name, stockmin, stock, createdAt, updatedAt)
	VALUES ($1, $2, $3, $4, $5, $6)`

	statement, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(
		good.ID,
		good.Name,
		good.StockMin,
		good.Stock,
		good.CreatedAt,
		good.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *goodRepositoryPostgres) Update(id string, good *model.Good) error {
	err := IsExist(id, r)
	if err != nil {
		// fmt.Println("woy")
		return err
	} else {
		query := `
		UPDATE goods
		SET name=$1, stockmin=$2, stock=$3, updatedAt=$4
		WHERE id=$5`

		statement, err := r.db.Prepare(query)
		if err != nil {
			return err
		}

		defer statement.Close()

		_, err = statement.Exec(
			good.Name,
			good.StockMin,
			good.Stock,
			good.UpdatedAt,
			id,
		)
		if err != nil {
			return err
		}

		return nil
	}
}

func (r *goodRepositoryPostgres) Delete(id string) error {
	err := IsExist(id, r)
	if err != nil {
		return err
	} else {
		query := `DELETE FROM goods WHERE id=$1`

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

func (r *goodRepositoryPostgres) FindById(id string) (*model.Good, error) {
	query := `SELECT * FROM goods WHERE id=$1`
	var good model.Good

	statement, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	err = statement.QueryRow(id).Scan(
		&good.ID,
		&good.Name,
		&good.StockMin,
		&good.Stock,
		&good.CreatedAt,
		&good.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &good, nil
}

func (r *goodRepositoryPostgres) FindAll() (model.Goods, error) {
	query := `SELECT * FROM goods`
	var goods model.Goods

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var good model.Good

		err = rows.Scan(
			&good.ID,
			&good.Name,
			&good.StockMin,
			&good.Stock,
			&good.CreatedAt,
			&good.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		goods = append(goods, good)
	}

	return goods, nil
}