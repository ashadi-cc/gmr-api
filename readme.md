## Simple API with jwt token auth
- cp .env-example .env 
- make server

## Create test 
gotests -all -template testify token.go > token_test.go

## Linter 
https://freshman.tech/linting-golang/