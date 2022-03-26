# ET Service

## The task

Develop an application. This app will download some data at regular intervals and place it to the database. The app is also a web-server. It will serve users requests to retrieve the data (received earlier) from the database.

### Data structure

Data source - <https://api.blockchain.com/v3/exchange/tickers>

Input data (retrieved from the data source):

```js
{<symbol>: {price: <price_24h>, volume: <volume_24h>, last_trade: <last_trade_price>}...}
```

Output data (user requests from our server):

```js
{<symbol>: {price: <price_24h>, volume: <volume_24h>, last_trade: <last_trade_price>}...}
```

#### Example

Input data:

```json
[
    {
        "symbol":"XLM-EUR",
        "price_24h":0.25685,
        "volume_24h":49644.7076291,
        "last_trade_price":0.24
    }
]
```

Output data:

```json
{
    "XLM-EUR": {
        "price": 0.25685,
        "volume": 49644.7076291,
        "last_trade":0.24
    }
}
```

## Deploying the app

The app is deployed with [*docker-compose*](https://docs.docker.com/compose/).

Move to the *deployments* folder. All further commands is meant to be run under this folder.

```bash
cd deployments
```

### .env file

Inside the *deployments* folder you'll find the *.env.default* file. This file contains configuration variables of the app. It already filled with the default values.  
Make a copy of *.env.default* file and name it

```
.env
```

Without this file (name should be exact the same) **the app will not deploy(!)**.  
You can change variables according to your taste or leave the default values.

### docker-compose

After setting the *.env* file, run *docker-compose* command (may need root-privileges, *sudo* the command below if needed). It will build and run the server.

```bash
docker-compose up -d
```

Wait the docker to be deployed.

To stop the server use:

```bash
docker-compose up stop
```

## Usage

You can send a request and retrieve the data while the server is running.

To do this just make a GET-request (you can do it inside the browser) to the localhost:port. Port is specified in the *.env* file (default 6911).

For example:

```
http://127.0.0.1:6911/
```

## Project structure

Project is structured according to [this recommendation](https://github.com/golang-standards/project-layout).

## Third-party modules

- [gin](https://github.com/gin-gonic/gin)
- [gorm](https://github.com/go-gorm/gorm)
- [gocron](https://github.com/go-co-op/gocron)
- [flag](https://github.com/namsral/flag)
- [govalidator](https://github.com/asaskevich/govalidator)

And [CompileDaemon](https://github.com/githubnemo/CompileDaemon) is used to track "go" files and rebuild the app inside the docker if they change.

## What can be improved?

Improvements that can be done for this project:

- **Tests**

- **Error wrapping** - currently errors are returned "as is". It's better to wrap every error in place of function use. For instance:

```go
err := fmt.Errorf("access denied: %w", ErrPermission)
```

- **Improve logging** - currently only errors and some third-party libraries messages are presented in logs. Logs don't have levels (they're not divided into *INFO*, *ERROR* etc.).

- **Divide one app into two services: grabber and server**

- **CompileDaemon is not working on Windows host**

- **Static documentation generation for the module**

- **Check with a linter (static code analysis tool)**
