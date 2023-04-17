package main

import "fmt"

func main() {

    var temperatura, finalP float64
    var escalai, escalaf string

    fmt.Print("Digite a temperatura inicial: ")
    fmt.Scan(&temperatura)
    fmt.Print("Digite a escala inicial: ")
    fmt.Scan(&escalai)
    fmt.Print("Digite a escala final: ")
    fmt.Scan(&escalaf)

    finalP, _ = ConvertTemperature(temperatura, escalai, escalaf)

    fmt.Print("Temperatura final: ", finalP)

}

func ConvertTemperature(temp float64, fromScale string, toScale string) (float64, error) {

    var tempFinal float64

    if fromScale == "C" && toScale == "F" {

        tempFinal = temp*9/5 + 32

    } else if fromScale == "C" && toScale == "K" {

        tempFinal = temp + 273.15

    } else if fromScale == "F" && toScale == "C" {

        tempFinal = (temp - 32) * 5/9

    } else if fromScale == "F" && toScale == "K" {

        tempFinal = (temp - 32) * 5/9 + 273.15

    } else if fromScale == "K" && toScale == "C" {

        tempFinal = temp - 273.15

    } else if fromScale == "K" && toScale == "F" {

        tempFinal = (temp - 273.15) * 9/5 + 32

    } else if fromScale != "K" || toScale != "F" && toScale != "K" && toScale != "C" && fromScale != "C" && fromScale != "F" {

        return 0, fmt.Errorf("Informacoes forneciadas erradas")

    }

    return tempFinal, nil
}
