package service

import (
	"testing"
	"unitest/repo/mock"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUserSvc_FindUserCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockIUser(ctrl)

	Convey("first step", t, func() {
		m.EXPECT().Get(gomock.Any()).Return(100, nil)

		count, err := New(m).FindUserCount("abc")
		So(err, ShouldEqual, nil)
		So(count, ShouldEqual, 100)
	})

}

func TestUserSvc_DelUser(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockIUser(ctrl)

	ApplyFunc(PrivateFunc, func(a, b, c string) string { return "china" })
	ApplyFunc(Add, func(a, b int) int { return 2 })

	Convey("删除用户", t, func() {

		res := New(m).DelUser("abc")
		So(res, ShouldEqual, "china")
	})
}
