To reproduce:

```sh
~/d/s/g/j/golang-import-error ❯❯❯ GOPATH=$PWD go install one
~/d/s/g/j/golang-import-error ❯❯❯ GOPATH=$PWD gotype ./src/one                                                                                                                ⏎
src/one/one.go:3:8: could not import two (reading export data: /Users/jonasi/dev/src/github.com/jonasi/golang-import-error/pkg/darwin_amd64/two.a: import error /Users/jonasi/dev/src/github.com/jonasi/golang-import-error/pkg/darwin_amd64/two.a:4:99 (byte offset = 158): encoding/json package not found)
src/one/one.go:5:8: undeclared name: two
```
