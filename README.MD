# protoc-gen-go-json

This tool generates Go language bindings of services in protobuf definition files for json.

# Installation
```shell
$ go install github.com/x64fun/protoc-gen-go-json
```

# Usage
```
protoc --go_out=. --go_opt=paths=source_relative \
       --go-json_out=. --go-json_opt=paths=source_relative \
       *.proto
```