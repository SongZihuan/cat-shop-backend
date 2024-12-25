package updateuserpassword

type Query struct {
	OldPassword string `form:"oldpassword"`
	NewPassword string `form:"newpassword"`
}
