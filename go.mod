module golang-study

go 1.12

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190506204251-e1dfcc566284
	golang.org/x/net => github.com/golang/net v0.0.0-20190506115046-65e2d4e15006
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190506115046-112230192c
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190506115046-ca7f33d4116e
	golang.org/x/text => github.com/golang/text v0.3.0
)

require (
	github.com/gin-gonic/gin v1.4.0 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.0.3 // indirect
	golang.org/x/crypto v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/sync v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/text v0.0.0-00010101000000-000000000000 // indirect
)
