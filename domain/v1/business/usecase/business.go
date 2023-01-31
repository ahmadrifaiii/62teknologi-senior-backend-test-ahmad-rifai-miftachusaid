package usecase

import (
	"62tech.co/service/model"

	"62tech.co/service/pkg/thirdparty/ypl"

	"62tech.co/service/config"
	business_model "62tech.co/service/domain/v1/business/model"
	business_repo "62tech.co/service/domain/v1/business/repo"
)

func BusinessInitialize(conf config.Configuration) (result interface{}, err error) {
	var payload *[]business_model.Business
	var bs []business_model.Business
	res, err := ypl.GetData()
	if err != nil {
		return nil, err
	}

	tx, err := conf.MysqlDB.Begin()

	for _, b := range res.Businesses {
		bs = append(bs, business_model.Business{
			ID:           b.ID,
			Alias:        b.Alias,
			Name:         b.Name,
			ImageURL:     b.ImageURL,
			IsClose:      b.IsClose,
			URL:          b.URL,
			ReviewCount:  b.ReviewCount,
			Rating:       b.Rating,
			Price:        b.Price,
			Phone:        b.Phone,
			DisplayPhone: b.DisplayPhone,
			Distance:     b.Distance,
		})
	}

	payload = &bs

	_, err = business_repo.CreateNewBulkBusiness(tx, payload)

	tx.Commit()

	return result, nil
}

func BusinessesList(conf config.Configuration) (result []business_model.Business, err error) {
	db := conf.MysqlDB
	r, err := business_repo.GetBusinessesList(db)
	if err != nil {
		return nil, err
	}

	if r == nil {
		return nil, nil
	}

	return r, nil
}

// get business detail
func BusinessDetail(conf config.Configuration, id int) (resp *business_model.Business, err error) {
	db := conf.MysqlDB
	return business_repo.GetBusinessDetail(db, id)
}

// create new business
func CreateNewBusiness(conf config.Configuration, payload *business_model.Business) (res *model.IsSuccess, err error) {
	tx, err := conf.MysqlDB.Begin()
	if err != nil {
		return nil, err
	}

	_, err = business_repo.CreateNewBusiness(tx, payload)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return &model.IsSuccess{
		IsSuccess: true,
	}, nil
}

// update business
func BusinessUpdate(conf config.Configuration, in *business_model.Business) (resp *business_model.Business, err error) {
	var (
		payload = business_model.Business{}
	)
	tx, err := conf.MysqlDB.Begin()
	if err != nil {
		return resp, err
	}

	_, err = business_repo.UpdateBusiness(tx, &payload)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return resp, nil
}

// delete business
func BusinessDelete(conf config.Configuration, in *business_model.Business) (resp *business_model.Business, err error) {
	var (
		payload = business_model.Business{}
	)
	tx, err := conf.MysqlDB.Begin()
	if err != nil {
		return resp, err
	}

	payload.ID = in.ID
	payload.Deleted = true

	_, err = business_repo.DeleteBusiness(tx, &payload)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return resp, nil
}
