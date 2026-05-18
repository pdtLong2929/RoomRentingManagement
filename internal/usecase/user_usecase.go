package usecase

import "context"

type UserUsecase interface {
	GetDetail(ctx *context.Context)
}
