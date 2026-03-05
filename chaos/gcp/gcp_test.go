package gcp

import (
	"github.com/splieth/chaos-pong/chaos"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGCPInstanceTerminateChaos_Terminate(t *testing.T) {
	g := GCPInstanceTerminateChaos{}
	result := g.Terminate()
	assert.False(t, result.Success)
	assert.Equal(t, "GCP chaos not yet implemented", result.Message)
}

func TestNewGCPChaos_AllActions(t *testing.T) {
	cfg := chaos.ProviderConfig{
		Enabled: true,
		Options: map[string]string{
			"project": "test-project",
			"zone":    "us-central1-a",
		},
	}
	fns := newGCPChaos(cfg)
	assert.Equal(t, 1, len(fns))
}

func TestNewGCPChaos_FilteredActions(t *testing.T) {
	cfg := chaos.ProviderConfig{
		Enabled: true,
		Actions: []string{"nonexistent-action"},
		Options: map[string]string{},
	}
	fns := newGCPChaos(cfg)
	assert.Equal(t, 0, len(fns))
}
