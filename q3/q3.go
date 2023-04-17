package main

import "fmt"

func main(){
    
    maior, menor, media, err := FindMinMaxAverage([]int{1})
    fmt.Printf("%d %d %f %s", maior, menor, media, err)
}

func FindMinMaxAverage(numbers []int) (int, int, float64, error) {
    
    maior := 0
    menor := 999999
    media := 0.
    x := 0
    total := 0
    
    if len(numbers) == 0 {
        
        return 0, 0, 0, fmt.Errorf("Lista vazia!")
        
    }
    
    for x = 0; x < len(numbers); x++{
        
        total += numbers[x]
        
        if numbers[x] > maior{
            
            maior = numbers[x]
            
        } 
        if numbers[x] < menor{
            
            menor = numbers[x]
            
        }
        
    }

    total = total - (maior+menor)
    
    media = float64(total) / float64(x)
    
    return maior, menor, media, nil
}
