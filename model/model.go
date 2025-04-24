package model

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/blacksheepaul/templateToGo/core/config"
	"github.com/blacksheepaul/templateToGo/core/logger"

	"github.com/patrickmn/go-cache"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gl "gorm.io/gorm/logger"
)

var (
	dao  *Dao
	once sync.Once
	log  logger.Logger
)

type DBProvider interface {
	Db() *gorm.DB
}

type Dao struct {
	db    *gorm.DB
	RawDB *sql.DB
	cache *cache.Cache
}

func (d *Dao) Db() *gorm.DB {
	return d.db
}

func (d *Dao) Begin() *TxDao {
	tx := d.db.Begin()
	if tx.Error != nil {
		panic(tx.Error)
	}
	return &TxDao{Dao: d, txDb: tx}
}

type TxDao struct {
	*Dao
	txDb *gorm.DB
}

func (tx *TxDao) Db() *gorm.DB {
	return tx.txDb
}

func (tx *TxDao) Commit() error {
	return tx.txDb.Commit().Error
}

func (tx *TxDao) Rollback() error {
	return tx.txDb.Rollback().Error
}

var _ DBProvider = (*Dao)(nil)
var _ DBProvider = (*TxDao)(nil)

func InitDao(cfg *config.Config, loggerInstance logger.Logger) {
	once.Do(func() {
		if db, err := gorm.Open(sqlite.Open(cfg.Database.Host), &gorm.Config{
			Logger: gl.Default.LogMode(gl.LogLevel(cfg.Log.ORMLogLevel)),
		}); err != nil {
			panic(err)
		} else {
			log = loggerInstance
			raw, _ := db.DB()
			dao = &Dao{db: db, RawDB: raw}
		}

		dao.cache = cache.New(5*time.Minute, 10*time.Minute)

	})
}

func GetDao() *Dao {
	if dao == nil {
		panic("dao is nil, please call InitDao first")
	}
	return dao
}

type Model struct {
	*gorm.Model
}

func (d *Dao) WriteCache(key string, value any, seconds int64) {
	exp := time.Duration(seconds) * time.Second
	d.cache.Set(key, value, exp)
}

func (d *Dao) GetCache(key string) (any, bool) {
	return d.cache.Get(key)
}

func (d *Dao) AdminGetAllCache() {
	items := d.cache.Items()

	log.Debugw("list cache items",
		"count", len(items),
	)

	for k, v := range items {
		log.Debug(fmt.Sprintf("cache item [ %s --- %v ]", k, v))
	}
}
