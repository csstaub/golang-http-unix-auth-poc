package main

import (
	"net"
	"os/user"
	"strconv"
	"syscall"
)

func fileDescriptor(conn *net.UnixConn) int {
	file, err := conn.File()
	if err != nil {
		panic(err)
	}
	return int(file.Fd())
}

func authUser(conn *net.UnixConn) (pid int32, username, group string, err error) {
	cred, err := syscall.GetsockoptUcred(fileDescriptor(conn), syscall.SOL_SOCKET, syscall.SO_PEERCRED)
	if err != nil {
		return -1, "", "", err
	}

	userLookup, err := user.LookupId(strconv.Itoa(int(cred.Uid)))
	if err != nil {
		return -1, "", "", err
	}

	groupLookup, err := user.LookupGroupId(strconv.Itoa(int(cred.Gid)))
	if err != nil {
		return -1, "", "", err
	}

	return cred.Pid, userLookup.Name, groupLookup.Name, nil
}
