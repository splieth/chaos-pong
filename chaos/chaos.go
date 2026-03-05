package chaos

import (
	"fmt"
	"math/rand"
)

type Chaos interface {
	Terminate() Result
}

type Result struct {
	Success bool
	Message string
}

// ProviderFactory creates chaos functions from a provider config.
// It receives the provider-specific config and returns the chaos
// actions that should be registered.
type ProviderFactory func(ProviderConfig) []Chaos

var providers = map[string]ProviderFactory{}
var functions []Chaos

// RegisterProvider registers a provider factory under the given name.
// Providers call this from their init() functions.
func RegisterProvider(name string, factory ProviderFactory) {
	providers[name] = factory
}

// LoadAndRegister loads a config file and registers chaos functions
// from all enabled providers.
func LoadAndRegister(configPath string) error {
	cfg, err := LoadConfig(configPath)
	if err != nil {
		return fmt.Errorf("loading chaos config: %w", err)
	}

	for name, providerCfg := range cfg.Providers {
		if !providerCfg.Enabled {
			continue
		}
		factory, ok := providers[name]
		if !ok {
			return fmt.Errorf("unknown chaos provider: %s", name)
		}
		fns := factory(providerCfg)
		functions = append(functions, fns...)
	}

	return nil
}

func Random() {
	if len(functions) == 0 {
		return
	}
	fn := functions[rand.Intn(len(functions))]
	fn.Terminate()
}
