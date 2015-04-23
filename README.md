# Simple bot to spam diego's email inbox

# to get dependencies
cd $GOPATH/src/github.com/brenol/annoying-diego
go get

# To run:
> sh start.sh

# Example .env file

> export FROM_ADDR="me@me.com"
> export TO_ADDR="me@me.com,you@you.com"
> export MANDRILL_KEY="my-super-secret-key"
