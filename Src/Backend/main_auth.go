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

	"github.com/gin-contrib/cors"
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
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"*"}
	// r.Use(cors.New(config))
	r.Use(cors.Default())
	r.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))
	r.GET("/", home)
	r.POST("/login", login)
	r.GET("/logout", logout)
	r.POST("/registerUser", registerUser)
	r.GET("/loginHTML", loginHTML)
	r.GET("/loginJS", loginJS)
	r.GET("/loginCSS", loginCSS)
	r.GET("/addItemHTML", addItemHTML)
	r.GET("/addItemJS", addItemJS)
	r.GET("/addItemCSS", addItemCSS)
	r.GET("/deleteItemHTML", deleteItemHTML)
	r.GET("/deleteItemJS", deleteItemJS)
	r.GET("/deleteItemCSS", deleteItemCSS)
	r.GET("/buyItemHTML", buyItemHTML)
	r.GET("/buyItemJS", buyItemJS)
	r.GET("/buyItemCSS", buyItemCSS)
	r.GET("/orderHTML", orderHTML)
	r.GET("/orderJS", orderJS)
	r.GET("/orderCSS", orderCSS)


	private := r.Group("/api")
	private.Use(AuthRequired)
	{
		private.GET("/statusUser", status)
		private.GET("/getItems", getItems)
		private.POST("/addItem", addItem)
		private.POST("/deleteItem", deleteItem)
		private.POST("/registerOrder", registerOrder)
		private.POST("/makeOrder", makeOrder)
		private.GET("/getOrders", getOrders)
		private.POST("/changePassword", changePassword)
		private.POST("/requestItem", requestItem)
		private.GET("/getRequests", getRequests)
		private.POST("/deleteRequest", deleteRequest)
	}
	return r
}

