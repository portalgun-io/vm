package vm

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"syscall"
	"time"
)

// QEMU VirtualMachine
type VirtualMachine struct {
	Cores          int    // Number of CPU cores
	Memory         uint64 // RAM quantity in megabytes. TODO: Should switch to memory being modules, dimm/etc, size, etc
	RemovableMedia string // TODO: Should likely be removable media
	Display        string
	VNC            string
	Monitor        string
	DiskDevices    []DiskDevice
	NetworkDevices []NetworkDevice
	Custom         [][]string
	CustomOptions  map[string]string
	//Storage    []StorageDevice
	//Networking []NetworkDevice
}

func NewVM(cores int, memory uint64) (machine Machine) {
	machine.Cores = cores
	machine.Memory = memory
	machine.DiskDevices = make([]DiskDevice, 0)

	return machine
}

// TODO: Is this attaching the device or inserting the disk? These should be separate
// actions and functions
func (vm *VirtualMachine) AttachCDRom(device string) {
	vm.RemovableMedia = device
}

// AddDrive attaches a new hard drive to
// the virtual machine
//TODO: Should it not be insert disk? Or at least .AttachDiskDevice()
func (vm *VirtualMachine) AttachDiskDevice(d Disk) {
	vm.drives = append(vm.Disks, d)
}

// AddDriveImage attaches the specified Image to
// the virtual machine
func (vm *VirtualMachine) AddDiskImage(i Image) {
	vm.Disks = append(vm.Disks, Disk{i.Path, i.Format})
}

func (vm *VirtualMachine) AttachNetworkDevice(netdev NetworkDevice) {
	vm.NetworkDevices = append(vm.NetworkDevices, netdev)
}

func (vm *VirtualMachine) SetDisplay(mode string) {
	vm.display = mode
}

// AddVNC attaches a VNC server to
// the virtual machine, bound to the specified address and port
// If wsPort is not 0, VNC will work over WebSocket on that port
// TODO: VNC is fairly complex and should be moved into its own struct/interface
// especially considering we want to abandon it for spice essentially exclusively
func (vm *VirtualMachine) EnableVNC(address string, port, wsPort int, passwd bool) {
	vm.VNC = fmt.Sprintf("%s:%d", address, port)
	// TODO: Why not just make this uint?
	if wsPort > 0 {
		vm.VNC = fmt.Sprintf("%s,websocket=%d", vm.VNC, wsPort)
		if passwd {
			vm.VNC = fmt.Sprintf("%s,password", vm.VNC)
		}
	}
}

// AddMonitor redirects the QEMU monitor
// to the specified unix socket file
func (vm *VirtualMachine) AddMonitorUnix(dev string) {
	vm.Monitor = dev
}

// TODO: Why not just use a map instead of a 2D Array?
func (vm *VirtualMachine) AddOption(opt, val string) {
	vm.Custom = append(vm.Custom, []string{opt, val})
}

// Start stars the machine
// The 'kvm' bool specifies if KVM should be used
// It returns the QEMU process and an error (if any)
// TODO: This function needs A LOT of work

// TODO: Find all available hypervisors to launch VM with

// TODO: Creating the command-line should be separate from the execution of said command
func (vm *VirtualMachine) Start(arch string, kvm bool, stderrCb func(s string)) (*os.Process, error) {
	// TODO: arch should be attribute of VM
	qemu := fmt.Sprintf("qemu-system-%s", arch)
	args := []string{"-smp", strconv.Itoa(vm.Cores), "-m", strconv.FormatUint(vm.Memory, 10)}

	if kvm {
		args = append(args, "-enable-kvm")
	}

	if len(vm.RemovableMedia) > 0 {
		args = append(args, "-cdrom")
		args = append(args, vm.RemovableMedia)
	}

	for _, d := range vm.DiskDevices {
		args = append(args, "-drive")
		args = append(args, fmt.Sprintf("file=%s,format=%s", d.Path, d.Format))
	}

	if len(vm.NetworkDevices) == 0 {
		args = append(args, "-net")
		args = append(args, "none")
	}

	for _, iface := range vm.NetworkDevices {
		s := fmt.Sprintf("%s,id=%s", iface.Type, iface.ID)
		if len(iface.Name) > 0 {
			s = fmt.Sprintf("%s,ifname=%s", s, iface.Name)
		}

		args = append(args, "-netdev")
		args = append(args, s)

		s = fmt.Sprintf("virtio-net,netdev=%s", iface.ID)
		// TODO: Need to add realistic mac address generation
		if len(iface.MAC) > 0 {
			s = fmt.Sprintf("%s,mac=%s", s, iface.MAC)
		}

		args = append(args, "-device")
		args = append(args, s)
	}

	if len(vm.VNC) > 0 {
		args = append(args, "-vnc")
		args = append(args, vm.VNC)
	} else if len(vm.Display) == 0 {
		args = append(args, "-display")
		args = append(args, "none")
	}

	if len(vm.Display) > 0 {
		args = append(args, "-display")
		args = append(args, vm.Display)
	}

	if len(vm.Monitor) > 0 {
		args = append(args, "-qmp")
		args = append(args, fmt.Sprintf("unix:%s,server,nowait", vm.Monitor))
	}

	for _, c := range vm.Custom {
		args = append(args, c[0])
		args = append(args, c[1])
	}

	cmd := exec.Command(qemu, args...)

	// TODO: Set env just for this exec using cmd.Env = append(os.Environ(), "FOO=value", "OTHERFOO=thing")

	cmd.SysProcAttr = new(syscall.SysProcAttr)
	cmd.SysProcAttr.Setsid = true

	stderr, err := cmd.StderrPipe()
	if err == nil {
		go func() {
			s, err := ioutil.ReadAll(stderr)
			if err != nil {
				return
			}

			stderrCb(string(s))
		}()
	}

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	proc := cmd.Process
	errc := make(chan error)

	go func() {
		err := cmd.Wait()
		if err != nil {
			errc <- fmt.Errorf("'qemu-system-%s': %s", arch, err)
			return
		}
	}()

	time.Sleep(50 * time.Millisecond)

	var vmerr error
	select {
	case vmerr = <-errc:
		if vmerr != nil {
			return nil, vmerr
		}
	default:
		break
	}

	return proc, nil
}
