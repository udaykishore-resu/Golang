# Recommended Workflow:

# 1. Build all binaries
mage build:all

# 2. Test Docker build locally
mage docker:build

# 3. Create full release (including Docker push)
mage release:full