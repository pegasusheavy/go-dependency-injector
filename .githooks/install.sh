#!/bin/bash
#
# Git hooks installation script for go-dependency-injector
#
# This script configures git to use the hooks in .githooks/
# Run this once after cloning the repository.
#

set -e

YELLOW='\033[1;33m'
GREEN='\033[0;32m'
RED='\033[0;31m'
CYAN='\033[0;36m'
NC='\033[0m'

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(dirname "$SCRIPT_DIR")"

echo -e "${CYAN}╔══════════════════════════════════════════════════════════════╗${NC}"
echo -e "${CYAN}║           Git Hooks Installation                            ║${NC}"
echo -e "${CYAN}║           go-dependency-injector                            ║${NC}"
echo -e "${CYAN}╚══════════════════════════════════════════════════════════════╝${NC}"
echo ""

# Check if we're in a git repository
if ! git rev-parse --git-dir > /dev/null 2>&1; then
    echo -e "${RED}Error: Not a git repository${NC}"
    exit 1
fi

# Make hooks executable
echo -e "${YELLOW}Making hooks executable...${NC}"
chmod +x "$SCRIPT_DIR/pre-commit" 2>/dev/null || true
chmod +x "$SCRIPT_DIR/pre-push" 2>/dev/null || true
chmod +x "$SCRIPT_DIR/commit-msg" 2>/dev/null || true

# Configure git to use .githooks directory
echo -e "${YELLOW}Configuring git hooks path...${NC}"
git config core.hooksPath .githooks

echo ""
echo -e "${GREEN}✓ Git hooks installed successfully!${NC}"
echo ""
echo -e "${CYAN}Installed hooks:${NC}"
echo -e "  ${GREEN}pre-commit${NC}  - Formatting, vetting, and tests"
echo -e "  ${GREEN}pre-push${NC}    - Full test suite with race detection"
echo -e "  ${GREEN}commit-msg${NC}  - Conventional commit message validation"
echo ""
echo -e "${CYAN}To uninstall, run:${NC}"
echo -e "  git config --unset core.hooksPath"
echo ""
echo -e "${CYAN}To skip hooks temporarily:${NC}"
echo -e "  git commit --no-verify"
echo -e "  git push --no-verify"
echo ""


