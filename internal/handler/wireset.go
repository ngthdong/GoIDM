package handler

import (
	"github.com/google/wire"
	"github.com/ngthdong/GoIDM/internal/handler/grpc"
	"github.com/ngthdong/GoIDM/internal/handler/http"
	"github.com/ngthdong/GoIDM/internal/handler/consumers"
	"github.com/ngthdong/GoIDM/internal/handler/jobs"
)

var WireSet = wire.NewSet(
	grpc.WireSet,
	http.WireSet,
	consumers.WireSet,
	jobs.WireSet,
)
