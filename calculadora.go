package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func suma() int {
	var suma = 0
	var numero1, numero2 = leerValores()

	var primerValor, err1 = strconv.Atoi(numero1)
		if err1 != nil {
			fmt.Println("No puedes introducir letras solo numeros")
		}
	
	var segundoValor, err2 = strconv.Atoi(numero2)
		if err2 != nil {
			fmt.Println("No puedes introducir letras solo numeros")
		}

	suma = primerValor + segundoValor
	return suma
}

func resta() int {
	var resta = 0
	var numero1, numero2 = leerValores()
	var primerValor, err1 = strconv.Atoi(numero1)
		if err1 != nil {
			fmt.Println("No puedes introducir letras solo numeros")
		}
	var segundoValor, err2 = strconv.Atoi(numero2)
		if err2 != nil {
			fmt.Println("No puedes introducir letras solo numeros")
		}
	resta = primerValor - segundoValor
	return resta
}

func multiplicacion() int {
	var multiplicacion = 0
	var numero1, numero2 = leerValores()
	var primerValor, err1 = strconv.Atoi(numero1)
		if err1 != nil {
			fmt.Println("No puedes introducir letras solo numeros")
		}
	var segundoValor, err2 = strconv.Atoi(numero2)
		if err2 != nil {
			fmt.Println("No puedes introducir letras solo numeros")
		}
	multiplicacion = primerValor * segundoValor
	return multiplicacion
}

func division() int {
	var division = 0
	var numero1, numero2 = leerValores()
	var primerValor, err1 = strconv.Atoi(numero1)
		if err1 != nil {
			fmt.Println("No puedes introducir letras solo numeros")
		}
	var segundoValor, err2 = strconv.Atoi(numero2)
		if err2 != nil {
			fmt.Println("No puedes introducir letras solo numeros")
		}
	
		if segundoValor == 0 {
			fmt.Println("No puedes dividir por 0")
		} else {
			division = primerValor - segundoValor
		}

	return division
}

func leerValores() (string, string) {
	fmt.Println("Introduzca su primer valor")
	var primerScanner = bufio.NewScanner(os.Stdin)
		primerScanner.Scan()
	var primerValor = primerScanner.Text()

	fmt.Println("Introduzca su segundo valor")
	var segundoScanner = bufio.NewScanner(os.Stdin)
		segundoScanner.Scan()
	var segundoValor = segundoScanner.Text()
	return primerValor, segundoValor

}

func menu() {
	fmt.Println("Ingrese la operacion a realizar: ")
	fmt.Println("Ingrese (1) Para realizar una suma(+)")
	fmt.Println("Ingrese (2) Para realizar una resta(-)")
	fmt.Println("Ingrese (3) Para realizar una multiplicacion(*)")
	fmt.Println("Ingrese (4) Para realizar una División (/)")
}

func main() {
	menu()
	var scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
	
	var operacion = scanner.Text()
	switch operacion {
	case "1":
		fmt.Println("El Resultado de la suma es:", suma())
	case "2":
		fmt.Println("El Resultado de la resta es:", resta())
	case "3": 
		fmt.Println("El Resultado de la suma es:", multiplicacion())
	case "4":
		fmt.Println("El Resultado de la division es:", division())
	default:
		fmt.Println("No es una operacion correcta")
	}
	fmt.Println("Pragrama terminado")
}