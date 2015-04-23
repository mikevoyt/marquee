all:
	GOARCH=arm GOARM=5 GOOS=linux go build marquee.go