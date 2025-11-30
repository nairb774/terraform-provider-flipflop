# Tri-State Flipflop Example
#
# The tri-state flipflop resource maintains three values (a, b, c) and can handle
# unknown values at plan time. This is useful when the trigger value depends
# on external data or computed values that aren't known until apply time.
#
# Unlike the standard two-state flipflop, the tri-state version provides three
# indices: top_index, middle_index, and bottom_index, allowing for more complex
# rotation patterns.

# Example using external data that may be unknown at plan time
data "external" "rotation_trigger" {
  program = ["bash", "-c", "echo '{\"timestamp\":\"'$(date +%s)'\"}'"]
}

# The tri-state flipflop can handle unknown values
resource "flipflop_tri" "example" {
  value = data.external.rotation_trigger.result.timestamp
}

# All three states are tracked
output "state_a" {
  value = flipflop_tri.example.a
}

output "state_b" {
  value = flipflop_tri.example.b
}

output "state_c" {
  value = flipflop_tri.example.c
}

# The three indices indicate rotation order
output "top_index" {
  value       = flipflop_tri.example.top_index
  description = "Most recently updated state (0=a, 1=b, 2=c)"
}

output "middle_index" {
  value       = flipflop_tri.example.middle_index
  description = "Second most recently updated state"
}

output "bottom_index" {
  value       = flipflop_tri.example.bottom_index
  description = "Least recently updated state"
}
