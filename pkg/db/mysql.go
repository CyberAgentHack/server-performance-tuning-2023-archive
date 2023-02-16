package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/db/config"
	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/errcode"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbTimeZone = "Asia/Tokyo"
)

func NewMySQL(cfg *config.DBConfig) (*sql.DB, error) {
	if cfg.SecretsManagerDBConfig != nil {
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("ap-northeast-1"),
		})
		if err != nil {
			return nil, errcode.New(err)
		}

		ss := secretsmanager.New(sess)
		secret, err := ss.GetSecretValue(&secretsmanager.GetSecretValueInput{
			SecretId: aws.String(cfg.SecretsManagerDBConfig.SecretID),
		})
		if err != nil {
			return nil, errcode.New(err)
		}

		var rawDBCfg config.RawDBConfig
		err = json.Unmarshal([]byte(*secret.SecretString), &rawDBCfg)
		if err != nil {
			return nil, errcode.New(err)
		}

		cfg.RawDBConfig = &rawDBCfg
	}

	// time_zoneの指定はシングルクォート込みでescapeが必要
	// https://github.com/go-sql-driver/mysql/blob/5a8a207333b3cbdd6f50a31da2d448658343637e/README.md#system-variables
	tz := url.QueryEscape("'" + dbTimeZone + "'")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local&time_zone=%s",
		cfg.RawDBConfig.Username,
		cfg.RawDBConfig.Password,
		cfg.RawDBConfig.Host,
		cfg.RawDBConfig.Port,
		cfg.RawDBConfig.DB,
		tz,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, errcode.New(err)
	}

	db.SetMaxIdleConns(50)

	if err := db.Ping(); err != nil {
		return nil, errcode.New(err)
	}

	return db, nil
}
