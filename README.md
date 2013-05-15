# Golang library for the Twitter 1.1 API

This project aims to be a [Go][1] library for the Twitter [1.1][2] API.

## Want to help?

Register your APP at https://dev.twitter.com/apps to get your Key and Secret.

Leave "Callback URL" blank, you don't need it for the command line tests.

```sh
# Getting the source
go get github.com/xiam/twitter

# Installing the twitter command
go install github.com/xiam/twitter/cli/twitter

# Go to the pkg directory.
cd $GOPATH/src/github.com/xiam/twitter

# Use the twitter command to retrive an API token.
# This will ask you for a PIN and will give you your user credentials.
twitter -key $APP_KEY -secret $APP_SECRET

# Create a settings.yaml file with your app keys and user credentials.
cat settings.yaml
twitter:
  app:
    key: ZerGYGhZytwFrsaR4xAse
    secret: PCadfTgdxAsercATs4Asr5dAx
  user:
    token: 12345678-rOaRx4saKTTuNJlhiuI7ehumzOV5xSp6dOtlk1Rs
    secret: fmt5pMcEbXer4DmmRFls7KesjXcQ4utgqrTf0KcR8

# Run the tests.
go test

# Hack what you need and send pull requests :-).
vim main.go
```

This is not production ready.

See the docs at http://godoc.org/github.com/xiam/twitter

[1]: http://golang.org
[2]: https://dev.twitter.com/docs/api/1.1

