# glassnode-challenge

## Ranking API


## Performance
* Protobuf can be used in inter-communication for better performance. JSON is used because of its simplicity and compatibility with `Coinmarket` API.
* Ranking can be cached with `Time to live (TTL)` because in general, it does not change as frequently as the price. Cache can be achieved using Redis/Etcd, or in process memory.


 