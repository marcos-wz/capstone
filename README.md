# CAPSTONE - Fruit Store
The purpouse of this system is to manage a fruit store stock.

# Features
These are the main features system:
- Solid, clean, readable, maintainable, flexible, and testeable code.
- Clean architecture, one method per interface.
- Every layer unit tests implementation. Testify and Mockery used libraries.
- Input data integrity validations, go through validator library.
- Error propagation management.
- Echo http framework based.
- Directories explanation:
    - Docs directory contains flow chart diagrams.
    - Build directory contains deployments files, docker, kubernetes, makefiles, scripts.
    - Data directory contains production and testing data, csv and json files.
    - Internal directory contains domain code: controllers, services, repositories, and entities.

# First Delivery

## Filter Controller
- Always return json reponses
- Active filters managed by entity tags using validator library. Current active filters: id name color country.
Fruit Filter Error evaluation process:
- 422 Unprocessable Entity : Params filter errors
- 500 Internal Server : CSV File error
- 206 Partial Content : CSV Parse Fruit Error, 
    filter fruit partial data with errors validations in JSON Format.

## Filter Service

## Reader Repository 
Parser Fruit Reader: Go through 2 libraries(csv,validate) to guarantee csv data integrity. Input data method
The ParseFruitRecord unit test, validates by all tags validations defined in the fruit entity struct
Full errors manager
