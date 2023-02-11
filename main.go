package main

import (
	"fmt"
	"time"

	"github.com/martinlindhe/notify"
)

func main() {
	fmt.Print("Mensagem de notificação: ")
	var message string
	fmt.Scan(&message)

	fmt.Print("Daqui a quantos segundos: ")
	var setTimeDuration int
	fmt.Scan(&setTimeDuration)

	time.Sleep(time.Duration(setTimeDuration) * time.Second)
	notify.Alert("_", "ALERT!", "Lembrete: "+message, "path/to/icon.png")

}
