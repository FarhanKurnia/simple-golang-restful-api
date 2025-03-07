package web

type CategoryUpdateRequest struct {
	Id   int64  `validate:"requiered"`
	Name string `validate:"required, max=200, min=1"`
}
