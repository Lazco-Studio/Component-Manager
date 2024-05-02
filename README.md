<div align="center">
<h1>Component Manager</h1>
Seamlessly manage and integrate your JS/TS components with ease.<br>
<h3><code>cm add</code></h3>
<br>
</div>

---

Component Manager (cm) is a sophisticated tool developed in Golang, specifically tailored for managing reusable JavaScript (JS), TypeScript (TS), and React components sourced from remote repositories.

### Manage your components with ease
It simplifies the development process by allowing developers to effortlessly download and integrate these components into their projects.

### No need to worry about dependencies
It will automatically download and install the necessary dependencies for the selected component. It also supports multiple package manager, automatically adapting to utilize pnpm, bun, yarn, or npm as required.

# Installation
Copy and paste the following command to your terminal.
```bash
bash <(wget -qO- https://short.on-cloud.tw/cm-install-script)
```

Supported Platforms
- *Linux*: x86_64
- *MacOS*: x86_64, arm64
- *Windows*: x86_64

# Advance Installation / Contributing
## Requirements
- [Go (1.22)](https://go.dev/doc/install)
- [upx (4.2.3)](https://github.com/upx/upx/releases/latest)

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
