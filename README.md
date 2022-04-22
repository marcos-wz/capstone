# Capstone - Fruit Store
The purpouse of this system is to manage a fruit store stock. 

# System Features
These are the main features system:
- Solid, secure, maintainable, and testeable code.
- Clean architecture, and separate files structuration for coding maintenance.
- Each layer(controllers,services,and repositories) contains its own unit tests. Testify and Mockery libraries are used.
- Input data integrity validations from controllers and repositories are based on tags through validator library.
- Error propagation implementation.
- Echo HTTP framework based.
- SSL secure conections
- Server configurations by yaml file 
- Debugging faeuture by command line flags invockations 
- Directories explanation:
    - Docs directory contains deliveries flow chart.
    - Build directory contains deployments files, docker, kubernetes, makefiles, scripts.
    - Data directory contains production and testing data, csv and json files.
    - Internal directory contains domain code: controllers, services, repositories, and entities.

# First Delivery
The first delivery returns a list of filtered fruits, based on filter. It involves 3 main components, all with their own entities.
1) Filter Controller: internal/controller/filter.go
2) Filter Service: internal/service/filter.go
3) Reader Repository: internal/repository/reader.go

## Filter Controller: internal/controller/filter.go
Filter controller interface, manage the filter data integrity from user, and the returned responses.
Repository and service error propagation support. Full errors management.
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

## Filter Service: internal/service/filter.go
Filter service interface, returns filtered fruits list from the repository. Invalid filter returns empty fruits list with error.
- Refactoring by filter. 
- Repository error propagation support. Full errors management.
- Error propagation: If repository parse reader error exists, returns partial fruits list, with default values, and parser error validations

## Reader Repository: internal/repository/reader.go
Reader repository interface, read all fruit records from the csv file. 
If parse fruit error ocurrs, returns the parse errors list, with partial records and their errors validation descriptions.
This interface is based on encoding/csv and playground/validator libraries.
- if empty fields or error field validation ocurrs, the default data type value is set.
- if parse validation errors on required fields exists, the record is ommited.
- It guarantees fruit data integrity and always returns a fruits instances.
- The Parse Fruit Record unit test, validates all tags validations use cases defined in the fruit entity struct
- Full errors management
