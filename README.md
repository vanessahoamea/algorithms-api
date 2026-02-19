## About
Algorithms API provides solutions to various computational problems. It is queried over HTTP/HTTPS and accepts data in JSON format.

This repository also contains a performance test script developed with k6, which can be configured to function as a load, spike or stress test. It can also be integrated with a Jenkins server that will execute the pipeline defined in the `Jenkinsfile`.

## Setup
### Local environment
Environment variables are stored inside an `.env` file in the root directory.

```
# server vars
APP_ENV=dev
APP_PORT=<YOUR_APP_PORT>
APP_BASE_URL=http://localhost:<YOUR_APP_PORT>/v1

# test vars
K6_OPTIONS_FILE=<YOUR_K6_OPTIONS_FILE> # "load.json", "spike.json" or "stress.json"
```

The easiest way to compile and start the server and/or the k6 suite is by using the shorthand commands defined in the `Makefile`:

```
# server command
make build && make run

# test command
make k6
```

### Docker
Alternatively, the app and its corresponding performance tests can be containerized with Docker Compose.

```
# run the two containers in detached mode
docker compose up --build -d

# stop and remove the containers
docker compose down
```

This setup requires setting two more environment variables: the port for the `test` container, and the URL of the API running inside the `server` container.

```
K6_PORT=<YOUR_K6_PORT>
K6_BASE_URL=http://server:<YOUR_APP_PORT>/v1
```

---

The pre-built Docker images are also publicly available and can be pulled from Docker Hub:

```
docker pull vanessahoamea/algorithms-api-server:latest
docker pull vanessahoamea/algorithms-api-test:latest
```

The containers should be on a shared network, so that the k6 test suite can send requests to the API's address. The environment variables can be set with the `-e` flag.

```
docker run -d --name server \
    --network <NETWORK> \
    -p <APP_PORT>:<APP_PORT> \
    -e APP_ENV=staging \
    -e PORT=<APP_PORT> \
    -e BASE_URL=<APP_BASE_URL> \
    vanessahoamea/algorithms-api-server:latest
```

```
docker run -d --name test \
    --network <NETWORK> \
    -p <K6_PORT>:<K6_PORT> \
    -e BASE_URL=<K6_BASE_URL> \
    -e OPTIONS_FILE=<K6_OPTIONS_FILE>.json \
    vanessahoamea/algorithms-api-test:latest
```

## API Usage
> The documentation is also available in OpenAPI format, and can be accessed at `/v1/swagger/index.html`.

### POST `/v1/n-queens`
Solves the given N Queens problem instance.

The request body should specify the number of queens and the blocked squares on the chessboard:

```
{
    "n": 4,
    "blocked": [[0, 0], [1,1], [3,2]]
}
```

### POST `/v1/knapsack`
Solves the given Knapsack problem instance. Both binary and fractional variants are considered.

The request body should specify the values and weights of each object, as well as the maximum weight that the knapsack can hold:

```
{
    "values": [19, 4, 1, 16, 16],
    "weights": [32, 37, 24, 49, 27],
    "capacity": 87
}
```

### POST `/v1/shortest-path`
Solves the given Shortest Path problem instance. Only accepts positive weights. Edges are treated as directed.

The request body should specify the number of nodes in the graph, the edges along with their respective weights, and the source node.

```
{
    "n": 6,
    "edges": [
        [0, 1, 2], [0, 2, 4], [1, 2, 1],
        [1, 3, 7], [2, 4, 3], [3, 5, 1],
        [4, 3, 2], [4, 5, 5]
    ],
    "source": 0,
}
```