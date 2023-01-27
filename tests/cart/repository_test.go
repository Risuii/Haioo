package cart_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Risuii/helpers/constant"
	"github.com/Risuii/internal/cart"
	"github.com/Risuii/models/filter"
	"github.com/Risuii/models/product"
	"github.com/Risuii/tests/mock"
	"github.com/stretchr/testify/assert"
)

var currentTime = time.Date(2021, 12, 12, 0, 0, 0, 0, &time.Location{})
var productStruct = product.Product{
	ID:         1,
	Nama:       "test",
	KodeProduk: "test",
	Kuantitas:  1,
	CreatedAt:  currentTime,
	UpdateAt:   currentTime,
}

func TestAddRepository(t *testing.T) {
	t.Run("Add Product Success", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := cart.NewCartRepositoryImpl(db, constant.TableCart)

		defer db.Close()

		query := fmt.Sprintf(`INSERT INTO %s`, constant.TableCart)
		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectExec().WithArgs(productStruct.Nama, productStruct.KodeProduk, productStruct.Kuantitas, productStruct.CreatedAt).WillReturnResult(sqlmock.NewResult(1, 1))

		ID, err := repo.Add(ctx, productStruct)

		assert.Equal(t, int64(1), ID)
		assert.NoError(t, err)
	})

	t.Run("Add Product Error", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := cart.NewCartRepositoryImpl(db, constant.TableCart)

		defer db.Close()

		query := fmt.Sprintf(`INSERT INTO %s`, constant.TableCart)
		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectExec().WithArgs(productStruct.Nama, productStruct.KodeProduk, productStruct.Kuantitas, productStruct.CreatedAt).WillReturnResult(sqlmock.NewResult(0, 0))

		ID, err := repo.Add(ctx, productStruct)

		assert.Equal(t, int64(0), ID)
		assert.NoError(t, err)
	})
}

func TestUpdateKuantitasRepository(t *testing.T) {
	t.Run("Update Kuantitas Success", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := cart.NewCartRepositoryImpl(db, constant.TableCart)

		defer db.Close()

		query := fmt.Sprintf(`UPDATE %s SET`, constant.TableCart)

		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectExec().WithArgs(productStruct.Kuantitas).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.UpdateKuantitas(ctx, productStruct.ID, productStruct)

		assert.NoError(t, err)
	})

	t.Run("Update Kuantitas Error", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := cart.NewCartRepositoryImpl(db, constant.TableCart)

		defer db.Close()

		query := fmt.Sprintf(`UPDATE %s SET`, constant.TableCart)

		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectExec().WithArgs(productStruct.Kuantitas).WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.UpdateKuantitas(ctx, productStruct.ID, productStruct)

		assert.Error(t, err)
	})
}

func TestFindByKodeProdukRepository(t *testing.T) {
	t.Run("Find By Kode Produk Success", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := cart.NewCartRepositoryImpl(db, constant.TableCart)

		defer db.Close()

		query := fmt.Sprintf(`SELECT id, nama, kodeProduk, kuantitas, created_at, update_at FROM %s WHERE kodeProduk = ?`, constant.TableCart)
		rows := sqlmock.NewRows([]string{"id", "nama", "kodeProduk", "kuantitas", "created_at", "update_at"}).AddRow(productStruct.ID, productStruct.Nama, productStruct.KodeProduk, productStruct.Kuantitas, productStruct.CreatedAt, productStruct.UpdateAt)

		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectQuery().WithArgs(productStruct.KodeProduk).WillReturnRows(rows)

		productStruct, err := repo.FindByKodeProduk(ctx, productStruct.KodeProduk)

		assert.NotNil(t, productStruct)
		assert.NoError(t, err)
	})

	t.Run("Find By Kode Produk Error", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := cart.NewCartRepositoryImpl(db, constant.TableCart)

		defer db.Close()

		query := fmt.Sprintf(`SELECT id, nama, kodeProduk, kuantitas, created_at, update_at FROM %s WHERE kodeProduk = ?`, constant.TableCart)
		rows := sqlmock.NewRows([]string{"id", "nama", "kodeProduk", "kuantitas", "created_at", "update_at"})

		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectQuery().WithArgs(productStruct.KodeProduk).WillReturnRows(rows)

		productStruct, err := repo.FindByKodeProduk(ctx, productStruct.KodeProduk)

		assert.Empty(t, productStruct)
		assert.Error(t, err)
	})
}

