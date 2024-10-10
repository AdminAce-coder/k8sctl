package v1

import "github.com/gogf/gf/v2/frame/g"

type ResourceReq struct {
	g.Meta `path:"/resource" method:"post" summary:"配置信息" tags:"配置"`
}
type ResourceRes struct{}
