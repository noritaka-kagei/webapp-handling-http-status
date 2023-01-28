[![GitHub release](https://img.shields.io/github/release/noritaka-kagei/webapp-handling-http-status.svg)](https://github.com/noritaka-kagei/webapp-handling-http-status/releases/latest)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/noritaka-kagei/webapp-handling-http-status.svg)](https://github.com/noritaka-kagei/webapp-handling-http-status/blob/main/server4go)
[![License: Apache-2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/noritaka-kagei/webapp-handling-http-status/blob/main/LICENSE)


# webapp-handling-http-status
sapmle web application for handling HTTP response status based on request path of URL  

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

### How to Run web application
#### server4go
```
1. Build application
  $ ./scripts/build.sh
2. Run
  $ ./bin/run-linux-amd64
```
