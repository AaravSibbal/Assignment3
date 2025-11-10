package server

import (
	"context"
	"fmt"
	"database/sql"
	// "fmt"
	"log"
	"net/http"
	"os"
	"time"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

/*
	application struct for things we need to a lot of times
*/
type application struct {
	ctx            *context.Context
	db             *sql.DB
	infoLog        *log.Logger
	errorLog       *log.Logger
}

func Run() {
	/*
		reading the .env file
	*/
	envFile, err := godotenv.Read(".env")
	if err != nil {
		log.Fatalf("Couldn't read the .env file %v", err)
	}

	/*
		getting the db connection
	*/
	db, err := getDBConnection(envFile)
	if err != nil {
		log.Fatalf("could not connect to the DB: %v", err)
	}

	/*
		setting up my loggers
	*/
	infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	errLogger := log.New(os.Stdout, "ERROR:\t", log.Ltime|log.Ldate|log.Lshortfile)

	ctx := context.Background()
	app := &application{
		ctx: &ctx,
		db: db,
		infoLog: infoLog,
		errorLog: errLogger,

	}
	addr := fmt.Sprintf("%s:%s",envFile["ADDRESS"], envFile["PORT"])
	/*
		starting my server
	*/
	srv := &http.Server{
		Addr:         addr,
		ErrorLog:     errLogger,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      app.Routes(),
	}

	err = srv.ListenAndServe()

	log.Fatal(err)
	
}


/*
	getting the db connection
*/
func getDBConnection(envFile map[string]string) (*sql.DB, error){
	
	// reading all the relevent parameters for the connection
	host := envFile["POSTGRES_HOST"]
	postPort := envFile["POSTGRES_PORT"]
	user := envFile["POSTGRES_USER"]
	password := envFile["POSTGRES_PASSWORD"]
	dbName := envFile["POSTGRES_NAME"]

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, postPort, user, password, dbName)

	fmt.Println(psqlInfo)

	/*
		opening the connection to the db
	*/
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalln("there was an error getting the db connection", err)
		return nil, err
	}

	/*
		pinging it just to be sure
	*/
	if err = db.Ping(); err != nil {
		log.Fatalln("we couldn't ping the db for some reason", err)
	}

	fmt.Println("db was connected successfuly")
	// success 
	return db, nil
}
