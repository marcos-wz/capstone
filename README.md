
# Main Features
- Echo http framework based
- Clean, Readable, Maintainable, Fexible, Testeable code
- Clean architecture, one method per interface.
- Every layer has its own unit tests, libraries Testify and Mockery
- Input data validations through Validator library
- Error propagation management
- Directories explanation:
    - Docs directory contains flow chart diagrams.
    - Build directory contains deployments files, docker, kubernetes, makefiles, scripts.
    - Data directory contains production and testing data, csv and json files.
    - Internal directory contains domain code: controllers, services, repositories, and entities.

# Reader Repository 
    Parser Fruit Reader: Go through 2 libraries(csv,validate) to guarantee csv data integrity. Input data method
    The ParseFruitRecord unit test, validates by all tags validations defined in the fruit entity struct
    Full errors manager

# Filter Service

# Filter Controller
- Always return json reponses
- Active filters managed by entity tags using validator library. Current active filters: id name color country.
Fruit Filter Error evaluation process:
- 422 Unprocessable Entity : Params filter errors
- 500 Internal Server : CSV File error
- 206 Partial Content : CSV Parse Fruit Error, 
    filter fruit partial data with errors validations in JSON Format.