func TestFindAllRepository(t *testing.T) {

	t.Run("Get All Items Success", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := cart.NewCartRepositoryImpl(db, constant.TableCart)

		defer db.Close()

		query := fmt.Sprintf(`SELECT id, nama, kodeProduk, kuantitas, created_at, update_at FROM %s`, constant.TableCart)
		rows := sqlmock.NewRows([]string{"id", "nama", "kodeProduk", "kuantitas", "created_at", "update_at"}).AddRow(productStruct.ID, productStruct.Nama, productStruct.KodeProduk, productStruct.Kuantitas, productStruct.CreatedAt, productStruct.UpdateAt)

		mock.ExpectQuery(query).WillReturnRows(rows)

		productStruct, err := repo.FindAll()

		assert.NotEmpty(t, productStruct)
		assert.NoError(t, err)
	})

	t.Run("Get All Items Error", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := cart.NewCartRepositoryImpl(db, constant.TableCart)

		defer db.Close()

		query := fmt.Sprintf(`SELECT id, nama, kodeProduk, kuantitas, created_at, update_at FROM %s`, constant.TableCart)
		rows := sqlmock.NewRows([]string{"id", "nama", "kodeProduk", "kuantitas", "created_at", "update_at"})

		mock.ExpectQuery(query).WillReturnRows(rows)

		productStruct, err := repo.FindAll()

		assert.Empty(t, productStruct)
		assert.NoError(t, err)
	})
}

func TestDeleteRepository(t *testing.T) {
	t.Run("Delete Items Success", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := cart.NewCartRepositoryImpl(db, constant.TableCart)

		defer db.Close()

		query := fmt.Sprintf(`DELETE FROM %s WHERE id = ?`, constant.TableCart)

		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectExec().WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.Delete(ctx, productStruct.ID)

		assert.NoError(t, err)
	})

	t.Run("Delete Items Error", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := cart.NewCartRepositoryImpl(db, constant.TableCart)

		defer db.Close()

		query := fmt.Sprintf(`DELETE FROM %s WHERE id = ?`, constant.TableCart)

		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectExec().WithArgs().WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.Delete(ctx, productStruct.ID)

		assert.Error(t, err)
	})
}

