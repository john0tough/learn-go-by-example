package helloworld

import "fmt"

const EnglishHelloPrefix = "Hello, "
const SpanishHelloPrefix = "Hola, "

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	if lang == "spanish" {
		return SpanishHelloPrefix + name
	}
	return EnglishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Gio", ""))

}
