package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "sweet/api/pb/v1"
	"sweet/internal/biz"
)

// SweetService
type SweetService struct {
	pb.UnimplementedSweetServer
	uc  *biz.SweetUsecase
	log *log.Helper
}

// NewGreeterService
func NewSweetService(uc *biz.SweetUsecase, logger log.Logger) *SweetService {
	return &SweetService{uc: uc, log: log.NewHelper(logger)}
}

// Ping healthy check
func (s *SweetService) Ping(ctx context.Context, req *pb.BaseReq) (*pb.BaseReply, error) {
	return &pb.BaseReply{Code: 200, Msg: ""}, nil
}
