# HexPlate

HexPlate is a boilerplate template for Golang.

## Features

* Hexagonal Architecture design pattern
* Code generation from OpenAPI Schema
* Runs in Docker
* Wiremock for dependency injection

## Commands

| Command         | Description                                       |
|-----------------|---------------------------------------------------|
| `./bin/start`   | Starts the service and dependencies in watch mode |
| `./bin/stop`    | Stops the service and dependencies                |
| `./bin/gen-api` | Generates the API from the OpenAPI spec           |

## Folder Structure
* `internal` - *code relating to the application*
  * `api` - *api endpoint operations and helpers*
    * `[resource]` - *endpoint operations for the resource*
      * `provider.go` *the Wire provider for dependency injection*
      * `request.go` *helpers for handling requests such as DTO models*
      * `[CRUD]_[resource].go` *endpoint operations. Eg: `create_pet.go`*
  * `domain` - *the core. Everything calls in to this layer*
    * `[resource]` - *service methods for the resource*
      * `model.go` - *resource definitions*
      * `repository.go` - *interface for the repository layer*
      * `service.go` - *interface for the service that the resource must adhere to*
  * `repository` - *repository logic*
    * `[repository]` - *the repository used, Eg: `mongo`*
      * `[resource]` - *repository methods for the resource*
        * `provider.go` - *the Wire provider for dependency injection. Provides the repository and collection*
    * `operations.go` - *the interface repository operations must adhere to*
    * `provider.go` - *the Wire provider for dependency injection. Provides the operations*

The domain is the centre of the application. 
The API layer sends calls to the domain layer and gets back results. The domain sends information to the repository layer and bubbles the results back to the API as needed.
