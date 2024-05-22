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

function not_supported() {
  echo -e "${RED}${CROSSMARK} ${BLUE}$1${YELLOW} not supported.${NC}"
  echo -e "${RED}${CROSSMARK} Installation aborted.${NC}"
  exit 1
}

case $(uname -ms) in
'Linux x86_64')
  target=linux_amd64
  echo -e "${GREEN}${CHECKMARK} Operating system ${BLUE}Linux x86_64 ${GREEN}detected.${NC}"
  ;;
'Darwin x86_64')
  target=darwin_amd64
  echo -e "${GREEN}${CHECKMARK} Operating system ${BLUE}macOS x86_64 ${GREEN}detected.${NC}"
  ;;
'Darwin arm64')
  target=darwin_amd64
  echo -e "${GREEN}${CHECKMARK} Operating system ${BLUE}macOS arm64 ${GREEN}detected.${NC}"
  ;;
'Linux aarch64' | 'Linux arm64')
  not_supported "Linux arm64"
  ;;
*)
  not_supported "$(uname -ms)"
  ;;
esac

# Variables
app_name=cm
download_link=https://github.com/lazco-studio/Component-Manager/releases/latest/download/cm-cli_$target

function install() {
  sudo -v
  echo -e -n "${GREEN}"
  sudo wget $download_link -q --show-progress --progress=bar:force -O /usr/local/bin/$app_name
  echo -e -n "${NC}"
  sudo chmod +x /usr/local/bin/$app_name
  app_version=$(/usr/local/bin/$app_name --version)
  echo -e "${GREEN}${CHECKMARK} Successfully installed ${BLUE}$app_version${GREEN}.${NC}"
}

if [ -f /usr/local/bin/$app_name ]; then
  echo -e -n "${YELLOW}${CROSSMARK} Warning: App named $app_name already exists in ${BLUE}/usr/local/bin${YELLOW}. Do you want to overwrite it? (y/N): ${NC}"
  read -n 1 -r
  echo

  if [[ $REPLY =~ ^[Yy]$ ]]; then
    echo -e "${MAGENTA}${ROCKET} Updating ${BLUE}$app_name ${MAGENTA}...${NC}"
    install
  else
    echo -e "${RED}${CROSSMARK} Installation aborted.${NC}"
    exit 1
  fi
else
  echo -e "${MAGENTA}${ROCKET} Installing ${BLUE}$app_name ${MAGENTA}...${NC}"
  install
fi
