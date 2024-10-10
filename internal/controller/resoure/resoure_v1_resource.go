package resoure

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/AdminAce-coder/k8sctl/api/resoure/v1"
)

func (c *ControllerV1) Resource(ctx context.Context, req *v1.ResourceReq) (res *v1.ResourceRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
