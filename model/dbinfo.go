/*
@Time    : 2020/5/28
@Author  : Wangcq
@File    : dbinfo.go
@Software: GoLand
*/

package model

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// db 的信息
type (
	DbinfoResponse struct {
		ID        uint      `gorm:"primary_key" json:"id"`
		CreatedAt time.Time `json:"created_time"`
		UpdatedAt time.Time `json:"updated_time"`
		Name      string    `json:"name"`
		DbHost    string    `json:"db_host"`
		DbPort    string    `json:"db_port"`
		DbUser    string    `json:"db_user"`
	}
	DbInfo struct {
		DbPassword string `json:"db_password"`
		DbinfoResponse
	}
)

// sql 执行后结果集
type ResultMessage struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"updated_time"`

	// stage 当前语句已经进行到的阶段，包括CHECKED、EXECUTED、RERUN、NONE
	Stage string `json:"stage"`

	//错误级别。返回值为非0的情况下，说明是有错的。0表示审核通过。1表示警告，不影响执行，2表示严重错误，必须修改
	ErrorLevel string `json:"error_level"`
	// 执行状态
	StageStatus string `json:"stage_status"`
	// 失败时 错误详细信息
	ErrorMessage string `json:"error_message"`
	// 执行的sql
	Sql string `json:"sql"`
	// 影响行数
	AffectedRows string `json:"affected_rows"`
	// 唯一ID
	OpidTime string `json:"opid_time"`
	// 备份的库名
	BackupDbname string `json:"backup_dbname"`
	// 执行时长
	ExecuteTime string `json:"execute_time"`
	//Sqlsha1 string `json:"sqlsha_1"`
	// 备份用时
	BackupTime string `json:"backup_time"`
	// 回滚Sql
	Rollsql string `json:"rollsql"`
	// 关联执行的库
	ExcuteDb uint `json:"excute_db"`

	// 是否执行过回滚语句
	IsExcuteRollsql bool `json:"is_excute_rollsql"`
}
