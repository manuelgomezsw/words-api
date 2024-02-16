package mysql

import (
	"cloud.google.com/go/cloudsqlconn"
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"net"
	"os"
	"sync"
	"words-api/internal/util/constant"
)

var (
	ClientDB *sql.DB
	once     sync.Once
)

func init() {
	once.Do(func() {
		ClientDB = mustConnect()
	})
}

func mustConnect() *sql.DB {
	var err error
	var mysqlDB *sql.DB

	if os.Getenv(constant.InstanceConnectionName) != "" {
		if os.Getenv(constant.DbUser) == "" {
			log.Fatal("Warning: DB_USER must be defined")
		}

		mysqlDB, err = connectWithConnector()
		if err != nil {
			log.Fatalf("connectConnector: unable to connect: %s", err)
		}
	}

	if mysqlDB == nil {
		log.Fatal("Missing database connection type. Please define one of INSTANCE_HOST, INSTANCE_UNIX_SOCKET, or INSTANCE_CONNECTION_NAME")
	}

	return mysqlDB
}

func connectWithConnector() (*sql.DB, error) {
	var (
		dbUser                 = os.Getenv(constant.DbUser)
		dbPwd                  = os.Getenv(constant.DbPassword)
		dbName                 = os.Getenv(constant.DbName)
		instanceConnectionName = os.Getenv(constant.InstanceConnectionName)
	)

	d, err := cloudsqlconn.NewDialer(context.Background())
	if err != nil {
		return nil, fmt.Errorf("cloudsqlconn.NewDialer: %w", err)
	}

	var opts []cloudsqlconn.DialOption

	mysql.RegisterDialContext("cloudsqlconn",
		func(ctx context.Context, addr string) (net.Conn, error) {
			return d.Dial(ctx, instanceConnectionName, opts...)
		})

	dbURI := fmt.Sprintf("%s:%s@cloudsqlconn(localhost:3306)/%s?parseTime=true",
		dbUser, dbPwd, dbName)

	dbPool, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}
	return dbPool, nil
}
