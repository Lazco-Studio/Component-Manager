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

# Variables
app_name=cm
download_link=https://github.com/lazco-studio/Component-Manager/releases/latest/download/cm-cli_linux_amd64

function install() {
  sudo wget $download_link -q --show-progress --progress=bar:force -O /usr/local/bin/$app_name
  sudo chmod +x /usr/local/bin/$app_name
}

if [ -f /usr/local/bin/$app_name ]; then
  echo -e -n "${YELLOW}Warning: App named $app_name already exists in /usr/local/bin. Do you want to overwrite it? (y/N): ${NC}"
  read -n 1 -r
  echo

  sudo -v

  if [[ $REPLY =~ ^[Yy]$ ]]; then
    echo -e "${GREEN}Updating $app_name...${NC}"
    install
  else
    echo -e "${RED}Installation aborted.${NC}"
    exit 1
  fi
else
  echo -e "${GREEN}Installing $app_name...${NC}"
  install
fi
