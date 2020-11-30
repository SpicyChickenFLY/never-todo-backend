package shell

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/romberli/log"
)

var out bytes.Buffer
var stderr bytes.Buffer

// ExecCommand is a func for go to call shell command
// TODO: let user decide if he/she need sudoer premission
func ExecCommand(cmdStr string) error {
	// sudo mv srcFile dstFile
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		log.Warnf("cmd-[%s]%s:%s\n",
			cmdStr, err, stderr.String())
		fmt.Printf("cmd-[%s]%s:%s\n",
			cmdStr, err, stderr.String())
		return err
	}
	return nil
}

// Mv is a shell command for moving file
func Mv(srcFile, dstFile string) error {
	// sudo mv srcFile dstFile
	return ExecCommand(
		fmt.Sprintf("sudo mv %s %s", srcFile, dstFile))
}

// Cp is a shell command for copying file
func Cp(srcFile, dstFile string) error {
	// sudo cp -r srcFile dstFile
	return ExecCommand(
		fmt.Sprintf("sudo cp -r %s %s", srcFile, dstFile))
}

// Chown is a shell command for changing owner of dir/file
func Chown(dirPath, userName, groupName string) error {
	// sudo chown -R userName.groupName dirPath
	return ExecCommand(
		fmt.Sprintf("sudo chown -R %s.%s %s", userName, groupName, dirPath))

}

// Chmod is a shell command for changing
// access premission for file/dir
func Chmod(dirPath string, fileMode uint32) error {
	// sudo chmod -R 755 dirPath
	return ExecCommand(
		fmt.Sprintf("sudo chmod -R %d %s", fileMode, dirPath))

}

// Mkdir is a shell command for making dir
func Mkdir(dirPath string) error {
	// sudo mkdir -p dirPath
	return ExecCommand(
		fmt.Sprintf("sudo mkdir -p %s", dirPath))
}

// Tar is a shell command for tar/untar file
func Tar(srcFile string, dstPath string) error {
	// tar -zxvf srcFile -C dstPath
	return ExecCommand(
		fmt.Sprintf("sudo tar -zxvf %s -C %s ", srcFile, dstPath))
}

// Useradd is a shell command for adding user in linux
func Useradd(userName string) error {
	return ExecCommand(
		fmt.Sprintf("sudo useradd -M %s", userName))
}

// Groupadd is a shell command for adding group in linux
func Groupadd(groupName string) error {
	return ExecCommand(
		fmt.Sprintf("sudo groupadd %s", groupName))

}

// UseraddWithGroup is a function to create user in specified group
func UseraddWithGroup(groupName, userName string) error {
	return ExecCommand(
		fmt.Sprintf("sudo useradd -M -g %s %s", groupName, userName))
}
