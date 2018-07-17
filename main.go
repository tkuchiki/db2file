package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func openDB(dbuser, dbpass, dbhost, dbname, socket string, port int) (*sql.DB, error) {
	userpass := fmt.Sprintf("%s:%s", dbuser, dbpass)
	var conn string
	if socket != "" {
		conn = fmt.Sprintf("unix(%s)", socket)
	} else {
		conn = fmt.Sprintf("tcp(%s:%d)", dbhost, port)
	}

	return sql.Open("mysql", fmt.Sprintf("%s@%s/%s", userpass, conn, dbname))
}

func tmpDir(dir string) string {
	return filepath.Join(os.TempDir(), dir)
}

func fpath(dir, filename string) string {
	return filepath.Join(dir, filename)
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func main() {
	var app = kingpin.New("db2file", "Database dump to file")

	var dbuser = app.Flag("dbuser", "Database user").Default("root").String()
	var dbpass = app.Flag("dbpass", "Database password").String()
	var dbhost = app.Flag("dbhost", "Database host").Default("localhost").String()
	var dbport = app.Flag("dbport", "Database port").Default("3306").Int()
	var dbsock = app.Flag("dbsock", "Database socket").String()
	var dbname = app.Flag("dbname", "Database name").Required().String()
	var query = app.Flag("query", "SQL").Required().String()
	var dump = app.Flag("dump", "Dump file from database column").Required().String()
	var filename = app.Flag("filename", "filename column").Required().String()
	var outDir = app.Flag("out-dir", "Output directory").Default(tmpDir("db2file")).PlaceHolder("$TMPDIR/db2file").String()
	var overwrite = app.Flag("overwrite", "Overwrite file same filename").Bool()

	app.Version("0.1.0")

	kingpin.MustParse(app.Parse(os.Args[1:]))

	db, err := openDB(*dbuser, *dbpass, *dbhost, *dbname, *dbsock, *dbport)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(*query)
	if err != nil {
		log.Fatal(err)
	}
	cols, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	colNames := make(map[string]struct{})
	for _, col := range cols {
		colNames[col] = struct{}{}
	}

	if _, ok := colNames[*dump]; !ok {
		log.Fatal("--dump-column is required")
	}

	if _, ok := colNames[*filename]; !ok {
		log.Fatal("--filename is required")
	}

	b := make([][]byte, len(cols))

	row := make([]interface{}, len(cols))
	for i, _ := range b {
		row[i] = &b[i]
	}

	if !exists(*outDir) {
		if err := os.MkdirAll(*outDir, 0755); err != nil {
			log.Fatal(err)
		}
	}

	for rows.Next() {
		if err := rows.Scan(row...); err != nil {
			log.Fatal(err)
		}

		values := make(map[string][]byte)
		for i, val := range b {
			values[cols[i]] = val
		}

		dumpFile := fpath(*outDir, string(values[*filename]))

		if !*overwrite && exists(dumpFile) {
			log.Println(fmt.Sprintf("[skip] %s already exists", dumpFile))
			continue
		}

		fp, err := os.OpenFile(dumpFile, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer fp.Close()

		buf := bytes.NewReader(values[*dump])
		_, err = io.Copy(fp, buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(fmt.Sprintf("[dump] %s", dumpFile))
	}
}
