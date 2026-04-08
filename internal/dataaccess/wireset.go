package dataaccess

import (
	"github.com/google/wire"
	"github.com/ntdong/GoIDM/internal/dataaccess/database"
)

var WireSet = wire.NewSet(
	database.WireSet,
)
