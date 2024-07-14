package postgresutils

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

type PageRequest struct {
	Page uint     `in:"query=page"`
	Size uint     `in:"query=size"`
	Sort []string `in:"query=sort"`
}

func (p PageRequest) ApplyPagination(q *goqu.SelectDataset) *goqu.SelectDataset {
	q = q.Offset(p.Page * p.Size).Limit(p.Size)
	orderExps := make([]exp.OrderedExpression, 0, len(p.Sort))

	for _, sort := range p.Sort {
		orderExps = append(orderExps, parseSort(sort))
	}

	q = q.Order(orderExps...)
	return q
}

func parseSort(sort string) exp.OrderedExpression {
	column, order, ok := strings.Cut(sort, ",")
	if !ok {
		order = "asc"
	}
	if order == "asc" {
		return goqu.C(column).Asc()
	}
	return goqu.C(column).Desc()
}

type Page[T any] struct {
	Content       []T
	Page          uint
	Size          uint
	TotalElements uint
	TotalPages    uint
}

func NewPage[T any](content []T, page, size, totalElements uint) Page[T] {
	return Page[T]{
		Content:       content,
		Page:          page,
		Size:          size,
		TotalElements: totalElements,
		TotalPages:    totalPages(totalElements, size),
	}
}

func FromOtherPage[SrcPage, DstPage any](
	srcPage Page[SrcPage],
	mapFunc func(SrcPage) DstPage,
) Page[DstPage] {
	dstContent := make([]DstPage, 0, len(srcPage.Content))
	for _, src := range srcPage.Content {
		dstContent = append(dstContent, mapFunc(src))
	}
	return NewPage(dstContent, srcPage.Page, srcPage.Size, srcPage.TotalElements)
}

func totalPages(totalElements, size uint) uint {
	if totalElements == 0 {
		return 0
	}
	if size == 0 {
		return 1
	}
	return (totalElements-1)/size + 1
}

func countTotal(ctx context.Context, db *sql.DB, q *goqu.SelectDataset) (uint, error) {
	query, args, err := q.Select(goqu.COUNT("*")).Order().ToSQL()
	if err != nil {
		return 0, err
	}

	var total uint
	err = db.QueryRowContext(ctx, query, args...).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}

type Scanner interface {
	Scan(dest ...interface{}) error
}

type RowMapper[T any] func(rows *sql.Rows) (T, error)

func ListPaginated[T any](
	ctx context.Context,
	db *sql.DB,
	q *goqu.SelectDataset,
	pageReq *PageRequest,
	mapRow RowMapper[T],
) (Page[T], error) {
	total, err := countTotal(ctx, db, q)
	if err != nil {

		println("COUNT ERROR")
		return Page[T]{}, err
	}
	if total == 0 {
		return NewPage[T](nil, pageReq.Page, pageReq.Size, total), nil
	}

	q = pageReq.ApplyPagination(q)

	data, err := listData(ctx, db, q, mapRow)
	if err != nil {
		println("LIST DATA ERROR")
		return Page[T]{}, err
	}

	return NewPage[T](data, pageReq.Page, pageReq.Size, total), nil
}

func listData[T any](
	ctx context.Context,
	db *sql.DB,
	q *goqu.SelectDataset,
	mapRow RowMapper[T],
) (_ []T, err error) {
	query, args, err := q.ToSQL()
	if err != nil {
		println("TO SQL ERROR")
		return nil, err
	}

	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		println("QueryContext ERROR")
		return nil, err
	}

	defer func() {
		err = errors.Join(err, rows.Err(), rows.Close())
	}()

	var result []T
	for rows.Next() {
		row, err := mapRow(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, row)
	}

	return result, nil
}
