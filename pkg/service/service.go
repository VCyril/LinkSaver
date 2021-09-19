package service

import "LinkSaver/pkg/repository"

type Authorization interface {

}

type LinkList interface {

}

type Link interface {

}

type Service struct {
	Authorization
	LinkList
	Link
}

func NewService(repo *repository.Repository) *Service{
	return &Service{}
}