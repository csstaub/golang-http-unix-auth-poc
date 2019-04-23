package main

import (
	"fmt"
	"net"
	"time"
)

type authenticatedListener struct {
	listener *net.UnixListener
}

type authenticatedConn struct {
	conn       net.Conn
	remoteAddr authenticatedUser
}

type authenticatedUser struct {
	remoteAddr  net.Addr
	pid         int32
	user, group string
}

func (a *authenticatedListener) Accept() (net.Conn, error) {
	conn, err := a.listener.Accept()
	if err != nil {
		return nil, err
	}

	pid, user, group, err := authUser(conn.(*net.UnixConn))
	if err != nil {
		return nil, err
	}

	return &authenticatedConn{
		conn: conn,
		remoteAddr: authenticatedUser{
			remoteAddr: conn.LocalAddr(),
			pid:        pid,
			user:       user,
			group:      group,
		},
	}, nil
}

func (a *authenticatedListener) Close() error {
	return a.listener.Close()
}

func (a *authenticatedListener) Addr() net.Addr {
	return a.listener.Addr()
}

func (a *authenticatedConn) Read(b []byte) (n int, err error) {
	return a.conn.Read(b)
}

func (a *authenticatedConn) Write(b []byte) (n int, err error) {
	return a.conn.Write(b)
}

func (a *authenticatedConn) Close() error {
	return a.conn.Close()
}

func (a *authenticatedConn) LocalAddr() net.Addr {
	return a.conn.LocalAddr()
}

func (a *authenticatedConn) RemoteAddr() net.Addr {
	return &a.remoteAddr
}

func (a *authenticatedConn) SetDeadline(t time.Time) error {
	return a.conn.SetDeadline(t)
}

func (a *authenticatedConn) SetReadDeadline(t time.Time) error {
	return a.conn.SetReadDeadline(t)
}

func (a *authenticatedConn) SetWriteDeadline(t time.Time) error {
	return a.conn.SetWriteDeadline(t)
}

func (a *authenticatedUser) Network() string {
	return a.remoteAddr.Network()
}

func (a *authenticatedUser) String() string {
	return fmt.Sprintf("%s:%s:%s", a.remoteAddr.String(), a.user, a.group)
}
