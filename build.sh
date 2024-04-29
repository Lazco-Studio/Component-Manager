#!/bin/bash

# Colors
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
RED='\033[0;31m'
BLUE='\033[0;34m'
MAGENTA='\033[0;35m'
NC='\033[0m' # No Color

# Symbols
CHECKMARK='\xE2\x9C\x94'
ROCKET='\xF0\x9F\x9A\x80'
CROSSMARK='\xE2\x9C\x98'
HOURGLASS='\xE2\x8C\x9B'

# Check if GITHUB_TOKEN exists
if ! [ -f env/GITHUB_TOKEN ]; then
  echo -e "${RED}${CROSSMARK} ${BLUE}GITHUB_TOKEN ${YELLOW}not found under env folder. Please add it and try again.${NC}"
  echo -e "${RED}${CROSSMARK} Installation aborted.${NC}"
  exit 1
else
  echo -e "${BLUE}${CHECKMARK} GITHUB_TOKEN ${GREEN}found under env folder.${NC}"
  GITHUB_TOKEN=$(cat env/GITHUB_TOKEN)
fi

# Create necessary folders
mkdir -p dist/uncompressed
mkdir -p dist/compressed

# Build binaries
echo -e "${MAGENTA}Building binaries...${NC}"
# linux/amd64
echo -e "${MAGENTA}${ROCKET} Processing ${BLUE}linux/amd64${MAGENTA}...${NC}"
CGO_ENABLED=0  GOOS=linux    GOARCH=amd64 go build -ldflags "-s -w -X main.GITHUB_TOKEN=$GITHUB_TOKEN" -o ./dist/uncompressed/cm-cli_linux_amd64
echo -e "${GREEN}${CHECKMARK} File saved as ${BLUE}./dist/uncompressed/cm-cli_linux_amd64${NC}"
# darwin/amd64
echo -e "${MAGENTA}${ROCKET} Processing ${BLUE}darwin/amd64${MAGENTA}...${NC}"
CGO_ENABLED=0  GOOS=darwin   GOARCH=amd64 go build -ldflags "-s -w -X main.GITHUB_TOKEN=$GITHUB_TOKEN" -o ./dist/uncompressed/cm-cli_darwin_amd64
echo -e "${GREEN}${CHECKMARK} File saved as ${BLUE}./dist/uncompressed/cm-cli_darwin_amd64${NC}"
# windows/amd64
echo -e "${MAGENTA}${ROCKET} Processing ${BLUE}windows/amd64${MAGENTA}...${NC}"
CGO_ENABLED=0  GOOS=windows  GOARCH=amd64 go build -ldflags "-s -w -X main.GITHUB_TOKEN=$GITHUB_TOKEN" -o ./dist/uncompressed/cm-cli_windows_amd64
echo -e "${GREEN}${CHECKMARK} File saved as ${BLUE}./dist/uncompressed/cm-cli_windows_amd64${NC}"

# Compress binaries
echo -e "${MAGENTA}Compressing binaries...${NC}"
# linux/amd64
echo -e "${MAGENTA}${ROCKET} Processing ${BLUE}linux/amd64${MAGENTA}...${NC}"
upx -f --best --lzma ./dist/uncompressed/cm-cli_linux_amd64 -o ./dist/compressed/cm-cli_linux_amd64
echo -e "${GREEN}${CHECKMARK} File saved as ${BLUE}./dist/compressed/cm-cli_linux_amd64${NC}"
# darwin/amd64
echo -e "${MAGENTA}${ROCKET} Processing ${BLUE}darwin/amd64${MAGENTA}...${NC}"
upx -f --best --lzma ./dist/uncompressed/cm-cli_darwin_amd64 -o ./dist/compressed/cm-cli_darwin_amd64 --force-macos
echo -e "${GREEN}${CHECKMARK} File saved as ${BLUE}./dist/compressed/cm-cli_darwin_amd64${NC}"
# windows/amd64
echo -e "${MAGENTA}${ROCKET} Processing ${BLUE}windows/amd64${MAGENTA}...${NC}"
upx -f --best --lzma ./dist/uncompressed/cm-cli_windows_amd64 -o ./dist/compressed/cm-cli_windows_amd64
echo -e "${GREEN}${CHECKMARK} File saved as ${BLUE}./dist/compressed/cm-cli_windows_amd64${NC}"
