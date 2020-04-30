# Go training

```
docker run --name go_training -ti -v $(pwd):/go/app -w /go/app -p 4000:4000 golang:1.14 bash
```

```
go run main.go
```

Access your web by `http://localhost:4000`

# Exercise
Base on API and build a website with the same UI https://getbootstrap.com/docs/4.4/examples/pricing/
