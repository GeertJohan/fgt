## fgt
fgt runs any command for you and exits with exitcode 1 when the child process sent anythingto stdout or stderr

### Installation
`go install github.com/GeertJohan/fgt`


### Usage
Some examples:

`fgt echo` will return successfully

`fgt echo hi` will return with exitcode 1

`fgt gofmt -l <yourpackage>` will return with exitcode 1 when gofmt needs something changed