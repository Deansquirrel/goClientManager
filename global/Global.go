package global

import (
	"context"
	"github.com/Deansquirrel/goClientManager/object"
)

var SysConfig *object.SysConfig
var Ctx context.Context
var Cancel func()