func TestFindByFilterRepository(t *testing.T) {

	t.Run("Get Item By Filter All Params Not Nil Success", func(t *testing.T) {

		filter := filter.Filter{
			Nama:      "test",
			Kuantitas: 1,
		}

		db, mock := mock.NewMock()
		repo := cart.NewCartRepositoryImpl(db, constant.TableCart)

		defer db.Close()

		query := fmt.Sprintf(`SELECT id, nama, kodeProduk, kuantitas, created_at, update_at FROM %s WHERE nama = '%s' AND kuantitas = '%d'`, constant.TableCart, productStruct.Nama, productStruct.Kuantitas)
		rows := sqlmock.NewRows([]string{"id", "nama", "kodeProduk", "kuantitas", "created_at", "update_at"}).AddRow(productStruct.ID, productStruct.Nama, productStruct.KodeProduk, productStruct.Kuantitas, productStruct.CreatedAt, productStruct.UpdateAt)
		ctx := context.TODO()

		mock.ExpectQuery(query).WillReturnRows(rows)

		productStruct, err := repo.FindByFilter(ctx, filter)

		assert.NotEmpty(t, productStruct)
		assert.NoError(t, err)
	})

	t.Run("Get Item By Filter All Params Not Nil Error", func(t *testing.T) {

		filter := filter.Filter{
			Nama:      "test",
			Kuantitas: 1,
		}

		db, mock := mock.NewMock()
		repo := cart.NewCartRepositoryImpl(db, constant.TableCart)

		defer db.Close()

		query := fmt.Sprintf(`SELECT id, nama, kodeProduk, kuantitas, created_at, update_at FROM %s WHERE nama = '%s' AND kuantitas = '%d'`, constant.TableCart, productStruct.Nama, productStruct.Kuantitas)
		rows := sqlmock.NewRows([]string{"id", "nama", "kodeProduk", "kuantitas", "created_at", "update_at"})
		ctx := context.TODO()

		mock.ExpectQuery(query).WillReturnRows(rows)

		productStruct, err := repo.FindByFilter(ctx, filter)
		fmt.Println("INI DATA", filter)

		assert.Empty(t, productStruct)
		assert.Error(t, err)
	})

	t.Run("Get Item By Filter Just Nama Params Success", func(t *testing.T) {

		filter := filter.Filter{
			Nama:      "test",
			Kuantitas: 0,
		}

		db, mock := mock.NewMock()
		repo := cart.NewCartRepositoryImpl(db, constant.TableCart)

		defer db.Close()

		query := fmt.Sprintf(`SELECT id, nama, kodeProduk, kuantitas, created_at, update_at FROM %s WHERE nama = '%s'`, constant.TableCart, productStruct.Nama)
		rows := sqlmock.NewRows([]string{"id", "nama", "kodeProduk", "kuantitas", "created_at", "update_at"}).AddRow(productStruct.ID, productStruct.Nama, productStruct.KodeProduk, productStruct.Kuantitas, productStruct.CreatedAt, productStruct.UpdateAt)
		ctx := context.TODO()

		mock.ExpectQuery(query).WillReturnRows(rows)

		productStruct, err := repo.FindByFilter(ctx, filter)

		assert.NotEmpty(t, productStruct)
		assert.NoError(t, err)
	})

	t.Run("Get Item By Filter Just Nama Params Error", func(t *testing.T) {
		filter := filter.Filter{
			Nama:      "test",
			Kuantitas: 0,
		}

		db, mock := mock.NewMock()
		repo := cart.NewCartRepositoryImpl(db, constant.TableCart)

		defer db.Close()

		query := fmt.Sprintf(`SELECT id, nama, kodeProduk, kuantitas, created_at, update_at FROM %s WHERE nama = '%s'`, constant.TableCart, productStruct.Nama)
		rows := sqlmock.NewRows([]string{"id", "nama", "kodeProduk", "kuantitas", "created_at", "update_at"})
		ctx := context.TODO()

		mock.ExpectQuery(query).WillReturnRows(rows)

		productStruct, err := repo.FindByFilter(ctx, filter)

		assert.Empty(t, productStruct)
		assert.Error(t, err)
	})

	t.Run("Get Items By Filter Just Kuantitas Params Success", func(t *testing.T) {
		filter := filter.Filter{
			Nama:      "",
			Kuantitas: 1,
		}

		db, mock := mock.NewMock()
		repo := cart.NewCartRepositoryImpl(db, constant.TableCart)

		defer db.Close()

		query := fmt.Sprintf(`SELECT id, nama, kodeProduk, kuantitas, created_at, update_at FROM %s WHERE kuantitas = '%d'`, constant.TableCart, productStruct.Kuantitas)
		rows := sqlmock.NewRows([]string{"id", "nama", "kodeProduk", "kuantitas", "created_at", "update_at"}).AddRow(productStruct.ID, productStruct.Nama, productStruct.KodeProduk, productStruct.Kuantitas, productStruct.CreatedAt, productStruct.UpdateAt)
		ctx := context.TODO()

		mock.ExpectQuery(query).WillReturnRows(rows)

		productStruct, err := repo.FindByFilter(ctx, filter)

		assert.NotEmpty(t, productStruct)
		assert.NoError(t, err)
	})

	t.Run("Get Items By Filter Just Kuantitas Params Error", func(t *testing.T) {
		filter := filter.Filter{
			Nama:      "",
			Kuantitas: 1,
		}

		db, mock := mock.NewMock()
		repo := cart.NewCartRepositoryImpl(db, constant.TableCart)

		defer db.Close()

		query := fmt.Sprintf(`SELECT id, nama, kodeProduk, kuantitas, created_at, update_at FROM %s WHERE kuantitas = '%d'`, constant.TableCart, productStruct.Kuantitas)
		rows := sqlmock.NewRows([]string{"id", "nama", "kodeProduk", "kuantitas", "created_at", "update_at"})
		ctx := context.TODO()

		mock.ExpectQuery(query).WillReturnRows(rows)

		productStruct, err := repo.FindByFilter(ctx, filter)

		assert.Empty(t, productStruct)
		assert.Error(t, err)
	})
}
