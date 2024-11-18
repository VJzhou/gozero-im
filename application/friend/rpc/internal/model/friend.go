package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// 好友表
type Friend struct {
	Id        int64
	UserId    int64 // 用户ID
	FriendId  int64
	Remark    string // 备注
	Status    int    // 状态
	GroupId   int64
	CreatedAt time.Time // 注册时间
}

func (m *Friend) TableName() string {
	return "contact"
}

type FriendModel struct {
	db *gorm.DB
}

func NewFriendModel(db *gorm.DB) *FriendModel {
	return &FriendModel{
		db: db,
	}
}

func (m *FriendModel) Insert(ctx context.Context, data *ApplyFriend) error {
	return m.db.WithContext(ctx).Create(data).Error
}

func (m *FriendModel) FindById(ctx context.Context, id int64) (*Friend, error) {
	var row Friend
	err := m.db.WithContext(ctx).Where("id = ?", id).First(&row).Error
	return &row, err
}

func (m *FriendModel) FindByUidAndFid(ctx context.Context, userId, friendId int64) (*Friend, error) {
	var row Friend
	err := m.db.WithContext(ctx).Where("user_id = ? and friend_id = ?", userId, friendId).First(&row).Error
	return &row, err
}

func (m *FriendModel) DeleteById(ctx context.Context, id int64) error {
	return m.db.WithContext(ctx).Where("id = ?", id).Delete(&Friend{}).Error
}
