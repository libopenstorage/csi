package api

/*
  #include <libcgroup.h>
  #cgo LDFLAGS: -lcgroup
*/
import "C"

// Cgroup is the structure describing one or more control groups. The structure
// is opaque to applications.
type Cgroup struct {
	g *C.struct_cgroup
}

// Device structure represents the type storage being provided to the
// data service.
type Device struct {
	// Type could be a string such as "block", "ebs", "nfs" etc.  It is
	// up to the data service to interpret the device type.
	Type string

	// Metadata contains device type specific constraints and information.
	// For example, for an EBS volume type, it can contain the AWS access keys.
	Metadata map[string]string
}

// Geography physical location of the node.
type Geography struct {
	Zone string
	Rack int
}

// Node contains details regarding a specific host.  This information
// will be provided to the service being deployed.  This information is
// provided as a file on the host in yaml format.
type Node struct {

	// ID is unique node identifier.
	ID string

	// IPs is the list of IPs for this node.
	IPs []string

	// Geography is the physical location of the node.
	Geography Geogrpahy

	// Devices is the list of devices.
	Devices []Device

	// Constraints are cgroup restriuctions on the service container.
	Constraints Cgroup

	// ClusterID uniquely identifies the cluster that this
	// data service is part of.
	ClusterID string

	// Metadata provides arbitrary name value pairs.
	Metadata map[string]string
}

// Bootstrap contains information for the scheduler.  It instructs the scheduler
// to deploy a given service on a set of nodes from a container image.
type Bootstrap struct {
	// DataServiceName data service name.
	DataServiceName string

	// Image name to execute on the host.
	Image string
}

// Installer will manage the provisioning of data services on a set of
// machines as per the bootstrap and node information.
type Installer interface {
	// Deploy data service to set of nodes.
	Deploy(b *Bootstrap, Nodes []Node)

	// Upgrade data service on set of nodes.
	Upgrade(b *Bootstrap, Nodes []Node)

	// Destroy data service on set of nodes.
	Destroy(b *Bootstrap, Nodes []Node)
}
