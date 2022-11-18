package v1

import (
	"github.com/samandar2605/practice_api/api/models"
	"github.com/samandar2605/practice_api/config"
	"github.com/samandar2605/practice_api/storage"
)

type handlerV1 struct {
	cfg     *config.Config
	storage storage.StorageI
}

type HandlerV1Options struct {
	Cfg     *config.Config
	Storage storage.StorageI
}

func New(options *HandlerV1Options) *handlerV1 {
	return &handlerV1{
		cfg:     options.Cfg,
		storage: options.Storage,
	}
}

func errorResponse(err error) *models.ErrorResponse {
	return &models.ErrorResponse{
		Error: err.Error(),
	}
}
