package chaos

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	yaml := `providers:
  aws:
    enabled: true
    region: "eu-central-1"
    profile: "default"
    actions:
      - ec2-instance-terminate
  gcp:
    enabled: false
    project: "my-project"
    zone: "us-central1-a"
    actions:
      - instance-terminate
`
	path := filepath.Join(t.TempDir(), "chaos.yaml")
	os.WriteFile(path, []byte(yaml), 0644)

	cfg, err := LoadConfig(path)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(cfg.Providers))

	aws := cfg.Providers["aws"]
	assert.True(t, aws.Enabled)
	assert.Equal(t, "eu-central-1", aws.Options["region"])
	assert.Equal(t, "default", aws.Options["profile"])
	assert.Equal(t, []string{"ec2-instance-terminate"}, aws.Actions)

	gcp := cfg.Providers["gcp"]
	assert.False(t, gcp.Enabled)
	assert.Equal(t, "my-project", gcp.Options["project"])
	assert.Equal(t, "us-central1-a", gcp.Options["zone"])
}

func TestLoadConfig_FileNotFound(t *testing.T) {
	_, err := LoadConfig("/nonexistent/path.yaml")
	assert.Error(t, err)
}

func TestLoadAndRegister_DisabledProvider(t *testing.T) {
	yaml := `providers:
  testprovider:
    enabled: false
    actions:
      - some-action
`
	path := filepath.Join(t.TempDir(), "chaos.yaml")
	os.WriteFile(path, []byte(yaml), 0644)

	// Reset global state
	functions = nil

	err := LoadAndRegister(path)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(functions))
}

func TestLoadAndRegister_UnknownProvider(t *testing.T) {
	yaml := `providers:
  unknown:
    enabled: true
    actions:
      - some-action
`
	path := filepath.Join(t.TempDir(), "chaos.yaml")
	os.WriteFile(path, []byte(yaml), 0644)

	// Reset global state
	functions = nil

	err := LoadAndRegister(path)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unknown chaos provider")
}

func TestLoadAndRegister_WithFactory(t *testing.T) {
	yaml := `providers:
  testprovider:
    enabled: true
    region: "us-east-1"
    actions:
      - test-action
`
	path := filepath.Join(t.TempDir(), "chaos.yaml")
	os.WriteFile(path, []byte(yaml), 0644)

	// Reset global state
	functions = nil
	oldProviders := providers
	providers = map[string]ProviderFactory{}
	defer func() { providers = oldProviders }()

	RegisterProvider("testprovider", func(cfg ProviderConfig) []Chaos {
		assert.Equal(t, "us-east-1", cfg.Options["region"])
		assert.Equal(t, []string{"test-action"}, cfg.Actions)
		return []Chaos{testChaos{}}
	})

	err := LoadAndRegister(path)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(functions))
}

type testChaos struct{}

func (t testChaos) Terminate() Result {
	return Result{Success: true, Message: "test"}
}
