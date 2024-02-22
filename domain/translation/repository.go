package translation

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (s *Repository) GetMessage() Model {
	return Model{Text: "Hello World"}
}
