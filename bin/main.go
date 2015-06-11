package main

import (
	"flag"
	"fmt"
	"strings"

	fightsdb "github.com/colingagnon/weakgame-api/db/mysql/fights"
	monstersdb "github.com/colingagnon/weakgame-api/db/mysql/monsters"
	productsdb "github.com/colingagnon/weakgame-api/db/mysql/products"
	usersdb "github.com/colingagnon/weakgame-api/db/mysql/users"
	"github.com/colingagnon/weakgame-api/http/fights"
	"github.com/colingagnon/weakgame-api/http/monsters"
	"github.com/colingagnon/weakgame-api/http/products"
	"github.com/colingagnon/weakgame-api/http/users"

	"database/sql"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/howeyc/gopass"
	"github.com/colingagnon/weakgame-api/db"
	_ "github.com/ziutek/mymysql/godrv"
)

const (
	// Ugly way to check to see if they passed in a password
	// chance of collision with a GUID is very low
	defaultPassValue = "5e7dc6f6a1a94c39be95b88a47c2458b"
)

var (
	dbPort  string
	dbHost  string
	dbName  string
	dbUname string
	dbPass  string
)

func init() {
	flag.StringVar(&dbPort, "P", "3306", "port")
	flag.StringVar(&dbPass, "p", defaultPassValue, "password")
	flag.StringVar(&dbHost, "h", "localhost", "host")
	flag.StringVar(&dbName, "d", "", "database name")
	flag.StringVar(&dbUname, "u", "root", "username")
	flag.Parse()
}

func main() {
	pass := dbPass
	if pass == defaultPassValue {
		fmt.Print("Password:")
		pass = strings.TrimSpace(string(gopass.GetPasswdMasked()))
	}

	if strings.TrimSpace(dbPort) == "" {
		fmt.Println("Missing port")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if strings.TrimSpace(dbHost) == "" {
		fmt.Println("Missing host")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if strings.TrimSpace(dbName) == "" {
		fmt.Println("Missing database name")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if strings.TrimSpace(dbUname) == "" {
		fmt.Println("Missing username")
		flag.PrintDefaults()
		os.Exit(1)
	}

	dbConn, err := sql.Open("mymysql", fmt.Sprintf("tcp:%s:3306*%s/%s/%s", dbHost, dbName, dbUname, pass))
	if err != nil {
		panic(err)
	}
	db.MustValidateChecksum(dbConn, dbName)

	fightsdb.DB = dbConn

	monstersdb.DB = dbConn

	productsdb.DB = dbConn

	usersdb.DB = dbConn
//
	r := mux.NewRouter()
	g := r.Methods("GET").Subrouter()
	po := r.Methods("POST").Subrouter()
	pu := r.Methods("PUT").Subrouter()
	op := r.Methods("OPTIONS").Subrouter()
	//d := r.Methods("DELETE").Subrouter()

	// Disabled routes we don't want to expose
	op.HandleFunc("/fights/random", users.Pre)
	g.HandleFunc("/fights/random", fights.GetRandom)
	
	op.HandleFunc("/fights/round", users.Pre)
	g.HandleFunc("/fights/round", fights.GetRound)
	//g.HandleFunc("/fights", fights.List)
	//po.HandleFunc("/fights", fights.Post)

	//g.HandleFunc("/fights/{id}", fights.Get)
	//pu.HandleFunc("/fights/{id}", fights.Put)
	//d.HandleFunc("/fights/{id}", fights.Delete)

	//g.HandleFunc("/monsters", monsters.List)
	//po.HandleFunc("/monsters", monsters.Post)
	
	op.HandleFunc("/monsters/{id}", users.Pre)
	g.HandleFunc("/monsters/{id}", monsters.Get)
	//pu.HandleFunc("/monsters/{id}", monsters.Put)
	//d.HandleFunc("/monsters/{id}", monsters.Delete)
	
	op.HandleFunc("/products", users.Pre)
	g.HandleFunc("/products", products.List)
	//po.HandleFunc("/products", products.Post)

	pu.HandleFunc("/products/purchase/{id}", products.Purchase)
	op.HandleFunc("/products/purchase/{id}", users.Pre)
	g.HandleFunc("/products/{id}", products.Get)
	//pu.HandleFunc("/products/{id}", products.Put)
	//d.HandleFunc("/products/{id}", products.Delete)

	//g.HandleFunc("/users", users.List)
	g.HandleFunc("/users/login", users.GetLoginInfo)
	op.HandleFunc("/users/login", users.Pre)
	po.HandleFunc("/users/login", users.Login)
	
	g.HandleFunc("/users/revive", users.GetRevive)
	op.HandleFunc("/users/revive", users.Pre)
	po.HandleFunc("/users/revive", users.PostRevive)
	
	po.HandleFunc("/users", users.Post)
	
	op.HandleFunc("/users", users.Pre)
	g.HandleFunc("/users/{id}", users.Get)
	pu.HandleFunc("/users/{id}", users.Put)
	//d.HandleFunc("/users/{id}", users.Delete)

	//g.HandleFunc("/swagger.json", swaggerresponse)
	err = http.ListenAndServe(":8080", r)
	//err = http.ListenAndServe("weakgame.maxxjs.local:8080", r)
	//err = http.ListenAndServe("weakgame.maxxjs.com:8080", r)
	if err != nil {
		fmt.Println(err)
	}
	
}

// Token middleware
func AuthMiddleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("middleware", r.URL)
        
        h.ServeHTTP(w, r)
    })
}
