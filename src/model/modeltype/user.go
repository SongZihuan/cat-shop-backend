package modeltype

type UserType int

const (
	NormalUserType    UserType = 1
	AdminUserType     UserType = 2
	RootAdminUserType UserType = 3
)

type UserStatus int

const (
	NormalUserStatus UserStatus = 1
	FreezeUserStatus UserStatus = 2
	DeleteUserStatus UserStatus = 3
)
