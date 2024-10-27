package initialize

import (
	"fmt"
	"gin-demo/global"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"log"
	"strings"
)

type casbinPolicy struct {
	Sub string `json:"sub"`
	Obj string `json:"obj"`
	Act string `json:"act"`
}

func KeyMatchFunc(args ...interface{}) (interface{}, error) {
	pattern := args[0].(string)
	key := args[1].(string)
	// 没有 '*'，直接比较
	if !strings.Contains(pattern, "*") {
		return pattern == key, nil
	}
	// 如果模式为 '*'，表示可以匹配任意路径
	if pattern == "*" {
		return true, nil
	}
	// keyMatch函数用于匹配路径和模式，其中'*'代表匹配任意字符
	// 如果模式包含 '*'
	// 将 '*' 替换为正则表达式的 `.*`，匹配任意字符
	pattern = strings.ReplaceAll(pattern, "*", ".*")
	// 使用 HasPrefix 来匹配通配符前缀的内容
	return strings.HasPrefix(key, strings.TrimSuffix(pattern, ".*")), nil
}

func InitCasbin() {
	// 定义模型策略
	modelConf := `
		[request_definition]        
		r = sub, obj, act
		[policy_definition]         
		p = sub, obj, act
		[role_definition]           
		g = _, _
		[policy_effect]              
		e = some(where (p.eft == allow))
		[matchers]                 
		m = g(r.sub, p.sub) && keyMatch(r.obj, p.obj) && (r.act == p.act || p.act == "*")
	`
	// 初始化 Casbin 的 GORM 适配器
	adapter, err := gormadapter.NewAdapterByDB(global.DB)
	if err != nil {
		fmt.Printf("failed to initialize adapter: %v\n", err)
	}

	// 初始化 Casbin
	enforcer, err := casbin.NewEnforcer(modelConf, adapter)
	if err != nil {
		fmt.Printf("failed to initialize casbin: %v\n", err)
	}

	// 同步数据库中的策略
	err = enforcer.LoadPolicy()
	if err != nil {
		log.Printf("failed to load policy: %v\n", err)
	}

}
