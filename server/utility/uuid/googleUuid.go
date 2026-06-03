package uuid

import (
	"github.com/gogf/gf/v2/container/gvar"
	googleUUID "github.com/google/uuid"
)

var (
	GoogleUuid googleUUID.UUID
)

func InitUuid() {
	GoogleUuid = googleUUID.New()
}

func GetUuid() string {
	return gvar.New(SnowFlakeNo.NextID()).String()
	//return GoogleUuid.String()
}
