# Calcolatrice-go

App console semplice in Go che implementa una calcolatrice interattiva.

## Cosa fa

Il programma legge i comandi da tastiera e supporta:

- `+` somma tra due numeri
- `-` sottrazione tra due numeri
- `*` moltiplicazione tra due numeri
- `/` divisione tra due numeri (gestisce divide-by-zero)
- `pow` potenza (base^esponente)
- `sqrt` radice quadrata (numeri non negativi)
- `exit` per uscire

Per ogni operazione il programma chiede i valori necessari, esegue il calcolo e mostra il risultato formattato con due decimali. In caso di errore (divisione per zero, radice di negativo, input non valido) mostra un messaggio di errore.

## Come avviare

Assicurati di avere Go installato (versione 1.20+ consigliata).

1. Clona il repository:

```bash
git clone https://github.com/<tuo-username>/Calcolatrice-go.git
cd Calcolatrice-go
```

2. Compila e avvia:

```bash
go run main.go
```

In alternativa puoi compilare prima e poi eseguire il binario:

```bash
go build -o calcolatrice
./calcolatrice    # su Windows: calcolatrice.exe
```

## Uso

Dopo l'avvio vedrai il menu. Inserisci l'operazione e fornisci i numeri quando richiesto.

Esempio:

- `+` -> primo numero, secondo numero -> output `a + b = risultato`
- `sqrt` -> numero -> output `√x = risultato`

3. Per uscire, digita `exit`.

## Note

- Gestione input tramite `fmt.Scanln`, quindi usa il punto come separatore decimale (`3.14`).
- Il codice è tutto in `main.go`.

