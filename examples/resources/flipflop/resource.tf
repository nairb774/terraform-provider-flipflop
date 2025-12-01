# Basic flipflop usage
# The flipflop resource tracks state changes and maintains two values (a and b)
# When the value changes, one of the outputs updates while the other remains stable
resource "flipflop" "example" {
  value = "trigger-value"
}

# Outputs demonstrate the two-state tracking
output "state_a" {
  value = flipflop.example.a
}

output "state_b" {
  value = flipflop.example.b
}

# The index indicates which state is currently active (0 = a, 1 = b)
output "active_index" {
  value = flipflop.example.index
}
