package uuid

import (
	"APT/utility/charset"
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"time"
)

var (
	startTimeStamp int64 = 1717193166
	timeStampLog   int64
	sequence       int64 = 0
	maxI           int64 = 999
	machineId            = 1
)

func CreatePayCode(bin string) (Code string) {
	var (
		nowTime      = time.Now().Unix()
		timestamp    = nowTime - startTimeStamp
		timestampStr string
		sequenceStr  = ""
	)
	if timeStampLog == nowTime {
		sequence++
		if sequence > maxI {
			// 当前毫秒内生成的序号已经超出最大范围，等待下一毫秒重新生成
			for {
				nowTime = time.Now().Unix()
				if nowTime > timeStampLog {
					timestamp = nowTime - startTimeStamp
					timeStampLog = nowTime
					sequence = 0
					break
				}
			}
		}
	} else {
		timeStampLog = nowTime
		sequence = 0
	}
	sequenceStr = fmt.Sprintf("%04d", sequence)
	timestampStr = fmt.Sprintf("%010d", timestamp)
	Code = fmt.Sprintf("%s%d%s%s", bin, machineId, timestampStr, sequenceStr)
	return
}

func CreateOrderCode(prefix string) (Code string) {
	var (
		currentTime = gtime.Now().Format("ymdhisu")
		RoundNum    = charset.RandomCreateText(2, "0123456789")
	)
	return prefix + currentTime + string(RoundNum)
}
