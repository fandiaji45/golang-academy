package handler

import (
	"net/http"

	"github.com/fandiaji45/golang-academy/code-with-skeleton/go-skeleton/internal/usecase/todo_list_category/entity"
	todo_list_category_usecase "github.com/fandiaji45/golang-academy/code-with-skeleton/go-skeleton/internal/usecase/todo_list_handler_category"
	"github.com/rahmatrdn/go-skeleton/internal/http/middleware"
	"github.com/rahmatrdn/go-skeleton/internal/parser"
	"github.com/rahmatrdn/go-skeleton/internal/presenter/json"

	fiber "github.com/gofiber/fiber/v2"
)

type TodoListCategoryHandler struct {
	parser                      parser.Parser
	presenter                   json.JsonPresenter
	todoListCategoryCrudUsecase todo_list_category_usecase.ICrudTodoListCategoryUsecase
}

func NewTodoListHandler(
	parser parser.Parser,
	presenter json.JsonPresenter,
	todoListCategoryCrudUsecase todo_list_category_usecase.ICrudTodoListCategoryUsecase,
) *TodoListHandler {
	return &TodoListHandler{parser, presenter, todoListCategoryCrudUsecase}
}

func (w *TodoListCategoryHandler) Register(app fiber.Router) {
	app.Get("/todo-lists-category/:id", middleware.VerifyJWTToken, w.GetByID)
	app.Get("/todo-lists-category", middleware.VerifyJWTToken, w.GetByUserID)
	app.Post("/todo-lists-category", middleware.VerifyJWTToken, w.Create)
	app.Put("/todo-lists-category/:id", middleware.VerifyJWTToken, w.Update)
	app.Delete("/todo-lists-category/:id", middleware.VerifyJWTToken, w.Delete)
}

// @Summary         Get Todo List Category by ID
// @Description     Get a Todo List Category by its ID
// @Tags            Todo List Category
// @Accept          json
// @Produce         json
// @Security        Bearer
// @Param           id path int true "ID of the Todo List Category"
// @Success			201 {object} entity.GeneralResponse{data=entity.TodoListCategoryResponse} "Success"
// @Failure			401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure			422 {object} entity.CustomErrorResponse "Invalid Request Body"
// @Failure			500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router			/api/v1/todo-lists-category/{id} [get]
func (w *TodoListCategoryHandler) GetByID(c *fiber.Ctx) error {
	id, err := w.parser.ParserIntIDFromPathParams(c)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	data, err := w.todoListCategoryCrudUsecase.GetByID(c.Context(), id)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, data, "Success", http.StatusOK)
}

// @Summary         Retrieve Todo Lists Category by User ID
// @Description     Retrieve a list of Todo Lists Category belonging to a user by their User ID
// @Tags            Todo List Category
// @Accept			json
// @Produce			json
// @Security 		Bearer
// @Success			201 {object} entity.GeneralResponse{data=[]entity.TodoListCategoryResponse} "Success"
// @Failure			401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure			422 {object} entity.CustomErrorResponse "Invalid Request Body"
// @Failure			500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router			/api/v1/todo-list-category [get]
func (w *TodoListCategoryHandler) GetByUserID(c *fiber.Ctx) error {
	userID, err := w.parser.ParserUserID(c)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	data, err := w.todoListCategoryCrudUsecase.GetByUserID(c.Context(), userID)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, data, "Success", http.StatusOK)
}

// @Summary			Create a new Todo List Category
// @Description		Create a new Todo List Category
// @Tags			Todo List Category
// @Accept			json
// @Produce			json
// @Security 		Bearer
// @Param			req body entity.TodoListCategoryReq true "Payload Request Body"
// @Success			201 {object} entity.GeneralResponse{data=entity.TodoListCategoryReq} "Success"
// @Failure			401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure			422 {object} entity.CustomErrorResponse "Invalid Request Body"
// @Failure			500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router			/api/v1/todo-list-category [post]
func (w *TodoListCategoryHandler) Create(c *fiber.Ctx) error {
	var req entity.TodoListCategoryReq

	err := w.parser.ParserBodyRequestWithUserID(c, &req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	data, err := w.todoListCategoryCrudUsecase.Create(c.Context(), req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, data, "Success", http.StatusOK)
}

// @Summary         Update an existing Todo List Category by ID
// @Description     Update an existing Todo List Category
// @Tags            Todo List Category
// @Accept          json
// @Produce         json
// @Security        Bearer
// @Param           id path int true "ID of the todo list category"
// @Param			req body entity.TodoListReq true "Payload Request Body"
// @Success			201 {object} entity.GeneralResponse "Success"
// @Failure			401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure			422 {object} entity.CustomErrorResponse "Invalid Request Body"
// @Failure			500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router			/api/v1/todo-list-category [put]
func (w *TodoListCategoryHandler) Update(c *fiber.Ctx) error {
	var req entity.TodoListCategoryReq
	err := w.parser.ParserBodyWithIntIDPathParamsAndUserID(c, &req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	err = w.todoListCategoryCrudUsecase.UpdateByID(c.Context(), req)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, nil, "Success", http.StatusOK)
}

// @Summary         Delete Todo List Category by ID
// @Description     Delete an existing Todo List Category by its ID
// @Tags			Todo List Category
// @Accept			json
// @Produce			json
// @Security 		Bearer
// @Param           id path int true "ID of the todo list category"
// @Success			201 {object} entity.GeneralResponse "Success"
// @Failure			401 {object} entity.CustomErrorResponse "Unauthorized"
// @Failure			422 {object} entity.CustomErrorResponse "Invalid Request Body"
// @Failure			500 {object} entity.CustomErrorResponse "Internal server Error"
// @Router			/api/v1/todo-lists-category/{id} [delete]
func (w *TodoListCategoryHandler) Delete(c *fiber.Ctx) error {
	id, err := w.parser.ParserIntIDFromPathParams(c)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	err = w.todoListCategoryCrudUsecase.DeleteByID(c.Context(), id)
	if err != nil {
		return w.presenter.BuildError(c, err)
	}

	return w.presenter.BuildSuccess(c, nil, "Success", http.StatusOK)
}
