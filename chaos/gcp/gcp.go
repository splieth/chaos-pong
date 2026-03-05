package gcp

import (
	"slices"

	"github.com/splieth/chaos-pong/chaos"
)

func init() {
	chaos.RegisterProvider("gcp", newGCPChaos)
}

func newGCPChaos(cfg chaos.ProviderConfig) []chaos.Chaos {
	project := cfg.Options["project"]
	zone := cfg.Options["zone"]

	var fns []chaos.Chaos
	if len(cfg.Actions) == 0 || slices.Contains(cfg.Actions, "instance-terminate") {
		fns = append(fns, GCPInstanceTerminateChaos{project: project, zone: zone})
	}
	return fns
}

// GCPInstanceTerminateChaos is a placeholder for GCP Compute Engine instance termination.
type GCPInstanceTerminateChaos struct {
	project string
	zone    string
}

func (g GCPInstanceTerminateChaos) Terminate() chaos.Result {
	// TODO: implement with cloud.google.com/go/compute SDK
	return chaos.Result{Success: false, Message: "GCP chaos not yet implemented"}
}
