package action

import "github.com/SongZihuan/cat-shop-backend/src/database/action/internal"

func ConnectToMySQL() error {
	return internal.ConnectToMySQL()
}

func CloseMySQL() {
	internal.CloseMySQL()
}

func IsReady() bool {
	return internal.IsReady()
}
