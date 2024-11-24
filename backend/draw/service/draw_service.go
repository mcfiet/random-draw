package service

import (
	"github.com/google/uuid"
	"github.com/mcfiet/goDo/draw/model"
	drawRepository "github.com/mcfiet/goDo/draw/repository"
)

type DrawService struct {
	repo drawRepository.DrawRepository
}

func NewDrawService(repo drawRepository.DrawRepository) *DrawService {
	return &DrawService{repo}
}

func (service *DrawService) GetAllDraws() ([]model.DrawResult, error) {
	return service.repo.GetAllDraws()
}

func (service *DrawService) CreateDraw(draw *model.DrawResult) error {
	return service.repo.CreateDraw(draw)
}

func (service *DrawService) GetDrawByGiverId(id uuid.UUID) (model.DrawResult, error) {
	return service.repo.GetDrawByGiverId(id)
}
