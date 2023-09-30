## Important docker commands:

```shell
# build image
docker build . -t whoabhisheksah/go-server

# push image
docker login 
docker push whoabhisheksah/go-server:latest

# run image locally
docker run -i --expose=8080 -p 8080:8080 whoabhisheksah/go-server 
```

