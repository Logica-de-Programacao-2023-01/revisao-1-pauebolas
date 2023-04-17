package main

import "fmt"

func main(){
    
    var preco, precoFinal float64
    var estado string
    var imposto int
    
    fmt.Println("Digite o valor do produto: ")
    fmt.Scan(&preco)
    fmt.Println("Digite o estado destido (CAPSLOCK): ")
    fmt.Scan(&estado)
    fmt.Print("Digite o codigo do imposto (1 a 3): ")
    fmt.Scanln(&imposto)
    
    precoFinal, _ = CalculateFinalPrice(preco, estado, imposto)
    
    fmt.Print("O preco final é: ", precoFinal)
    
}

func CalculateFinalPrice(basePrice float64, state string, taxCode int) (float64, error) {
	
	var comFrete, comImposto, finalPrice float64
	
	if basePrice <= 0 {
        return 0, fmt.Errorf("Preco deve ser maior do que zero!")
    }
    if taxCode < 1 || taxCode > 3 {
        return 0, fmt.Errorf("Imposto deve ser 1, 2 ou 3!")
    }
	
	
	if state == "SP" {

        comFrete = basePrice * 1.10
	    
	} else if state == "RJ" {
	    
        comFrete = basePrice * 1.15
	    
	} else if state == "MG" {
	    
	    comFrete = basePrice * 1.20
	    
	} else if state == "ES" {
	    
	    comFrete = basePrice * 1.25
	    
	} else {
	    
	    comFrete = basePrice * 1.30
	    
	}
	
	if taxCode == 1 {
	    
	    comImposto = basePrice * 1.05
	    
	} else if taxCode == 2 {
	    
	    comImposto = basePrice * 1.10
	    
	} else if taxCode == 3 {
	    
	    comImposto = basePrice * 1.15
	    
	}
	
	
	finalPrice = (comFrete - basePrice) + comImposto
	
	
	return finalPrice, nil
}




