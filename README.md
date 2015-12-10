To reproduce:

```sh
$ export GOPATH=$PWD
$ rm -rf $GOPATH/pkg
$ go install one # success
$ go run repro.go
ERROR reading export data: /Users/jonasi/dev/src/github.com/jonasi/golang-import-error/pkg/darwin_amd64/one.a: import error /Users/jonasi/dev/src/github.com/jonasi/golang-import-error/pkg/darwin_amd64/one.a:3:86 (byte offset = 115): encoding/json package not found
```
