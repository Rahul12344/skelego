package sqlservice

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Rahul12344/skelego"

	"github.com/Rahul12344/skelego/services/storage"

	// postgres gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//DatabaseService Interface for SQL database.
type DatabaseService interface {
	storage.Store
	DB() (*sql.DB, error)
}

//db example postgres interface
type db struct {
	client  *sql.DB
	openssl bool
	done    chan error
}

//New Implements DatabaseService; defaults to postgres client
func New() DatabaseService {
	return &db{
		done: make(chan error),
	}
}

//Configurifier Configs for RDBMS connection.
func (s *db) Configurifier(conf skelego.Config) {
	conf.DefaultSetting("storage.engine", "postgres")
	conf.DefaultSetting("storage.host", "5432")
	conf.DefaultSetting("storage.username", "")
	conf.DefaultSetting("storage.password", "")
}

//Connect connects to database instance.
//TODO: add connection pool
func (s *db) Connect(ctx context.Context, config skelego.Config, logger skelego.Logging) {
	//databaseType := config.Get("storage.engine").(string)
	port := config.Get("storage.host")
	database := config.Get("storage.name")
	username := config.Get("storage.username")
	password := config.Get("storage.password")

	db, err := sql.Open("postgres", s.dbURI(port, username, database, password))
	if err != nil {
		print(err.Error())
	}
	logger.LogEvent("Successfully connected!", db)
	s.client = db
	if err := db.Ping(); err != nil {
		logger.LogFatal("Error in database connection; restart app; error: %s", err.Error())
	}
}

func (s *db) dbURI(databaseHost, username, databaseName, password interface{}) string {
	return fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", databaseHost, username, databaseName, password)
}

//Start empty
func (s *db) Start(ctx context.Context, logger skelego.Logging) {
	logger.LogEvent("Starting database...")
	return
}

//Shuts down database... to be run on server shutdown
func (s *db) Stop(ctx context.Context, logger skelego.Logging) {
	s.shutdown(logger)
}

func (s *db) shutdown(logger skelego.Logging) error {
	if err := s.client.Close(); err != nil {
		logger.LogError(err.Error())
		return err
	}
	logger.LogEvent("Shutting down database connection...")
	return nil
}

func (s *db) alive() {
	s.done <- s.client.Ping()
}

//TODO: add error handling
func (s *db) Check(ctx context.Context) {
	s.client.PingContext(ctx)
}

//DB returns database information
func (s *db) DB() (*sql.DB, error) {
	return s.client, nil
}
