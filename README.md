## Goys
It is a project that sets and gets keys using the rest api system.

## Init

Make sure you have golang installed on your computer. After pulling the project, enter the src folder and run the following command.
```shell
go run .
```

If you have docker on your computer, you can run the following command.
```shell
docker-compose up -d
```
## Test

You can run following command in ```src``` folder,if you want to test.
```
go test -v ./api
```
If it is run successfully, you can see the response of the health check function at **[localhost:8889](localhost:8889)**. (This will return you a blank page with 200 code.)

##Uri and parameters

|  Endpoint |  Method | Request Body
| ------------ | ------------ | ------------ |
|  /set | POST  |  `{"key": "set-key","value": "set-value"} ` |
|  /get?key={key} | GET  |      |
| /flush | GET,POST | |
| / | GET | |
