package main

import (
	"fmt"

	"github.com/TwiN/go-color"
	"github.com/martinlindhe/notify"
)

func main() {
	fmt.Print("Mensagem de notificação: ")
	var message string
	fmt.Scan(&message)
	notify.Alert("_", "ALERT!", "Lembrete: "+message, "path/to/icon.png")
	fmt.Print(color.Ize(color.Cyan, "Test:"))

}
