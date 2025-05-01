package security.policies

default allow = false

# Example policy: Only allow participants with role "admin" to perform "delete" actions
allow {
    input.action == "delete"
    input.user.role == "admin"
}

# Example policy: Allow read actions for all authenticated users
allow {
    input.action == "read"
    input.user.authenticated == true
}

# Example policy: Deny all other actions by default
deny {
    not allow
}
