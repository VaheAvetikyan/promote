# promotions

## Preface
We receive some records in a CSV file (example ​promotions.csv​attached) every 30
minutes. We would like to store these objects in a way to be accessed by an endpoint.
Given an ID the endpoint should return the object, otherwise, return not found.
Eg:
curl http://localhost:1321/promotions/1
{"id":"172FFC14-D229-4C93-B06B-F48B8C095512", "price":9.68, "expiration_date": "2022-06-04 06:01:20"}

## Install
Perform the following steps:
1. Download and install [Golang](https://golang.org/).
1. Install Postgres database [PostgreSQL](https://www.postgresql.org/).
1. Get the source code of this repository by the following command.
    ```bash
    go install github.com/VaheAvetikyan/promote@latest
    ```
## Configuration
There are 2 methods for configuration file:
1. You can modify default file in the project path configuration/promotions.config.yaml
1. You can create your local config yaml file and specify name and path for that file as enviroment variables. 

    Example:
    ```bash
    export CONFIG_FILE_NAME=promotions.config.yaml
    ```
    ```bash
    export CONFIG_FILE_PATH=~/Documents/
    ```
    
## Starting Server
1. Starting this web application by the following command.
    ```bash
    go run main.go
    ```
    
|Service Name|HTTP Method|URL|Parameter|Summary|
|:---|:---:|:---|:---|:---|
|Get Promotion|GET|``/promotions/[id]``|id|Get a promotion data.|
|List Promotions|GET|``/promotions?offset=[page]&limit=[size]``|page, size|Get a list of promotions.|
|Add Promotion|POST|``/promotions``|{
    "id": "d018ef0b-dbd9-48f1-ac1a-eb4d90e57118",
    "price": 60.683466,
    "expirationDate": "2018-08-04 05:32:31 +0200 CEST"
}|Regist a book data.|
|Upload Promotions|POST|``/api/books``|Book|Edit a book data.|
