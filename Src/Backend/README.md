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
    Name: string,
    Amount: int
}
@Return "Response message"
```
### Register an order
### Make an order
### Check an order
### Submit login form
### Submit request form
### Check request list
### Delete request
