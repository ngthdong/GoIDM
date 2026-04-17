package mq

import (
	"github.com/google/wire"

	"github.com/ngthdong/GoIDM/internal/dataaccess/mq/consumer"
	"github.com/ngthdong/GoIDM/internal/dataaccess/mq/producer"
)

var WireSet = wire.NewSet(
	consumer.WireSet,
	producer.WireSet,
)
