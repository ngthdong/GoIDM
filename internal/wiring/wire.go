//go:build wireinject
// +build wireinject

//
//go:generate go run github.com/google/wire/cmd/wire	
package wiring

import (
	"github.com/google/wire"
	"github.com/ntdong/GoIDM/internal/configs"
	"github.com/ntdong/GoIDM/internal/dataaccess"
	"github.com/ntdong/GoIDM/internal/handler"
	"github.com/ntdong/GoIDM/internal/handler/grpc"
	"github.com/ntdong/GoIDM/internal/logic"
)

var WireSet = wire.NewSet(
	configs.WireSet,
	dataaccess.WireSet,
	logic.WireSet,
	handler.WireSet,
)

func InitializeGRPCServer(configFilePath configs.ConfigFilePath) (grpc.Server, func(), error) {
	wire.Build(WireSet)

	return nil, nil, nil
}
