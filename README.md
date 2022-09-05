# verify_test


## Endpoints
```
GET - Return all users 
/users/


GET - Return 1 user
/users/USER-ID


POST - Create a user
/users/
{
    "name": "Some name",
    "password": "somepassword",
    "email": "someemail@gmail.com",
    "age": "18",
}

DELETE - Delete a user
/users/USER-ID


PUT - Update a user
/users/USER-ID
{
    "name": "Some name2",
    "email": "someemai2l@gmail.com",
    "age": "20",
}
```


## How to run?
```
<!-- create your .env base on the sample.env -->

<!-- RUN -->
go mod download
make run_local

<!-- OR on Docker-->
docker build . -t verify
docker run -d --name verify -p 8080:8080 verify
```