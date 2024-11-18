package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// user 用户表
type User struct {
	Id        int64     // 用户ID
	Mobile    string    // 手机号
	Nickname  string    // 用户昵称
	Avatar    string    // 用户头像
	Gender    int       // 用户性别[0:未知;1:男 ;2:女;]
	Password  string    // 用户密码
	Salt      string    // 密码盐值
	Motto     string    // 用户座右铭
	Email     string    // 用户邮箱
	Birthday  string    // 生日
	IsRobot   int       // 是否机器人[0:否;1:是;]
	CreatedAt time.Time // 注册时间
	UpdatedAt time.Time // 更新时间
}

func (m *User) TableName() string {
	return "user"
}

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{
		db: db,
	}
}

func (m *UserModel) Insert(ctx context.Context, data *User) error {
	return m.db.WithContext(ctx).Create(data).Error
}

func (m *UserModel) FindById(ctx context.Context, id int64) (*User, error) {
	var u User
	err := m.db.WithContext(ctx).Where("id = ?", id).First(&u).Error
	return &u, err
}

func (m *UserModel) FindByEmail(ctx context.Context, email string) (*User, error) {
	var u User
	err := m.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &u, err
}

func (m *UserModel) FindByMobile(ctx context.Context, mobile string) (*User, error) {
	var u User
	err := m.db.WithContext(ctx).Where("mobile = ?", mobile).First(&u).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &u, err
}

func (m *UserModel) Update(ctx context.Context, data *User) error {
	return m.db.WithContext(ctx).Save(data).Error
}
