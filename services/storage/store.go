package storage

import (
	"context"

	"github.com/Rahul12344/skelego"
)

//Store Interface for dealing with SQL or NoSQL database. Functionality wise, enforces basic CRUD implementation features, which may be
//constricting. Attempting to enforce returning of structs to force SQL queries to be wrapped into objects.
type Store interface {
	Check(context.Context) //Check for DB health
	skelego.Service
}

//Schema Interface for an object in a database.
type Schema interface {
	Migrate()
	TableName() string
}
