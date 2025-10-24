package repository

import (
	"encoding/csv"
	"io"
	"os"
	"threatintelapi/model"
)

type OpenPhishRepository interface {
	GetAll() ([]model.OpenPhishRecord, error)
}

type openPhishRepositoryImpl struct {
	filePath string
}

func NewOpenPhishRepository(filePath string) OpenPhishRepository {
	return &openPhishRepositoryImpl{
		filePath: filePath,
	}
}

func (r *openPhishRepositoryImpl) GetAll() ([]model.OpenPhishRecord, error) {
	file, err := os.Open(r.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Skip header
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	var records []model.OpenPhishRecord
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if len(row) >= 18 {
			record := model.OpenPhishRecord{
				URL:             row[0],
				Brand:           row[1],
				IP:              row[2],
				ASN:             row[3],
				ASNName:         row[4],
				CountryCode:     row[5],
				CountryName:     row[6],
				TLD:             row[7],
				DiscoverTime:    row[8],
				FamilyID:        row[9],
				Host:            row[10],
				ISOTime:         row[11],
				PageLanguage:    row[12],
				SSLCertIssuedBy: row[13],
				SSLCertIssuedTo: row[14],
				SSLCertSerial:   row[15],
				IsSpear:         row[16],
				Sector:          row[17],
			}
			records = append(records, record)
		}
	}

	return records, nil
}
