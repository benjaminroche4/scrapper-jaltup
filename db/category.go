package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"scrapperjaltup/model"
	"strconv"
)

func (thiz *DB) CountCategories() (int, error) {
	count := 0
	ctx, cancel := context.WithTimeout(context.Background(), thiz.timeout)
	defer cancel()
	row := thiz.db.QueryRowContext(ctx, "SELECT count(*) from `category`")

	if err := row.Scan(&count); err != nil {
		return 0, fmt.Errorf("DB(CountCategories): %w", err)
	}
	if row.Err() != nil {
		return 0, fmt.Errorf("DB(CountCategories): %w", row.Err())
	}
	return count, nil
}

func (thiz *DB) SelectCategories() ([]model.Category, error) {
	categories := []model.Category{}

	ctx, cancel := context.WithTimeout(context.Background(), thiz.timeout)
	defer cancel()
	rows, err := thiz.db.QueryContext(ctx, "SELECT * FROM `category`")
	if err != nil {
		return categories, fmt.Errorf("DB(SelectCategories): %w", err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return categories, fmt.Errorf("DB(SelectCategories): %w", err)
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Printf("DB(SelectCategories): err = %v", err)
			continue
		}

		category := model.Category{}
		for i, value := range values {
			if value != nil {
				switch columns[i] {
				case "id":
					category.ID, _ = strconv.ParseInt(string(value), 10, 64)

				case "public_id":
					category.PublicID = string(value)

				case "name":
					category.Name = string(value)

				case "slug":
					category.Slug = string(value)
				}
			}
		}
		categories = append(categories, category)
	}
	if err = rows.Err(); err != nil {
		return categories, fmt.Errorf("DB(SelectCategories): %w", err)
	}

	return categories, nil
}

func (thiz *DB) InsertCategories(categories []model.Category) error {
	stmt :=
		"INSERT INTO `category`" +
			"(public_id,name,slug) " +
			" VALUES (?, ?, ?)"

	ctx, cancel := context.WithTimeout(context.Background(), thiz.timeout)
	defer cancel()
	tx, err := thiz.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("DB(InsertCategories): %w", err)
	}

	for i := range categories {
		category := categories[i]

		args := []interface{}{
			category.PublicID,
			category.Name,
			category.Slug,
		}
		res, err := tx.Exec(stmt, args...)
		if err != nil {
			log.Printf("DB(InsertCategories): err = %v", err)
			continue
		}

		id, err := res.LastInsertId()
		if err == nil {
			categories[i].ID = id
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("DB(InsertCategories): %w", err)
	}

	return nil
}

func (thiz *DB) CleanCategories() error {
	ctx, cancel := context.WithTimeout(context.Background(), thiz.timeout)
	defer cancel()

	_, err := thiz.db.ExecContext(ctx, "DELETE FROM `category`")
	if err != nil {
		return fmt.Errorf("DB(CleanCategories): %w", err)
	}

	return nil
}
