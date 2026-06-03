package uuid

import (
	"fmt"
	"sync"
	"time"
)

type Snowflake struct {
	timestampShift uint
	machineIdShift uint
	sequenceShift  uint
	machineId      uint64
	sequence       uint64
	timestamp      uint64
	mutex          sync.Mutex
}

const (
	twepoch       = uint64(1577836800000) // 2020-01-01T00:00:00.000Z 起点时间
	timestampMask = (1<<41 - 1) << 22     // 时间戳掩码
	machineIdMask = (1<<5 - 1) << 12      // 机器ID掩码
	sequenceMask  = 1<<12 - 1             // 序列号掩码
	workerIdMax   = 1<<5 - 1              // 最大机器ID
	sequenceIdMax = 1<<12 - 1             // 最大序列ID
)

var (
	SnowFlakeNo *Snowflake
)

func InitSnowFlake() {
	var (
		err error
	)
	if SnowFlakeNo, err = NewSnowflake(1); err != nil {
		panic(err)
	}
}

func NewSnowflake(machineId uint64) (*Snowflake, error) {
	if machineId > workerIdMax {
		return nil, fmt.Errorf("machineId is too large, must be less than %d", workerIdMax)
	}
	return &Snowflake{
		timestampShift: 22,
		machineIdShift: 12,
		sequenceShift:  0,
		machineId:      machineId,
		sequence:       0,
		timestamp:      0,
	}, nil
}

func (s *Snowflake) NextID() uint64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	currentTime := uint64(time.Now().UnixNano()/1e6) + twepoch
	if s.timestamp == currentTime {
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			for currentTime <= s.timestamp {
				currentTime = uint64(time.Now().UnixNano()/1e6) + twepoch
			}
		}
	} else {
		s.sequence = 0
	}
	s.timestamp = currentTime

	id := (currentTime << s.timestampShift) | (s.machineId << s.machineIdShift) | s.sequence
	return id
}
