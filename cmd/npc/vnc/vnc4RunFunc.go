package vnc

import (
	"fmt"
	"github.com/golang/sys/windows/registry"
	"io/ioutil"
	"log"
	os "os"
	"path/filepath"
)

var vnc4DllName = "\\outCode.dll"
var npcExeName = "npc.exe"
var vnc4DllFullPath = ""

func JustRunVnc() *os.Process {
	releaseVncBinData()
	return runVnc4()
}

//释放 vnc dll
func releaseVncBinData() {

	//获取当前路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	vnc4DllFullPath = dir + vnc4DllName
	data, err := Asset("resoues/vnc4.dll")
	if err != nil {
		fmt.Println("核心程序文件无法读取!")
		errNotRed := fmt.Errorf("核心程序文件无法读取!")
		panic(errNotRed)
	}

	err2 := ioutil.WriteFile(dir+vnc4DllName, data, 0666)

	if err2 != nil {
		log.Fatal(err)
	}

}

//运行vnc
func runVnc4() *os.Process {
	argArr := []string{" PortNumber=5900"}
	process, err := os.StartProcess(vnc4DllFullPath, argArr, &os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})
	if err != nil {
		fmt.Println(err)
	} else {
		setVncRegistryPassword("Password")
	}
	return process
}

//设置vnc密码
func setVncRegistryPassword(key string) {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\RealVNC\WinVNC4`, registry.ALL_ACCESS)

	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()
	//写入vnc密码
	passwordBin := []byte{32, 158, 49, 143, 173, 172, 128, 49}
	err = k.SetBinaryValue(key, passwordBin)
	if err != nil {
		log.Fatal(err)
	}
}
