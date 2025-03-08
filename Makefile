.PHONY: build
## build: Compile the executables
build:
	go build -ldflags="-s -w -X main.Version=$$(git rev-parse --short HEAD)" -o bin/donezo



.PHONY: sqlc
## sqlc: Generate repository using sqlc
sqlc:
	@sqlc generate



.PHONY: run
## run: Build and run in development mode
run:
	@go run main.go



.PHONY: install
## install: Install TUI
install: build
	@mkdir -p $(HOME)/.local/bin
	@cp ./bin/donezo $(HOME)/.local/bin
	@echo "Installed donezo to '$(HOME)/.local/bin'. Please add '$(HOME)/.local/bin' to your PATH."



.PHONY: uninstall
## uninstall: Uninstall TUI
uninstall:
	@rm $(HOME)/.local/bin/donezo



.PHONY: clean
## clean: Clean project and previous builds
clean:
	@rm bin/*



.PHONY: deps
## deps: Download modules
deps:
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest && \
	@go mod download



.PHONY: help
all: help
# help: show help message
help: Makefile
	@echo
	@echo " Choose a command to run in "$(NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
