package consumers

import (
	"context"

	"go.uber.org/zap"

	"github.com/ngthdong/GoIDM/internal/dataaccess/mq/producer"
	"github.com/ngthdong/GoIDM/internal/logic"
	"github.com/ngthdong/GoIDM/internal/utils"
)

type DownloadTaskCreated interface {
	Handle(ctx context.Context, event producer.DownloadTaskCreated) error
}

type downloadTaskCreated struct {
	downloadTaskLogic logic.DownloadTask
	logger            *zap.Logger
}

func NewDownloadTaskCreated(
	downloadTaskLogic logic.DownloadTask,
	logger *zap.Logger,
) DownloadTaskCreated {
	return &downloadTaskCreated{
		downloadTaskLogic: downloadTaskLogic,
		logger:            logger,
	}
}

func (d downloadTaskCreated) Handle(ctx context.Context, event producer.DownloadTaskCreated) error {
	logger := utils.LoggerWithContext(ctx, d.logger).With(zap.Any("event", event))
	logger.Info("download task created event received")

	return nil
}
