package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/edubarbieri/julius/entity"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	insertNfeSql string = `
		INSERT INTO nfe	
			(access_key, url, issue_date, store_name, store_cnpj, total, discount)	
			VALUES($1, $2, $3, $4, $5, $6, $7);
	`
	insertNfeItemSql string = `
		INSERT INTO nfe_item
			(access_key, description, quantity, unit_measure, unit_price, total_price)
			VALUES($1, $2, $3, $4, $5, $6);
	`
)

type PgNfeRepository struct {
	dbpool *pgxpool.Pool
}

func NewPgNfeRepository(databaseUrl string) (PgNfeRepository, error) {
	dbpool, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		return PgNfeRepository{}, err
	}
	return PgNfeRepository{
		dbpool: dbpool,
	}, nil
}

func (p PgNfeRepository) ExistByAccessKey(ctx context.Context, accessKey string) (bool, error) {
	var count int64
	error := p.dbpool.QueryRow(ctx, "select count(1) from nfe where access_key = $1", accessKey).Scan(&count)
	return count > 0, error
}

func (p PgNfeRepository) Save(ctx context.Context, nfe entity.Nfe) (entity.Nfe, error) {

	tx, err := p.dbpool.Begin(ctx)
	if err != nil {
		return entity.Nfe{}, err
	}
	_, err = tx.Exec(ctx, insertNfeSql,
		nfe.AccessKey, nfe.Url, nfe.Date, nfe.StoreName, nfe.StoreCnpj, nfe.Total, nfe.Discount)
	if err != nil {
		return entity.Nfe{}, err
	}
	for _, item := range nfe.Items {
		_, err = tx.Exec(ctx, insertNfeItemSql,
			nfe.AccessKey, item.Description, item.Quantity, item.UnitOfMeasure, item.UnityPrice, item.TotalPrice)
		if err != nil {
			return entity.Nfe{}, err
		}
	}
	err = tx.Commit(ctx)
	return entity.Nfe{}, err
}
