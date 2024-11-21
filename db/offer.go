package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"scrapperjaltup/model"
	"strconv"
	"time"
)

func (thiz *DB) CountOffers() (int, error) {
	count := 0
	ctx, cancel := context.WithTimeout(context.Background(), thiz.timeout)
	defer cancel()

	row := thiz.db.QueryRowContext(ctx, "SELECT count(*) from `offer`")

	if err := row.Scan(&count); err != nil {
		return 0, fmt.Errorf("DB(CountOffers): %w", err)
	}
	if row.Err() != nil {
		return 0, fmt.Errorf("DB(CountOffers): %w", row.Err())
	}
	return count, nil
}

func (thiz *DB) SelectOffers() ([]model.Offer, error) {
	offers := []model.Offer{}
	companiesMap, _ := thiz.getCompanies()

	ctx, cancel := context.WithTimeout(context.Background(), thiz.timeout)
	defer cancel()

	rows, err := thiz.db.QueryContext(ctx, "SELECT * FROM `offer`")
	if err != nil {
		return offers, fmt.Errorf("DB(SelectOffers): %w", err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return offers, fmt.Errorf("DB(SelectOffers): %w", err)
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Printf("DB(SelectOffers): err = %v", err)
			continue
		}

		offer := model.Offer{}
		for i, value := range values {
			if value != nil {
				var (
					place model.Place
					job   model.Job
					tag   []string
				)

				switch columns[i] {
				case "id":
					offer.ID, _ = strconv.ParseInt(string(value), 10, 64)

				case "company_id":
					companyID, _ := strconv.ParseInt(string(value), 10, 64)
					company, ok := companiesMap[companyID]
					if ok {
						offer.Company = company
					} else {
						offer.Company = model.Company{ID: companyID}
					}

				case "public_id":
					offer.PublicID = string(value)

				case "title":
					offer.Title = string(value)

				case "place":
					if err := json.Unmarshal(value, &place); err == nil {
						offer.Place = place
					}

				case "job":
					if err := json.Unmarshal(value, &job); err == nil {
						offer.Job = job
					}

				case "url":
					offer.URL = string(value)

				case "tag":
					if err := json.Unmarshal(value, &tag); err == nil {
						offer.Tag = tag
					}

				case "status":
					offer.Status = string(value)

				case "created_at":
					offer.CreatedAt, _ = time.Parse(time.DateTime, string(value))

				case "end_date":
					offer.EndAt, _ = time.Parse("2006-01-02", string(value))

				case "slug":
					offer.Slug = string(value)

				case "premium":
					offer.Premium = string(value) == "1"

				case "external_id":
					offer.ExternalID = string(value)

				case "service_name":
					offer.ServiceName = string(value)
				}
			}
		}
		err = thiz.getCategories(&offer)
		if err != nil {
			log.Printf("DB(SelectOffers): err = %v", err)
		}
		offers = append(offers, offer)
	}
	if err = rows.Err(); err != nil {
		return offers, fmt.Errorf("DB(SelectOffers): %w", err)
	}

	return offers, nil
}

func (thiz *DB) InsertOffers(offers []model.Offer) error {
	stmt :=
		"INSERT INTO `offer`(" +
			"company_id," +
			"public_id," +
			"title," +
			"place," +
			"job," +
			"url," +
			"tag," +
			"status," +
			"created_at," +
			"end_date," +
			"slug," +
			"premium," +
			"external_id," +
			"service_name" +
			") VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	ctx, cancel := context.WithTimeout(context.Background(), thiz.timeout)
	defer cancel()

	tx, err := thiz.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("DB(InsertOffers): %w", err)
	}

	for i := range offers {
		offer := offers[i]

		var companyID *int64

		place, err := json.Marshal(offer.Place)
		if err != nil {
			log.Printf("DB(InsertOffers): err = %v", err)
			continue
		}
		job, err := json.Marshal(offer.Job)
		if err != nil {
			log.Printf("DB(InsertOffers): err = %v", err)
			continue
		}
		tag, err := json.Marshal(offer.Tag)
		if err != nil {
			log.Printf("DB(InsertOffers): err = %v", err)
			continue
		}
		premium := int(0)
		if offer.Premium {
			premium = int(1)
		}
		if offer.Company.ID != 0 {
			companyID = &offer.Company.ID
		}

		args := []interface{}{
			companyID,
			offer.PublicID,
			offer.Title,
			place,
			job,
			offer.URL,
			tag,
			offer.Status,
			offer.CreatedAt.Format(time.DateTime),
			offer.EndAt.Format("2006-01-02"),
			offer.Slug,
			premium,
			offer.ExternalID,
			offer.ServiceName,
		}
		res, err := tx.Exec(stmt, args...)
		if err != nil {
			log.Printf("DB(InsertOffers): err = %v", err)
			continue
		}

		id, err := res.LastInsertId()
		if err == nil {
			offer.ID = id
			offers[i].ID = id
		}

		thiz.insertCategories(tx, &offer)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("DB(InsertOffers): %w", err)
	}

	return nil
}

func (thiz *DB) CleanOffers() error {
	ctx, cancel := context.WithTimeout(context.Background(), thiz.timeout)
	defer cancel()

	_, err := thiz.db.ExecContext(ctx, "DELETE FROM `offer`")
	if err != nil {
		return fmt.Errorf("DB(CleanOffers): %w", err)
	}

	return nil
}

// Unexported functions

func (thiz *DB) insertCategories(tx *sql.Tx, offer *model.Offer) {
	stmt :=
		"INSERT INTO `offer_category`" +
			"(offer_id,category_id) " +
			" VALUES (?, ?)"

	for i := range offer.Categories {
		category := offer.Categories[i]

		if category.ID == 0 {
			continue
		}

		args := []interface{}{
			offer.ID,
			category.ID,
		}
		_, err := tx.Exec(stmt, args...)
		if err != nil {
			log.Printf("DB(insertCategories): err = %v", err)
			continue
		}
	}
}

func (thiz *DB) getCompanies() (map[int64]model.Company, error) {
	companiesMap := map[int64]model.Company{}

	companies, err := thiz.SelectCompanies()
	if err != nil {
		return companiesMap, fmt.Errorf("DB(getCompanies): %w", err)
	}
	for i := range companies {
		company := companies[i]
		companiesMap[company.ID] = company
	}
	return companiesMap, nil
}

func (thiz *DB) getCategories(offer *model.Offer) error {
	stmt := "SELECT" +
		"  `category`.id," +
		"  `category`.public_id," +
		"  `category`.name," +
		"  `category`.slug" +
		" FROM `category` " +
		" JOIN `offer_category` ON `category`.id=`offer_category`.category_id " +
		" WHERE `offer_category`.offer_id=?"

	ctx, cancel := context.WithTimeout(context.Background(), thiz.timeout)
	defer cancel()
	rows, err := thiz.db.QueryContext(ctx, stmt, offer.ID)
	if err != nil {
		return fmt.Errorf("DB(getCategories): %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		category := model.Category{}

		err = rows.Scan(&category.ID, &category.PublicID, &category.Name, &category.Slug)
		if err != nil {
			continue
		}

		offer.Categories = append(offer.Categories, category)
	}

	if err = rows.Err(); err != nil {
		return fmt.Errorf("DB(getCategories): %w", err)
	}

	return nil
}
