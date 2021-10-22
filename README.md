# Project BRI QMS

Goal : Sebagai penyedia layanan bank, kami ingin memberikan nasabah layanan perbankan tanpa harus lama menghabiskan waktu menunggu antrian.

Main Developer : Zazhil Adhafi  
Contributor : Rizqi Pratama, Rifky Tedianto, Luh Gede Dyah Pradnyadari  

- [Project BRI QMS](#project-bri-qms)
  - [Configuration](#configuration)
  - [API Tables](#api-tables)
  - [Cara Akses API](#cara-akses-api)
  - [POST `/register`](#post-register)
  - [POST `/login`](#post-login)
  - [POST `/book/create`](#post-bookcreate)
    - [Booking Berhasil](#booking-berhasil)
    - [Booking Penuh](#booking-penuh)
  - [GET `/bank`](#get-bank)
  - [GET `/bank/detail/1`](#get-bankdetail1)
  - [GET `/book/:id`](#get-bank)
  - [GET `/book/detail/:id`](#get-bankdetail1)
  - [PUT `/book/selesai/:id`](#delete-bookselesai3)
  - [DELETE `/book/selesai/:id`](#delete-bookselesai3)
  - [Tech Stack](#tech-stack)

## Configuration

Didokumentasikan didalam dokumentasi lain.

## API Tables

We define routes for handling operations:

| Method | Route           | Action              |
| ------ | --------------- | ------------------- |
| POST   | /register       | create account      |
| POST   | /login          | login to app |
| POST   | /book/create    | create book         |
| POST   | /book/create    | create book (full book)         |
| GET    | /bank           | get list bank                 |
| GET    | /bank/detail/:id           | get bank detail by BankID                 |
| GET    | /book/:id  | get all list book detail by user login
| GET    | /book/detail/:id  | get detail booking by Booking ID                |
| PUT | /book/selesai/:id | update status to done by Booking ID         |
| DELETE | /book/selesai/3 | soft delete by Booking ID              |

## Cara Akses API

Access API via `http://localhost:8080/{route}`


## POST `/register`

Request Body:

```json
{"username": "dts2021@tes.com", "password": "dtsitp"}
```

Response:
status code : 200

```json
{
    "data": {
        "id_user": 5,
        "username": "dts2021@tes.com"
    },
    "message": "registrasi berhasil",
    "status": 200
}
```

## POST `/login`

Request Body:

``` json
{"username": "dts2021@tes.com", "password": "dtsitp"}
```

Response: status code: 200

```json
{
    "data": {
        "id_user": 5,
        "username": "dts2021@tes.com"
    },
    "message": "login berhasil",
    "status": 200
}
```

## POST `/book/create`

### Booking Berhasil

Request Body:

```json
{
    "id_bank_tujuan": 1,
    "keperluan_layanan": 1,
    "id_user": 5
}
```

Response: status code: 200

```json
{
    "data": {
        "id_booking": 9,
        "tanggal_pelayanan": "23-10-2021",
        "jam_pelayanan": "20:50",
        "keperluan_layanan": 1,
        "status": "",
        "id_bank_tujuan": 1,
        "id_user": 5,
        "Bank": {
            "id_bank_tujuan": 0,
            "nama_bank": "",
            "alamat": "",
            "kapasitas": 0
        },
        "User": {
            "id_user": 0,
            "username": ""
        },
        "DeletedAt": null
    },
    "message": "berhasil",
    "status": 200
}
```

### Booking Penuh

Request Body:

```json
{
    "id_bank_tujuan": 1,
    "keperluan_layanan": 3,
    "id_user": 5
}
```

Response: status code: 201

```json
{
    "message": "booking penuh",
    "status": 201
}
```

## GET `/bank`

Response: status code: 200

```json
{
    "data": [
        {
            "id_bank_tujuan": 1,
            "nama_bank": "BANK KCP SOREANG",
            "alamat": "Jl.Soreang No.180",
            "kapasitas": 0
        },
        {
            "id_bank_tujuan": 2,
            "nama_bank": "BANK KCP Banjaran",
            "alamat": "Jl.Banjaran No.181",
            "kapasitas": 0
        }
    ],
    "message": "berhasil",
    "status": 200
}
```

## GET `/bank/detail/:id`


Response: status code: 200

```json
{
    "error": false,
    "message": "berhasil",
    "result": [
        {
            "id_booking": 3,
            "tanggal_pelayanan": "22-10-2021",
            "jam_pelayanan": "13:51",
            "keperluan_layanan": 2,
            "status": "",
            "id_bank_tujuan": 1,
            "id_user": 2,
            "Bank": {
                "id_bank_tujuan": 0,
                "nama_bank": "",
                "alamat": "",
                "kapasitas": 0
            },
            "User": {
                "id_user": 0,
                "username": ""
            },
            "DeletedAt": null
        },
        {
            "id_booking": 6,
            "tanggal_pelayanan": "23-10-2021",
            "jam_pelayanan": "13:51",
            "keperluan_layanan": 2,
            "status": "",
            "id_bank_tujuan": 1,
            "id_user": 2,
            "Bank": {
                "id_bank_tujuan": 0,
                "nama_bank": "",
                "alamat": "",
                "kapasitas": 0
            },
            "User": {
                "id_user": 0,
                "username": ""
            },
            "DeletedAt": null
        },
        {
            "id_booking": 8,
            "tanggal_pelayanan": "23-10-2021",
            "jam_pelayanan": "13:51",
            "keperluan_layanan": 3,
            "status": "",
            "id_bank_tujuan": 1,
            "id_user": 2,
            "Bank": {
                "id_bank_tujuan": 0,
                "nama_bank": "",
                "alamat": "",
                "kapasitas": 0
            },
            "User": {
                "id_user": 0,
                "username": ""
            },
            "DeletedAt": null
        },
        {
            "id_booking": 9,
            "tanggal_pelayanan": "23-10-2021",
            "jam_pelayanan": "20:50",
            "keperluan_layanan": 1,
            "status": "",
            "id_bank_tujuan": 1,
            "id_user": 1,
            "Bank": {
                "id_bank_tujuan": 0,
                "nama_bank": "",
                "alamat": "",
                "kapasitas": 0
            },
            "User": {
                "id_user": 0,
                "username": ""
            },
            "DeletedAt": null
        },
        {
            "id_booking": 12,
            "tanggal_pelayanan": "23-10-2021",
            "jam_pelayanan": "20:50",
            "keperluan_layanan": 2,
            "status": "",
            "id_bank_tujuan": 1,
            "id_user": 5,
            "Bank": {
                "id_bank_tujuan": 0,
                "nama_bank": "",
                "alamat": "",
                "kapasitas": 0
            },
            "User": {
                "id_user": 0,
                "username": ""
            },
            "DeletedAt": null
        },
        {
            "id_booking": 13,
            "tanggal_pelayanan": "23-10-2021",
            "jam_pelayanan": "20:50",
            "keperluan_layanan": 1,
            "status": "",
            "id_bank_tujuan": 1,
            "id_user": 5,
            "Bank": {
                "id_bank_tujuan": 0,
                "nama_bank": "",
                "alamat": "",
                "kapasitas": 0
            },
            "User": {
                "id_user": 0,
                "username": ""
            },
            "DeletedAt": null
        }
    ],
    "status": 200
}
```

## GET `/book/:id`

Response: status code: 200

```json
{
    "error": false,
    "result": [
        {
            "id_booking": 11,
            "tanggal_pelayanan": "23-10-2021",
            "jam_pelayanan": "20:50",
            "keperluan_layanan": 2,
            "status": "",
            "id_bank_tujuan": 2,
            "id_user": 5,
            "Bank": {
                "id_bank_tujuan": 0,
                "nama_bank": "",
                "alamat": "",
                "kapasitas": 0
            },
            "User": {
                "id_user": 0,
                "username": ""
            },
            "DeletedAt": null
        },
        {
            "id_booking": 12,
            "tanggal_pelayanan": "23-10-2021",
            "jam_pelayanan": "20:50",
            "keperluan_layanan": 2,
            "status": "",
            "id_bank_tujuan": 1,
            "id_user": 5,
            "Bank": {
                "id_bank_tujuan": 0,
                "nama_bank": "",
                "alamat": "",
                "kapasitas": 0
            },
            "User": {
                "id_user": 0,
                "username": ""
            },
            "DeletedAt": null
        },
        {
            "id_booking": 13,
            "tanggal_pelayanan": "23-10-2021",
            "jam_pelayanan": "20:50",
            "keperluan_layanan": 1,
            "status": "",
            "id_bank_tujuan": 1,
            "id_user": 5,
            "Bank": {
                "id_bank_tujuan": 0,
                "nama_bank": "",
                "alamat": "",
                "kapasitas": 0
            },
            "User": {
                "id_user": 0,
                "username": ""
            },
            "DeletedAt": null
        }
    ]
}
```

## GET `/book/detail/:id`

Response: status code: 200

```json
{
    "error": false,
    "result": {
        "id_booking": 5,
        "tanggal_pelayanan": "23-10-2021",
        "jam_pelayanan": "13:51",
        "keperluan_layanan": 1,
        "status": "",
        "id_bank_tujuan": 2,
        "id_user": 2,
        "Bank": {
            "id_bank_tujuan": 0,
            "nama_bank": "",
            "alamat": "",
            "kapasitas": 0
        },
        "User": {
            "id_user": 0,
            "username": ""
        },
        "DeletedAt": null
    }
}
```

## PUT `/book/selesai/:id`

Request Body:

``` json
{
    "id_booking" : 2
}
```

Response: status code: 200

```json
{
    "error": false,
    "msg": "success update data",
    "result": {
        "id_booking": 2,
        "tanggal_pelayanan": "23-10-2021",
        "jam_pelayanan": "20:50",
        "keperluan_layanan": 3,
        "status": "done",
        "id_bank_tujuan": 2,
        "id_user": 1,
        "Bank": {
            "id_bank_tujuan": 0,
            "nama_bank": "",
            "alamat": "",
            "kapasitas": 0
        },
        "User": {
            "id_user": 0,
            "username": ""
        },
        "DeletedAt": null
    }
}
```

## DELETE `/book/selesai/:id`

Response:
status code: 201

```json
{
    "message": "berhasil",
    "status": 201
}
```

## Tech Stack

- [Golang] - programming language
- [Fiber] - web framework with zero memory allocation and performance
- [MySQL] - open source database
- [GORM] - The fantastic ORM library for Golang

[golang]: https://golang.org/
[fiber]: https://github.com/gofiber/fiber/
[mysql]: https://www.mysql.com/
[gorm]: https://gorm.io/
