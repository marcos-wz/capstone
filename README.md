# CAPSTONE - Fruit Store
The purpouse of this system is to manage a fruit store stock. 

# System Features
These are the main features system:
- Solid, secure, maintainable, and testeable code.
- Clean architecture, one method per interface.
- Every layer unit tests implementation. Testify and Mockery used library.
- Input data integrity validations based on tags through validator library.
- Error propagation management.
- Echo http framework based.
- Directories explanation:
    - Docs directory contains basics flow chart.
    - Build directory contains deployments files, docker, kubernetes, makefiles, scripts.
    - Data directory contains production and testing data, csv and json files.
    - Internal directory contains domain code: controllers, services, repositories, and entities.

# First Delivery

## Filter Controller
- Always return json reponses.
- Only valid filters and values are allowed.
- Valid filters are managed by entity tags using validator library. 
- Current active filters: id, name, color, country.
Fruit Filter responses:
- 200 Status OK: returns fruits filtered list
- 206 Partial Content: returns fruits filtered list and reader parser errors(Invalid CSV file!!)	
- 422 Unprocessable Entity : returns param filter and value errors
- 500 Internal Server : returns reader CSV File error (critical!)
- 400 Bad Request: default errors

## Filter Service


## Reader Repository 
Parser Fruit Reader: Go through 2 libraries(csv,validate) to guarantee csv data integrity. Input data method
The ParseFruitRecord unit test, validates by all tags validations defined in the fruit entity struct
Full errors manager
