package adminupdateuserpassword

type Query struct {
	NewPassword string `form:"newpassword"`
}
