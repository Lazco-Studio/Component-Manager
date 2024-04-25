# Installation
Copy and paste the following command to your terminal.
```bash
bash <(wget -qO- https://short.on-cloud.tw/cm-install-script)
```

# Advance Installation / Contributing
## Requirements
- [Go (1.22)](https://go.dev/doc/install)
- [upx](https://github.com/upx/upx/releases/latest)

## Installation Steps
Download the source files.
```bash
git clone https://github.com/lazco-studio/Component-Manager.git
```

Set the necessary environment variables.
```bash
# Create env folder
mkdir env

# Set GITHUB_TOKEN, replace "github_token" with your github token
echo "github_token" > env/GITHUB_TOKEN
```

Run the build script.
```bash
./build.sh
```

Then copy the executable file to `/usr/local/bin`.
```bash
sudo cp ./dist/cm-cli_linux_amd64 /usr/local/bin/cm
```

Grant execute permission for `/usr/local/bin/cm`.
```bash
sudo chmod +x /usr/local/bin/cm
```
