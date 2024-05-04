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

function linux_amd64() {
  echo -e "${MAGENTA}${ROCKET} Building ${BLUE}linux/amd64${MAGENTA}...${NC}"
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -X main.GITHUB_TOKEN=$GITHUB_TOKEN" -o ./dist/uncompressed/cm-cli_linux_amd64
  echo -e "${GREEN}${CHECKMARK} File saved as ${BLUE}./dist/uncompressed/cm-cli_linux_amd64${NC}"

  echo -e "${MAGENTA}${ROCKET} Compressing ${BLUE}linux/amd64${MAGENTA}...${NC}"
  upx -f --best --lzma ./dist/uncompressed/cm-cli_linux_amd64 -o ./dist/compressed/cm-cli_linux_amd64
  echo -e "${GREEN}${CHECKMARK} File saved as ${BLUE}./dist/compressed/cm-cli_linux_amd64${NC}"
}

function darwin_amd64() {
  echo -e "${MAGENTA}${ROCKET} Building ${BLUE}darwin/amd64${MAGENTA}...${NC}"
  CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w -X main.GITHUB_TOKEN=$GITHUB_TOKEN" -o ./dist/uncompressed/cm-cli_darwin_amd64
  echo -e "${GREEN}${CHECKMARK} File saved as ${BLUE}./dist/uncompressed/cm-cli_darwin_amd64${NC}"

  echo -e "${MAGENTA}${ROCKET} Compressing ${BLUE}darwin/amd64${MAGENTA}...${NC}"
  upx -f --best --lzma ./dist/uncompressed/cm-cli_darwin_amd64 -o ./dist/compressed/cm-cli_darwin_amd64 --force-macos
  echo -e "${GREEN}${CHECKMARK} File saved as ${BLUE}./dist/compressed/cm-cli_darwin_amd64${NC}"
}

function windows_amd64() {
  echo -e "${MAGENTA}${ROCKET} Building ${BLUE}windows/amd64${MAGENTA}...${NC}"
  CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -X main.GITHUB_TOKEN=$GITHUB_TOKEN" -o ./dist/uncompressed/cm-cli_windows_amd64.exe
  echo -e "${GREEN}${CHECKMARK} File saved as ${BLUE}./dist/uncompressed/cm-cli_windows_amd64.exe${NC}"

  echo -e "${MAGENTA}${ROCKET} Compressing ${BLUE}windows/amd64${MAGENTA}...${NC}"
  upx -f --best --lzma ./dist/uncompressed/cm-cli_windows_amd64.exe -o ./dist/compressed/cm-cli_windows_amd64.exe
  echo -e "${GREEN}${CHECKMARK} File saved as ${BLUE}./dist/compressed/cm-cli_windows_amd64.exe${NC}"
}

if [ "$1" == "linux_amd64" ]; then
  linux_amd64
elif [ "$1" == "darwin_amd64" ]; then
  darwin_amd64
elif [ "$1" == "windows_amd64" ]; then
  windows_amd64
elif [ "$1" != "" ]; then
  echo -e "${RED}${CROSSMARK} ${BLUE}$1${YELLOW} not supported.${NC}"
  echo -e "${RED}${CROSSMARK} Nothing to build.${NC}"
  exit 1
else
  linux_amd64
  darwin_amd64
  windows_amd64
fi
