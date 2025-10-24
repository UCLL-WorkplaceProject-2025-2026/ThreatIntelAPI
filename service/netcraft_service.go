package service

import (
	"log"
	"threatintelapi/model"
	"threatintelapi/repository"
)

type NetcraftService interface {
	GetAllRecords() ([]model.NetcraftRecord, error)
}

type netcraftServiceImpl struct {
	repo repository.NetcraftRepository
}

func NewNetcraftService(repo repository.NetcraftRepository) NetcraftService {
	return &netcraftServiceImpl{
		repo: repo,
	}
}

func (s *netcraftServiceImpl) GetAllRecords() ([]model.NetcraftRecord, error) {
	records, err := s.repo.GetAll()
	if err != nil {
		log.Printf("Error fetching Netcraft records: %v", err)
		return nil, err
	}

	log.Printf("Retrieved %d Netcraft records", len(records))
	return records, nil
}
