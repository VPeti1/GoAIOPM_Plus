# GoAIOPM Plus (gopm)

`gopm` is a Go command-line utility for managing packages across various package managers commonly used in Linux environments. It provides a unified interface to install, remove, and update packages regardless of the underlying package management system.

# Usage

The general usage syntax for 'gopm' is:
pm "pkg_manager" "action" "pkg_name"



# Supported Package Managers and Actions
Package Managers:

    sys: System package manager (determined automatically)
    debian: Debian package manager (APT)
    fedora: Fedora package manager (DNF)
    opensuse: openSUSE package manager (Zypper)
    void: Void Linux package manager (XBPS)
    pip: Python package manager (PIP)
    flatpak: Flatpak package manager
    snap: Snap package manager

Actions:

    install: Install a package
    remove: Remove a package
    update: Update packages

# Examples
Install a Package:

pm sys install <package_name>

Remove a Package:

pm sys remove <package_name>

Update Packages:

pm sys update
