>  kitex -module doushengV4 -service user user.thrift

>  kitex -module doushengV4 -service publish publish.thrift

>  kitex -module doushengV4 -service interact interact.thrift

>  go mod edit -replace=github.com/apache/thrift=github.com/apache/thrift@v0.13.0
