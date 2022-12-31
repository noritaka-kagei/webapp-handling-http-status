# webapp-handling-http-status
sapmle web application for handling HTTP status based on request path of URL  

#### example
```
$ curl -i localhost:8080/200
HTTP/1.1 **200** OK
Date: Sat, 31 Dec 2022 14:10:10 GMT
Content-Length: 0

$ curl -i localhost:8080/300
HTTP/1.1 **300** Multiple Choices
Date: Sat, 31 Dec 2022 14:10:07 GMT
Content-Length: 0
```
