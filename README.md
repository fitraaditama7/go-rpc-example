# go-rpc-example

### Requirements
- Golang (>= 1.11) - Download and Install [Golang](https://golang.org/dl/)
- MongoDB - Download and Install [MongoDB](https://www.mongodb.com/)
- Protocol Buffer - Download and Install [Protoc](https://developers.google.com/protocol-buffers/)

### Setting Up Application
Download Golang Package
```
make download
```

Running Server
```
make server
```

### Command
After server running you can running commands with format
```
go run client/main.go [command] [-args] [value]
```

List of Commands

| Commands  | Description |
| ------------- | ------------- |
| list  | Get list data from blogpost  |
| create  | Creating new blogpost  |
| read  | Get data by ID  |
| update  | Update data by ID  |
| delete  | Delete data by ID  |

List of Arguments

| Arguments  | Description |
| ------------- | ------------- |
| -i  | MongoDB ID  |
| -a  | Author Name from Blogpost  |
| -t  | Title from Blogpost  |
| -c  | Content from Blogpost  |

List Blog Post
```
go run client/main.go list
```

Create Blog Post
```
go run client/main.go create -a "Fitra Aditama" -t "Learning Go" -c "Go is fun"
```

Read Blog Post
```
go run client/main.go read -i "5f312dd51e9f7aad22cf6158"
```

Update Blog Post 
```
go run client/main.go read -i "5f312dd51e9f7aad22cf6158" -a "Mochammad Fitra Aditama" -t "Learning Golang" -c "Golang is fun"
```

Delete Blog Post 
```
go run client/main.go delete -i "5f312dd51e9f7aad22cf6158"
```
