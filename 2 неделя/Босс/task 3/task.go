//package main
//
//import (
//	"bufio"
//	"fmt"
//	"net"
//	"os"
//	"os/signal"
//	"strings"
//	"syscall"
//	"time"
//)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).
*/
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// go run telnet.go --timeout=10s localhost 8080

const commandError = `Ошибка при вызове команды!:
	example: go-telnet --timeout=10s host port`

type Telnet struct {
	Timeout time.Duration
	Host    string
	Port    string
}

// createClient создает новый telnet client
func (t *Telnet) createClient(conn net.Conn) {
	fmt.Println("Welcome to Telnet!")
	console := bufio.NewReader(os.Stdin)
	connReader := bufio.NewReader(conn) // считывает данные из соединения
	for {
		fmt.Print("Сообщение на сервер > ")
		text, err := console.ReadString('\n') //  аналог сканф , присваеваем введеные символы
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading string: %v", err)
			return
		}
		text = strings.TrimSpace(text) // если в начале и или в конце есть пробелы - удаляем

		fmt.Fprintf(conn, text+"\n")
		if text == "exit" {
			fmt.Fprintf(os.Stdout, "%s\n", "connection closed")
			return
		}

		message, err := connReader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading from conn: %v\n", err)
			os.Exit(1)
		}
		message = strings.TrimSpace(message)
		fmt.Printf("Сообщение от сервера < %s\n", message)
	}
}

// parseArgs обрабатывает полученные на вход флаги
func (t *Telnet) parseArgs() bool {
	if len(os.Args) == 3 {
		t.Host = os.Args[1]
		t.Port = os.Args[2]
		return true
	}
	if len(os.Args) == 4 {
		arg := os.Args[1] // вся командная строка
		substr := "--timeout="
		if strings.Contains(arg, substr) { //Contains("seafood", "foo")  --> true
			timeDuration := strings.TrimPrefix(arg, substr)  // --timeout= вычитает  и оставляет 10s
			timeout, err := time.ParseDuration(timeDuration) //10.5 s => 10s, 30ms
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
			}
			t.Timeout = timeout
		} else {
			return false
		}
		t.Host = os.Args[2] // localhost
		t.Port = os.Args[3]
		return true
	}
	return false
}

func main() {
	t := &Telnet{}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGINT) // SIGQUIT это чудо принимает команду "^\" , SIGINT --> ^C

	ok := t.parseArgs() // проверка на адекватность командной строки
	if !ok {
		fmt.Fprintf(os.Stderr, "%s\n", commandError)
		os.Exit(1)
	}
	d := net.Dialer{Timeout: t.Timeout}           //  время ожидания для завершения ?
	conn, err := d.Dial("tcp", t.Host+":"+t.Port) //подключается к адресу в названной сети.
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	/*Оператор defer вызывает функцию, выполнение которой отложено до момента возврата
	из окружающей функции, либо из-за того, что окружающая функция выполнила оператор return,
	достигла конца тела своей функции, либо из-за паники соответствующей горутины.*/

	go t.createClient(conn)

	select {
	case <-quit:
		fmt.Fprintf(os.Stdout, "%s\n", "programm finished")
		os.Exit(0)
	}
}
