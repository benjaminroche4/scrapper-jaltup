package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"scrapperjaltup/model"
	"strconv"
	"time"
)

func (thiz *DB) CountCompanies() (int, error) {
	count := 0
	ctx, cancel := context.WithTimeout(context.Background(), thiz.timeout)
	defer cancel()

	row := thiz.db.QueryRowContext(ctx, "SELECT count(*) from `company`")

	if err := row.Scan(&count); err != nil {
		return 0, fmt.Errorf("DB(CountCompanies): %w", err)
	}
	if row.Err() != nil {
		return 0, fmt.Errorf("DB(CountCompanies): %w", row.Err())
	}
	return count, nil
}

func (thiz *DB) SelectCompanies() ([]model.Company, error) {
	companies := []model.Company{}

	ctx, cancel := context.WithTimeout(context.Background(), thiz.timeout)
	defer cancel()

	rows, err := thiz.db.QueryContext(ctx, "SELECT * FROM `company`")
	if err != nil {
		return companies, fmt.Errorf("DB(SelectCompanies): %w", err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return companies, fmt.Errorf("DB(SelectCompanies): %w", err)
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Printf("DB(SelectCompanies): err = %v", err)
			continue
		}

		company := model.Company{}
		for i, value := range values {
			if value != nil {
				switch columns[i] {
				case "id":
					company.ID, _ = strconv.ParseInt(string(value), 10, 64)

				case "public_id":
					company.PublicID = string(value)

				case "name":
					company.Name = string(value)

				case "siret":
					company.Siret = string(value)

				case "contact_email":
					company.ContactEmail = string(value)

				case "phone_number":
					company.PhoneNumber = string(value)

				case "website_url":
					company.WebSiteURL = string(value)

				case "logo":
					company.Logo = string(value)

				case "created_at":
					company.CreatedAt, _ = time.Parse(time.DateTime, string(value))

				case "slug":
					company.Slug = string(value)

				case "verified":
					company.Verified = string(value) == "1"
				}
			}
		}
		companies = append(companies, company)
	}
	if err = rows.Err(); err != nil {
		return companies, fmt.Errorf("DB(SelectCompanies): %w", err)
	}

	return companies, nil
}

func (thiz *DB) InsertCompanies(companies []model.Company) error {
	stmt :=
		"INSERT INTO `company`" +
			"(public_id,name,siret,contact_email,phone_number,website_url,logo,created_at,slug,verified) " +
			" VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	ctx, cancel := context.WithTimeout(context.Background(), thiz.timeout)
	defer cancel()

	tx, err := thiz.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("DB(InsertCompanies): %w", err)
	}

	for i := range companies {
		company := companies[i]

		verified := int(0)
		if company.Verified {
			verified = int(1)
		}

		args := []interface{}{
			company.PublicID,
			company.Name,
			company.Siret,
			company.ContactEmail,
			company.PhoneNumber,
			company.WebSiteURL,
			company.Logo,
			company.CreatedAt.Format(time.DateTime),
			company.Slug,
			verified,
		}
		res, err := tx.Exec(stmt, args...)
		if err != nil {
			log.Printf("DB(InsertCompanies): err = %v", err)
			continue
		}

		id, err := res.LastInsertId()
		if err == nil {
			companies[i].ID = id
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("DB(InsertCompanies): %w", err)
	}

	return nil
}

func (thiz *DB) CleanCompanies() error {
	ctx, cancel := context.WithTimeout(context.Background(), thiz.timeout)
	defer cancel()

	_, err := thiz.db.ExecContext(ctx, "DELETE FROM `company`")
	if err != nil {
		return fmt.Errorf("DB(CleanCompanies): %w", err)
	}

	return nil
}
