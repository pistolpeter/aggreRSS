# aggreRSS

An RSS Feed aggregation project.

## Usage

1. Clone the repository

```bash
$ git clone github.com/pistolpeter/aggreRSS
```

2. Start a Postgresql server and add the connection string and your server's port number to a .env file in the repo.

```.env
# .env

PORT=<Port number>
CONNECTION=<Postgresql connection string>
```

3. Start the feed aggregation server

```bash
$ go build && ./aggreRSS
```
