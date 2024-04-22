# Installation
## Build from source
### Requirements
- [Go (1.22)](https://go.dev/doc/install)
- [upx](https://github.com/upx/upx/releases/latest)

### Installation Steps
Download the source files.
```bash
git clone git@github.com:LAZCO-STUDIO-LTD/Component-Manager.git
```

Set the necessary environment variables.
```bash
mkdir env

# Set GITHUB_TOKEN, replace "github token" with your github token
echo "github_token" > env/GITHUB_TOKEN
```

Run the build script.
```bash
./build.sh
```

Then copy the executable file to `/usr/local/bin`
```bash
sudo cp ./dist/cm-cli_linux_amd64 /usr/local/bin/cm
```