# Simple User Management using GRPC and HTTP Gateway

## Usage

go run /backend/core/cmd/main.go

## Supported method

GET, POST, PATCH, DELETE

CREATE, RETRIEVE, UPDATE, DELETE, LIST (support pagination, filters and sort)

## TODO

~~Implement filter in LIST~~

~~Fix swagger or protobuf to allow filter (change get to post method)~~

~~Validate user input~~

~~Manage user login~~

~~Validate user authenticated~~

Handle DB error when init

~~Making a mapping layer to reduce manual mapping request data to model~~

~~Need define validate layer to handle custom validate data~~
