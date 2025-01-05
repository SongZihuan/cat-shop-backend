package action

import (
	"github.com/SongZihuan/cat-shop-backend/src/database/action/internal"
	"github.com/SongZihuan/cat-shop-backend/src/model"
)

func SendMsgByUser(msg string, user *model.User) error {
	return SendMsgByUserID(msg, user.ID)
}

func SendMsgByUserID(msg string, userID uint) error {
	return internal.DB().Create(model.NewMsg(userID, msg)).Error
}
