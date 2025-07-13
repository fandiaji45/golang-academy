package entity

type TodoListCategoryReq struct {
	ID          int64  `json:"id,omitempty" swaggerignore:"true"`
	Name        int64  `json:"name,omitempty" validate:"required"`
	Description string `json:"description" validate:"required" name:"Deskripsi"`
	CreatedAt   string `json:"created_at" validate:"required" name:"Tanggal Created_at"`
	CreatedBy   int64  `json:"created_by" validate:"required" name:"Create_by"`
}

type TodoListCategoryResponse struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	CreatedBy   int64  `json:"created_by"`
}

func (r *TodoListCategoryReq) SetID(ID int64) {
	r.ID = ID
}

func (r *TodoListCategoryReq) SetCreatedBy(CreatedBy int64) {
	r.CreatedBy = CreatedBy
}
