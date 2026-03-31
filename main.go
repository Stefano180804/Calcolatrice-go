package main

import (
    "fmt"
    "math"
)

const (
    ColoreReset  = "\033[0m"
    ColoreRosso  = "\033[31m"
    ColoreVerde  = "\033[32m"
    ColoreBlu    = "\033[34m"
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
        fmt.Print(ColoreBlu + prompt + ColoreReset)
        _, err := fmt.Scanln(&n)
        if err == nil {
            return n
        }
        fmt.Println(ColoreRosso + "Input non valido. Riprova." + ColoreReset)
    }
}

func menu() string {
    var scelta string
    fmt.Println(ColoreBlu + "\n=== CALCOLATRICE ===" + ColoreReset)
    fmt.Println(ColoreBlu + "Operazioni disponibili:" + ColoreReset)
    fmt.Println(ColoreBlu + " +    : somma" + ColoreReset)
    fmt.Println(ColoreBlu + " -    : sottrazione" + ColoreReset)
    fmt.Println(ColoreBlu + " *    : moltiplicazione" + ColoreReset)
    fmt.Println(ColoreBlu + " /    : divisione" + ColoreReset)
    fmt.Println(ColoreBlu + " pow  : potenza" + ColoreReset)
    fmt.Println(ColoreBlu + " sqrt : radice quadrata" + ColoreReset)
    fmt.Println(ColoreBlu + " exit : uscita dal programma" + ColoreReset)
    fmt.Print(ColoreBlu + "Scegli un'operazione: " + ColoreReset)
    fmt.Scanln(&scelta)
    return scelta
}

func main() {
    for {
        operazione := menu()

        if operazione == "exit" {
            fmt.Println(ColoreBlu + "Uscita dal programma..." + ColoreReset)
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
            fmt.Println(ColoreRosso + "Operazione non valida. Riprova." + ColoreReset)
            continue
        }

        var risultato float64
        var err error

        switch operazione {
        case "+":
            risultato = somma(a, b)
            fmt.Printf(ColoreVerde+"\n%.2f + %.2f = %.2f\n"+ColoreReset, a, b, risultato)
        case "-":
            risultato = sottrai(a, b)
            fmt.Printf(ColoreVerde+"\n%.2f - %.2f = %.2f\n"+ColoreReset, a, b, risultato)
        case "*":
            risultato = moltiplica(a, b)
            fmt.Printf(ColoreVerde+"\n%.2f * %.2f = %.2f\n"+ColoreReset, a, b, risultato)
        case "/":
            risultato, err = dividi(a, b)
            if err == nil {
                fmt.Printf(ColoreVerde+"\n%.2f / %.2f = %.2f\n"+ColoreReset, a, b, risultato)
            }
        case "pow":
            risultato = potenza(a, b)
            fmt.Printf(ColoreVerde+"\n%.2f ^ %.2f = %.2f\n"+ColoreReset, a, b, risultato)
        case "sqrt":
            risultato, err = radice(a)
            if err == nil {
                fmt.Printf(ColoreVerde+"\n√%.2f = %.2f\n"+ColoreReset, a, risultato)
            }
        }

        if err != nil {
            fmt.Println(ColoreRosso + "Errore: " + err.Error() + ColoreReset)
        }

        fmt.Println(ColoreBlu + "------------------------------" + ColoreReset)
    }
}