package uuid

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
)

// CreateSnowUUID 生成雪花算法UUID
func CreateSnowUUID(ctx context.Context) (UUID string, err error) {
	var (
		snowflake     *Snowflake
		uuidMachineId uint64
	)
	uuidMachineId = g.Cfg().MustGet(ctx, "uuid.snow.machineId").Uint64()
	if snowflake, err = NewSnowflake(uuidMachineId); err != nil {
		return
	}
	return gvar.New(snowflake.NextID()).String(), nil
}
