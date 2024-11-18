package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

const (
	StatusPending = iota
	StatusPass
	StatusReject
	StatusExpried
)

// user 用户表
type ApplyFriend struct {
	Id        int64
	UserId    int64 // 用户ID
	FriendId  int64
	Remark    string    // 备注
	Status    int       // 状态
	CreatedAt time.Time // 注册时间
}

func (m *ApplyFriend) TableName() string {
	return "contact_apply"
}

type ApplyFriendModel struct {
	db *gorm.DB
}

func NewApplyFriendModel(db *gorm.DB) *ApplyFriendModel {
	return &ApplyFriendModel{
		db: db,
	}
}

func (m *ApplyFriendModel) Insert(ctx context.Context, data *ApplyFriend) error {
	return m.db.WithContext(ctx).Create(data).Error
}

func (m *ApplyFriendModel) FindById(ctx context.Context, id int64) (*ApplyFriend, error) {
	var row ApplyFriend
	err := m.db.WithContext(ctx).Where("id = ?", id).First(&row).Error
	return &row, err
}

func (m *ApplyFriendModel) FindByUidAndFid(ctx context.Context, userId, friendId int64) (*ApplyFriend, error) {
	var row ApplyFriend
	err := m.db.WithContext(ctx).Where("user_id = ? and friend_id = ?", userId, friendId).First(&row).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &row, nil
}

func (m *ApplyFriendModel) UpdateStatus(ctx context.Context, id int64, status int) error {
	return m.db.WithContext(ctx).Where("id = ? and status = ?", id, StatusPending).Update("status", status).Error
}
