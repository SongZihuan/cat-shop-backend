package updateuserpassword

type Query struct {
	OldPassword string `form:"oldPassword"`
	NewPassword string `form:"newPassword"`
}
