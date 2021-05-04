package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hello("Taro", ""))
}

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(lang) + name
}

// (prefix string)はnamed return value と呼ばれ、zero値が割り当てられる。return を呼び出すだけで済むようになる
func greetingPrefix(lang string) (prefix string) {

	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
