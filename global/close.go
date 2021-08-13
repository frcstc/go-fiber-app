package global

import "fmt"

func CloseGlobal() {


	// 关闭redis
	err := Redis.Close()
	if err != nil {
		SLog.Error("Redis close err", err)
		fmt.Println("Redis close err", err)
	}
	// 关闭数据库
	sqlDB, err := DB.DB()
	err = sqlDB.Close()
	if err != nil {
		SLog.Error("mysql close err", err)
		fmt.Println("mysql close err", err)
	}
	err = LogFile.Close()
	if err != nil {
		SLog.Error("logfile close err", err)
		fmt.Println("logfile close err", err)
	}

}
