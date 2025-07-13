package mysql

import (
	"context"

	"github.com/rahmatrdn/go-skeleton/config"
	"github.com/rahmatrdn/go-skeleton/internal/helper"
	"github.com/rahmatrdn/go-skeleton/internal/repository/mysql/entity"

	"github.com/fandiaji45/golang-academy/code-with-skeleton/go-skeleton/internal/repository/mysql/entity"

	apperr "github.com/rahmatrdn/go-skeleton/error"

	errwrap "github.com/pkg/errors"
	"gorm.io/gorm"
)

type ITodoListCategoryRepository interface {
	TrxSupportRepo
	GetByUserID(ctx context.Context, ID int64) (result []*entity.TodoListCategory, err error)
	GetByID(ctx context.Context, ID int64) (result *entity.TodoListCategory, err error)
	Create(ctx context.Context, dbTrx TrxObj, params *entity.TodoListCategory, nonZeroVal bool) error
	LockByID(ctx context.Context, dbTrx TrxObj, ID int64) (result *entity.TodoListCategory, err error)
	Update(ctx context.Context, dbTrx TrxObj, params *entity.TodoListCategory, changes *entity.TodoListCategory) (err error)
	DeleteByID(ctx context.Context, dbTrx TrxObj, id int64) error
}

type TodoListCategoryRepository struct {
	GormTrxSupport
}

func NewTodoListCategoryRepository(mysql *config.Mysql) *TodoListCategoryRepository {
	return &TodoListCategoryRepository{GormTrxSupport{db: mysql.DB}}
}

func (r *TodoListCategoryRepository) GetCreatedBy(ctx context.Context, CreatedBy int64) (result []*entity.TodoListCategory, err error) {
	funcName := "TodoListCategoryRepository.GetCreatedBy"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	err = r.db.Raw("SELECT * FROM todo_list_categories WHERE created_by = ?", CreatedBy).Scan(&result).Error
	if errwrap.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperr.ErrRecordNotFound()
	}

	return result, err
}

func (r *TodoListCategoryRepository) GetByID(ctx context.Context, ID int64) (result *entity.TodoListCategory, err error) {
	funcName := "TodoListCategoryRepository.GetByID"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	err = r.db.Raw("SELECT * FROM todo_list_categories WHERE id = ? LIMIT 1", ID).Scan(&result).Error
	if errwrap.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperr.ErrRecordNotFound()
	}

	return result, err
}

func (r *TodoListCategoryRepository) Create(ctx context.Context, dbTrx TrxObj, params *entity.TodoListCategory, nonZeroVal bool) error {
	funcName := "TodoListCategoryRepository.Create"

	if err := helper.CheckDeadline(ctx); err != nil {
		return errwrap.Wrap(err, funcName)
	}

	cols := helper.NonZeroCols(params, nonZeroVal)
	return r.Trx(dbTrx).Select(cols).Create(&params).Error
}

func (r *TodoListCategoryRepository) LockByID(ctx context.Context, dbTrx TrxObj, ID int64) (result *entity.TodoListCategory, err error) {
	funcName := "TodoListCategoryRepository.LockByID"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	err = r.Trx(dbTrx).
		Raw("SELECT * FROM todo_list_categories WHERE id = ? FOR UPDATE", ID).
		Scan(&result).Error

	if errwrap.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperr.ErrRecordNotFound()
	}

	return result, err
}

func (r *TodoListCategoryRepository) Update(ctx context.Context, dbTrx TrxObj, params *entity.TodoListCategory, changes *entity.TodoListCategory) (err error) {
	funcName := "TodoListCategoryRepository.Update"

	if err := helper.CheckDeadline(ctx); err != nil {
		return errwrap.Wrap(err, funcName)
	}

	db := r.Trx(dbTrx).Model(params)
	if changes != nil {
		err = db.Updates(*changes).Error
	} else {
		err = db.Updates(helper.StructToMap(params, false)).Error
	}

	if err != nil {
		return errwrap.Wrap(err, funcName)
	}

	return nil
}

func (r *TodoListCategoryRepository) DeleteByID(ctx context.Context, dbTrx TrxObj, id int64) error {
	funcName := "TodoListCategoryRepository.DeleteByID"

	if err := helper.CheckDeadline(ctx); err != nil {
		return errwrap.Wrap(err, funcName)
	}

	err := r.Trx(dbTrx).Where("id = ?", id).Delete(&entity.TodoList{}).Error
	if err != nil {
		return err
	}

	return nil
}
