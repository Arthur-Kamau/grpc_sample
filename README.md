# sample grpc
* two projects that use grpc for communication (web & backend)
benefits fast, have your models auto generated.
* down side http2 not natively supported by browsers but android/ios & desktop apps work okay.
* check envoy for config


# prerequistes
-> latest lts of each.

1. node & npm
2. envoy
3. golang 


# runnig the system
open three instances of the terminal.

1. in project root run `envoy -c envoy/envoy.yaml`
2. in api folder run `go run main.go`
3. in web folder run `npm run dev`