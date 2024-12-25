package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/qahvazor/qahvazor/internal/dto"
	"github.com/qahvazor/qahvazor/internal/model"
	"gorm.io/gorm"
)

type ImageRepositoryImpl struct {
	db *gorm.DB
}

func NewImageRepository(db *gorm.DB) ImageRepository {
	return &ImageRepositoryImpl{db: db}
}

func (r *ImageRepositoryImpl) GetImageById(id uint) (*dto.ImageDTO, error){
	var imageModel model.ImageModel
	if err := r.db.Where("id = ?", id).First(&imageModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &dto.ImageDTO{
		FileName: imageModel.Filename,
		FilePath: imageModel.Path,
		FileExt: imageModel.Ext,
	}, nil
}

func (r *ImageRepositoryImpl) CreateImage(data dto.ImageDTO) (int, error) {
	image := model.ImageModel{
		Filename: data.FileName,
		Path:     data.FilePath,
		Ext:      data.FileExt,
		HashUID:  uuid.New(),
	}

	query := r.db.Create(&image)
	if query.Error != nil {
		return 0, query.Error
	}

	return int(image.ID), nil
} 
