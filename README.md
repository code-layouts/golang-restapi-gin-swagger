# golang-restapi-gin-swagger
golang-gin 및 Swagger 를 활용하여 Restful API 구현하는 샘플 프로젝트 입니다.

## Project-layout
golang 프로젝트 레이아웃을 참고 합니다.
[project-layout](https://github.com/golang-standards/project-layout)


## Prerequisite
swagger golang 모듈 설치
```
go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag@latest
swag --version
```

## Project Init
macos 에 go 를 설치 합니다.
```shell
git clone https://github.com/code-layouts/golang-restapi-gin-swagger.git
cd golang-restapi-gin-swagger
go mod init example/apiserver/v1
go mod tidy 
```

## Build
```
go build -o target/service ./main.go
```

## Generate Sample Data
```
cat <<EOF >> ./data/users.json
[
  {
    "id": 1,
    "firstName": "Frank",
    "lastName": "Murphy",
    "email": "frank.murphy@rustvale.com",
    "title": "Mr",
    "role": "User",
    "usercode": "fr1234",
    "createDts": "2021-04-08T05:33:05.184Z",
    "updateDts": "2021-10-28T00:02:47.249Z"
  },
  {
    "id": 2,
    "firstName": "melon",
    "lastName": "Fruit",
    "email": "melon@gmail.com",
    "title": "Mr",
    "role": "Admin",
    "usercode": "222222",
    "createDts": "2021-10-25T06:45:31.210Z",
    "updateDts": "2021-10-25T07:57:34.201Z"
  }
]
EOF
```

## Run
```
## Very first time you need to generate sample data
target/service
```

## Check
```
curl -v -L -X GET 'http://localhost:8080/users' -H 'Content-Type: application/json'
```


## Appendix
- [OpenAPI in Local](http://localhost:8080/swagger/index.html)