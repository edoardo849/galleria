package storage

//https://github.com/GoogleCloudPlatform/golang-samples/blob/master/getting-started/bookshelf/db_mysql.go
import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var createTableStatements = []string{
	`CREATE DATABASE IF NOT EXISTS gallery DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci';`,
	`USE gallery;`,
	`CREATE TABLE IF NOT EXISTS images (
		id CHAR(40) NOT NULL,
		url VARCHAR(255) NOT NULL,
		description TEXT NULL,
		title VARCHAR(255) NULL,
		format CHAR(3) NOT NULL,
		PRIMARY KEY (id)
	)`,
}

// mysqlDB persists books to a MySQL instance.
type mysqlDB struct {
	conn *sql.DB

	list   *sql.Stmt
	listBy *sql.Stmt
	insert *sql.Stmt
	get    *sql.Stmt
	update *sql.Stmt
	delete *sql.Stmt
}

// Ensure mysqlDB conforms to the ImageDatabase interface.
var _ ImageDatabase = &mysqlDB{}

// MySQLConfig is the configuration for a mySQL db
type MySQLConfig struct {
	// Optional.
	Username, Password string

	// Host of the MySQL instance.
	//
	// If set, UnixSocket should be unset.
	Host string

	// Port of the MySQL instance.
	//
	// If set, UnixSocket should be unset.
	Port int

	// UnixSocket is the filepath to a unix socket.
	//
	// If set, Host and Port should be unset.
	UnixSocket string
}

// dataStoreName returns a connection string suitable for sql.Open.
func (c MySQLConfig) dataStoreName(databaseName string) string {
	var cred string
	// [username[:password]@]
	if c.Username != "" {
		cred = c.Username
		if c.Password != "" {
			cred = cred + ":" + c.Password
		}
		cred = cred + "@"
	}

	if c.UnixSocket != "" {
		return fmt.Sprintf("%sunix(%s)/%s", cred, c.UnixSocket, databaseName)
	}
	return fmt.Sprintf("%stcp([%s]:%d)/%s", cred, c.Host, c.Port, databaseName)
}

// newMySQLDB creates a new ImageDatabase backed by a given MySQL server.
func newMySQLDB(config MySQLConfig) (ImageDatabase, error) {
	// Check database and table exists. If not, create it.
	if err := config.ensureTableExists(); err != nil {
		return nil, err
	}

	conn, err := sql.Open("mysql", config.dataStoreName("gallery"))
	if err != nil {
		return nil, fmt.Errorf("mysql: could not get a connection: %v", err)
	}
	if err := conn.Ping(); err != nil {
		conn.Close()
		return nil, fmt.Errorf("mysql: could not establish a good connection: %v", err)
	}

	db := &mysqlDB{
		conn: conn,
	}

	// Prepared statements. The actual SQL queries are in the code near the
	// relevant method (e.g. addImage).
	if db.list, err = conn.Prepare(listStatement); err != nil {
		return nil, fmt.Errorf("mysql: prepare list: %v", err)
	}

	if db.get, err = conn.Prepare(getStatement); err != nil {
		return nil, fmt.Errorf("mysql: prepare get: %v", err)
	}
	if db.insert, err = conn.Prepare(insertStatement); err != nil {
		return nil, fmt.Errorf("mysql: prepare insert: %v", err)
	}

	if db.delete, err = conn.Prepare(deleteStatement); err != nil {
		return nil, fmt.Errorf("mysql: prepare delete: %v", err)
	}

	return db, nil
}

// Close closes the database, freeing up any resources.
func (db *mysqlDB) Close() {
	db.conn.Close()
}

// rowScanner is implemented by sql.Row and sql.Rows
type rowScanner interface {
	Scan(dest ...interface{}) error
}

