package service

import (
	"log"
	"threatintelapi/model"
	"threatintelapi/repository"
)

type OpenPhishService interface {
	GetAllRecords() ([]model.OpenPhishRecord, error)
}

type openPhishServiceImpl struct {
	repo repository.OpenPhishRepository
}

func NewOpenPhishService(repo repository.OpenPhishRepository) OpenPhishService {
	return &openPhishServiceImpl{
		repo: repo,
	}
}

func (s *openPhishServiceImpl) GetAllRecords() ([]model.OpenPhishRecord, error) {
	records, err := s.repo.GetAll()
	if err != nil {
		log.Printf("Error fetching OpenPhish records: %v", err)
		return nil, err
	}

	log.Printf("Retrieved %d OpenPhish records", len(records))
	return records, nil
}
