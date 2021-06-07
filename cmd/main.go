package main

import (
	"database/sql"
	"github.com/SonnLarissa/gosql/cmd/server/app"
	"github.com/SonnLarissa/gosql/pkg/customers"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	host := "0.0.0.0"
	port := "9999"
	dsn := "postgres://app:pass@localhost:5432/db"

	if err := execute(host, port, dsn); err != nil {
		log.Println(err)
		os.Exit(1)
	}

}

func execute(host string, port string, dsn string) (err error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := db.Close(); cerr != nil {
			if err == nil {
				err = cerr
				return
			}
			log.Println(err)
		}
	}()

	mux := http.NewServeMux()
	customerSvg := customers.NewService(db)
	server := app.NewServer(mux, customerSvg)
	server.Init()

	srv := &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: server,
	}
	return srv.ListenAndServe()
}

//type Customer struct {
//	ID      int64
//	Name    string
//	Phone   string
//	Active  bool
//	Created time.Time
//}
//
//func main() {
//	dsn := "postgres://app:pass@localhost:5432/db"
//	db, err := sql.Open("pgx", dsn)
//	if err != nil {
//		log.Println(err)
//		os.Exit(1)
//		return
//	}
//	defer func() {
//		if err := db.Close(); err != nil {
//			log.Println(err)
//		}
//	}()
//
//	ctx := context.Background()
//	_, err = db.ExecContext(ctx, `
//CREATE TABLE IF NOT EXISTS  customers(
//    id BIGSERIAL PRIMARY KEY,
//    name TEXT NOT NULL,
//    phone TEXT NOT NULL UNIQUE ,
//    active BOOLEAN NOT NULL DEFAULT TRUE,
//    created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
//)`)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	customer:=&Customer{}
//	err=db.QueryRowContext(ctx,`
//SELECT id, name, phone,active, created FROM customers WHERE id=1`).Scan(&customer.ID,&customer.Name,&customer.Phone, &customer.Active, &customer.Created)
//	id:=1
//	newPhone:="+992000000099"
//	err=db.QueryRowContext(ctx,`
//UPDATE customers SET phone=$2 WHERE id=$1 RETURNING id, name, phone, active, created`,
//id, newPhone).Scan(&customer.ID, &customer.Name, &customer.Phone, &customer.Active, &customer.Created)
//	if errors.Is(err, sql.ErrNoRows){
//		log.Println("NO rows")
//		return
//	}
//if err!=nil{
//		log.Println(err)
//		return
//	}
//	log.Printf("%#v", customer)
//	log.Print(customer)

//	name := "Petya"
//	phone := "992000000001"
//	result, err := db.ExecContext(ctx, `
//INSERT INTO customers(name, phone) VALUES ($1, $2) ON CONFLICT (phone) DO UPDATE SET name =excluded.name;`, name, phone)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	log.Println(result.RowsAffected())
//	log.Println(result.LastInsertId())
//}

/*package main

import (
	"github.com/Shahlojon/crud/cmd/app"
	"github.com/Shahlojon/crud/pkg/customers"
	"net/http"
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"os"
)

func main() {
	host :="0.0.0.0"
	port := "9999"
	dbConnectionString :="postgres://app:pass@localhost:5432/db"
	if err := execute(host, port, dbConnectionString); err != nil{
		log.Print(err)
		os.Exit(1)
	}
}

func execute(host, port, dbConnectionString string) (err error){
	db, err := sql.Open("pgx", dbConnectionString)
	if err !=nil{
		return err
	}
	defer db.Close()

	mux := http.NewServeMux()
	customerService := customers.NewService(db)
	server := app.NewServer(mux, customerService)
	server.Init()

	httpServer := &http.Server{
		Addr:host+":"+port,
		Handler: server,
	}

	return httpServer.ListenAndServe()
}*/
