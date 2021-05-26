package vsphere_vm

type VMS map[string]*VMExportData

type VMExportData struct {
	UUid string
	IPAddress string
	Addresses []string
}