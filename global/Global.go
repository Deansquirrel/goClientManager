package global

import (
	"context"
	"github.com/Deansquirrel/goClientManager/config"
)

var ServerConfig *config.ServerConfig
var ClientConfig *config.ClientConfig
var Ctx context.Context
var Cancel func()
