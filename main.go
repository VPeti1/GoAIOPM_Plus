package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func handleSys(action, pkg string) {
	switch action {
	case "install":
		fmt.Println("Installing pacman pkg:", pkg)
		runCommand("sudo", "pacman", "-S", pkg)
	case "remove":
		fmt.Println("Removing pacman pkg:", pkg)
		runCommand("sudo", "pacman", "-R", pkg)
	case "update":
		fmt.Println("Updating pacman pkgs")
		runCommand("sudo", "pacman", "-Syu")
	default:
		fmt.Println("Invalid argument!")
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func handlePkgManager(pkgManager, action, pkg string) {
	switch pkgManager {
	case "sys":
		handleSys(action, pkg)
	case "debian":
		handleDebian(action, pkg)
	case "fedora":
		handleFedora(action, pkg)
	case "opensuse":
		handleOpenSUSE(action, pkg)
	case "void":
		handleVoid(action, pkg)
	default:
		fmt.Println("Invalid pkg manager")
	}
}

func handleDebian(action, pkg string) {
	switch action {
	case "install":
		fmt.Println("Installing APT pkg:", pkg)
		runCommand("sudo", "apt", "install", pkg)
	case "remove":
		fmt.Println("Removing APT pkg:", pkg)
		runCommand("sudo", "apt", "remove", pkg)
	case "update":
		fmt.Println("Updating APT pkgs")
		runCommand("sudo", "apt", "update")
		runCommand("sudo", "apt", "upgrade")
	default:
		fmt.Println("Invalid argument!")
	}
}

func handleFedora(action, pkg string) {
	switch action {
	case "install":
		fmt.Println("Installing dnf pkg:", pkg)
		runCommand("sudo", "dnf", "install", pkg)
	case "remove":
		fmt.Println("Removing dnf pkg:", pkg)
		runCommand("sudo", "dnf", "remove", pkg)
	case "update":
		fmt.Println("Updating dnf pkgs")
		runCommand("sudo", "dnf", "update")
	default:
		fmt.Println("Invalid argument!")
	}
}

func handleOpenSUSE(action, pkg string) {
	switch action {
	case "install":
		fmt.Println("Installing zypper pkg:", pkg)
		runCommand("sudo", "zypper", "install", pkg)
	case "remove":
		fmt.Println("Removing zypper pkg:", pkg)
		runCommand("sudo", "zypper", "remove", pkg)
	case "update":
		fmt.Println("Updating zypper pkgs")
		runCommand("sudo", "zypper", "update")
	default:
		fmt.Println("Invalid argument!")
	}
}

func handleVoid(action, pkg string) {
	switch action {
	case "install":
		fmt.Println("Installing XBPS pkg:", pkg)
		runCommand("sudo", "xbps-install", pkg)
	case "remove":
		fmt.Println("Removing XBPS pkg:", pkg)
		runCommand("sudo", "xbps-remove", pkg)
	case "update":
		fmt.Println("Updating XBPS pkgs")
		runCommand("sudo", "xbps-install", "-Su")
	default:
		fmt.Println("Invalid argument!")
	}
}

func runCommand(command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin // Attach stdin to the command
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func getInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func fst() {
	clear()
	fmt.Println("Welcome to AIOPM Setup!")
	fmt.Println("What distro are you using?")
	fmt.Println("(arch,debian,fedora,opensuse or void)")
	fmt.Println("(Derivatives included)")
	fmt.Print("> ")
	input := getInput("")
	switch input {
	case "arch", "Arch":
		fmt.Println("Setting configuration for Arch")
		exec.Command("sudo", "mkdir", "/etc/aiopm").Run()
		exec.Command("sudo", "touch", "/etc/aiopm/a1.cw").Run()
	case "debian", "Debian":
		fmt.Println("Setting configuration for Debian")
		exec.Command("sudo", "mkdir", "/etc/aiopm").Run()
		exec.Command("sudo", "touch", "/etc/aiopm/a2.cw").Run()
	case "fedora", "Fedora":
		fmt.Println("Setting configuration for Fedora")
		exec.Command("sudo", "mkdir", "/etc/aiopm").Run()
		exec.Command("sudo", "touch", "/etc/aiopm/a3.cw").Run()
	case "opensuse", "Opensuse":
		exec.Command("sudo", "mkdir", "/etc/aiopm").Run()
		exec.Command("sudo", "touch", "/etc/aiopm/a4.cw").Run()
	case "void", "Void":
		exec.Command("sudo", "mkdir", "/etc/aiopm").Run()
		exec.Command("sudo", "touch", "/etc/aiopm/a5.cw").Run()
	default:
		fmt.Println("Invalid input! Retrying")
		systemPause()
		fst()
	}
}

func systemPause() {
	fmt.Println("Press Enter to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func main() {
	// Splash screen
	fmt.Println(` $$$$$$\  $$$$$$\  $$$$$$\  $$$$$$$\  $$\      $$\ 
$$  __$$\ \_$$  _|$$  __$$\ $$  __$$\ $$$\    $$$ |
$$ /  $$ |  $$ |  $$ /  $$ |$$ |  $$ |$$$$\  $$$$ |
$$$$$$$$ |  $$ |  $$ |  $$ |$$$$$$$  |$$\$$\$$ $$ |
$$  __$$ |  $$ |  $$ |  $$ |$$  ____/ $$ \$$$  $$ |
$$ |  $$ |  $$ |  $$ |  $$ |$$ |      $$ |\$  /$$ |
$$ |  $$ |$$$$$$\  $$$$$$  |$$ |      $$ | \_/ $$ |
\__|  \__|\______| \______/ \__|      \__|     \__|
By VPeti`)
	time.Sleep(1 * time.Second)

	// Validate arguments
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Too few arguments! Usage: pm <pkg_manager> <action> <pkg_name>")
		os.Exit(1)
	}

	pkgManager := os.Args[1]

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Insufficient arguments! Usage: pm <pkg_manager> <action> <pkg_name>")
		os.Exit(1)
	}

	action := os.Args[2]

	// Handle different pkg managers
	switch pkgManager {
	case "pip":
		// Handle pip commands
		if action == "install" || action == "remove" || action == "update" {
			if len(os.Args) < 4 {
				fmt.Fprintln(os.Stderr, "Insufficient arguments for pip action!")
				os.Exit(1)
			}
			pkgArg := os.Args[3]
			fmt.Println("Performing action on pip pkg:", pkgArg)
			cmd := exec.Command("pip", strings.Join([]string{action, pkgArg}, " "))
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error executing pip command:", err)
				os.Exit(1)
			}
		} else {
			fmt.Fprintln(os.Stderr, "Invalid argument for pip!")
			os.Exit(1)
		}
	case "flatpak":
		// Handle flatpak commands
		if action == "install" || action == "remove" || action == "update" {
			if len(os.Args) < 4 {
				fmt.Fprintln(os.Stderr, "Insufficient arguments for flatpak action!")
				os.Exit(1)
			}
			pkgArg := os.Args[3]
			fmt.Println("Performing action on flatpak pkg:", pkgArg)
			cmd := exec.Command("flatpak", strings.Join([]string{action, pkgArg}, " "))
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error executing flatpak command:", err)
				os.Exit(1)
			}
		} else {
			fmt.Fprintln(os.Stderr, "Invalid argument for flatpak!")
			os.Exit(1)
		}
	case "snap":
		// Handle snap commands
		if action == "install" || action == "remove" {
			if len(os.Args) < 4 {
				fmt.Fprintln(os.Stderr, "Insufficient arguments for snap action!")
				os.Exit(1)
			}
			pkgArg := os.Args[3]
			fmt.Println("Performing action on snap pkg:", pkgArg)
			cmd := exec.Command("snap", strings.Join([]string{action, pkgArg}, " "))
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error executing snap command:", err)
				os.Exit(1)
			}
		} else {
			fmt.Fprintln(os.Stderr, "Invalid argument for snap!")
			os.Exit(1)
		}
	case "sys":
		if len(os.Args) == 3 && action == "update" {
			// Handle system update without package name
			handleSys(action, "")
		} else if len(os.Args) >= 4 {
			pkg := os.Args[3]
			if fileExists("/etc/aiopm/a1.cw") {
				handlePkgManager("sys", action, pkg)
			} else if fileExists("/etc/aiopm/a2.cw") {
				handlePkgManager("debian", action, pkg)
			} else if fileExists("/etc/aiopm/a3.cw") {
				handlePkgManager("fedora", action, pkg)
			} else if fileExists("/etc/aiopm/a4.cw") {
				handlePkgManager("opensuse", action, pkg)
			} else if fileExists("/etc/aiopm/a5.cw") {
				handlePkgManager("void", action, pkg)
			} else {
				fst()
			}
		} else {
			fmt.Fprintln(os.Stderr, "Invalid action or insufficient arguments for sys!")
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid pkg manager.")
		os.Exit(1)
	}
}
