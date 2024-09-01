package mocks

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GetDBMock returns a mock GORM DB connection and SQL mock object for testing
func GetDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	// sqlmock を使ってモック DB を作成
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	// GORM v2 のドライバ設定を使用してモック DB を GORM に接続
	dialector := mysql.New(mysql.Config{
		Conn:                      db,   // sqlmock の DB 接続を設定
		SkipInitializeWithVersion: true, // バージョンチェックをスキップ
	})

	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	return gormDB, mock, nil
}
