package database

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	InitializeDB,
	InitializeGoquDB,
	wire.Bind(new(Database), new(*goqu.Database)),
	NewAccountDatabaseAccessor,
	NewAccountPasswordDataAccessor,
)
