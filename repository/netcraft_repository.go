package repository

import (
	"encoding/json"
	"io"
	"os"
	"threatintelapi/model"
)

type NetcraftRepository interface {
	GetAll() ([]model.NetcraftRecord, error)
}

type netcraftRepositoryImpl struct {
	filePath string
}

func NewNetcraftRepository(filePath string) NetcraftRepository {
	return &netcraftRepositoryImpl{
		filePath: filePath,
	}
}

func (r *netcraftRepositoryImpl) GetAll() ([]model.NetcraftRecord, error) {
	file, err := os.Open(r.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var records []model.NetcraftRecord
	if err := json.Unmarshal(data, &records); err != nil {
		return nil, err
	}

	return records, nil
}
