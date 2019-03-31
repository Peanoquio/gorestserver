# gorestserver
REST server written in Go

## Introduction

This is a HTTP REST API server

### Starting the Server

To start the server, execute these commands in order in the terminal/console:
```shell
go build

./gorestserver
```

### Making API Calls

To send requests to the server, use this URL path.
Take note that you can replace `{testParam}` with any value.
If you want to test different request methods (POST, DELETE, PUT, etc.), you can install and use the Postman tool.
```shell
http://127.0.0.1:8080/v1/api/test/{testParam}
```
