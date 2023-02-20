
curl --header "Content-Type: application/json" --request POST --data '{"username":"xyz","password":"xyz"}' http://localhost:9999/api/login
curl -H "Content-Type: application/json" -X POST -D '{"username":"xyz","password":"xyz"}' http://localhost:9999/api/login

## Testing steps

* GET All
```
curl http://localhost:9999/movies
```
* GET Single
```
curl http://localhost:9999/movies/1
```
* CREATE - no need for ID cause it will create a random one
```
curl -H "Content-Type: application/json" -X POST -D '{"isbn":"3425823","title":"Movie Seven","director":{"firstname":"Jake","lastname":"Wallace"}}' http://localhost:9999/movies
```
* UPDATE
```
curl -H "Content-Type: application/json" -X PUT -D '{"isbn":"3425823","title":"Movie Seven","director":{"firstname":"Chris","lastname":"Wallace"}}' http://localhost:9999/movies/{id}
```
* DELETE
```
curl -H "Content-Type: application/json" -X DELETE http://localhost:9999/movies/{id}
```
