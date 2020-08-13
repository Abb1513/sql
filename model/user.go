/*
@Time    : 2020/8/1
@Author  : Wangcq
@File    : user.go
@Software: GoLand
*/

package model

import "time"

// 用户信息
type Users struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"updated_time"`
	Name      string    `gorm:"unique" json:"name"`
	Password  string    `json:"password"`
}
