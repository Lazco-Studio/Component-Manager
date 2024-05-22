# Colors
$GREEN = [ConsoleColor]::Green
$YELLOW = [ConsoleColor]::Yellow
$RED = [ConsoleColor]::Red

# Variables
$app_name = "cm"
$download_link = "https://github.com/lazco-studio/Component-Manager/releases/latest/download/cm-cli_windows_amd64.exe"
$install_path = "C:\Program Files\$app_name\"
$app_path = "C:\Program Files\$app_name\$app_name.exe"

if (!([Security.Principal.WindowsPrincipal][Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole] "Administrator")) { Start-Process powershell.exe " -NoExit -NoProfile -ExecutionPolicy Bypass -File `"$PSCommandPath`"" -Verb RunAs; exit }

function Test-Admin {
  $currentUser = New-Object Security.Principal.WindowsPrincipal $([Security.Principal.WindowsIdentity]::GetCurrent())
  $currentUser.IsInRole([Security.Principal.WindowsBuiltinRole]::Administrator)
}

if ((Test-Admin) -eq $false)  {
  if ($elevated) {
  } else {
    Start-Process powershell.exe -Verb RunAs -ArgumentList ('-noprofile -noexit -file "{0}" -elevated' -f ($myinvocation.MyCommand.Definition))
  }
  exit
}

function Add-To-Path {
  $currentPath = [Environment]::GetEnvironmentVariable("PATH", "Machine")

  if ($currentPath -split ";" -notcontains $install_path) {
    $newPath = $currentPath + ";" + $install_path
    [Environment]::SetEnvironmentVariable("PATH", $newPath, "Machine")
    Write-Host "Directory added to PATH variable successfully." -ForegroundColor Green
  } else {
    Write-Host "Directory is already in PATH variable." -ForegroundColor Yellow
  }
}

function Install-Tool {
  if (!(Test-Path "$install_path")) {
    New-Item -Path "$install_path" -ItemType Directory
  }

  Invoke-WebRequest $download_link -OutFile $app_path

  Add-To-Path

  $env:Path = [System.Environment]::GetEnvironmentVariable("Path","Machine") + ";" + [System.Environment]::GetEnvironmentVariable("Path","User") 
  $version = & "$app_path" --version

  Write-Host "Successfully installed $version." -ForegroundColor $GREEN
}


if (Test-Path "$app_path") {
  Write-Host -NoNewline "Warning: App named $app_name already exists in $install_path. Do you want to overwrite it? (y/N): " -ForegroundColor $YELLOW
  $response = Read-Host

  if ($response -eq "Y" -or $response -eq "y") {
    Write-Host "Updating $app_name ..." -ForegroundColor $GREEN
    Install-Tool
  } else {
    Write-Host "Installation aborted." -ForegroundColor $RED
    exit 1
  }
} else {
  Write-Host "Installing $app_name ..." -ForegroundColor $GREEN
  Install-Tool
}