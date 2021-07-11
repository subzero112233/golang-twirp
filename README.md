# twirp-app

for the REST api do:
```
export API_TYPE=rest
```

# Start the application:
```
   go run cmd/main.go
```


# Initiate an rpc GetStats request:
```echo 'player_name:"Jae Crowder"' \
 | protoc - encode stats.GetStatsRequest ./rpc/stats/stats.proto \
 | curl -s - request POST \
 - header "Content-Type: application/protobuf" \
 - data-binary @- \
 http://localhost:8000/twirp/stats.StatsService/GetStats \
 | protoc - decode stats.GetStatsResponse ./rpc/stats/stats.proto
```

# Initiate both GetStats and AddStats using the client:
```
go run example/example.go
``` 

# Benchmark:
```
go test -bench=. -benchtime=5000x
```
