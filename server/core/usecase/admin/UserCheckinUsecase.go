package admin

import (
	"server/core/entity"
	"server/core/errors"
	queryservice "server/core/infra/queryService"
	"server/core/infra/queryService/types"
	"server/core/infra/repository"

	"github.com/google/uuid"
)

type UserCheckinUsecase struct {
	checkInRepository repository.ICheckinRepository
	userCheckinQuery  queryservice.ICheckinQueryService
}

func NewUserCheckinUsecase(
	checkInRepository repository.ICheckinRepository,
	userCheckinQuery queryservice.ICheckinQueryService,
) *UserCheckinUsecase {
	return &UserCheckinUsecase{
		checkInRepository: checkInRepository,
		userCheckinQuery:  userCheckinQuery,
	}
}

func (u *UserCheckinUsecase) GetAllRecent(limit int) ([]*entity.Checkin, *errors.DomainError) {
	if limit == 0 {
		limit = 30
	}
	pager := types.NewPageQuery(
		nil,
		&limit,
	)
	checkins, err := u.userCheckinQuery.GetAllUserAllCheckin(pager)
	if err != nil {
		return nil, errors.NewDomainError(errors.QueryError, err.Error())
	}
	return checkins, nil
}

func (u *UserCheckinUsecase) GetUserLog(userID uuid.UUID, pager *types.PageQuery) ([]*entity.Checkin, *types.PageResponse, *errors.DomainError) {
	checkins, pageResponse, err := u.userCheckinQuery.GetMyAllCheckin(userID, pager)
	if err != nil {
		return nil, nil, errors.NewDomainError(errors.QueryError, err.Error())
	}

	return checkins, pageResponse, nil
}
