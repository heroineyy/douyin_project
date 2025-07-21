package user_info

import (
	"byte_douyin_project/cache"
	"byte_douyin_project/models"
	"errors"
)

const (
	FOLLOW = 1
	CANCEL = 2
)

var (
	ErrIvdAct    = errors.New("未定义操作")
	ErrIvdFolUsr = errors.New("关注用户不存在")
)

func PostFollowAction(userId, userToId int64, actionType int) error {
	//由于userId是经过乐token鉴权故不需要check，只需要检查userToId
	if !models.NewUserInfoDAO().IsUserExistById(userToId) {
		return ErrIvdFolUsr
	}
	if actionType != FOLLOW && actionType != CANCEL {
		return ErrIvdAct
	}
	//自己不能关注自己
	if userId == userToId {
		return ErrIvdAct
	}
	userDAO := models.NewUserInfoDAO()

	switch actionType {
	case FOLLOW:
		if err := userDAO.AddUserFollow(userId, userToId); err != nil {
			return err
		}
		//更新redis的关注信息
		cache.NewProxyIndexMap().UpdateUserRelation(userId, userToId, true)
	case CANCEL:
		if err := userDAO.CancelUserFollow(userId, userToId); err != nil {
			return err
		}
		cache.NewProxyIndexMap().UpdateUserRelation(userId, userToId, false)
	default:
		return ErrIvdAct
	}
	return nil
}
