package data

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/timurkash/kratos-layout/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data is struct for Data component
type Data struct {
	log *log.Helper
}

// NewData is constructor for Data component
func NewData(confData *conf.Data, logger log.Logger) (*Data, error) {
	if confData == nil {
		return nil, fmt.Errorf("bad config data")
	}
	return &Data{
		log: log.NewHelper(logger),
	}, nil
}
