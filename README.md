# query-builders

Example of SQL Query builders in GOlang to create API REST

[![Build Status](https://travis-ci.org/Demitroi/query-builders.svg?branch=master)](https://travis-ci.org/Demitroi/query-builders)
[![codecov](https://codecov.io/gh/Demitroi/query-builders/branch/master/graph/badge.svg)](https://codecov.io/gh/Demitroi/query-builders)
[![Go Report Card](https://goreportcard.com/badge/github.com/Demitroi/query-builders)](https://goreportcard.com/report/github.com/Demitroi/query-builders)

### Usage

We need a mysql/mariadb database server. Create a database and the table from [dabatase.sql](database.sql) file. Execute indicating the query builder, for example.

``` sh
$ go run main.go gendry
Now listening on: http://0.0.0.0:1081
Application started. Press CTRL+C to shut down.
```

There are 3 query builders available.

* [gendry](https://github.com/didi/gendry)

* [goqu](https://github.com/doug-martin/goqu)

* [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx)

Also we can pass database params via args.

``` sh
$ go run main.go --help
query-builder is an example of mysql databases in golang

Usage:
  query-builder [flags] [builder]

Examples:
query-builder gendry
query-builder goqu
query-builder dbx

Flags:
      --database-address string    Database connection address (default "localhost") 
      --database-name string       Database name (default "query_builders")
      --database-password string   Database user password
      --database-port int          Database connection port (default 3306)
      --database-protocol string   Database connection protocol (default "tcp")      
      --database-user string       Database user (default "root")
  -h, --help                       help for query-builder
  -p, --port int                   Port to serve (default 1081)
```

### API

#### List persons

GET /api/v1/persons

Query string Params

| Field          | Type   |
| -------------- | ------ |
| id             | string |
| name           | string |
| city           | string |
| birth_date_eq  | string |
| birth_date_gte | string |
| birth_date_lte | string |
| weight_eq      | number |
| weight_gte     | number |
| weight_lte     | number |
| height_eq      | number |
| height_gte     | number |
| height_lte     | number |

Example response

``` json
[
  {
    "id": "1",
    "name": "Demitroi Marshall",
    "city": "Ejido Mezquiral",
    "birth_date": "1995-06-07T08:00:00Z",
    "weight": 85,
    "height": 181
  },
  {
    "id": "6",
    "name": "John Cena",
    "city": "Tampa",
    "birth_date": "1977-04-23T08:00:00Z",
    "weight": 112,
    "height": 184
  }
]
```

#### Get person

GET /api/v1/persons/:id

Example response

``` json
{
  "id":"1",
  "name":"Demitroi Marshall",
  "city":"San Luis Rio Colorado",
  "birth_date":"1995-06-07T08:00:00Z",
  "weight":85,
  "height":181
}
```

#### Add person

POST /api/v1/persons

Example request

``` json
{
  "name": "John Cena",
  "city": "Tampa",
  "birth_date": "1977-04-23T08:00:00Z",
  "weight": 112,
  "height": 184
}
```

Example Response

``` json
{
  "id": "6"
}
```

#### Update person

POST /api/v1/persons/6

Example request

``` json
{
  "name": "John Cena",
  "city": "Tampa",
  "birth_date": "1977-04-23T08:00:00Z",
  "weight": 112,
  "height": 184
}
```

Example Response

``` json
{
  "msg": "OK"
}
```

#### Delete person

Delete /api/v1/persons/:id

Example response

``` json
{
  "msg": "OK"
}
```
