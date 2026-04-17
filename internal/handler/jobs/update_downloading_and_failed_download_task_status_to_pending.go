package jobs

import (
	"context"

	"github.com/ngthdong/GoIDM/internal/logic"
)

type UpdateDownloadingAndFailedDownloadTaskStatusToPending interface {
	Run(context.Context) error
}

type updateDownloadingAndFailedDownloadTaskStatusToPending struct {
	downloadTaskLogic logic.DownloadTask
}

func NewUpdateDownloadingAndFailedDownloadTaskStatusToPending(
	downloadTaskLogic logic.DownloadTask,
) UpdateDownloadingAndFailedDownloadTaskStatusToPending {
	return &updateDownloadingAndFailedDownloadTaskStatusToPending{
		downloadTaskLogic: downloadTaskLogic,
	}
}

func (u updateDownloadingAndFailedDownloadTaskStatusToPending) Run(ctx context.Context) error {
	panic("Unimplemented")
}