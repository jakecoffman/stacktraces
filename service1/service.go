package service1

import "github.com/jakecoffman/stacktraces/service2"

func DoStuff1() {
	service2.DoStuff2()
}
