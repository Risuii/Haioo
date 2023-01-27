package cart

import (
	"context"
	"time"

	"github.com/Risuii/helpers/exception"
	"github.com/Risuii/helpers/response"
	"github.com/Risuii/models/filter"
	"github.com/Risuii/models/product"
)

type (
	CartUseCase interface {
		AddItems(ctx context.Context, params product.Product) response.Response
		GetItems(ctx context.Context, params filter.Filter) response.Response
		DeleteItems(ctx context.Context, kodeProduk string) response.Response
	}

	cartUseCaseImpl struct {
		repo CartRepository
	}
)

func NewCartUseCaseImpl(repo CartRepository) CartUseCase {
	return &cartUseCaseImpl{
		repo: repo,
	}
}

func (cu *cartUseCaseImpl) AddItems(ctx context.Context, params product.Product) response.Response {
	data, err := cu.repo.FindByKodeProduk(ctx, params.KodeProduk)
	if err == nil {
		data = product.Product{
			ID:         data.ID,
			Nama:       data.Nama,
			KodeProduk: data.KodeProduk,
			Kuantitas:  data.Kuantitas + params.Kuantitas,
			UpdateAt:   time.Now(),
		}

		err := cu.repo.UpdateKuantitas(ctx, data.ID, data)
		if err != nil {
			return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
		}

		return response.Success(response.StatusOK, data)
	}

	item := product.Product{
		ID:         params.ID,
		Nama:       params.Nama,
		KodeProduk: params.KodeProduk,
		Kuantitas:  params.Kuantitas,
		CreatedAt:  time.Now(),
	}

	ID, err := cu.repo.Add(ctx, item)
	if err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
	}

	item.ID = ID

	return response.Success(response.StatusCreated, item)
}

func (cu *cartUseCaseImpl) GetItems(ctx context.Context, params filter.Filter) response.Response {

	if params.Nama != "" || params.Kuantitas != 0 {
		data, err := cu.repo.FindByFilter(ctx, params)

		if err == exception.ErrNotFound {
			return response.Error(response.StatusNotFound, exception.ErrNotFound)
		}

		if err != nil {
			return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
		}

		return response.Success(response.StatusOK, data)
	}

	data, err := cu.repo.FindAll()

	if err == exception.ErrNotFound {
		return response.Error(response.StatusNotFound, exception.ErrNotFound)
	}

	if err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
	}

	return response.Success(response.StatusOK, data)
}

func (cu *cartUseCaseImpl) DeleteItems(ctx context.Context, kodeProduk string) response.Response {
	user, err := cu.repo.FindByKodeProduk(ctx, kodeProduk)
	if err == exception.ErrNotFound {
		return response.Error(response.StatusNotFound, exception.ErrNotFound)
	}

	if err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
	}

	if err := cu.repo.Delete(ctx, user.ID); err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
	}

	msg := "Success Delete Data"

	return response.Success(response.StatusOK, msg)
}
