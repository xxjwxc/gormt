all: 
	make windows 
	make mac
	make linux
	make clear
windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o gormt.exe main.go 
	tar czvf gormt_windows.zip gormt.exe config.yml
mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o gormt main.go 
	tar czvf gormt_mac.zip gormt config.yml
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gormt main.go 
	tar czvf gormt_linux.zip gormt config.yml
clear:
	rm gormt
	rm gormt.exe
	rm -rf  model/* 
	rm -rf err/ 

