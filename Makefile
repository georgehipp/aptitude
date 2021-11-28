build:
	@echo "Compile binary for Windows, Darwin, Linux"
	@[ -e bin ] || mkdir bin
	@GOARCH=amd64 GOOS=windows go build -o aptitude cmd/*.go
	@mv aptitude.exe bin/win_aptitude.exe
	@GOARCH=amd64 GOOS=darwin go build -o aptitude cmd/*.go
	@mv aptitude bin/dar_aptitude
	@GOARCH=amd64 GOOS=linux go build -o aptitude cmd/*.go
	@mv aptitude bin/lin_aptitude

run:
	@echo "Execute from Source"
	@cd cmd; go run *.go; cd ..

godocs:
	@echo "Documentation Generated with go doc"
	@cd internal/number; go doc -all; cd ..
	@cd internal/perception; go doc -all; cd ..
	@cd internal/reason; go doc -all; cd ..
	@cd internal/spatial; go doc -all; cd ..
	@cd internal/utility; go doc -all; cd ..
	@cd internal/word; go doc -all; cd ..
	@cd cmd; go doc -all; cd ..

gotest:
	@echo "Test golang code"
	@cd internal/**; go test -v
