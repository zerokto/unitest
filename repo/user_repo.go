package repo

//go:generate mockgen -destination=mock/mock_user.go -package=mock . IUser
type IUser interface {
	Get(key string) (int, error)
}

type User struct {
}

func (u *User) Get(key string) (int, error) {
	if key == "" {
		return 0, nil
	}
	return 1, nil
}
