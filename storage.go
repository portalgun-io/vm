package vm

type StorageDevice interface {
	AttachDevice()
	DetachDevice()
}

type RemovableMedia struct {
}

type RemovableMediaDevice struct {
}

type DiskDevice struct {
	Path   string // Image file path
	Format string // Image format
}
