package repository

import "github.com/jmoiron/sqlx"

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

func NewRepository(db *sqlx.DB) *Repository{
	return &Repository{}
}