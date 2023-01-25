# Ecommerce API

## Summary

E-commerce refers to the buying and selling of goods or services using the internet, and the transfer of money and data to execute these transactions. This program is a mini-project to create a back-end REST API for a simple e-commerce website using Go programming languange.

## Problem & Motivation

With digitalization of marketplace ever increasing in popularity, the needs for developing the software have also increased. This mini-project aims to create a back-end API for a simple e-commerce website using Go programming languange.

## Features

- Create listings for products to sell
- List available products to buy with filtering
- Add products to a shopping cart and checkout

## How to Run

Initialize Go in workspace

```
go mod init
```

Import main dependencies:

- Echo http framework
- PostgreSQL GORM

```
go get github.com/labstack/echo/v4
go get gorm.io/gorm
go get gorm.io/driver/postgres
```

To run program locally, in `.env` file:

- uncomment `"DB_HOST=127.0.0.1"`,
- comment `"DB_HOST=ecommerce-postgres"`,
- change `"DB_USER"` and `"DB_PASSWORD"` to your own postgres,

then run command:

```
go run main.go
```

To run program in docker compose, revert previous comments in `.env` file then run commands:

```
docker-compose build
docker-compose up -d
```

To stop the program running in docker containers, run:

```
docker-compose down
```

## API Documentation

Google Sheets: https://docs.google.com/spreadsheets/d/1m7vjL-PufRdo4NbXtVnx8B70KasnpftmWPlb7RrH3hE/edit?usp=sharing

Postman collection: https://documenter.getpostman.com/view/24354393/2s8YeoQZNX
