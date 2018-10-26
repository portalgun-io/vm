package main

import (
	"fmt"
	"log"

	. "github.com/portalgun-io/vm/constants"
	qemu "github.com/portalgun-io/vm"
)

func main() {
	fmt.Println("Example working with images using qemu-img")

	create()
	snapshots()
}

func snapshots() {
	img, err := qemu.OpenImage("vm.qcow2")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("base image", img.Path, "format", img.Format, "size", img.Size)

	err = img.CreateSnapshot("backup")
	if err != nil {
		log.Fatal(err)
	}

	snaps, err := img.Snapshots()
	if err != nil {
		log.Fatal(err)
	}

	for _, snapshot := range snaps {
		fmt.Println(snapshot.Name, snapshot.Date)
	}
}

func create() {
	img := qemu.NewImage("vm.qcow2", qemu.ImageFormatQCOW2, 5 * Gigabyte)
	img.SetBackingFile("vm.qcow2")

	err := img.Create()
	if err != nil {
		log.Fatal(err)
	}
}

