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

Example usage to use `time_rotating` to trigger automatic rotation of a
`aws_iam_access_key` resource:

```terraform

resource "aws_iam_user" "rotation" {
  name = "some-user"
}

resource "time_rotating" "rotation" {
  rotation_days = 7
}

resource "flipflop" "rotation" {
  value = time_rotating.rotation.id
}

locals {
  rotation_values = [flipflop.rotation.a, flipflop.rotation.b]
}

resource "null_resource" "rotation" {
  count = length(local.rotation_values)
  triggers = {
    user  = aws_iam_user.rotation.name
    value = local.rotation_values[count.index]
  }
}

resource "aws_iam_access_key" "rotation" {
  count = length(local.rotation_values)
  user  = null_resource.rotation[count.index].triggers.user
}

locals {
  rotated_access_key = aws_iam_access_key.rotation[flipflop.rotation.index]
}

```

While this is quite verbose, the complication comes from needing to get
Terraform to only invalidate one leg of the dependency graph at a time.

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
