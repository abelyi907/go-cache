package go_cache

import (
	"encoding/json"
	"fmt"
)

func ToString(value any) string {
	switch v := value.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case nil:
		return ""
	default:
		// 尝试JSON序列化结构体
		if data, err := json.Marshal(v); err == nil {
			return string(data)
		}
		// 回退到默认格式
		return fmt.Sprintf("%v", v)
	}
}
