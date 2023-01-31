package repo

import (
	"database/sql"

	"62tech.co/service/model"

	database "62tech.co/service/config/database"

	business_model "62tech.co/service/domain/v1/business/model"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

// get businesses list
func GetBusinessesList(sqlx *sqlx.DB) (resp []business_model.Business, err error) {
	var Model business_model.Business

	// sql builder
	st := sqlbuilder.NewStruct(Model)
	sb := st.SelectFrom(business_model.TableBusiness)

	sqlStatement, args := sb.Build()

	stmt, err := sqlx.Prepare(sqlStatement)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(args...)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(err)
			return nil, err
		}
		return nil, err
	}

	for rows.Next() {
		var py business_model.Business
		if err := rows.Scan(st.Addr(&py)...); err != nil {
			log.Error(err)
			continue
		}

		resp = append(resp, py)
	}

	return
}

// get user detail
func GetBusinessDetail(sqlx *sqlx.DB, id int) (resp *business_model.Business, err error) {
	var Model business_model.Business

	// sql builder
	st := sqlbuilder.NewStruct(Model)
	sb := st.SelectFrom(business_model.TableBusiness)
	sb.Where(
		sb.Equal("id", id),
	)

	sqlStatement, args := sb.Build()

	stmt, err := sqlx.Prepare(sqlStatement)

	if err != nil {
		return resp, err
	}

	row := stmt.QueryRow(args...)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(err)
			return resp, err
		}
		return resp, err
	}

	row.Scan(st.Addr(&resp))

	return
}

// crate new business
func CreateNewBusiness(tx *sql.Tx, p *business_model.Business) (result sql.Result, err error) {
	st := sqlbuilder.NewStruct(business_model.Business{})
	sb := st.InsertIntoForTag(business_model.TableBusiness, "insert", *p)

	sqlStatement, args := sb.Build()

	stmt, err := tx.Prepare(sqlStatement)
	if err != nil {
		return nil, database.Error(err)
	}

	result, err = stmt.Exec(args...)

	err = database.Error(err)

	return
}

// crate new business
func CreateNewBulkBusiness(tx *sql.Tx, p *[]business_model.Business) (result sql.Result, err error) {
	sb := sqlbuilder.NewInsertBuilder()
	sb.InsertIgnoreInto(business_model.TableBusiness)
	sb.Cols("id",
		"alias",
		"name",
		"image_url",
		"is_closed",
		"url",
		"review_count",
		"rating",
		"price",
		"phone",
		"display_phone",
		"distance",
	)

	for _, b := range *p {
		sb.Values(b.ID,
			b.Alias,
			b.Name,
			b.ImageURL,
			b.IsClose,
			b.URL,
			b.ReviewCount,
			b.Rating,
			b.Price,
			b.Phone,
			b.DisplayPhone,
			b.Distance,
		)
	}

	sqlStatement, args := sb.Build()

	stmt, err := tx.Prepare(sqlStatement)
	if err != nil {
		return nil, database.Error(err)
	}

	result, err = stmt.Exec(args...)

	err = database.Error(err)

	return
}

// update business
func UpdateBusiness(tx *sql.Tx, p *business_model.Business) (result sql.Result, err error) {
	st := sqlbuilder.NewStruct(model.UserPayload{})
	sb := st.UpdateForTag(business_model.TableBusiness, "update", *p)

	sb.Where(
		sb.Equal("id", p.ID),
	)

	sqlStatement, args := sb.Build()

	stmt, err := tx.Prepare(sqlStatement)
	if err != nil {
		return nil, database.Error(err)
	}

	result, err = stmt.Exec(args)

	err = database.Error(err)

	return
}

// delete business
func DeleteBusiness(tx *sql.Tx, p *business_model.Business) (result sql.Result, err error) {
	st := sqlbuilder.NewStruct(business_model.Business{})
	sb := st.UpdateForTag(business_model.TableBusiness, "delete", *p)
	sb.Where(
		sb.Equal("id", p.ID),
	)

	sqlStatement, args := sb.Build()

	stmt, err := tx.Prepare(sqlStatement)
	if err != nil {
		return nil, database.Error(err)
	}

	result, err = stmt.Exec(args)

	err = database.Error(err)

	return
}
