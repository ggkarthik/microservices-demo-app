#!/bin/bash

echo "============================================"
echo "Investigating Workflow Issues"
echo "============================================"
echo ""

# Check if we're in the right directory
if [ ! -d ".github/workflows" ]; then
    echo "❌ Error: Not in the repository root"
    exit 1
fi

echo "1. Checking GitHub Container Registry Setup..."
echo "-----------------------------------------------"
echo "Repository: ggkarthik/microservices-demo-app"
echo "Registry: ghcr.io"
echo ""
echo "Required permissions for GHCR:"
echo "  ✓ packages: write (in workflow)"
echo "  ✓ contents: read (in workflow)"
echo "  ✓ Repository visibility should be public or have package permissions set"
echo ""

echo "2. Checking workflow files..."
echo "-----------------------------------------------"
cd .github/workflows

for workflow in *.yml *.yaml; do
    if [ -f "$workflow" ]; then
        echo "✓ $workflow exists"
    fi
done
cd ../..
echo ""

echo "3. Checking for common workflow issues..."
echo "-----------------------------------------------"

# Check if Trivy will fail on vulnerabilities
echo "Checking Trivy configuration..."
if grep -q "exit-code: '1'" .github/workflows/build-publish.yml; then
    echo "⚠️  WARNING: Trivy is set to fail builds on CRITICAL/HIGH vulnerabilities"
    echo "   Location: .github/workflows/build-publish.yml"
    echo "   This may cause builds to fail if vulnerabilities are found"
    echo ""
fi

# Check for security vulnerabilities
echo "4. Known Issues:"
echo "-----------------------------------------------"
echo "⚠️  GitHub reported 28 vulnerabilities in the repository:"
echo "   - 3 HIGH severity"
echo "   - 21 MODERATE severity"
echo "   - 4 LOW severity"
echo ""
echo "   URL: https://github.com/ggkarthik/microservices-demo-app/security/dependabot"
echo ""
echo "   These vulnerabilities may cause Trivy scans to fail the build."
echo ""

echo "5. Common Solutions:"
echo "-----------------------------------------------"
echo "Option 1: Temporarily disable Trivy exit code"
echo "  - Change exit-code: '1' to exit-code: '0' in build-publish.yml"
echo "  - This will allow builds to continue despite vulnerabilities"
echo ""
echo "Option 2: Fix vulnerabilities before building"
echo "  - Review Dependabot alerts"
echo "  - Update vulnerable dependencies"
echo "  - Rebuild images with updated dependencies"
echo ""
echo "Option 3: Use ignore-unfixed and continue-on-error"
echo "  - Modify workflow to not fail on vulnerabilities"
echo "  - Log vulnerabilities but continue publishing"
echo ""

echo "6. Checking GHCR Package Visibility Settings:"
echo "-----------------------------------------------"
echo "To publish to GHCR, ensure:"
echo "  1. Go to: https://github.com/settings/packages"
echo "  2. Package visibility should allow public access"
echo "  3. Repository has package write permissions enabled"
echo ""
echo "To check published packages:"
echo "  https://github.com/ggkarthik?tab=packages"
echo ""

echo "7. Workflow Monitoring:"
echo "-----------------------------------------------"
echo "Check workflow status at:"
echo "  https://github.com/ggkarthik/microservices-demo-app/actions"
echo ""
echo "Look for the 'Build and Publish' workflow run triggered by tag v1.0.0"
echo ""

echo "============================================"
echo "Investigation Complete"
echo "============================================"
