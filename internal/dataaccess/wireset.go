package dataaccess

import (
	"github.com/google/wire"
	"github.com/ngthdong/GoIDM/internal/dataaccess/cache"
	"github.com/ngthdong/GoIDM/internal/dataaccess/database"
)

var WireSet = wire.NewSet(
	cache.WireSet,
	database.WireSet,
)
