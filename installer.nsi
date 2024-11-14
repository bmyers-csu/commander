# Name of installer
Outfile "commander.exe"

# Installation directory
InstallDir $PROGRAMFILES\Commander

# Request admin privileges
RequestExecutionLevel admin

Section "Install"
    # The output path for this file
    SetOutPath $INSTDIR

    # Copy the executable to the output path
    File "windowscommander.exe"

    # Define uninstaller name
    WriteUninstaller "$INSTDIR\uninstall.exe"

    # Add registry entry to run at boot
    WriteRegStr HKCU "Software\Microsoft\Windows\CurrentVersion\Run" "Commander" "$INSTDIR\windowscommander.exe"
SectionEnd

Section "Uninstall"
    # Remove files and shortcuts
    Delete "$INSTDIR\windowscommander.exe"

    # Remove uninstaller
    Delete "$INSTDIR\uninstall.exe"

    # Remove the installation directory
    RMDir "$INSTDIR"

SectionEnd
