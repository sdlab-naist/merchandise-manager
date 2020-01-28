# REST API
```
go version go1.13.6 darwin/amd64
go get -u github.com/gin-gonic/gin
go get -u  github.com/gin-contrib/cors
```
## IP Address
```
163.221.29.46:13131
```
### Request item list
```
GET /getItems
@Return {
    {
      ID:   int,  
      Name: string,
      Price: double,
      Cost: double,
      Amount: int
    }, . . .
}
```
### Submit add form
```
POST /addItem
@Body {
    Name: string,
    Price: double,
    Cost: dobule,
    Amount: int
}
@Return "Response message"
```
### Submit delete form
```
POST /deleteItem
@Body {
    ID: int,
}
@Return "Response message"
```
### Register an order
```
POST /registOrder
@Body {
    OrderID: string (optional),
    ItemID: string,
    Amount: int
}
@Return "Response message"
```
### Make an order
```
POST /makeOrder
@Body {
    OrderID: string,
}
@Return "Response message"
```
### Check an order
```
GET /getOrders
@Return {
    {
      ID:   int,  
      OrderID: string,
      ItemID: string,
      Amount: int
    }, . . .
}
```
### User registration
```
POST /registerUser
@Body {
    Username: string,
    Password: string,
    Firstname: string,
    Lastname: string,
    Role: string,
    Email: string
}
@Return "Response message"
```
### Login
```
POST /login
@Body {
    Username: string,
    Password: string
}
@Return "Response message"
```
### Change password
```
POST /changePassword
@Body {
    Username: string,
    Email: string,
    Password: string
}
@Return "Response message"
```
### Submit request form
```
POST /requestItem
@Body {
    Username: string,
    Itemname: string,
    Amount: int
}
@Return "Response message"
```
### Check request list
```
GET /getRequests
@Return {
    {
      ID:   int,  
      Username: string,
      Itemname: string,
      Amount: int,
      Status: string
    }, . . .
}
```
### Delete request
```
POST /deleteRequest
@Body {
    ID: int,
}
@Return "Response message"
```