//go:build wireinject
// +build wireinject

//
//go:generate go run github.com/google/wire/cmd/wire
package wiring

import (
	"github.com/google/wire"
	"github.com/ngthdong/GoIDM/internal/configs"
	"github.com/ngthdong/GoIDM/internal/dataaccess"
	"github.com/ngthdong/GoIDM/internal/handler"
	"github.com/ngthdong/GoIDM/internal/handler/grpc"
	"github.com/ngthdong/GoIDM/internal/logic"
	"github.com/ngthdong/GoIDM/internal/utils"
)

var WireSet = wire.NewSet(
	configs.WireSet,
	utils.WireSet,
	dataaccess.WireSet,
	logic.WireSet,
	handler.WireSet,
)

func InitializeGRPCServer(configFilePath configs.ConfigFilePath) (grpc.Server, func(), error) {
	wire.Build(WireSet)

	return nil, nil, nil
}
