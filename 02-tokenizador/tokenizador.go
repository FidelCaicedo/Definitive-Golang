package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

// CountTokens recibe textos crudos y devuelve cuántas veces aparece cada token.
// No imprime, no lee argumentos, no toca el mundo exterior: solo transforma
// entrada en salida. Por eso es testeable.
func CountTokens(inputs []string) map[string]int {
	counts := make(map[string]int)

	for _, input := range inputs {
		tokens := strings.FieldsFunc(strings.ToLower(input), func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		})

		for _, token := range tokens {
			counts[token]++
		}
	}

	return counts
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, `uso: tokenizador "texto a tokenizar" [...]`)
		os.Exit(1)
	}

	counts := CountTokens(args)

	// El recorrido de un map no tiene orden garantizado, así que ordenamos
	// las claves para que la salida sea estable entre ejecuciones.
	keys := make([]string, 0, len(counts))
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%-15s %d\n", k, counts[k])
	}
}
