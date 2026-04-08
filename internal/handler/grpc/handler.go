package grpc

import (
	"context"

	"github.com/ntdong/GoIDM/internal/generated/grpc/go_load"
	"github.com/ntdong/GoIDM/internal/logic"
	"google.golang.org/grpc"
)

type Handler struct {
	go_load.UnimplementedGoLoadServiceServer
	accountLogic logic.Account
}

func NewHandler(
	accountLogic logic.Account,
) go_load.GoLoadServiceServer {
	return &Handler{
		accountLogic: accountLogic,
	}
}

func (h *Handler) CreateAccount(ctx context.Context, request *go_load.CreateAccountRequest) (*go_load.CreateAccountResponse, error) {
	output, err := h.accountLogic.CreateAccount(ctx, logic.CreateAccountParams{
		AccountName: request.GetAccountName(),
		Password: request.GetPassword(),
	}) 
	if err != nil {
		return nil, err
	}

	return &go_load.CreateAccountResponse{
		AccountId: output.ID,
	}, nil
}

func (h *Handler) CreateSession(_ context.Context, _ *go_load.CreateSessionRequest) (*go_load.CreateSessionResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (h *Handler) CreateDownloadTask(_ context.Context, _ *go_load.CreateDownloadTaskRequest) (*go_load.CreateDownloadTaskResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (h *Handler) GetDownloadTaskList(_ context.Context, _ *go_load.GetDownloadTaskListRequest) (*go_load.GetDownloadTaskListResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (h *Handler) UpdateDownloadTask(_ context.Context, _ *go_load.UpdateDownloadTaskRequest) (*go_load.UpdateDownloadTaskResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (h *Handler) DeleteDownloadTask(_ context.Context, _ *go_load.DeleteDownloadTaskRequest) (*go_load.DeleteDownloadTaskResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (h *Handler) GetDownloadTaskFile(_ *go_load.GetDownloadTaskFileRequest, _ grpc.ServerStreamingServer[go_load.GetDownloadTaskFileResponse]) error {
	panic("not implemented") // TODO: Implement
}

func (h *Handler) mustEmbedUnimplementedGoLoadServiceServer() {
	panic("not implemented") // TODO: Implement
}
