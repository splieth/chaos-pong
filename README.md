# chaos-pong

> A classic Pong game with a twist: lose a round and it wreaks havoc on your cloud infrastructure.

## Motivation

Chaos engineering is the practice of deliberately introducing failures into systems to build confidence in their resilience. Tools like Netflix's Chaos Monkey have shown that randomly terminating infrastructure can expose weaknesses before they cause real outages.

But chaos engineering can feel abstract and low-stakes when it's just a cron job running in the background. Chaos Pong makes it personal. Every missed ball triggers real destruction in your cloud account -- terminated EC2 instances, deleted EBS volumes, or whatever chaos actions you've configured. Suddenly, your Pong skills directly determine whether your infrastructure survives.

It's chaos engineering meets gamification: defend your cloud with your paddle.

## Features

- Classic two-player Pong with keyboard controls
- Multiplayer mode via WebSocket (server/client architecture)
- Configurable chaos actions triggered on score
- Multi-provider support (AWS, GCP) via a plugin-based provider system
- YAML configuration for providers, regions, and action selection

## Prerequisites

- Go 1.22+
- AWS credentials configured (for AWS chaos actions)

## Getting Started

```bash
# Build and run the game
make run

# Run tests
make test
```

## Controls

| Key    | Action              |
|--------|---------------------|
| W / S  | Move left paddle    |
| Up / Down | Move right paddle |
| Escape | Quit                |

## Multiplayer

Start the server, then connect two clients:

```bash
# Terminal 1: start the server
make server

# Terminal 2 & 3: connect clients
make client
```

The game begins once two clients are connected.

## Chaos Configuration

Configure which cloud providers and actions to use in `chaos.yaml`:

```yaml
providers:
  aws:
    enabled: true
    region: "eu-central-1"
    profile: "default"
    actions:
      - ec2-instance-terminate
      - ebs-destroy
  gcp:
    enabled: false
    project: "my-project"
    zone: "europe-west1-b"
    actions:
      - instance-terminate
```

### Adding a Provider

Providers register themselves via `init()` using the factory pattern. To add a new provider, create a sub-package under `chaos/` that implements the `Chaos` interface and calls `chaos.RegisterProvider()`.


## License

MIT
