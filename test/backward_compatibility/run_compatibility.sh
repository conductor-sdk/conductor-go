#!/bin/bash
set -e

echo "🚀 Starting Backward Compatibility Tests"
echo "========================================"

# Function to get latest release version
get_latest_version() {
    LATEST_RELEASE=$(curl -s https://api.github.com/repos/conductor-sdk/conductor-go/releases/latest | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/' 2>/dev/null || echo "")

    if [ ! -z "$LATEST_RELEASE" ] && [ "$LATEST_RELEASE" != "null" ]; then
        echo "$LATEST_RELEASE"
        return
    fi

    echo "v1.5.4"
}

echo "🔍 Detecting latest released version..."
LATEST_RELEASE=$(get_latest_version)
echo "✓ Detected latest version: $LATEST_RELEASE"

CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)

echo ""
echo "📋 Test Configuration:"
echo "   Released Version: $LATEST_RELEASE"
echo "   Current Branch:   $CURRENT_BRANCH"
echo ""

# Phase 1: Test with released SDK
echo "📦 Phase 1: Testing with Released SDK ($LATEST_RELEASE)"
echo "--------------------------------------------------------"

cd releasedVersion

# Copy shared source code
echo "📋 Copying shared source code..."
cp -r ../shared/* .
echo "✓ Source code copied"

# Create go.mod for released SDK
echo "📝 Creating go.mod for released SDK..."
cat > go.mod << EOF
module conductor-backward-compatibility-test

go 1.17

require (
	github.com/conductor-sdk/conductor-go $LATEST_RELEASE
	github.com/sirupsen/logrus v1.9.3
	github.com/antihax/optional v1.0.0
)
EOF

echo "📥 Downloading released SDK dependencies..."
go mod tidy
echo "✓ Dependencies resolved for released SDK ($LATEST_RELEASE)"

echo ""
echo "🧪 Running compatibility test with released SDK..."
echo "================================================="
go run compatibility.go
echo ""
echo "✅ Phase 1 PASSED: Released SDK ($LATEST_RELEASE) test successful"

# Cleanup copied files
rm -f compatibility.go
rm -rf src/

# Phase 2: Test with current code
echo ""
echo "🔧 Phase 2: Testing with Current Code ($CURRENT_BRANCH)"
echo "--------------------------------------------------------"

cd ../currentCodeVersion

# Copy shared source code
echo "📋 Copying shared source code..."
cp -r ../shared/* .
echo "✓ Source code copied"

# Create go.mod for current code
echo "📝 Creating go.mod for current code..."
cat > go.mod << EOF
module conductor-backward-compatibility-test

go 1.17

require (
	github.com/conductor-sdk/conductor-go v0.0.0
	github.com/sirupsen/logrus v1.9.3
	github.com/antihax/optional v1.0.0
)

replace github.com/conductor-sdk/conductor-go => ../../..
EOF

echo "📥 Setting up current code dependencies..."
go mod tidy
echo "✓ Dependencies resolved for current SDK"

echo ""
echo "🧪 Running compatibility test with current SDK..."
echo "================================================="
go run compatibility.go
echo ""
echo "✅ Phase 2 PASSED: Current SDK test successful"

# Cleanup copied files
rm -f compatibility.go
rm -rf src/

# Success
echo ""
echo "🎉 BACKWARD COMPATIBILITY CONFIRMED!"
echo "===================================="
echo "✓ Released SDK ($LATEST_RELEASE) tests passed"
echo "✓ Current branch ($CURRENT_BRANCH) tests passed"
echo "✓ No breaking changes detected"
echo ""
echo "🚀 Your changes are safe to merge!"