all: # 构建
	make tar 
windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o gormt.exe main.go 
mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o gormt main.go 
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gormt main.go 
tar: # 打包
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o gormt.exe main.go 
	tar czvf gormt_windows.zip gormt.exe config.yml
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o gormt main.go 
	tar czvf gormt_mac.zip gormt config.yml
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gormt main.go 
	tar czvf gormt_linux.zip gormt config.yml
clear:
	- rm -rf  model/* 
	- rm -rf err/ 
	- rm gormt
	- rm gormt.exe
	- rm gormt_linux.zip
	- rm gormt_mac.zip
	- rm gormt_windows.zip
