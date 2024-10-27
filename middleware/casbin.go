package middleware

import (
	"strings"
)

type casbinPolicy struct {
	Sub string `json:"sub"`
	Obj string `json:"obj"`
	Act string `json:"act"`
}

func KeyMatchFunc(args ...interface{}) (interface{}, error) {

	pattern := args[0].(string)
	// 如果模式为 '*'，表示可以匹配任意路径
	if pattern == "*" {
		return true, nil
	}
	key := args[1].(string)

	// keyMatch 函数用于匹配路径和模式，其中 '*' 代表匹配任意字符

	// 如果模式包含 '*'
	if strings.Contains(pattern, "*") {
		// 将 '*' 替换为正则表达式的 `.*`，匹配任意字符
		pattern = strings.ReplaceAll(pattern, "*", ".*")
		// 使用 HasPrefix 来匹配通配符前缀的内容
		return strings.HasPrefix(key, strings.TrimSuffix(pattern, ".*")), nil
	}

	// 没有通配符时直接匹配
	return pattern == key, nil
}

func casbinAuth() {
	// 初始化 Casbin 的 GORM 适配器
	//adapter, err := gormadapter.NewAdapterByDB(global.DB)
	//if err != nil {
	//	log.Fatalf("failed to initialize adapter: %v", err)
	//}

}
