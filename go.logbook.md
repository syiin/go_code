Thu 14 Nov 2019 11:24:39
- remember to set the $GOPATH environment variable before running `go run main.go`
- to check this, run `printenv | grep GOPATH`
- note that go expects your code in a src folder
- when sending requests to test the API (ie. auth_db_api), use something like this `curl -X POST -H 'Content-Type: application/json' -d "{\"Email\": \"admin@propertypricetag.com\", \"Password\": \"password\"}" localhost:8080/register`
> note the escape characters around quotations
> that its a POST request
> and that the content type is set
> prior attempts without this and with Postman resulted in empty fields (ie `Name:""`)
