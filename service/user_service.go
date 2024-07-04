package service

import (
	"fmt"
	"unitest/repo"
)

type UserSvc struct {
	UserRepo repo.IUser
}

func New(user repo.IUser) *UserSvc {
	return &UserSvc{UserRepo: user}
}

func (u *UserSvc) FindUserCount(userID string) (int, error) {
	if userID == "" {
		return 0, nil
	}

	return u.UserRepo.Get(userID)
}

func (u *UserSvc) DelUser(userID string) string {
	if userID == "" {
		return ""
	}
	tmp := PrivateFunc("one", "two", "three")
	aa := Add(1, 3)

	return tmp + fmt.Sprint(aa)
}

func PrivateFunc(a, b, c string) string {
	return a + b + c
}

func Add(a, b int) int {
	return a + b
}
