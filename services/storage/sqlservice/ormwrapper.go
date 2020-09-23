package sqlservice

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Rahul12344/skelego"

	"github.com/Rahul12344/skelego/services/storage"
	"github.com/jinzhu/gorm"

	// postgres gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//ORMWrapper Interface for ORM Wrapper for SQL database.
type ORMWrapper interface {
	storage.Store
	DB() (*sql.DB, error)
	ORM() *gorm.DB
}

//db example postgres interface
type orm struct {
	client  *gorm.DB
	openssl bool
	done    chan error
}

//NewORM Implements DatabaseService; defaults to postgres client
func NewORM() ORMWrapper {
	return &orm{
		done: make(chan error),
	}
}

//Configurifier Configs for RDBMS connection.
func (s *orm) Configurifier(conf skelego.Config) {
	conf.DefaultSetting("storage.engine", "postgres")
	conf.DefaultSetting("storage.host", "localhost")
	conf.DefaultSetting("storage.host", 5432)
	conf.DefaultSetting("storage.username", "")
	conf.DefaultSetting("storage.password", "")
}

//Connect connects to database instance.
//TODO: add connection pool
func (s *orm) Connect(ctx context.Context, config skelego.Config, logger skelego.Logging) {
	//databaseType := config.Get("storage.engine").(string)
	host := config.Get("storage.host")
	port := config.Get("storage.port")
	database := config.Get("storage.name")
	username := config.Get("storage.username")
	password := config.Get("storage.password")

	logger.LogEvent("Connecting to %s", s.dbURI(host, port, username, database, password))

	db, err := gorm.Open("postgres", s.dbURI(host, port, username, database, password))
	if err != nil {
		print(err.Error())
	}
	logger.LogEvent("Successfully connected!", db)
	s.client = db
	if err := db.DB().Ping(); err != nil {
		logger.LogFatal("Error in database connection; restart app; error: %s", err.Error())
	}
}

func (s *orm) dbURI(databaseHost, port, username, databaseName, password interface{}) string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", databaseHost, port, username, databaseName, password)
}

//Start empty
func (s *orm) Start(ctx context.Context, logger skelego.Logging) {
	logger.LogEvent("Starting database...")
	return
}

//Shuts down database... to be run on server shutdown
func (s *orm) Stop(ctx context.Context, logger skelego.Logging) {
	s.shutdown(logger)
}

func (s *orm) shutdown(logger skelego.Logging) error {
	if err := s.client.Close(); err != nil {
		logger.LogError(err.Error())
		return err
	}
	logger.LogEvent("Shutting down database connection...")
	return nil
}

func (s *orm) alive() {
	s.done <- s.client.DB().Ping()
}

//TODO: add error handling
func (s *orm) Check(ctx context.Context) {
	s.client.DB().PingContext(ctx)
}

//DB returns database information
func (s *orm) DB() (*sql.DB, error) {
	return s.client.DB(), nil
}

//DB returns database information
func (s *orm) ORM() *gorm.DB {
	return s.client
}
