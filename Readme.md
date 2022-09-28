# RAPI

## Context
A simple application illustrating how to create an api with golang, gin & postgres.

## Running the application
1. start the postgres database: `sudo service postgresql start`
2. run the application: `go run rapi`

## Endpoints & Usage

#### GET /v1/ping

This endpoint is used to check if the application is running.

##### Example

```bash
curl http://localhost:8888/ping
```

```json
{
  "message": "pong"
}
```

#### GET /v1/items

This endpoint is used to get all items.

##### Example
```bash
curl  -X GET \
      http://localhost:8888/v1/items
```
response
```json
[
    {
        "id": 1,
        "name": "copegus",
        "reservation_id": "986"
    },
    {
        "id": 2,
        "name": "esbriet",
        "reservation_id": "652"
    },
    {
        "id": 3,
        "name": "fuzeon",
        "reservation_id": "453"
    }
]
```

#### POST /v1/items

Because the reservation service has a delay (I added here 10 seconds for debugging), the http response is sent back before the item entity is created in the database. In the case of a web application, there shall be a way to communicate this change to the user.

##### Example
```bash
curl  -X POST \
	-H "Content-Type: application/json" \
      -d '{"name": "item_name"}' \
      http://localhost:8888/v1/items
```
logs

```
[GIN] 2022/09/28 - 07:01:56 | 200 | 112.71µs | 127.0.0.1 | POST "/v1/items"
2022/09/28 07:01:56 Sleeping for 10 seconds... 
2022/09/28 07:02:06 reservationIdChannel:  81
2022/09/28 07:02:06 1 rows inserted
```

## Development
According to go idiomatic guidelines:
- https://go.dev/doc/faq7
- https://go.dev/ref/spec


## Missing parts in application code
- [ ] config management for environment variables
- [ ] dependency injection library
- [ ] lazy object instantiation
- [ ] generic service library for CRUDO operations
- [ ] refactoring using the `gorm` orm: https://gorm.io/docs/
- [ ] enhancing the middleware with security and logging features (jwt, etc.)

## Missing parts in the infrastructure
- [ ] setting up the dev container with all required components (db + app)
- [ ] writing helm charts of the application
- [ ] orchestrating the deployment of the microservice with kubernetes (db & api)
- [ ] ci / cd pipeline with concurse ci (https://concourse-ci.org/pipelines.html)
- [ ] introduction of sonarqube for code scanning
- [ ] introduction of monitoring tools (prometheus, grafana, etc.)

## Parts that are not there but `not missing`
- [ ] Pure OOP
- [ ] Adding new packages for components which are not reusable
- [ ] Introduction of data mappers between dtos & models only at a later stage of the project when repositores & services are diverging.
