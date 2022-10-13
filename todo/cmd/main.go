package main
import (
"github.com/Serminaz/GoRun/todo"
"log"
)

func main() {
	srv := new(todo.Server)
	if err := strv.Run("8000"); err != nil {
		log.Fatal( "error initializing configs: %s", err.Error())
	}
}