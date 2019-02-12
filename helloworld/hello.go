package helloworld

import "fmt"

const englishPrefix = "Hello "
const spanishPrefix = "Hola "
const frenchPrefix = "Bonjour "
const french = "French"
const spanish = "Spanish"

func Hello(name string, language string) string {
	switch name {
	case "":
		name = "World"
	case "+":
		name = "WORLD"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string){
	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main(){
	fmt.Println(Hello("World", "english"))
}
