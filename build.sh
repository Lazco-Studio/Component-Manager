GITHUB_TOKEN=$(cat env/GITHUB_TOKEN)

go install

go build -ldflags "-s -w -X main.GITHUB_TOKEN=$GITHUB_TOKEN" -o ./dist/cm-cli_linux_amd64_uncompressed

upx -f --best --lzma ./dist/cm-cli_linux_amd64_uncompressed -o ./dist/cm-cli_linux_amd64
