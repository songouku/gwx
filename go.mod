module gwx

go 1.12

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190513172903-22d7a77e9e5f
	golang.org/x/net => github.com/golang/net v0.0.0-20190514140710-3ec191127204
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190509141414-a5b02f93d862
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190420181800-aa740d480789
	google.golang.org/appengine => github.com/golang/appengine v1.6.0
	stathat.com/c/consistent => github.com/stathat/consistent v1.0.0
)

require github.com/gogf/gf v1.11.4
