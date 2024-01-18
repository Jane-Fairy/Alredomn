package logic

import (
	"context"

	"Alredomn/cmd/internal/svc"
	"Alredomn/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderCreateLogic {
	return &OrderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderCreateLogic) OrderCreate(req *types.OrderCreateRequest) (resp *types.OrderCreateReply, err error) {
	// todo: add your logic here and delete this line

	return
}
