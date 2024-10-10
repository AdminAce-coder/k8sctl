// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package resoure

import (
	"context"

	"github.com/AdminAce-coder/k8sctl/api/resoure/v1"
)

type IResoureV1 interface {
	Resource(ctx context.Context, req *v1.ResourceReq) (res *v1.ResourceRes, err error)
}
