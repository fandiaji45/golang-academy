package todo_list_category_usecase

//	"github.com/rahmatrdn/go-skeleton/internal/repository/mysql"
import (
	"github.com/fandiaji45/golang-academy/code-with-skeleton/go-skeleton/internal/repository/mysql"
)

type CrudTodoListCategoryUsecase struct {
	todoListRepo mysql.ITodoListCategoryRepository
}

func NewCrudTodoListCategoryUsecase(
	todoListRepo mysql.ITodoListCategoryRepository,
) *CrudTodoListCategoryUsecase {
	return &CrudTodoListCategoryUsecase{todoListCategoryRepo}
}
