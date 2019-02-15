package global

import (
	"context"
	"github.com/Deansquirrel/goClientManager/object"
)

var ServerConfig *object.ServerConfig
var ClientConfig *object.ClientConfig
var Ctx context.Context
var Cancel func()
