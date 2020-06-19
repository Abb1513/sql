package utools

import (
	"database/sql"
	"fmt"
	"goinception/model"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var goinceptionHost, goinceptionPort, rollbakdb_host, rollbakdb_port, rollbakdb_passwd, rollbakdb_user string

func InitGoinception() {
	goinceptionHost = viper.GetString("goinception.host")
	goinceptionPort = viper.GetString("goinception.port")
	rollbakdb_host = viper.GetString("goinception.rollbakdb_host")
	rollbakdb_port = viper.GetString("goinception.rollbakdb_port")
	rollbakdb_passwd = viper.GetString("goinception.rollbakdb_passwd")
	rollbakdb_user = viper.GetString("goinception.rollbakdb_user")
}

// 执行Sql
func Exec(dbinfo model.DbInfo, sqll string) []model.ResultMessage {
	goinceptionUrl := fmt.Sprintf("r:r@tcp(%s:%s)/", goinceptionHost, goinceptionPort)
	db, err := sql.Open("mysql", goinceptionUrl)
	defer db.Close()
	sql := fmt.Sprintf(`/*--user=%s;--password=%s;--host=%s;--port=%s;--execute=1;backup=1*/
						inception_magic_start;
						%s
						inception_magic_commit;
						`, dbinfo.DbUser, DecodeDbPasswd(dbinfo.DbPassword), dbinfo.DbHost, dbinfo.DbPort, sqll)

	//sql := fmt.Sprintf(`/*--user=root;--password=q10010;--host=120.77.210.151;--port=3306;--execute=1;backup=1*/
	//inception_magic_start;
	//use  test;
	//insert into test1 (myname) VALUES("%s");
	//insert into test1 (myname) VALUES("%s");
	//inception_magic_commit;`, time.Now().String(), time.Now().Format("2006-01-02 15:04:05"))
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var resultList []model.ResultMessage
	for rows.Next() {
		var orderId, affectedRows, stage, errorLevel, stageStatus, errorMessage, sql, sequence, backupDbname, executeTime, sqlsha1, backupTime []uint8
		err = rows.Scan(&orderId, &stage, &errorLevel, &stageStatus, &errorMessage, &sql, &affectedRows, &sequence, &backupDbname, &executeTime, &sqlsha1, &backupTime)
		if err != nil {
			fmt.Println("scan, ", err)
		}
		resultmessage := model.ResultMessage{
			Stage:        string(stage),
			ErrorLevel:   string(errorLevel),
			StageStatus:  string(stageStatus),
			ErrorMessage: string(errorMessage),
			Sql:          string(sql),
			AffectedRows: string(affectedRows),
			OpidTime:     string(sequence),
			BackupDbname: string(backupDbname),
			ExecuteTime:  string(executeTime),
			BackupTime:   string(backupTime),
			ExcuteDb:     dbinfo.ID,
		}
		if string(errorLevel) == "2" {
			resultList = append(resultList, resultmessage)
			return resultList
		}
		if string(orderId) != "1" {
			resultList = append(resultList, resultmessage)
			//
			// fmt.Printf("执行ID: %s\n 影响行数: %s\n 执行操作: %s\n  错误等级: %s\n 操作状态:%s\n 错误信息: %s\n 执行SQL: %s\n Opit_time: %s\n 备份库名:%s\n 执行时间:%s\n", string(orderId), string(affectedRows), string(stage), string(errorLevel), string(stageStatus), string(errorMessage), string(sql), string(sequence), string(backupDbname), string(executeTime))
		}
	}
	return resultList
}

// 获取回滚Sql
func GetRollSql(bakdb, optime string) string {
	mysqlurl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", rollbakdb_user, rollbakdb_passwd, rollbakdb_host, rollbakdb_port, bakdb)
	db, _ := sql.Open("mysql", mysqlurl)
	err := db.Ping()
	if err != nil {
		fmt.Println("client, ", err)
	}
	selectTablename := fmt.Sprintf("select tablename from $_$Inception_backup_information$_$ where opid_time='%s'", optime)
	rows, err := db.Query(selectTablename)
	if err != nil {
		fmt.Println("from 89", err)
	}
	var tablename string
	for rows.Next() {
		err = rows.Scan(&tablename)
		if err != nil {
			fmt.Println("from 91 ", err)
		}
	}

	sql := fmt.Sprintf("select rollback_statement from %s where opid_time='%s';", tablename, optime)
	defer db.Close()
	rows, _ = db.Query(sql)
	defer rows.Close()
	var baksql string
	for rows.Next() {
		err = rows.Scan(&baksql)
		if err != nil {
			fmt.Println("from 107 ", err)
		}
	}
	return baksql

}

func ExcuteRollBak(res model.ResultMessage, tmp model.DbInfo) error {
	mysqlurl := fmt.Sprintf("%s:%s@tcp(%s:%s)/", tmp.DbUser, DecodeDbPasswd(tmp.DbPassword), tmp.DbHost, tmp.DbPort)
	db, _ := sql.Open("mysql", mysqlurl)
	_, err := db.Query(res.Rollsql)
	if err != nil {
		return err
	}
	return nil
}
