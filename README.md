# Golang App - Supply Chain Event Management 

Some cool technology using in this project:

- Gin gonic (Golang framework)
- Gorm (Golang Object-relational mapping)
- GRPC (Binary protocol)
- BuntDB (Key-value embedded database in Golang)
- Zeebe Client (Connect to Zeebe engine)
- Gorush Client (Connect to Gorush for notification) 

## Project Basic Needs

- Golang - (go1.15.5)
- Gcc - (gcc version 7.5.0) to run SQLite in Linux
- MinGW - (GCC-8.1.0 x86_64-posix-seh) [Sourceforge.net](https://sourceforge.net/projects/mingw-w64/files/) ro run SQLite in Window

## Project Advance Needs

- Docker Compose to run Zeebe and Gorush
- Redis to replace BuntDB (production)
- MySQL to replace SQLite (production)

**Configure Server Information with server environment file in cmd/scem/ !**

## Running Script

Go to folder cmd/scem/ and run the below script to start the server.

### `go run main.go`

