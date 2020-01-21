package main

import (	
	"database/sql"
	"gopkg.in/gorp.v1"
	"encoding/json"
	"os"
	"fmt"
	"net/http"
	"log"
	"math/rand"
	"time"
	"strings"
	"os/exec"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
    _ "github.com/go-sql-driver/mysql"
)

type Item struct {
	ID    int64`db:"ID" json:"ID"`
	Name string	`db:"Name" json:"Name"`
	Price float64 `db:"Price" json:"Price"`
	Cost float64 `db:"Cost" json:"Cost"`
	Amount int64 `db:"Amount" json:"Amount"`
}

type User struct {
	ID    int64`db:"ID" json:"ID"`
 	Username string	`db:"Username" json:"Username"`
	Password string `db:"Password" json:"Password"`
	TempPassword string `db:"TempPassword" json:"TempPassword"`
	Email string `db:"Email" json:"Email"`
	Firstname string `db:"Firstname" json:"Firstname"`
	Lastname string `db:"Lastname" json:"Lastname"`
	Role	string `db:"Role" json:"Role"`
}

type ConfigurationDB struct {
    Username    string
	Password    string
	Host 		string
	Port		string
	DB_name		string
}

func tempPass() string{
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
	addr := Cdb.Username+":"+Cdb.Password+"@tcp("+Cdb.Host+":"+Cdb.Port+")/"+Cdb.DB_name
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

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context){
		c.String(http.StatusOK,"Merchandise Manager Serving . . .")
	})

	//01 '/requestItemList'
	router.GET("/getItems", func(c *gin.Context){
		var items []Item
	    _, err := dbmap.Select(&items, "SELECT * FROM Items")
	    if err == nil {
		   c.JSON(200, items)
	    } else {
		   c.JSON(404, gin.H{"error": "Get Items Error"})
	    }
	})

	//02
	router.POST("/addItem", func(c *gin.Context){
		var itemNew Item
		var itemOld Item
		c.Bind(&itemNew)
		err := dbmap.SelectOne(&itemOld, "SELECT * FROM Items WHERE Name=?", itemNew.Name)
		if err == nil{ // exist
			if itemNew.Name != "" && itemNew.Amount != 0 {
				itemNew := Item{
					ID:        itemOld.ID,
					Name: 	   itemOld.Name,
					Price:     itemNew.Price,
					Cost:      itemNew.Cost,
					Amount:    itemOld.Amount+itemNew.Amount,
				}
				update, _ := dbmap.Exec(`UPDATE Items SET Price=?, Cost=?, Amount=? WHERE ID=? AND Name=?`,itemNew.Price, itemNew.Cost, itemNew.Amount, itemOld.ID, itemOld.Name); 
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
					        ID:        item_id,
							Name: 	   itemNew.Name,
							Price:     itemNew.Price,
							Cost:      itemNew.Cost,
							Amount:    itemNew.Amount,
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
	})

	//03
	router.POST("/deleteItem", func(c *gin.Context){
		// c.String(http.StatusOK,"Delete Item")
		var itemNew Item
		var itemOld Item
		c.Bind(&itemNew)
		err := dbmap.SelectOne(&itemOld, "SELECT * FROM Items WHERE Name=?", itemNew.Name)
		if err == nil{ // exist
			if itemNew.Amount > itemOld.Amount {
				itemOldAmount := fmt.Sprintf("The number of existing item is %d",itemOld.Amount)
				c.JSON(400, gin.H{"error": itemOldAmount})
			}
			if itemNew.Amount < itemOld.Amount {
				itemNew := Item{
					ID:        itemOld.ID,
					Name: 	   itemOld.Name,
					Price:     itemOld.Price,
					Cost:      itemOld.Cost,
					Amount:    itemOld.Amount-itemNew.Amount,
				}
				delete, _ := dbmap.Exec(`UPDATE Items SET Price=?, Cost=?, Amount=? WHERE ID=? AND Name=?`,itemNew.Price, itemNew.Cost, itemNew.Amount, itemOld.ID, itemOld.Name); 
				if delete != nil {
					c.JSON(200, itemNew)
				} else {
					checkErr(err, "Deleted failed")
				}
			}
			if itemNew.Amount == itemOld.Amount {
				delete, _ := dbmap.Exec(`DELETE FROM Items WHERE ID=? AND Name=?`, itemOld.ID, itemOld.Name); 
				if delete != nil {
					c.JSON(200, "The item is deleted")
				} else {
					checkErr(err, "Deleted failed")
				}
			}
		} else { // non-exist
			c.JSON(400, gin.H{"error": "The item is not existing"})
		}
	})

	//04
	router.POST("/registerOrder", func(c *gin.Context){
		c.String(http.StatusOK,"Register Order")
	})

	//05
	router.POST("/makeOrder", func(c *gin.Context){
		c.String(http.StatusOK,"Make Order")
	})

	//06
	router.GET("/checkOrder", func(c *gin.Context){
		c.String(http.StatusOK,"Check Order")
	})

	//07
	router.POST("/login", func(c *gin.Context){
		var userNew User
		var userOld User
		c.Bind(&userNew)
		err := dbmap.SelectOne(&userOld, "SELECT * FROM Users WHERE Username=?", userNew.Username)
		if err == nil { //exist
			correctPassword := userOld.Password
			if comparePasswords(correctPassword,userNew.Password){
				c.JSON(200, gin.H{"success": "Login success"})
			} else {
				c.JSON(400, gin.H{"error": "Incorrect password"})
			}
		} else {
			c.JSON(400, gin.H{"error": "Error"})
		}
	})

	//07
	router.POST("/registerUser", func(c *gin.Context){
		var userNew User
		var userOld User
		c.Bind(&userNew)
		err := dbmap.SelectOne(&userOld, "SELECT * FROM Users WHERE Email=?", userNew.Email)
		if err == nil { //exist
			c.JSON(400, gin.H{"error": "This email is already existing"})
		} else { //non-exist
			tempP := tempPass()
			pazz := hashAndSalt(userNew.Password)
			dbmap.Exec(`INSERT INTO Users (Username, Password, TempPassword, Email, Firstname, Lastname, Role) VALUES (?, ?, ?, ?, ?, ?, ?)`, userNew.Username, pazz, tempP, userNew.Email, userNew.Firstname, userNew.Lastname, userNew.Role);
			exec.Command("sendemail", userNew.Email, "< 'Thank you for your registration.'")
			c.JSON(200, gin.H{"success": "Register success"})
		}
	})

	//07
	router.POST("/changePassword", func(c *gin.Context){
		var userNew User
		var userOld User
		c.Bind(&userNew)
		err := dbmap.SelectOne(&userOld, "SELECT * FROM Users WHERE Email=?", userNew.Email)
		if err == nil { //exist
			pazz := hashAndSalt(userNew.Password)
			dbmap.Exec(`UPDATE Users SET Password=? WHERE Email=? AND Username=?`,pazz , userNew.Email, userNew.Username); 
			c.JSON(200, gin.H{"error": "Your password is updated"})
		} else { //non-exist
			c.JSON(400, gin.H{"error": "Incorrect information"})
		}
	})

	//07
	// router.POST("/forgotPassword", func(c *gin.Context){
	// 	c.String(http.StatusOK,"Forgot Password")
	// })

	//08
	router.POST("/requestForm", func(c *gin.Context){
		c.String(http.StatusOK,"Submit Request")
	})

	//09
	router.GET("/requestList", func(c *gin.Context){
		c.String(http.StatusOK,"Request List")
	})

	//10
	router.DELETE("/deleteRequest", func(c *gin.Context){
		c.String(http.StatusOK,"Delete Request")
	})

	router.Run(":13131")
}

func check(e error){
        if e != nil {
           panic(e)
        }
}
