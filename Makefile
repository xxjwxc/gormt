all: # 构建
	make tar 
windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o gormt.exe main.go 
mac:
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o gormt main.go 
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o gormt main.go 
tar: # 打包
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o gormt.exe main.go 
	tar czvf gormt_windows.zip gormt.exe config.yml
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o gormt main.go 
	tar czvf gormt_mac.zip gormt config.yml
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gormt main.go 
	tar czvf gormt_linux.zip gormt config.yml
clear:
	test ! -d model/ || rm -rf  model/*
	test ! -d err/ || rm -rf  err/
	test ! -f gormt || rm gormt
	test ! -f gormt.exe || rm gormt.exe
	test ! -f gormt_linux.zip || rm gormt_linux.zip
	test ! -f gormt_mac.zip || rm gormt_mac.zip
	test ! -f gormt_windows.zip || rm gormt_windows.zip
	