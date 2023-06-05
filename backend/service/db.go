package service

import (
	"context"

	pgx "github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

// DbIFace is an interface for the database
type DbIFace interface {
	GetConn() *pgx.Conn
  NewApp(string, string) (string, error)
	GetApplications() (pgx.Rows, error)
	UserApps() (pgx.Rows, error)
	GetApplication(string) (ApplicationExtended, error)
	UpdateApplication(string, ApplicationSimple) error
	DeleteApplication(string) error
}

// Db is a struct that contains a pointer to the database connection
type Db struct {
	db *pgx.Conn
}

type ApplicationSimple struct {
	Id          string `json:"id" db:"id"`
	Name        string `json:"name" db:"app_name"`
	Description string `json:"description" db:"description"`
	Owners      string `json:"owners" db:"owners"`
	Team        string `json:"teamName" db:"team_name"`
}

type ApplicationExtended struct {
	Id          string `json:"id" db:"id"`
	Name        string `json:"name" db:"app_name"`
	Description string `json:"description" db:"description"`
	Owners      string `json:"owners" db:"owners"`
	Team        string `json:"teamName" db:"team_name"`
	Api_key     string `json:"apiKey" db:"api_key"`
	Is_valid    bool   `json:"isValid" db:"is_valid"`
}

func initDb(databaseURL string, log LoggerIFace) *Db {
	ctx := context.Background()
	var err error
	conn, err = pgx.Connect(ctx, databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	err = prepareStatements(conn, ctx)
	if err != nil {
		log.Fatal(err)
	}

	return &Db{db: conn}
}

func prepareStatements(connection *pgx.Conn, ctx context.Context) (err error) {
	_, err = connection.Prepare(ctx, "insert_into_applications", "INSERT INTO applications (name, description, is_valid) VALUES ($1, $2, true) RETURNING id")
	_, err = connection.Prepare(ctx, "select_all_applications", "SELECT * FROM applications")
	if err != nil {
		return err
	}
	return nil
}

func (s *Db) GetConn() *pgx.Conn {
	return s.db
}

// NewApp stores the information about a new application in the database.

func (db *Db) NewApp(name string, desc string) (string, error) {
	ctx := context.Background()

	var id string

	err := db.db.QueryRow(ctx, "insert_into_applications", name, desc).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

// UserApps returns a list of all the applications that exist right now.
// Once accessing the users email is figured out, this function will return
// a list of all the applications that the user owns.
func (db *Db) GetApplications() (pgx.Rows, error) {
	ctx := context.Background()
	rows, err := db.db.Query(ctx, "SELECT * FROM applications")
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (db *Db) GetApplication(applicationId string) (ApplicationExtended, error) {
	row, _ := conn.Query(context.Background(), "SELECT * FROM applications WHERE id=$1", applicationId)

	var app ApplicationExtended

	for row.Next() {
		err := row.Scan(&app.Id,
			&app.Name,
			&app.Description,
			&app.Owners,
			// &app.Team,
			&app.Api_key,
			&app.Is_valid)
		if err != nil {
			return app, err
		}
	}

	return app, nil
}

func (db *Db) UpdateApplication(applicationId string, applicationInfo ApplicationSimple) error {
	_, err := conn.Exec(context.Background(), "UPDATE applications SET name=$1, description=$2, owners=$3 WHERE id=$4",
		applicationInfo.Name,
		applicationInfo.Description,
		applicationInfo.Owners,
		applicationId)
	return err
}

func (db *Db) DeleteApplication(applicationId string) error {
	_, err := conn.Exec(context.Background(), "DELETE FROM applications WHERE id=$1", applicationId)
	return err
}

// Checks for duplicated app IDs or api keys
// I easily could've used if statements but I've always wanted to use switch statements
// and realistically this could expand to more columns in the future.
func (db *Db) CheckDuplicate(column string, newGen string) (bool, error) {
	ctx := context.Background()
	var query string

	// This switch statement is used to determine which column to check for duplicates
	switch column {
	case "id":
		query = "SELECT * FROM applications WHERE id = $1"
	case "api_key":
		query = "SELECT * FROM applications WHERE api_key = $1"
	}

	// Querying the database for duplicates
	err := db.db.QueryRow(ctx, query, newGen).Scan()
	if err == pgx.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}
