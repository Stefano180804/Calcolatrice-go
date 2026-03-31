package main

import (
    "fmt"
    "math"
)

func somma(a, b float64) float64 { 
	return a + b 
}

func sottrai(a, b float64) float64 { 
	return a - b 
}

func moltiplica(a, b float64) float64 { 
	return a * b 
}

func dividi(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("divisione per zero")
    }
    return a / b, nil
}

func potenza(base, esponente float64) float64 { 
	return math.Pow(base, esponente) 
}

func radice(valore float64) (float64, error) {
    if valore < 0 {
        return 0, fmt.Errorf("radice quadrata di numero negativo")
    }
    return math.Sqrt(valore), nil
}

func leggiNumero(prompt string) float64 {
    var n float64
    for {
        fmt.Print(prompt)
        _, err := fmt.Scanln(&n)
        if err == nil {
            return n
        }
        fmt.Println("Input non valido. Riprova.")
    }
}

func menu() string {
    var scelta string
    fmt.Println("\n=== CALCOLATRICE ===")
    fmt.Println("Operazioni disponibili:")
    fmt.Println(" +    : somma")
    fmt.Println(" -    : sottrazione")
    fmt.Println(" *    : moltiplicazione")
    fmt.Println(" /    : divisione")
    fmt.Println(" pow  : potenza")
    fmt.Println(" sqrt : radice quadrata")
    fmt.Println(" exit : uscita dal programma")
    fmt.Print("Scegli un'operazione: ")
    fmt.Scanln(&scelta)
    return scelta
}

func main() {
    for {
        operazione := menu()

        if operazione == "exit" {
            fmt.Println("Uscita dal programma...")
            break
        }

        var a, b float64
        switch operazione {
        case "+", "-", "*", "/":
            a = leggiNumero("Inserisci il primo numero: ")
            b = leggiNumero("Inserisci il secondo numero: ")
        case "pow":
            a = leggiNumero("Inserisci la base: ")
            b = leggiNumero("Inserisci l'esponente: ")
        case "sqrt":
            a = leggiNumero("Inserisci il numero: ")
        default:
            fmt.Println("Operazione non valida. Riprova.")
            continue
        }

        var risultato float64
        var err error

        switch operazione {
        case "+":
            risultato = somma(a, b)
            fmt.Printf("\n%.2f + %.2f = %.2f\n", a, b, risultato)
        case "-":
            risultato = sottrai(a, b)
            fmt.Printf("\n%.2f - %.2f = %.2f\n", a, b, risultato)
        case "*":
            risultato = moltiplica(a, b)
            fmt.Printf("\n%.2f * %.2f = %.2f\n", a, b, risultato)
        case "/":
            risultato, err = dividi(a, b)
            if err == nil {
                fmt.Printf("\n%.2f / %.2f = %.2f\n", a, b, risultato)
            }
        case "pow":
            risultato = potenza(a, b)
            fmt.Printf("\n%.2f ^ %.2f = %.2f\n", a, b, risultato)
        case "sqrt":
            risultato, err = radice(a)
            if err == nil {
                fmt.Printf("\n√%.2f = %.2f\n", a, risultato)
            }
        }

        if err != nil {
            fmt.Println("Errore:", err)
        }

        fmt.Println("------------------------------")
    }
}