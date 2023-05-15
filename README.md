# financial-TEMPLATE-api
API to register financial TEMPLATEs (expenses, earnings and transfers)

## Tecnology stack

* [GoLang 1.20](https://golang.org/doc/install) Go language compiler

## Installation

[TODO]

### Config

List of environment variables

|         Variable                  |                   Description                 |   Type        | Mandatory     | Default value                 |
| --------------------------------- | --------------------------------------------  |:----------:   |:-----------:  |:----------------------------: |
| APP_NAME                          | Name for app deploy                           | String        | Yes           | financial-TEMPLATE-api     |
| ENVIRONMENT                       | Environment of this instance is running       | String        | Yes           | development                   |
| DATABASE_SERVER_ADDRESS           | Database server name/address                  | String        | Yes           |                               |
| DATABASE_SERVER_PORT              | Database server port number                   | Integer       | Yes           |                               |
| DATABASE_NAME                     | Database name                                 | String        | Yes           |                               |
| DATABASE_USER                     | Database connection user                      | String        | Yes           |                               |
| DATABASE_PASSWORD                 | Database connection password                  | String        | Yes           |                               |
| API_PORT                          | Port exposed for this API                     | Integer       | Yes           |                               |
| SERVER_CLOSEWAIT                  | Time waiting to server shutdown               | Integer       | No            | 10                            |
| LOG_ACCESS_FILE                   | Path to access log file (API requests)        | String        | No            | ./access.log                  |
| LOG_APP_FILE                      | Path to application log file                  | String        | No            | ./app.log                     |
| LOG_LEVEL                         | Log level (ex: INFO, DEBUG, ERROR, etc)       | String        | No            | INFO                          |

### Running

To run locally:
```
docker compose up
```

### App Metrics
```
http://localhost:8081/metrics
```