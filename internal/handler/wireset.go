package handler

import (
	"github.com/google/wire"
	"github.com/ntdong/GoIDM/internal/handler/grpc"
	"github.com/ntdong/GoIDM/internal/handler/http"
)

var WireSet = wire.NewSet(
	grpc.WireSet,
	http.WireSet,
)
