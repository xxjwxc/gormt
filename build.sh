go build -o=gormt_mac ./main.go
GOOS=windows GOARCH=amd64 go build  -o=gormt.exe ./main.go
GOOS=linux GOARCH=amd64 go build  -o=gormt ./main.go
