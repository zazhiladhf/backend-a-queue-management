# Project BRI QMS

Goal : Sebagai penyedia layanan bank, kami ingin memberikan nasabah layanan perbankan tanpa harus lama menghabiskan waktu menunggu antrian.

Main Developer : Zazhil Adhafi  
Contributor : Rizqi Pratama, Rifky Tedianto, Luh Gede Dyah Pradnyadari  

- [Project BRI QMS](#project-bri-qms)
  - [Configuration](#configuration)
  - [API Tables](#api-tables)
  - [Cara Akses API](#cara-akses-api)
  - [POST `/login`](#post-login)
  - [POST `/register`](#post-register)
  - [POST `/book/create`](#post-bookcreate)
    - [Booking Berhasil](#booking-berhasil)
    - [Booking Penuh](#booking-penuh)
  - [GET `/bank`](#get-bank)
  - [GET `/bank/detail/1`](#get-bankdetail1)
  - [DELETE `/book/selesai/3`](#delete-bookselesai3)
  - [Tech Stack](#tech-stack)

## Configuration

Didokumentasikan didalam dokumentasi lain.

## API Tables

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

## Cara Akses API

Access API via `http://localhost:8000/{route}`

## POST `/login`

Authorization: basic auth

Request Body:

``` json
{"username": "dts2021@tes.com", "password": "dtsitp"}
```

Response: status code: 200

```json
{
    "data":{
        "username": "dts2021@tes.com",
    "id_user":1
    "msg": "login berhasil"
}
```

## POST `/register`

Request Body:

```json
{"username": "dts2021@tes.com", "password": "dtsitp"}
```

Response:
status code : 200

```json
{
    "data":{
        "username": "dts2021@tes.com",
    "id_user":1
    "msg": "registrasi berhasil"
}
```

## POST `/book/create`

### Booking Berhasil

Authorization: basic auth

Request Body:

```json
{
    "id_bank_tujuan":1,
    "keperluan_layanan":1,
    "id_user":1
}
```

Response: status code: 200

```json
{
    "msg": "berhasil"
    "data":{
        "id_booking":3,
        "id_bank_tujuan":1,
        "keperluan_layanan":1,
        "id_user":1,
        "tanggal_pelayanan":"01/08/2021",
        "jam_pelayanan":"09.00-10.00"
    }
}
```

### Booking Penuh

Request Body:

```json
{
    "id_bank_tujuan":2,
    "keperluan_layanan":2,
    "id_user":2
}
```

Response: status code: 201

```json
{
    "msg": "booking penuh"
    "data":{
        "id_bank_tujuan":2,
        "keperluan_layanan":2,
        "id_user":2
    }
}
```

## GET `/bank`

Authorization: basic auth

Response: status code: 200

```json
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

## GET `/bank/detail/1`

Authorization: basic auth

Response: status code: 200

```json
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

## DELETE `/book/selesai/3`

Authorization: basic auth

Response:
status code: 200

```json
{
    "msg": "berhasil"
    "data":{
        "id_booking":3}
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
