package main

import (
	"net"
)

func authUser(conn *net.UnixConn) (pid int32, username, group string, err error) {
	// GetsockoptUcred is not supported on macOS, but we want to be able to compile
	// it anyway for development/testing purposes. Return some nonsense values here.
	return 0, "cs", "cs", nil
	//return -1, "", "", errors.New("unable to authenticate, SO_PEERCRED not supported on darwin")
}
