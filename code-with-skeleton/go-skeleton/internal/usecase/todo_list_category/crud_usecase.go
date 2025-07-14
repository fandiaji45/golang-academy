package todo_list_category_usecase

import (
	"context"
	"fmt"
	"time"

	errwrap "github.com/pkg/errors"
	generalEntity "github.com/rahmatrdn/go-skeleton/entity"
	"github.com/rahmatrdn/go-skeleton/internal/helper"
	"github.com/rahmatrdn/go-skeleton/internal/repository/mysql"
	mentity "github.com/rahmatrdn/go-skeleton/internal/repository/mysql/entity"
	"github.com/rahmatrdn/go-skeleton/internal/usecase"
	"github.com/rahmatrdn/go-skeleton/internal/usecase/todo_list_category/entity"
)

type CrudTodoListCategoryUsecase struct {
	todoListCategoryRepo mysql.ITodoListCategoryRepository
}

func NewCrudTodoListCategoryUsecase(
	todoListCategoryRepo mysql.ITodoListCategoryRepository,
) *CrudTodoListCategoryUsecase {
	return &CrudTodoListCategoryUsecase{todoListCategoryRepo}
}

type ICrudTodoListCategoryUsecase interface {
	GetByUserID(ctx context.Context, userID int64) (res []*entity.TodoListCategoryResponse, err error)
	GetByID(ctx context.Context, todoListCategoryID int64) (*entity.TodoListCategoryResponse, error)
	Create(ctx context.Context, todoListCategoryReq entity.TodoListCategoryReq) (*entity.TodoListCategoryResponse, error)
	UpdateByID(ctx context.Context, todoListCategoryReq entity.TodoListCategoryReq) error
	DeleteByID(ctx context.Context, todoListCategoryID int64) error
}

func (t *CrudTodoListCategoryUsecase) GetByUserID(ctx context.Context, userID int64) (res []*entity.TodoListCategoryResponse, err error) {
	funcName := "CrudTodoListUsecase.GetByUserID"
	captureFieldError := generalEntity.CaptureFields{
		"user_id": helper.ToString(userID),
	}

	result, err := t.todoListCategoryRepo.GetByUserID(ctx, userID)
	if err != nil {
		helper.LogError("todoListCategoryRepo.GetByUserID", funcName, err, captureFieldError, "")

		return nil, err
	}

	for _, v := range result {
		res = append(res, &entity.TodoListCategoryResponse{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			CreatedAt:   helper.ConvertToJakartaTime(v.CreatedAt),
			CreatedBy:   v.CreatedBy,
		})
	}

	return res, nil
}

func (t *CrudTodoListCategoryUsecase) GetByID(ctx context.Context, todoListID int64) (*entity.TodoListCategoryResponse, error) {
	funcName := "CrudTodoListUsecase.GetByID"
	captureFieldError := generalEntity.CaptureFields{
		"user_id": helper.ToString(todoListID),
	}

	data, err := t.todoListCategoryRepo.GetByID(ctx, todoListID)
	if err != nil {
		helper.LogError("todoListCategoryRepo.GetByID", funcName, err, captureFieldError, "")

		return nil, err
	}
	if data == nil {
		return nil, nil
	}

	return &entity.TodoListCategoryResponse{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		CreatedAt:   helper.ConvertToJakartaTime(data.CreatedAt),
		CreatedBy:   data.CreatedBy,
	}, nil
}

func (t *CrudTodoListCategoryUsecase) Create(ctx context.Context, todoListCategoryReq entity.TodoListCategoryReq) (*entity.TodoListCategoryResponse, error) {
	funcName := "CrudTodoListUsecase.Create"
	captureFieldError := generalEntity.CaptureFields{
		"created_by": helper.ToString(todoListCategoryReq.CreatedBy),
		"payload":    helper.ToString(todoListCategoryReq),
	}

	if errMsg := usecase.ValidateStruct(todoListCategoryReq); errMsg != "" {
		return nil, errwrap.Wrap(fmt.Errorf(generalEntity.INVALID_PAYLOAD_CODE), errMsg)
	}

	//doingAt, _ := helper.ParseDate(todoListCategoryReq.DoingAt)

	todoListCategoryPayload := &mentity.TodoListCategory{
		CreatedBy:   todoListCategoryReq.CreatedBy,
		Name:        todoListCategoryReq.Name,
		Description: todoListCategoryReq.Description,
		CreatedAt:   time.Now(),
	}

	err := t.todoListCategoryRepo.Create(ctx, nil, todoListCategoryPayload, false)
	if err != nil {
		helper.LogError("todoListCategoryRepo.Create", funcName, err, captureFieldError, "")

		return nil, err
	}

	return &entity.TodoListCategoryResponse{
		ID:          todoListCategoryPayload.ID,
		Name:        todoListCategoryPayload.Name,
		Description: todoListCategoryPayload.Description,
		CreatedAt:   helper.ConvertToJakartaTime(todoListCategoryPayload.CreatedAt),
		CreatedBy:   todoListCategoryPayload.CreatedBy,
	}, nil
}

func (t *CrudTodoListCategoryUsecase) UpdateByID(ctx context.Context, todoListCategoryReq entity.TodoListCategoryReq) error {
	funcName := "CrudTodoListCategoryUsecase.UpdateByID"
	todoListCategoryID := todoListCategoryReq.ID

	captureFieldError := generalEntity.CaptureFields{
		"created_by": helper.ToString(todoListCategoryReq.CreatedBy),
		"payload":    helper.ToString(todoListCategoryReq),
	}

	// Start DB Transaction
	if err := mysql.DBTransaction(t.todoListCategoryRepo, func(trx mysql.TrxObj) error {
		// Locking Data
		lockedData, err := t.todoListCategoryRepo.LockByID(ctx, trx, todoListCategoryID)
		if err != nil {
			helper.LogError("todoListCategoryRepo.LockByID", funcName, err, captureFieldError, "")

			return err
		}
		if lockedData == nil {
			return fmt.Errorf("DATA IS NOT EXIST")
		}

		// Process Update
		//doingAt, _ := helper.ParseDate(todoListReq.DoingAt)
		if err := t.todoListCategoryRepo.Update(ctx, trx, lockedData, &mentity.TodoListCategory{
			Name:        todoListCategoryReq.Name,
			Description: todoListCategoryReq.Description,
		}); err != nil {
			helper.LogError("todoListCategoryRepo.Update", funcName, err, captureFieldError, "")

			return err
		}

		return nil
	}); err != nil {
		helper.LogError("todoListCategoryRepo.DBTransaction", funcName, err, captureFieldError, "")

		return err
	}

	return nil
}

func (t *CrudTodoListCategoryUsecase) DeleteByID(ctx context.Context, todoListCategoryID int64) error {
	funcName := "CrudTodoListUsecase.DeleteByID"
	captureFieldError := generalEntity.CaptureFields{
		"todo_list_Category_id": helper.ToString(todoListCategoryID),
	}

	err := t.todoListCategoryRepo.DeleteByID(ctx, nil, todoListCategoryID)
	if err != nil {
		helper.LogError("todoListCategoryRepo.DeleteByID", funcName, err, captureFieldError, "")

		return err
	}

	return nil
}
