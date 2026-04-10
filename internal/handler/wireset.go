package handler

import (
	"github.com/google/wire"
	"github.com/ngthdong/GoIDM/internal/handler/grpc"
	"github.com/ngthdong/GoIDM/internal/handler/http"
)

var WireSet = wire.NewSet(
	grpc.WireSet,
	http.WireSet,
)