func AuthRequired(c *gin.Context) {
	// c.Header("Allow", "POST, GET, OPTIONS")
	// c.Header("Access-Control-Allow-Origin", "*")
	// c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
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
	// c.Header("Allow", "POST, GET, OPTIONS")
	// c.Header("Access-Control-Allow-Origin", "*")
	// c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	session := sessions.Default(c)
	// var userNew User
	var userOld User
	// c.Bind(&userNew)
	fmt.Println("-------")
	// fmt.Println(userNew.Username)
	// fmt.Println(userNew.Password)
	uid := c.PostForm("Username")
	pwd := c.PostForm("Password")
	fmt.Println(uid)
	fmt.Println(pwd)
	fmt.Println("-------")
	err := dbmap.SelectOne(&userOld, "SELECT * FROM Users WHERE Username=? AND Status='Active'", uid)
	fmt.Println(userOld)
	fmt.Println(err)
	if err == nil { //exist
		correctPassword := userOld.Password
		if comparePasswords(correctPassword, pwd) {
			session.Set(userkey, uid)
			if err := session.Save(); err != nil {
				fmt.Println("1")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
				// return
			} else {
				fmt.Println("2")
				c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
				// return
			}
		} else {
			fmt.Println("3")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
			// return
		}
	} else {
		fmt.Println("4")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Account is invalid"})
		// return
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

func addItem(c *gin.Context) {
	var itemNew Item
	var itemOld Item
	c.Bind(&itemNew)
	fmt.Println(itemNew.Name)
	err := dbmap.SelectOne(&itemOld, "SELECT * FROM Items WHERE Name=?", itemNew.Name)
	if err == nil { // exist
		if itemNew.Name != "" && itemNew.Amount != 0 {
			itemNew := Item{
				ID:     itemOld.ID,
				Name:   itemOld.Name,
				Price:  itemNew.Price,
				Cost:   itemNew.Cost,
				Amount: itemOld.Amount + itemNew.Amount,
			}
			update, _ := dbmap.Exec(`UPDATE Items SET Price=?, Cost=?, Amount=? WHERE ID=? AND Name=?`, itemNew.Price, itemNew.Cost, itemNew.Amount, itemOld.ID, itemOld.Name)
			if update != nil {
				c.JSON(200, itemNew)
			} else {
				checkErr(err, "Updated failed")
			}
		} else {
			c.JSON(400, gin.H{"error": "Fields are empty"})
		}
	} else { // non-exist
		if itemNew.Name != "" && itemNew.Amount != 0 {
			if insert, _ := dbmap.Exec(`INSERT INTO Items (Name, Price, Cost, Amount) VALUES (?, ?, ?, ?)`, itemNew.Name, itemNew.Price, itemNew.Cost, itemNew.Amount); insert != nil {
				item_id, err := insert.LastInsertId()
				if err == nil {
					itemNew := &Item{
						ID:     item_id,
						Name:   itemNew.Name,
						Price:  itemNew.Price,
						Cost:   itemNew.Cost,
						Amount: itemNew.Amount,
					}
					c.JSON(200, itemNew)
				} else {
					checkErr(err, "Add Item Error")
				}
			}
		} else {
			c.JSON(400, gin.H{"error": "Fields are empty"})
		}
	}
}

func deleteItem(c *gin.Context) {
	var itemNew Item
	var itemOld Item
	c.Bind(&itemNew)
	err := dbmap.SelectOne(&itemOld, "SELECT * FROM Items WHERE ID=?", itemNew.ID)
	if err == nil { // exist
		if itemNew.Amount > itemOld.Amount {
			itemOldAmount := fmt.Sprintf("The number of existing item is %d", itemOld.Amount)
			c.JSON(400, gin.H{"error": itemOldAmount})
		}
		if itemNew.Amount < itemOld.Amount {
			totalAmount := itemOld.Amount - itemNew.Amount
			dbmap.Exec(`UPDATE Items SET Price=?, Cost=?, Amount=? WHERE ID=?`, itemOld.Price, itemOld.Cost, totalAmount, itemOld.ID)
			c.JSON(200, "The item is deleted")
		}
		if itemNew.Amount == itemOld.Amount {
			dbmap.Exec(`DELETE FROM Items WHERE ID=?`, itemOld.ID)
			c.JSON(200, "The item is deleted")
		}
	} else { // non-exist
		c.JSON(400, gin.H{"error": "The item is not existing"})
	}
}

func registerOrder(c *gin.Context) {
	var ordNew Order
	var ordOld Order
	c.Bind(&ordNew)
	dbmap.SelectOne(&ordOld, "SELECT * FROM Orders WHERE OrderID=?", ordNew.OrderID)
	if len(ordNew.OrderID) == 0 { // non-exist
		ordID := tempPass()
		dbmap.Exec(`INSERT INTO Orders (OrderID, ItemID, Amount) VALUES (?, ?, ?)`, ordID, ordNew.ItemID, ordNew.Amount)
		c.JSON(200, ordID)
	} else { // exist
		dbmap.Exec(`INSERT INTO Orders (OrderID, ItemID, Amount) VALUES (?, ?, ?)`, ordNew.OrderID, ordNew.ItemID, ordNew.Amount)
		c.JSON(200, ordNew.OrderID)
	}
}

func makeOrder(c *gin.Context) {
	var ordNew Order
	var ordOld Order
	c.Bind(&ordNew)
	err := dbmap.SelectOne(&ordOld, `SELECT * FROM Orders WHERE OrderID=?`, ordNew.OrderID)
	if err != nil {
		c.String(http.StatusOK, "Make Order (OrderID:"+ordNew.OrderID+") failed")
	} else {
		dbmap.Exec(`DELETE FROM Orders WHERE OrderID=?`, ordNew.OrderID)
		c.String(http.StatusOK, "Make Order (OrderID:"+ordNew.OrderID+") success")
	}
}

func getOrders(c *gin.Context) {
	var ords []Order
	_, err := dbmap.Select(&ords, "SELECT * FROM Orders")
	if err == nil {
		c.JSON(200, ords)
	} else {
		c.JSON(404, gin.H{"error": "Get Orders Error"})
	}
}

func changePassword(c *gin.Context) {
	var userNew User
	var userOld User
	c.Bind(&userNew)
	err := dbmap.SelectOne(&userOld, "SELECT * FROM Users WHERE Email=?", userNew.Email)
	if err == nil { //exist
		pazz := hashAndSalt(userNew.Password)
		dbmap.Exec(`UPDATE Users SET Password=? WHERE Email=? AND Username=?`, pazz, userNew.Email, userNew.Username)
		c.JSON(200, gin.H{"error": "Your password is updated"})
	} else { //non-exist
		c.JSON(400, gin.H{"error": "Incorrect information"})
	}
}

func requestItem(c *gin.Context) {
	var reqNew Request
	var reqOld Request
	c.Bind(&reqNew)
	err := dbmap.SelectOne(&reqOld, "SELECT * FROM Requests WHERE Itemname=? AND Username=?", reqNew.Itemname, reqOld.Username)
	if err == nil { //exist
		totalAmount := reqOld.Amount + reqNew.Amount
		dbmap.Exec(`UPDATE Requests SET Amount=? WHERE Username=? AND Itemname=?`, totalAmount, reqNew.Username, reqNew.Itemname)
		c.JSON(200, gin.H{"success": "Your requested has already been added"})
	} else { //non-exist
		dbmap.Exec(`INSERT INTO Requests (Username, Itemname, Amount, Status) VALUES (?, ?, ?, ?)`, reqNew.Username, reqNew.Itemname, reqNew.Amount, "Added")
		c.JSON(200, gin.H{"success": "Your requested has already been added"})
	}
}

func getRequests(c *gin.Context) {
	var reqs []Request
	_, err := dbmap.Select(&reqs, "SELECT * FROM Requests")
	if err == nil {
		c.JSON(200, reqs)
	} else {
		c.JSON(404, gin.H{"error": "Get Requests Error"})
	}
}

func deleteRequest(c *gin.Context) {
	var reqNew Request
	var reqOld Request
	c.Bind(&reqNew)
	err := dbmap.SelectOne(&reqOld, "SELECT * FROM Requests WHERE ID=?", reqNew.ID)
	if err == nil { // exist
		dbmap.Exec(`UPDATE Requests SET Status=? WHERE ID=?`, "Deleted", reqNew.ID)
		c.JSON(200, gin.H{"success": "Your requested has already been deleted"})
	} else { // non-exist
		c.JSON(400, gin.H{"error": "The request is not existing"})
	}
}

func loginHTML(c *gin.Context) {
	c.File("../Frontend/Login/login_view.html")
}

func loginJS(c *gin.Context) {
	c.File("../Frontend/Login/login_management.js")
}

func loginCSS(c *gin.Context) {
	c.File("../Frontend/Login/login.css")
}

func addItemHTML(c *gin.Context) {
	c.File("../Frontend/AddItem/add_item_view.html")
}

func addItemJS(c *gin.Context) {
	c.File("../Frontend/AddItem/item_management.js")
}

func addItemCSS(c *gin.Context) {
	c.File("../Frontend/AddItem/add_item.css")
}

func deleteItemHTML(c *gin.Context) {
	c.File("../Frontend/DeleteItem/delete_item_view.html")
}

func deleteItemJS(c *gin.Context) {
	c.File("../Frontend/DeleteItem/item_management.js")
}

func deleteItemCSS(c *gin.Context) {
	c.File("../Frontend/DeleteItem/delete_item.css")
}

func buyItemHTML(c *gin.Context) {
	c.File("../Frontend/BuyItem/buy_item_view.html")
}

func buyItemJS(c *gin.Context) {
	c.File("../Frontend/BuyItem/item_management.js")
}

func buyitemCSS(c *gin.Context) {
	c.File("../Frontend/BuyItem/buy_item.css")
}

func orderHTML(c *gin.Context) {
	c.File("../Frontend/Order/order.html")
}

func orderJS(c *gin.Context) {
	c.File("../Frontend/Order/order.js")
}

func orderCSS(c *gin.Context) {
	c.File("../Frontend/Order/order.css")
}