package repository

import (
	"context"
	"math"

	"github.com/irfhakeem/go-fiber-clean-starter/dto"
	"gorm.io/gorm"
)

type (
	IBaseRepository[T any] interface {
		Raw(ctx context.Context, tx *gorm.DB, query string) ([]T, error)
		FindAll(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest, query any, args ...any) (PaginationResult[T], error)
		FindByID(ctx context.Context, tx *gorm.DB, id int64) (*T, error)
		FindFirst(ctx context.Context, tx *gorm.DB, query any, args ...any) (*T, error)
		Where(ctx context.Context, tx *gorm.DB, query any, args ...any) ([]T, error)
		WhereExisting(ctx context.Context, tx *gorm.DB, query any, args ...any) (bool, error)
		Create(ctx context.Context, tx *gorm.DB, entity *T) (*T, error)
		Update(ctx context.Context, tx *gorm.DB, entity *T) (*T, error)
		Delete(ctx context.Context, tx *gorm.DB, id int64) error
	}

	BaseRepository[T any] struct {
		db *gorm.DB
	}
)

func NewBaseRepository[T any](db *gorm.DB) IBaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

type PaginationResult[T any] struct {
	Data               []T
	PaginationResponse dto.PaginationResponse
}

func Paginate(page int, perPage int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * perPage
		return db.Offset(offset).Limit(perPage)
	}
}

func (r *BaseRepository[T]) Raw(ctx context.Context, tx *gorm.DB, query string) ([]T, error) {
	if tx == nil {
		tx = r.db
	}

	var entities []T
	if err := tx.WithContext(ctx).Raw(query).Scan(&entities).Error; err != nil {
		return nil, err
	}

	return entities, nil
}

func (r *BaseRepository[T]) FindAll(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest, query any, args ...any) (PaginationResult[T], error) {
	if tx == nil {
		tx = r.db
	}

	var entities []T
	var count int64
	if req.PerPage <= 0 {
		req.PerPage = 10
	}
	if req.Page <= 0 {
		req.Page = 1
	}

	var entity T
	q := tx.WithContext(ctx).Model(&entity)
	q = q.Where(query, args...)

	if err := q.Count(&count).Error; err != nil {
		return PaginationResult[T]{}, err
	}

	if err := q.Scopes(Paginate(req.Page, req.PerPage)).Order("id asc").Find(&entities).Error; err != nil {
		return PaginationResult[T]{}, err
	}

	totalPage := int(math.Max(1, math.Ceil(float64(count)/float64(req.PerPage))))

	result := PaginationResult[T]{
		Data: entities,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			MaxPage: totalPage,
			Count:   count,
		},
	}
	return result, nil
}

func (r *BaseRepository[T]) FindByID(ctx context.Context, tx *gorm.DB, id int64) (*T, error) {
	if tx == nil {
		tx = r.db
	}

	var entity T
	if err := tx.WithContext(ctx).Where("id = ?", id).Take(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *BaseRepository[T]) FindFirst(ctx context.Context, tx *gorm.DB, query any, args ...any) (*T, error) {
	if tx == nil {
		tx = r.db
	}

	var entity T
	if err := tx.WithContext(ctx).Where(query, args...).First(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *BaseRepository[T]) Where(ctx context.Context, tx *gorm.DB, query any, args ...any) ([]T, error) {
	if tx == nil {
		tx = r.db
	}

	var entities []T
	if err := tx.WithContext(ctx).Where(query, args...).Order("id asc").Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *BaseRepository[T]) WhereExisting(ctx context.Context, tx *gorm.DB, query any, args ...any) (bool, error) {
	if tx == nil {
		tx = r.db
	}

	var entity T
	err := tx.WithContext(ctx).Where(query, args...).First(&entity).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *BaseRepository[T]) Create(ctx context.Context, tx *gorm.DB, entity *T) (*T, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(entity).Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (r *BaseRepository[T]) Update(ctx context.Context, tx *gorm.DB, entity *T) (*T, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Updates(&entity).Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (r *BaseRepository[T]) Delete(ctx context.Context, tx *gorm.DB, id int64) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(new(T), "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}
