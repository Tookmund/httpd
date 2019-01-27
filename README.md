# httpd
Simple httpd in Golang

Serves the current directory on `:8080`

Logs every request

Change the address it listens on by passing it as a argument.

E.g. to have the server listen on `127.0.0.1:1234`
```
./httpd 127.0.0.1:1234
```

