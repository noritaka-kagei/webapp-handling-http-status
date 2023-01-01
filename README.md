# webapp-handling-http-status
sapmle web application for handling HTTP status based on request path of URL  

#### example
```
$ curl -i localhost:8080/200
HTTP/1.1 200 OK
Date: Sun, 01 Jan 2023 23:59:59 GMT
Content-Length: 0

$ curl -i localhost:8080/300
HTTP/1.1 300 Multiple Choices
Date: Sun, 01 Jan 2023 23:59:59 GMT
Content-Length: 0
```
