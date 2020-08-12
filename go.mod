module github.com/bhbosman/goEchoClient

go 1.15

require (
	github.com/bhbosman/gocommon v0.0.0-20200921215456-bfddd9bb050e
	github.com/bhbosman/gocomms v0.0.0-20200922191823-6fd510883c10
	github.com/bhbosman/gologging v0.0.0-20200921180328-d29fc55c00bc
	github.com/bhbosman/gomessageblock v0.0.0-20200921180725-7cd29a998aa3
	github.com/bhbosman/goprotoextra v0.0.1
	github.com/reactivex/rxgo/v2 v2.1.0
	go.uber.org/fx v1.13.1
)

replace github.com/reactivex/rxgo/v2 v2.1.0 => github.com/bhbosman/rxgo/v2 v2.1.1-0.20200922152528-6aef42e76e00