// scanImage reads an image from a sql.Row or sql.Rows
func scanImage(s rowScanner) (*Image, error) {
	var (
		id          sql.NullString
		url         sql.NullString
		description sql.NullString
		title       sql.NullString
		format      sql.NullString
	)
	if err := s.Scan(&id, &url, &description, &title, &format); err != nil {
		return nil, err
	}

	image := &Image{
		ID:          id.String,
		URL:         url.String,
		Description: description.String,
		Title:       title.String,
		Format:      format.String,
	}
	return image, nil
}

const listStatement = `SELECT * FROM images ORDER BY title`

// ListImages returns a list of images, ordered by title.
func (db *mysqlDB) ListImages() ([]*Image, error) {
	rows, err := db.list.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var images []*Image
	for rows.Next() {
		image, err := scanImage(rows)
		if err != nil {
			return nil, fmt.Errorf("mysql: could not read row: %v", err)
		}

		images = append(images, image)
	}

	return images, nil
}

const getStatement = "SELECT * FROM images WHERE id = ?"

// GetImage retrieves an image by its ID.
func (db *mysqlDB) GetImage(id string) (*Image, error) {
	image, err := scanImage(db.get.QueryRow(id))
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("mysql: could not find book with id %s", id)
	}
	if err != nil {
		return nil, fmt.Errorf("mysql: could not get book: %v", err)
	}
	return image, nil
}

const insertStatement = `
  INSERT INTO images (
    id, url, description, title, format
  ) VALUES (?, ?, ?, ?, ?, ?, ?)`

// AddImage saves a given image, assigning it a new ID.
func (db *mysqlDB) AddImage(i *Image) (id string, err error) {
	_, err = execAffectingOneRow(db.insert, i.ID, i.URL, i.Description,
		i.Title, i.Format)
	if err != nil {
		return "", err
	}

	return i.ID, nil
}

const deleteStatement = `DELETE FROM images WHERE id = ?`

// DeleteBook removes a given book by its ID.
func (db *mysqlDB) DeleteImage(id string) error {
	if id == "" {
		return errors.New("mysql: image with unassigned ID passed into deleteBook")
	}
	_, err := execAffectingOneRow(db.delete, id)
	return err
}

// ensureTableExists checks the table exists. If not, it creates it.
func (c MySQLConfig) ensureTableExists() error {
	conn, err := sql.Open("mysql", c.dataStoreName(""))
	if err != nil {
		return fmt.Errorf("mysql: could not get a connection: %v", err)
	}
	defer conn.Close()

	// Check the connection.
	if conn.Ping() == driver.ErrBadConn {
		return fmt.Errorf("mysql: could not connect to the database. " +
			"could be bad address, or this address is not whitelisted for access.")
	}

	if _, err := conn.Exec("USE gallery"); err != nil {
		// MySQL error 1049 is "database does not exist"
		if mErr, ok := err.(*mysql.MySQLError); ok && mErr.Number == 1049 {
			return createTable(conn)
		}
	}

	if _, err := conn.Exec("DESCRIBE images"); err != nil {
		// MySQL error 1146 is "table does not exist"
		if mErr, ok := err.(*mysql.MySQLError); ok && mErr.Number == 1146 {
			return createTable(conn)
		}
		// Unknown error.
		return fmt.Errorf("mysql: could not connect to the database: %v", err)
	}
	return nil
}

// createTable creates the table, and if necessary, the database.
func createTable(conn *sql.DB) error {
	for _, stmt := range createTableStatements {
		_, err := conn.Exec(stmt)
		if err != nil {
			return err
		}
	}
	return nil
}

// execAffectingOneRow executes a given statement, expecting one row to be affected.
func execAffectingOneRow(stmt *sql.Stmt, args ...interface{}) (sql.Result, error) {
	r, err := stmt.Exec(args...)
	if err != nil {
		return r, fmt.Errorf("mysql: could not execute statement: %v", err)
	}
	rowsAffected, err := r.RowsAffected()
	if err != nil {
		return r, fmt.Errorf("mysql: could not get rows affected: %v", err)
	} else if rowsAffected != 1 {
		return r, fmt.Errorf("mysql: expected 1 row affected, got %d", rowsAffected)
	}
	return r, nil
}
