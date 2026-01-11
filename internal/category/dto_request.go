package category

type CategoryRequestDTO struct {
	Name string `json:"name" validate:"min=5,max=100,required"`
}
