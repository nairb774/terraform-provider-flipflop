# Password Rotation with Flipflop
#
# This example shows how to use flipflop to implement automatic password
# rotation. By encoding all password parameters through the flipflop, we ensure
# that only one password changes at a time during rotation.

# Trigger rotation every 30 days
resource "time_rotating" "password_rotation" {
  rotation_days = 30
}

# Encode all password configuration through flipflop
resource "flipflop" "password" {
  value = jsonencode({
    trigger = time_rotating.password_rotation.id
    length  = 32
    special = true
  })
}

# Decode the configuration for both states
locals {
  password_configs = [
    jsondecode(flipflop.password.a),
    jsondecode(flipflop.password.b),
  ]
}

# Create two passwords, one for each state
resource "random_password" "rotating" {
  count = length(local.password_configs)

  length  = local.password_configs[count.index].length
  special = local.password_configs[count.index].special

  keepers = {
    trigger = local.password_configs[count.index].trigger
  }
}

# Output the currently active password
output "current_password" {
  value     = random_password.rotating[flipflop.password.index].result
  sensitive = true
}

# Output both passwords (useful during rotation period)
output "all_passwords" {
  value     = [for p in random_password.rotating : p.result]
  sensitive = true
}
