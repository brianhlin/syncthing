// Copyright (C) 2015 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package osutil

import (
	"net"
	"net/url"
	"time"
)

// TCPPing returns the duration required to establish a TCP connection
// to the given host. ICMP packets require root priviledges, hence why we use
// tcp.
func TCPPing(address string) (time.Duration, error) {
	dialer := net.Dialer{
		Deadline: time.Now().Add(time.Second),
	}
	start := time.Now()
	conn, err := dialer.Dial("tcp", address)
	if conn != nil {
		conn.Close()
	}
	return time.Since(start), err
}

// GetLatencyForURL parses the given URL, tries opening a TCP connection to it
// and returns the time it took to establish the connection.
func GetLatencyForURL(addr string) (time.Duration, error) {
	uri, err := url.Parse(addr)
	if err != nil {
		return 0, err
	}

	return TCPPing(uri.Host)
}
