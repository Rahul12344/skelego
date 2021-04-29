package nosql

import (
	"context"

	"github.com/Rahul12344/skelego"
	"github.com/Rahul12344/skelego/services/storage"
)

//DatabaseService Interface dealing with NoSQL
type DatabaseService interface {
	storage.Store
	DB() error
	Port(...storage.Schema)
}

type db struct {
}

//New Creates a new service for dealing with NoSQL.
func New() DatabaseService {
	return &db{}
}

//Configurifier Configs for NoSQL DBM connection.
func (s *db) Configurifier(conf skelego.Config) {
	conf.DefaultSetting("nosql_database_name", "mongodb")
	conf.DefaultSetting("nosql_database_port", "27017")
	conf.DefaultSetting("nosql_database_username", "")
	conf.DefaultSetting("nosql_database_password", "")
}

func (s *db) Connect(ctx context.Context, config skelego.Config) {

}

func (s *db) Start(ctx context.Context) {

}

func (s *db) Stop(ctx context.Context) {

}

func (s *db) Check(ctx context.Context) {

}

//Migrate migrates tables into database
func (s *db) Port(documents ...storage.Schema) {
	for _, document := range documents {
		document.Migrate()
	}
}

func (s *db) Create(rows ...storage.Schema) error {
	return nil
}

//Simple SELECT... FROM ... WHERE... statement
func (s *db) Read(output storage.Schema, query interface{}, args ...interface{}) error {
	return nil
}

//Simple update of schema
func (s *db) Update(mod storage.Schema) error {
	return nil
}

//Simple DELETE... FROM... WHERE... statement
func (s *db) Delete(rows ...storage.Schema) error {
	return nil
}

//DB returns database information
func (s *db) DB() error {
	return nil
}