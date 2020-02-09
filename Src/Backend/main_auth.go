package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gorp.v1"
)

type Request struct {
	ID       int64  `db:"ID" json:"ID"`
	Username string `db:Username json:"Username"`
	Itemname string `db:"Itemname" json:"Itemname"`
	Amount   int64  `db:"Amount" json:"Amount"`
	Status   string `db:"Status" json:"Status"`
}

type Order struct {
	ID      int64  `db:"ID" json:"ID"`
	OrderID string `db:"OrderID" json:"OrderID"`
	ItemID  string `db:"ItemID" json:"ItemID"`
	Amount  int64  `db:"Amount" json:"Amount"`
}

type Item struct {
	ID     int64   `db:"ID" json:"ID"`
	Name   string  `db:"Name" json:"Name"`
	Price  float64 `db:"Price" json:"Price"`
	Cost   float64 `db:"Cost" json:"Cost"`
	Amount int64   `db:"Amount" json:"Amount"`
}

type User struct {
	ID           int64  `db:"ID" json:"ID"`
	Username     string `db:"Username" json:"Username"`
	Password     string `db:"Password" json:"Password"`
	TempPassword string `db:"TempPassword" json:"TempPassword"`
	Email        string `db:"Email" json:"Email"`
	Firstname    string `db:"Firstname" json:"Firstname"`
	Lastname     string `db:"Lastname" json:"Lastname"`
	Role         string `db:"Role" json:"Role"`
	Status       string `db:"Status" json:"Status"`
}

type ConfigurationDB struct {
	Username string
	Password string
	Host     string
	Port     string
	DB_name  string
}

const (
	userkey = "user"
)

func tempPass() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
}

func comparePasswords(hashedPwd string, plainPwd string) bool {
	plainPwdZ := []byte(plainPwd)
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwdZ)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

var dbmap = initDb()

func initDb() *gorp.DbMap {
	fmt.Println("Establishing . . .")
	file, _ := os.Open("config_db.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	Cdb := ConfigurationDB{}
	err := decoder.Decode(&Cdb)
	addr := Cdb.Username + ":" + Cdb.Password + "@tcp(" + Cdb.Host + ":" + Cdb.Port + ")/" + Cdb.DB_name
	db, err := sql.Open("mysql", addr)
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	return dbmap
}

func hashAndSalt(pwd string) string {
	vec := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(vec, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	r := engine()
	r.Use(gin.Logger())
	if err := engine().Run(":13131"); err != nil {
		log.Fatal("Unable to start:", err)
	}
}

func engine() *gin.Engine {
	r := gin.New()
	r.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))
	r.GET("/", home)
	r.POST("/login", login)
	r.GET("/logout", logout)
	r.POST("/registerUser", registerUser)

	private := r.Group("/api")
	private.Use(AuthRequired)
	{
		private.GET("/statusUser", status)
		private.GET("/getItems", getItems)
		// private.POST("/addItem", addItem)
		// private.POST("/deleteItem", deleteItem)
		// private.POST("/registerOrder", registerOrder)
		// private.POST("/makeOrder", makeOrder)
		// private.GET("/getOrders", getOrders)
		// private.POST("/changePassword", changePassword)
		// private.POST("/requestItem", requestItem)
		// private.GET("/getRequests", getRequests)
		// private.POST("/deleteRequest", deleteRequest)
	}
	return r
}

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	c.Next()
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Merchandise manager system API"})
}

func login(c *gin.Context) {
	session := sessions.Default(c)
	var userNew User
	var userOld User
	c.Bind(&userNew)
	err := dbmap.SelectOne(&userOld, "SELECT * FROM Users WHERE Username=? AND Status='Active'", userNew.Username)
	fmt.Println(err)
	if err == nil { //exist
		correctPassword := userOld.Password
		if comparePasswords(correctPassword, userNew.Password) {
			session.Set(userkey, userNew.Username)
			if err := session.Save(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
				return
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
			return
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Account is invalid"})
		return
	}
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete(userkey)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func status(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in", "user": user})
}

func registerUser(c *gin.Context) {
	var userNew User
	var userOld User
	c.Bind(&userNew)
	err := dbmap.SelectOne(&userOld, "SELECT * FROM Users WHERE Username=?", userNew.Username)
	if err == nil || userOld.Username == userNew.Username || userOld.Email == userNew.Email { //exist
		c.JSON(200, gin.H{"error": "This username is already existing"})
	} else { //non-exist
		tempP := tempPass()
		pazz := hashAndSalt(userNew.Password)
		dbmap.Exec(`INSERT INTO Users (Username, Password, TempPassword, Email, Firstname, Lastname, Role) VALUES (?, ?, ?, ?, ?, ?, ?)`, userNew.Username, pazz, tempP, userNew.Email, userNew.Firstname, userNew.Lastname, userNew.Role)
		fmt.Println("python3 py_email.py " + userNew.Username + " " + userNew.Email + " " + tempP)
		exec.Command("python3", "py_email.py", userNew.Username, userNew.Email, tempP).Run()
		c.JSON(200, gin.H{"success": "Register success"})
	}
}

func getItems(c *gin.Context) {
	var items []Item
	_, err := dbmap.Select(&items, "SELECT * FROM Items")
	if err == nil {
		c.JSON(200, items)
	} else {
		c.JSON(404, gin.H{"error": "Get Items Error"})
	}
}
