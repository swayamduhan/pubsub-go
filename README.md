# Pub Sub architecture in Go using Redis
Send messages from publisher and listen via subscriber.  
This architecture is widely used in combination with websockets to send data to clients on a scale.  

Module used : `redis-go`  

# How to setup
1. Start up a new docker redis container using 
```bash
    docker run -d -p 6379:6379 redis
```
2. Open new terminal.
```bash
    cd /publisher
    go run main.go
```

3. Open another terminal
```bash
    cd /subscriber
    go run main.go
```

