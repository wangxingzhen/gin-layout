package data

import (
	"context"
	"fmt"
	"gin-layout/internal/biz"
	"gin-layout/internal/conf"
	"gin-layout/pkg/logx"
	"github.com/go-redis/redis"
	"github.com/google/wire"
	"github.com/pkg/errors"
	logs "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"log"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewDB,          // mysql连接
	NewRDB,         //redis连接
	NewData,        // data层
	NewTransaction, // 事务
	// ...example...
	NewUcUserRepo, // 注入用户相关 example...
	// ...
)

// Data .
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

// 用来承载事务的上下文
type contextTxKey struct{}

// NewTransaction .
func NewTransaction(d *Data) biz.Transaction {
	return d
}

// NewData .
func NewData(db *gorm.DB, rdb *redis.Client) (*Data, error) {
	return &Data{
		db:  db,
		rdb: rdb,
	}, nil
}

// InTx Transaction
func (d *Data) InTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

// DB 获取mysql
func (d *Data) DB(ctx context.Context) *gorm.DB {
	// 当前的db是不是使用事务
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}

	return d.db.Session(&gorm.Session{
		Logger:  logx.GenGormLogger(ctx),
		Context: ctx,
	})
}

func (d *Data) RDB() *redis.Client {
	return d.rdb
}

// NewDB mysql连接
func NewDB(appConf *conf.AppConfig, logger *logs.Logger) (*gorm.DB, error) {
	return newMysqlClient(appConf, logger)
}

// NewRDB redis连接
func NewRDB(appConf *conf.AppConfig) (*redis.Client, error) {
	redisClient := newRedisClient(appConf.RedisAPI)
	_, err := redisClient.Ping().Result()
	return redisClient, errors.WithStack(err)
}

// createMysqlDsn 生成dsn
func createMysqlDsn(d *conf.MysqlConf) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=%v&loc=Local",
		d.Username,
		d.Password,
		d.Hostname,
		d.Port,
		d.DatabaseName,
		d.ParseTime,
	)
}

// newMysqlClient 获取mysql连接
func newMysqlClient(appConf *conf.AppConfig, logger *logs.Logger) (*gorm.DB, error) {
	newLogger := gormlogger.New(
		log.New(logger.Writer(), "\r\n", log.LstdFlags), // io writer
		gormlogger.Config{
			IgnoreRecordNotFoundError: false, // 忽略ErrRecordNotFound（记录未找到）错误
		},
	)

	db, err := gorm.Open(mysql.Open(createMysqlDsn(appConf.DBAddress)), &gorm.Config{
		// 禁用默认的事务操作
		SkipDefaultTransaction: true,
		Logger:                 newLogger,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	conn, err := db.DB()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	conn.SetMaxOpenConns(appConf.DBAddress.MaxOpenConns)
	conn.SetMaxIdleConns(appConf.DBAddress.MaxIdleConns)

	// 开发环境和测试环境开启debug
	if appConf.Env == conf.EnvTest || appConf.Env == conf.EnvDev {
		db = db.Debug()
	}

	return db, nil
}

// RedisClient 初始化redis连接池
func newRedisClient(r *conf.RedisConf) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:         r.Address,
		DialTimeout:  time.Duration(r.DialTimeoutMillisecond) * time.Millisecond,
		ReadTimeout:  time.Duration(r.RWTimeoutMillisecond) * time.Millisecond,
		WriteTimeout: time.Duration(r.RWTimeoutMillisecond) * time.Millisecond,
		PoolSize:     r.PoolSize,
		Password:     r.Password,
	})
}
