# CAPSTONE - Fruit Store
The purpouse of this system is to manage a fruit store stock. 

# System Features
These are the main features system:
- Solid, secure, maintainable, and testeable code.
- Clean architecture, one method per interface.
- Each layer(controllers,services,and repositories) contains its own unit tests. Testify and Mockery libraries are used.
- Input data integrity validations based on tags through validator library.
- Error propagation implementation.
- Echo HTTP framework based.
- Directories explanation:
    - Docs directory contains deliveries flow chart.
    - Build directory contains deployments files, docker, kubernetes, makefiles, scripts.
    - Data directory contains production and testing data, csv and json files.
    - Internal directory contains domain code: controllers, services, repositories, and entities.

# First Delivery

## Filter Controller
- Always return json reponses.
- Only valid filters and values are allowed.
- Valid filters are managed by entity tags using validator library. 
- Current active filters: id, name, color, country.

Responses:
- 200 Status OK: returns fruits filtered list
- 206 Partial Content: returns fruits filtered  partial list and reader parser errors(invalid csv file data!!)	
- 422 Unprocessable Entity : returns param filter and value errors
- 500 Internal Server : returns reader CSV File error (critical!)
- 400 Bad Request: default errors

## Filter Service
Get Filtered Fruits from the repository. Invalid filter returns empty fruits list with error.
- Refactoring by filter. 
- Repository error propagation support
- If repository parse reader error exists, returns partial fruits list, with default values, and parser error validations

## Reader Repository 
Parser Fruit Reader: Go through 2 libraries(csv,validate) to guarantee csv data integrity. Input data method
The ParseFruitRecord unit test, validates by all tags validations defined in the fruit entity struct
Full errors manager
