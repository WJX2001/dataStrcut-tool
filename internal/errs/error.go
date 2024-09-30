package errs

import (
	"fmt"
	"time"
)

// NewErrIndexOutOfRange 创建一个代表下标超出范围的错误
func NewErrIndexOutOfRange(length int, index int) error {
	return fmt.Errorf("datastrcut-tool: 下标超出范围，长度 %d, 下标 %d", length, index)
}

// NewErrInvalidType 创建一个代表类型转换失败的错误
func NewErrInvalidType(want string, got any) error {
	return fmt.Errorf("datastrcut-tool: 类型转换失败，预期类型:%s, 实际值:%#v", want, got)
}

// 无效时间间隔
func NewErrInvalidIntervalValue(interval time.Duration) error {
	return fmt.Errorf("datastrcut-tool: 无效的时间间隔值: %d, 预期值应大于0", interval)
}

// 最大重试间隔
func NewErrInvalidMaxIntervalValue(maxInterval, initialInterval time.Duration) error {
	return fmt.Errorf("datastrcut-tool: 最大重试间隔的时间 [%d] 应大于等于初始重试的间隔时间 [%d] ", maxInterval, initialInterval)
}
