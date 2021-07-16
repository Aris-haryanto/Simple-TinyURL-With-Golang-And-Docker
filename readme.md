# How to Get in

### Run go mod tidy
```sh
$ go mod tidy
```

# Normal way

### Download third party
```sh
$ go mod download
```

### Run Testing
```sh
$ go test -v test/tiny_test.go
```

### Run Application
```sh
$ go run cmd/main.go
```

### Build Application to binary
```sh
$ go build -o tinyurl cmd/main.go 
```

# Working with Docker

### Run test with docker
- Build image 
```sh
$ docker build -t tinyurl-test -f Dockerfile.test .
```

- Run Container
```sh
$ docker run tinyurl-test
```

### Run Application with docker
- Build image 
```sh
$ docker build . -t tinyurl
```

- Run Container
```sh
$ docker run -p 3000:3000 tinyurl
```
