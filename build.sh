curl -o ca-certificates.crt https://raw.githubusercontent.com/bagder/ca-bundle/master/ca-bundle.crt
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
docker build . -t slack-typing-indicator
