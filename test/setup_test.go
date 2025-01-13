package integration_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/blacksheepaul/templateToGo/core/config"
	"github.com/blacksheepaul/templateToGo/core/logger"
	"github.com/blacksheepaul/templateToGo/model"
	"github.com/blacksheepaul/templateToGo/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func TestMain(m *testing.M) {
	cfg := config.GetConfig("../config-test.yml")
	gin.SetMode(gin.DebugMode)
	serviceForTest(cfg, FakeLogger{})
	daoForTest(cfg, FakeLogger{})

	if cfg.Test.Flush {
		flushDb()
	}

	os.Exit(m.Run())
}

func flushDb() {
	dao := model.GetDao()
	driver, err := sqlite3.WithInstance(dao.RawDB, &sqlite3.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://../model/migrations",
		"sqlite3", driver)
	if err != nil {
		panic(err)
	}
	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		panic(fmt.Errorf("Failed to drop tables: %v", err))
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(fmt.Errorf("Failed to apply migrations: %v", err))
	}
}

func daoForTest(cfg *config.Config, logi logger.Logger) {
	model.InitDao(cfg, logi)
}

func serviceForTest(cfg *config.Config, logi logger.Logger) {
	service.InitService(logi, cfg)
}

type FakeLogger struct{}

func (l FakeLogger) Debug(fields ...interface{}) {
	fmt.Println(fields...)
}

func (l FakeLogger) Debugw(msg string, keysAndValues ...interface{}) {
	fmt.Println(append([]interface{}{msg}, keysAndValues...)...)
}

func (l FakeLogger) Info(fields ...interface{}) {
	fmt.Println(fields...)
}

func (l FakeLogger) Infow(msg string, keysAndValues ...interface{}) {
	fmt.Println(append([]interface{}{msg}, keysAndValues...)...)
}

func (l FakeLogger) Warn(fields ...interface{}) {
	fmt.Println(fields...)
}

func (l FakeLogger) Warnw(msg string, keysAndValues ...interface{}) {
	fmt.Println(append([]interface{}{msg}, keysAndValues...)...)
}

func (l FakeLogger) Error(fields ...interface{}) {
	fmt.Println(fields...)
}

func (l FakeLogger) Errorw(msg string, keysAndValues ...interface{}) {
	fmt.Println(append([]interface{}{msg}, keysAndValues...)...)
}

func (l FakeLogger) Fatal(fields ...interface{}) {
	fmt.Println(fields...)
}

func (l FakeLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	fmt.Println(append([]interface{}{msg}, keysAndValues...)...)
}
