# Project BRI QMS

Goal : Sebagai penyedia layanan bank, kami ingin memberikan nasabah layanan perbankan tanpa harus lama menghabiskan waktu menunggu antrian.

## setup

1. clone this repository
2. type `make run` to run the apps
3. Hit the Server `localhost:8000/`
4. Unit Test `make coverage`
5. see coverage all test in html `make coverage-out`

## Prerequisite

To run this program, you will need

### App Dependencies

```$xslt
- Golang 1.12+
- Go mod enabled
```

## How to Run

### Setup App Config

```
cp .env.example .env
```

### Run Application

```
make run
```

## How to Test

```
make test
```

## How to Lint

```
make lint
```

## Deployment

### Build

```
make build
```

## Configuration

| NAME        | DESCRIPTION                      | TYPE   | VALUE       |
| ----------- | -------------------------------- | ------ | ----------- |
| APP_NAME    | Application name                 | string | alphabet    |
| APP_PORT    | Application port                 | int    | number      |
| LOG_LEVEL   | Mode for log level configuration | string | debug/info  |
| ENVIRONMENT | Application environment          | string | development |
| JWT_SECRET  | JWT Secret                       | string | alphabet    |

## Tasks

We define routes for handling operations:

| Method | Route           | Action              |
| ------ | --------------- | ------------------- |
| POST   | /login          | validate basic auth |
| POST   | /register       | create account      |
| POST   | /book/create    | create book         |
| POST   | /book/create    | create book         |
| GET    | /bank           | get                 |
| GET    | /book/detail/1  | get                 |
| DELETE | /book/selesai/3 | delete              |

Access API via `http://localhost:8000/{route}`

1. POST `/login`

Authorization: basic auth

Request Body:

```
{"username": "dts2021@tes.com", "password": "dtsitp"}
```

Response: status code: 200

```
{
    "data":{
        "username": "dts2021@tes.com",
    "id_user":1
    "msg": "login berhasil"
}
```

2. POST `/register `

Request Body:

```
{"username": "dts2021@tes.com", "password": "dtsitp"}
```

Response:
status code : 200

```
{
    "data":{
        "username": "dts2021@tes.com",
    "id_user":1
    "msg": "registrasi berhasil"
}
```

3. POST `/book/create`

Authorization: basic auth

Request Body:

```
{
    "id_bank_tujuan":1,
    "keperluan_layanan:1,
    "id_user":1
}
```

Response: status code: 200

```
{
    "msg": "berhasil"
    "data":{
        "id_booking":3,
        "id_bank_tujuan":1,
        "keperluan_layanan:1,
        "id_user":1,
        "tanggal_pelayanan":"01/08/2021",
        "jam_pelayanan":"09.00-10.00"
    }
}
```

4. POST `/book/create`

Authorization: basic auth

Request Body:

```
{
    "id_bank_tujuan":2,
    "keperluan_layanan:2,
    "id_user":2
}
```

Response: status code: 201

```
{
    "msg": "booking penuh"
    "data":{
        "id_bank_tujuan":2,
        "keperluan_layanan:2,
        "id_user":2
    }
}
```

5. GET `/bank`

Authorization: basic auth

Response: status code: 200

```
{
    "msg": "berhasil"
    "data":[{
        "id_bank":1,
        "nama_bank":"BANK KCP SOREANG",
        "alamat":"Jl.Soreang No.180"
    },
    {
        "id_bank":2,
        "nama_bank":"BANK KCP Banjaran",
        "alamat":"Jl.Banjaran No.181"
     }]
}
```

6. GET `/bank/detail/1`

Authorization: basic auth

Response: status code: 200

```
{
    "msg": "berhasil"
    "data":{
        "id_bank":1,
        "nama_bank":"BANK KCP SOREANG",
        "alamat":"Jl.Soreang No.180",
        "tanggal_antrian_saat_ini":"07/08/2021",
        "no_antrian_saat_ini":20,
        "waktu_pelayanan":"09.30"
    }
}
```

7. DELETE `/book/selesai/3`

Authorization: basic auth

Response:
status code: 200

```
{
    "msg": "berhasil"
    "data":{
        "id_booking":3}
}
```

### Tech Stack

- [Golang] - programming language
- [Fiber] - web framework with zero memory allocation and performance
- [MySQL] - open source database
- [GORM] - The fantastic ORM library for Golang

[golang]: https://golang.org/
[fiber]: https://github.com/gofiber/fiber/
[mysql]: https://www.mysql.com/
[gorm]: https://gorm.io/
