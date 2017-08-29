default:
	go build
pi:
	GOARCH=arm GOARM=5 GOOS=linux go build --tags "libsqlite3 linux"
