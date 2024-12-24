# mytheresa Promotions Test

## Description
The task is completed for Golang programming language using PostgreSQL as a database for store products data, Redis as a caching system and Docker for containerization.

Implemented features:
- main API endpoint `/products` for requesting products with possible filters from query parameters:
    - `category`,
    - `priceLessThan`,
- tests for `priceLessThan` validation and discounts calculation,

### Run application

```sh
docker compose up
```

This command runs the main routine:
- starts container with PostgreSQL, create table products and populate it with initial data,
- starts container with Redis,
- starts container with main application (web-server),

The API could be accessed from `localhost:8082` (application config located in `./config/conf.yaml` file).

### Requests examples

- `localhost:8082/products` - get all products,
- `localhost:8082/products?category=sneakers` - get all products for specified category
- `localhost:8082/products?priceLessThan=90000` - get all products for price less or equal than specified value,
- `localhost:8082/products?priceLessThan=90000&category=sneakers` - get all products for specified category AND for price less or equal than specified value,

### Tests
There are a few unit-tests (no needed filesystem or networking) for query parameter validation and discounts calculation. To run tests:
```sh
go test ./... -v
```

### Possible improvements
There are a lot of possible improvements on this task, like more comprehensive logging system, advanced testing (i.e. functional tests), validation etc.

Would really appreciate any feedback!
