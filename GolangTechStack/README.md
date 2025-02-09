# Recommended Workflow:
# 1. List mage commands
mage -l
# 2. Build all binaries
mage build:all         
# 3. Removes all binaries
mage build:clean       
# 4. Test Docker build locally
mage docker:build      
# 5. Creates a new release using goreleaser in snapshot mode
mage release:create    
# 6. Run go tests
mage test        
# 7. Create full release (including Docker push)
mage release:full      