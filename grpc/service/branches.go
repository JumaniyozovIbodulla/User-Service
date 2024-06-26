package service

import (
	"context"
	br "user/genproto/user_service/branches"
)

func (s *BranchesService) Create(context.Context, *br.CreateBranch) (*br.Branch, error) {
	panic("unimplemented")
}

func (s *BranchesService) Delete(context.Context, *br.BranchePrimaryKey) (*br.Empty, error) {
	panic("unimplemented")
}

func (s *BranchesService) GetAll(context.Context, *br.GetListBranchesRequest) (*br.GetListBranchesResponse, error) {
	panic("unimplemented")
}

func (s *BranchesService) GetById(context.Context, *br.BranchePrimaryKey) (*br.Branch, error) {
	panic("unimplemented")
}

func (s *BranchesService) Update(context.Context, *br.UpdateBranch) (*br.Branch, error) {
	panic("unimplemented")
}
