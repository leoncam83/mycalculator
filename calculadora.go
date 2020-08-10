package mycalculator

// go mod init github.com/leoncam83/mycalculator crea un archivo modulo.

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calc struct{}

func (Calc) Operate(entrada string, operador string) int {

	entradaLimpia := strings.Split(entrada, operador)

	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])

	var valor int = 0

	switch operador {
	case "+":
		valor = operador1 + operador2
	case "-":
		valor = operador1 - operador2
	case "*":
		valor = operador1 * operador2
	case "/":
		valor = operador1 / operador2
	default:
		valor = 0
	}

	return valor
}

func parsear(entrada string) int {
	opetador, _ := strconv.Atoi(entrada)
	return opetador
}

func LeerEntrada(texto string) string {
	fmt.Println(texto)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

/* func main() {

	c := calc{}

	operador := leerEntrada("Operador +,-,*,/")
	entrada := leerEntrada("Valores: ")

	valor := c.operate(entrada, operador)

	fmt.Println(valor)

} */
