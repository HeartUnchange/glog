glog-change
====

## ChangeLog
1. Remove ```flag``` require, now initialization just with default value
2. Add config API instead ```flag.Parse```

## Usage
just pass the arguments what passed by command line in the past to the new API
```go
glog.Config().SetLogDir("/var/log") // set log dir
```
