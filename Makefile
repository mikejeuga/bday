repo=$(shell basename "`pwd`")

gopher:
	@git init
	@touch .gitignore
	@touch README.md
	@touch main.go
	@go mod init github.com/mikejeuga/$(repo)
	@go get github.com/google/uui
	@go get github.com/adamluzsi/testcase
	@go get github.com/gorilla/mux
	@go mod tidy

run:
	@go run ./cmd/main.go

t: test
test:
	@make ut at

ut: unit-test
unit-test:
	@go test -v -tags=unit ./...

at: acceptance-test
acceptance-test:
	@docker-compose -f docker-compose.yml up -d
	@go test -v -tags=acceptance ./...
	@docker-compose down

ic: init
init:
	@gh repo create ${repo} --private
	@git add .
	@git commit -m "Initial commit"
	@git remote add origin git@github.com:mikejeuga/${repo}.git
	@git branch -M main
	@git push -u origin main

c: commit
commit:
	@git add .
	@git commit -m "$m"
	@git pull --rebase
	git push

privacy:
	@gh repo edit --visisbility=private

public:
	@gh repo edit --visisbility=public

destroy:
	rm -rf .git
	@gh repo delete ${repo}