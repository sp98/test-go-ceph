package main

import (
	"fmt"
	"os"

	"github.com/ceph/go-ceph/rados"
	"github.com/ceph/go-ceph/rbd"
)

func main() {
	conn, err := rados.NewConn()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error initiating connection: %v\n", err)
		os.Exit(1)
	}
	conn.ReadDefaultConfigFile()
	conn.Connect()

	ioctx, err := conn.OpenIOContext("data")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error in iotx: %v\n", err)
		os.Exit(1)
	}

	status, err := rbd.GetGlobalMirrorGroupStatus(ioctx, "grp1")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Status %+v", status)
}
