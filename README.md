# Calcolatrice-go

App console semplice in Go che implementa una calcolatrice interattiva.

## Funzionalità

Il programma legge i comandi da tastiera e supporta:

- `+` somma tra due numeri
- `-` sottrazione tra due numeri
- `*` moltiplicazione tra due numeri
- `/` divisione tra due numeri (gestisce divisione per zero e mostra messaggio d'errore)
- `pow` potenza (base^esponente tramite `math.Pow`)
- `sqrt` radice quadrata (numeri non negativi con controllo di errore)
- `exit` per uscire dal programma

L'output è colorato in ANSI (se il terminale lo supporta) tramite costanti `ColoreBlu`, `ColoreVerde`, `ColoreRosso` e `ColoreReset`.

Per ogni operazione il programma chiede i valori necessari, esegue il calcolo e mostra il risultato con due decimali. In caso di input non valido (es. lettera) il programma richiede l'inserimento corretto.

## Comandi principali

- `go run main.go` : esegue l'app direttamente
- `go build -o calcolatrice` : compila il binario
- `./calcolatrice` (Linux/macOS) / `calcolatrice.exe` (Windows) : esegui il binario dopo la compilazione

## Istruzioni per l'uso

1. Apri il prompt o terminale nella cartella del progetto
2. Avvia con `go run main.go` (o il binario compilato)
3. Scegli l'operazione sul menu
4. Inserisci il primo e il secondo numero (per `sqrt` serve solo un numero)
5. Vedi il risultato, poi ripeti oppure digita `exit` per chiudere

## Esempio di sessione

```
=== CALCOLATRICE ===
Operazioni disponibili:
 +    : somma
 -    : sottrazione
 *    : moltiplicazione
 /    : divisione
 pow  : potenza
 sqrt : radice quadrata
 exit : uscita dal programma
Scegli un'operazione: +
Inserisci il primo numero: 5
Inserisci il secondo numero: 3
5.00 + 3.00 = 8.00
------------------------------
Scegli un'operazione: sqrt
Inserisci il numero: 9
9.00 = 3.00
------------------------------
Scegli un'operazione: exit
Uscita dal programma...
```

## Note aggiuntive

- Input numerico: usa il punto come separatore decimale (`3.14`)
- Il codice è tutto in `main.go` e non richiede dipendenze esterne oltre alla libreria standard Go.
