## Simple JWT API Auth
- cp .env-example .env 
- make server

## Create test 
gotests -all -template testify token.go > token_test.go