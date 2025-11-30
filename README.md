# Terraform Flip-Flop Provider

This is a provider that can be helpful in implementing rolling rotations.

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command: 
```sh
$ go install
```

## Using the provider

The flipflop provider helps implement zero-downtime rolling rotations in Terraform.
It maintains two states (a and b) and tracks which one is currently active, allowing
you to rotate resources gradually instead of all at once.

### Basic Concept

When you change the `value` input:
1. The flipflop updates one output (`a` or `b`) to match the new value
2. The other output remains unchanged (retains the old value)
3. The `index` output indicates which is active (0 for `a`, 1 for `b`)

This enables you to create two instances of a resource and only invalidate one at a time.

### Simple IAM Access Key Rotation

```terraform
resource "aws_iam_user" "example" {
  name = "rotating-user"
}

resource "time_rotating" "rotation" {
  rotation_days = 7
}

resource "flipflop" "rotation" {
  value = time_rotating.rotation.id
}

# Create two access keys
resource "aws_iam_access_key" "rotation" {
  count = 2  # Always create exactly 2
  user  = aws_iam_user.example.name
}

# Use the currently active key
locals {
  active_key = aws_iam_access_key.rotation[flipflop.rotation.index]
}

output "access_key_id" {
  value = local.active_key.id
}

output "secret_access_key" {
  value     = local.active_key.secret
  sensitive = true
}
```

### Advanced Pattern: Encoding Configuration

For more complex scenarios, encode all configuration through the flipflop
using `jsonencode()`. This ensures configuration changes also rotate gradually:

```terraform
resource "flipflop" "config" {
  value = jsonencode({
    trigger = time_rotating.rotation.id
    length  = 32
    special = true
  })
}

locals {
  configs = [
    jsondecode(flipflop.config.a),
    jsondecode(flipflop.config.b),
  ]
}

resource "random_password" "rotating" {
  count   = length(local.configs)
  length  = local.configs[count.index].length
  special = local.configs[count.index].special

  keepers = local.configs[count.index]
}

output "current_password" {
  value     = random_password.rotating[flipflop.config.index].result
  sensitive = true
}
```

See the `examples/` directory for more detailed use cases including password rotation
and handling unknown values with the tri-state variant.

## Developing the Provider

If you wish to work on the provider, you'll first need
[Go](http://www.golang.org) installed on your machine (see
[Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put
the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

## Release Process

Releases are automatically created when commits are pushed to the `main` branch using [semantic-release](https://semantic-release.gitbook.io/). Commit messages must follow the [Conventional Commits](https://www.conventionalcommits.org/) format:

- `feat:` - New feature (minor version bump)
- `fix:` - Bug fix (patch version bump)
- `feat!:` or `BREAKING CHANGE:` - Breaking change (major version bump)
