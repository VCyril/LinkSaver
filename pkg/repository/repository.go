package repository


type Authorization interface {

}

type LinkList interface {

}

type Link interface {

}

type Repository struct {
	Authorization
	LinkList
	Link
}

func NewRepository() *Repository{
	return &Repository{}
}