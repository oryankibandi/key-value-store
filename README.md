# Key Value Store

This repo is a minimal durable key-value store that can be integrated with **[raft consensus algorithm](https://raft.github.io/)** as a state machine.

Depending on when you are reading this, this will be part of a larger implementation on raft consensus algorithm.

## How to run

1. Clone the repo

    ```bash
    git clone <url>
    ```

2. Build the binary

    ```bash
    go mod tidy \
    go mod download \
    go build
    ```
3. Run the binary:
   ```bash
   ./key_value_store <port>
   ```

   Ensure to specify port to run the server.

4. Query the `/getval` and `/add` endpoints.

    To add a value:
    ```bash
    curl -X POST "http://localhost:<port>/api/add" -H "Content-Type: application/json" -d '"key":"x","value":"5"}'
    ```

    To retrieve a value:
    ```bash
    curl "http://localhost:<port>/api/getval"
    ````

## Sample requests

### Create Entry
Request
 ```bash
curl -X POST "http://localhost:<port>/api/add" -H "Content-Type: application/json" -d '"key":"x","value":"5"}'
```

Response:

```bash
{"key":"x","value":"5"}
```

### Retrieve Value
Request
 ```bash
curl -X GET "http://localhost:<port>/api/getval" -H "Content-Type: application/json" -d '{"key":"x"}'
```

Response:

```bash
{"Key":"x","Value":"5"}
```