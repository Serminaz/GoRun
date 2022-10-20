/*
=== Взаимодействие с ОС ===
Необходимо реализовать собственный шелл
встроенные команды: cd/pwd/echo/kill/ps
*/

package main

import (
	"bufio"
	"fmt"
	gops "github.com/mitchellh/go-ps" // gops могут сообщать инфу текущая трассировка стека,
	// версия Go, статистика памяти и т. д.
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

type data struct {
	cd      string
	cdItems []string
	pids    []int
	delim   string
}

var curData data

func rightPath(dir string) string {
	var s string
	// возвращает копию строки s со всеми неперекрывающимися экземплярами old, замененными на new.
	s = strings.ReplaceAll(dir, "/", curData.delim)
	s = strings.ReplaceAll(dir, "\\", curData.delim)

	if strings.HasSuffix(s, curData.delim) { //проверяет, заканчивается ли строка s суффиксом curData.delim.
		s = s[:len(s)-1] // hello -> hell
	}

	return s
}

func showDir() {
	fmt.Println()                            //  для отображения перехода из папки в папку (cd)
	items, err := ioutil.ReadDir(curData.cd) // // для создания временных файлов TempFile создает новый временный файл в каталоге dir,
	// открывает файл для чтения и записи и возвращает полученный *os. File.
	pwd()

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("%-50v%-15v%v\n", "Имя", "Размер", "Дата изменения")

	var files []string

	for _, i := range items {
		files = append(files, i.Name())
		fmt.Printf("%-50v%-15v%v\n", i.Name(), i.Size(), i.ModTime())
	}

	curData.cdItems = files
}

func checkExists(dir string) bool {
	_, err := os.Stat(dir)
	return !os.IsNotExist(err) // IsNotExist возвращает логическое значение, указывающее, известно ли об ошибке сообщение
	// о том, что файл или каталог не существует.
	//Удовлетворяется ErrNotExist, а также некоторыми ошибками системного вызова.
}

func dirConraints(str string) bool {
	for _, f := range curData.cdItems {
		if f == str {
			return true
		}
	}

	return false
}

func cd(dir string) {
	if dir == ".." {
		curData.cd = filepath.Dir(curData.cd)
		showDir()
		return
	}

	if dirConraints(dir) {
		cd(curData.cd + "\\" + dir)
	}

	if checkExists(dir) {
		curData.cd = rightPath(dir)
	} else {
		fmt.Println("Указанная директория не существует")
	}

	showDir()
}

func pwd() {
	fmt.Println("Текущая директория - ", curData.cd)
}

func echo(str string) {
	fmt.Println(str)
}

func ps() {
	var pids []int

	ps, _ := gops.Processes()
	fmt.Printf("%-10v%-10v%v\n", "Pid", "PPid", "Executable")

	for _, p := range ps {
		pids = append(pids, p.PPid())
		fmt.Printf("%-10v%-30v%v\n", p.Pid(), p.PPid(), p.Executable())
	}

	curData.pids = pids
}

func kill(pid string) {
	procId, err := strconv.Atoi(pid)
	if err != nil {
		fmt.Println("PID неправильный ", pid)
		return
	}

	proc, err := os.FindProcess(procId)
	if err != nil {
		fmt.Println("Процесс %s не существует - \n", procId)
		return
	}

	err = proc.Kill()
	if err != nil {
		fmt.Println("Не удается убить процесс - ", err)
		return
	}
}

func fork() {
	bin, _ := os.Executable()
	args := []string{""}
	env := os.Environ()
	env = append(env, "cd="+curData.cd)
	err := syscall.Exec(bin, args, env)

	if err != nil {
		log.Fatal(err)
	}
}

func Shell() {
	curData.cd, _ = os.Getwd()

	if runtime.GOOS == "windows" {
		curData.delim = "\\"
	} else {
		curData.delim = "/"
	}

	showDir()

	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		var args string

		line := scan.Text()
		sep := strings.Index(line, " ")
		bash := strings.Split(line, " ")[0]

		if sep > -1 {
			args = line[sep+1:]
		}

		switch bash {
		case "cd":
			cd(args)
		case "echo":
			echo(args)
		case "pwd":
			pwd()
		case "ps":
			ps()
		case "kill":
			kill(args)
		case "fork":
			fork()
		}
	}
}

func main() {
	Shell()
}
