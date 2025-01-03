package action

import (
	"github.com/SuperH-0630/cat-shop-back/src/database/action/internal"
	"github.com/SuperH-0630/cat-shop-back/src/model"
)

func SendMsgByUser(msg string, user *model.User) error {
	return SendMsgByUserID(msg, user.ID)
}

func SendMsgByUserID(msg string, userID uint) error {
	return internal.DB().Create(model.NewMsg(userID, msg)).Error
}
