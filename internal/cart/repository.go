package cart

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Risuii/helpers/exception"
	"github.com/Risuii/models/filter"
	"github.com/Risuii/models/product"
)

type (
	CartRepository interface {
		Add(ctx context.Context, params product.Product) (int64, error)
		UpdateKuantitas(ctx context.Context, id int64, params product.Product) error
		FindByKodeProduk(ctx context.Context, kodeProduk string) (product.Product, error)
		FindByFilter(ctx context.Context, params filter.Filter) ([]product.Product, error)
		FindAll() ([]product.Product, error)
		Delete(ctx context.Context, id int64) error
	}

	cartRepositoryImpl struct {
		DB        *sql.DB
		tableName string
	}
)

func NewCartRepositoryImpl(db *sql.DB, tableName string) CartRepository {
	return &cartRepositoryImpl{
		DB:        db,
		tableName: tableName,
	}
}

func (cr *cartRepositoryImpl) Add(ctx context.Context, params product.Product) (int64, error) {
	query := fmt.Sprintf(`INSERT INTO %s (nama, kodeProduk, kuantitas, created_at) VALUES (?,?,?,?)`, cr.tableName)
	stmt, err := cr.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return 0, exception.ErrInternalServer
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
		params.Nama,
		params.KodeProduk,
		params.Kuantitas,
		params.CreatedAt,
	)

	if err != nil {
		log.Println(err)
		return 0, exception.ErrInternalServer
	}

	ID, _ := result.LastInsertId()

	return ID, nil
}

func (cr *cartRepositoryImpl) UpdateKuantitas(ctx context.Context, id int64, params product.Product) error {
	query := fmt.Sprintf(`UPDATE %s SET kuantitas = ? WHERE id = %d`, cr.tableName, id)
	stmt, err := cr.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return exception.ErrInternalServer
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
		params.Kuantitas,
	)

	if err != nil {
		log.Println(err)
		return exception.ErrInternalServer
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected < 1 {
		return exception.ErrNotFound
	}

	return nil
}

func (cr *cartRepositoryImpl) FindByKodeProduk(ctx context.Context, kodeProduk string) (product.Product, error) {
	var product product.Product

	query := fmt.Sprintf(`SELECT id, nama, kodeProduk, kuantitas, created_at, update_at FROM %s WHERE kodeProduk = ?`, cr.tableName)
	stmt, err := cr.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return product, exception.ErrInternalServer
	}

	defer stmt.Close()

	rows := stmt.QueryRowContext(ctx, kodeProduk)

	err = rows.Scan(
		&product.ID,
		&product.Nama,
		&product.KodeProduk,
		&product.Kuantitas,
		&product.CreatedAt,
		&product.UpdateAt,
	)

	if err != nil {
		log.Println(err)
		return product, exception.ErrNotFound
	}

	return product, nil
}

func (cr *cartRepositoryImpl) FindAll() ([]product.Product, error) {
	var products []product.Product

	rows, err := cr.DB.Query(fmt.Sprintf(`SELECT id, nama, kodeProduk, kuantitas, created_at, update_at FROM %s`, cr.tableName))
	if err != nil {
		log.Println(err)
		return products, exception.ErrInternalServer
	}

	defer rows.Close()

	for rows.Next() {
		var c product.Product
		if err := rows.Scan(
			&c.ID,
			&c.Nama,
			&c.KodeProduk,
			&c.Kuantitas,
			&c.CreatedAt,
			&c.UpdateAt,
		); err != nil {
			log.Println(err)
			return products, exception.ErrNotFound
		}
		products = append(products, c)
	}

	return products, nil
}

func (cr *cartRepositoryImpl) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = %d`, cr.tableName, id)
	stmt, err := cr.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return exception.ErrInternalServer
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
	)

	if err != nil {
		log.Println(err)
		return exception.ErrInternalServer
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected < 1 {
		return exception.ErrNotFound
	}

	return nil
}

func (cr *cartRepositoryImpl) FindByFilter(ctx context.Context, params filter.Filter) ([]product.Product, error) {
	var products []product.Product

	if params.Nama != "" && params.Kuantitas != 0 {
		var products []product.Product

		rows, err := cr.DB.Query(fmt.Sprintf(`SELECT id, nama, kodeProduk, kuantitas, created_at, update_at FROM %s WHERE nama = '%s' AND kuantitas = '%d'`, cr.tableName, params.Nama, params.Kuantitas))
		if err != nil {
			log.Println(err)
			return products, exception.ErrInternalServer
		}

		defer rows.Close()

		for rows.Next() {
			var c product.Product
			if err := rows.Scan(
				&c.ID,
				&c.Nama,
				&c.KodeProduk,
				&c.Kuantitas,
				&c.CreatedAt,
				&c.UpdateAt,
			); err != nil {
				log.Println(err)
				return products, exception.ErrNotFound
			}
			products = append(products, c)
		}

		if products == nil {
			return products, exception.ErrNotFound
		}

		return products, nil
	} else if params.Nama != "" && params.Kuantitas == 0 {
		var products []product.Product

		rows, err := cr.DB.Query(fmt.Sprintf(`SELECT id, nama, kodeProduk, kuantitas, created_at, update_at FROM %s WHERE nama = '%s'`, cr.tableName, params.Nama))
		if err != nil {
			log.Println(err)
			return products, exception.ErrInternalServer
		}

		defer rows.Close()

		for rows.Next() {
			var c product.Product
			if err := rows.Scan(
				&c.ID,
				&c.Nama,
				&c.KodeProduk,
				&c.Kuantitas,
				&c.CreatedAt,
				&c.UpdateAt,
			); err != nil {
				log.Println(err)
				return products, exception.ErrNotFound
			}
			products = append(products, c)
		}

		if products == nil {
			return products, exception.ErrNotFound
		}

		return products, nil
	} else if params.Nama == "" && params.Kuantitas != 0 {
		var products []product.Product

		rows, err := cr.DB.Query(fmt.Sprintf(`SELECT id, nama, kodeProduk, kuantitas, created_at, update_at FROM %s WHERE kuantitas = '%d'`, cr.tableName, params.Kuantitas))
		if err != nil {
			log.Println(err)
			return products, exception.ErrInternalServer
		}

		defer rows.Close()

		for rows.Next() {
			var c product.Product
			if err := rows.Scan(
				&c.ID,
				&c.Nama,
				&c.KodeProduk,
				&c.Kuantitas,
				&c.CreatedAt,
				&c.UpdateAt,
			); err != nil {
				log.Println(err)
				return products, exception.ErrNotFound
			}
			products = append(products, c)
		}

		if products == nil {
			return products, exception.ErrNotFound
		}

		return products, nil
	}

	return products, nil
}
