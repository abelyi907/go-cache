cd ../
go test ./... -v -covermode=count -coverprofile fib.out
go tool cover -html=fib.out -o fib.html