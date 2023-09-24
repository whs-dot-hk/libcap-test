package main

import (
	"archive/tar"
	"os"
	"io"
	"log"

	"github.com/docker/docker/pkg/system"
)

func main() {
	t, _ := os.Create("./test.tar")
	defer t.Close()
	tw := tar.NewWriter(t)
	defer tw.Close()
	i, _ := os.Lstat("./web");
	hdr, err := tar.FileInfoHeader(i, "")
	if err != nil {
		log.Fatal(err)
	}
	hdr.Xattrs = make(map[string]string)
	capability, err := system.Lgetxattr("./web", "security.capability")
	if err != nil {
		log.Fatal(err)
	}
	hdr.Xattrs["security.capability"] = string(capability)
	if err := tw.WriteHeader(hdr); err != nil {
		log.Fatal(err)
	}
	r, _ := os.Open("./web")
	defer r.Close()
	if _, err := io.Copy(tw, r); err != nil {
		log.Fatal(err)
	}
}
