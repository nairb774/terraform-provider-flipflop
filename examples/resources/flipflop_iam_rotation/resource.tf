# AWS IAM Access Key Rotation Example
#
# This example demonstrates how to use the flipflop resource to implement
# automatic rotation of AWS IAM access keys. The pattern ensures zero-downtime
# rotation by maintaining two valid keys during the transition period.

resource "aws_iam_user" "example" {
  name = "rotating-user"
}

# Trigger rotation every 7 days
resource "time_rotating" "rotation" {
  rotation_days = 7
}

# The flipflop resource tracks the rotation state
resource "flipflop" "rotation" {
  value = time_rotating.rotation.id
}

# Create two access keys, one for each flipflop state
locals {
  rotation_values = [flipflop.rotation.a, flipflop.rotation.b]
}

# Use null_resource to capture the rotation values as triggers
# This ensures the access keys are recreated when the flipflop values change
resource "null_resource" "rotation" {
  count = length(local.rotation_values)
  triggers = {
    user  = aws_iam_user.example.name
    value = local.rotation_values[count.index]
  }
}

resource "aws_iam_access_key" "rotation" {
  count = length(local.rotation_values)
  user  = null_resource.rotation[count.index].triggers.user

  # This lifecycle ensures keys are created before old ones are destroyed
  lifecycle {
    create_before_destroy = true
  }
}

# Output the currently active access key
output "active_access_key_id" {
  value = aws_iam_access_key.rotation[flipflop.rotation.index].id
}

output "active_secret_access_key" {
  value     = aws_iam_access_key.rotation[flipflop.rotation.index].secret
  sensitive = true
}
