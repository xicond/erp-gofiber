
docker build -t my-fiber-app .

docker run -p 8080:8080 -v $(pwd):/app my-fiber-app sh -c "go mod init example.com/greetings"

#go mod tidy
go mod graph


# -a to force a rebuild of all packages
# -trimpath removes file system paths from the compiled binary
# -p flag to specify the number of concurrent build processes
# CGO_ENABLED=0 using cgo (C bindings), it can slow down your builds significantly.
#     Try to minimize or avoid using cgo unless absolutely necessary
# -tags=your_tag, If your application has features that are not always used,
#     consider using build tags to include or exclude those features during the build process.
go build -ldflags="-s -w" -a  -p 4 main.go

# Generate a private key
openssl genrsa -out private.key 2048

# Generate a public key
openssl rsa -in private.key -outform PEM -pubout -out public.key
