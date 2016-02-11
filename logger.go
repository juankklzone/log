package log

import (
	"fmt"
	"os"
	"time"
)

var fileName = "errors.log"

//Init cambia el nombre del archivo que registra los errores
func Init(file string) {
	fileName = file
}

//RegistrarError se encarga de escribir en el archivo y en caso que sea un error fatal mata el proceso
func RegistrarError(err Throwable) {
	file, e := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0666)
	if e != nil {
		os.Create(fileName)
		file, e = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0666)
	}
	if err.Level() == LevelFatal {
		fmt.Println(err)
		defer os.Exit(-1)
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("%s :<NIVEL %s> [CODIGO %d]: %s \n", time.Now().String(), err.Level().String(), err.Code(), err.Error()))
}
