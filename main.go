//Write a function that converts a hexadecimal string to its decimal integer value Without using strconv.ParseInt directly ; use it, but explain what base and bitSize mean. // Convert "1E" -> 30, "FF" -> 255 func hexToDecimal(hexStr string) (int64, error) { HexStr := }

package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func hexToDecimal(hexStr string) (int64, error) {
// 	val, err := strconv.ParseInt(hexStr, 16, 64)

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return val, nil
// }

// func main() {
// 	val1, _ := hexToDecimal("1E")
// 	val2, _ := hexToDecimal("FF")

// 	fmt.Println(val1) // 30
// 	fmt.Println(val2) // 255
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func binToDecimal(binStr string) (int64, error) {

// 	result, err := strconv.ParseInt(binStr, 2, 64)

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return result, nil
// }

// func main() {
// 	val1, _ := binToDecimal("10")
// 	val2, _ := binToDecimal("1010")
// 	val3, _ := binToDecimal("11111111")

// 	fmt.Println(val1)
// 	fmt.Println(val2)
// 	fmt.Println(val3)
// }

//Write a function that capitalizes only the first letter of a word, lowercasing the rest // "hELLO" -> "Hello", "WORLD" -> "World"

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func capitalizeWord(word string) string {
// 	return strings.Title(strings.ToLower(word))

// }

// func main() {
// 	test1 := "hELLO" //"Hello"
// 	test2 := "WORLD" //"World"
// 	fmt.Println(capitalizeWord(test1) + "\n" + capitalizeWord(test2))
// }

//Write a function that takes a slice of words and applies uppercase to the last N words // words = ["this", "is", "so", "exciting"], n = 2 // -> ["this", "is", "SO", "EXCITING"]

//package main

// import (
// 	"fmt"
// )

// func uppercaseLastN(words []string, n int) []string {
// 	start := len(words) - n
// 	if start < 0 {
// 		start = 0
// 	}
// 	for i := start; i < len(words); i++ {
// 		words[i] = strings.ToUpper(words[i])
// 	}
// 	return words
// }

// func main() {
// 	n := 2
// 	vart := []string{"this", "is", "so", "exciting"}
// 	fmt.Println(uppercaseLastN(vart, n))
// }

//Write a function that checks if a string is a punctuation mark from the project's list // "," -> true, "!" -> true, "x" -> false

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func isPunctuation(word string) bool {
// 	return strings.ContainsRune("!?:;.,", rune(word[0]))
// }

// func main() {
// 	tests := []string{",", "!", "x"}

// 	for _, t := range tests {
// 		fmt.Printf("%q -> %v\n", t, isPunctuation(t))
// 	}
// }

// Write a function that fixes spacing around punctuation for a simple case Given a slice of tokens (words + punctuation), remove the space before punctuation marks. // ["hello", ",", "world", "!"] -> "hello, world!"

//package main

import (
	"fmt"
	"strings"
)

func spacebpunc(text string) string {
	var sb string.Builder

	for i := 0; i < len(text); i++ {
		if text[i] == ' ' && i+1 <= len(text) && strings.ContainsRune("!?:;,.", rune(text[i+1])) {
			continue
		}
		sb.WriteByte(text[i])
	}
	return sb.string()

}

func main() {
	tokens := []string{"hello", ",", "world", "!"}
	fmt.Println(joinWithPunctuation(tokens)) // hello, world!

	test2 := []string{"I", "am", "angry", "!"}
	fmt.Println(joinWithPunctuation(test2)) // I am angry!

	test3 := []string{"I", "just", "needed", "to", "remember", "slice", "syntax", "that", "is", "why", "?"}
	fmt.Println(joinWithPunctuation(test3))
}

//Write a function that determines whether to use "a" or "an" before a given word // "apple" -> "an" // "horse" -> "an" // "book" -> "a" // "honest" -> "an" (starts with h)

/*
package main

import (
"fmt"
"strings"
)

func aOrAn(nextWord string) string {
if nextWord == "" {
return "a"
}

word := strings.ToLower(nextWord)

// Special cases where 'h' is silent
silentH := []string{"honest", "hour", "horse", "honor", "heir"}

for _, w := range silentH {
if strings.HasPrefix(word, w) {
return "an"
}
}

// Words that start with vowel sounds
vowels := "aeiou"

if strings.ContainsRune(vowels, rune(word[0])) {
return "an"
}

return "a"
}

func main() {
words := []string{"apple", "horse", "book", "honest"}

for _, w := range words {
fmt.Printf("%s %s\n", aOrAn(w), w)
}
}
*/

//Write a function that processes a full sentence and fixes all "a" → "an" corrections

// "There it was. A amazing rock. A honest man. A book." -> "There it was. An amazing rock. An honest man. A book."

/*
package main

import (
"fmt"
"strings"
"unicode"
)

func isVowel(r byte) bool {
switch unicode.ToLower(rune(r)) {
case 'a', 'e', 'i', 'o', 'u':
return true
default:
return false
}
}

func isSilentH(word string) bool {
silent := []string{"honest", "hour", "honor", "heir"}
w := strings.ToLower(word)

for _, s := range silent {
if strings.HasPrefix(w, s) {
return true
}
}
return false
}

func stripPunct(word string) (string, string) {
// separates word from trailing punctuation like "rock." -> ("rock", ".")

i := len(word)
for i > 0 && unicode.IsPunct(rune(word[i-1])) {
i--
}
return word[:i], word[i:]
}

func fixArticles(text string) string {
words := strings.Fields(text)
if len(words) == 0 {
return text
}

for i := 0; i < len(words)-1; i++ {
if strings.EqualFold(words[i], "a") {
nextWord, _ := stripPunct(words[i+1])

if len(nextWord) == 0 {
continue
}
first := nextWord[0]
if isVowel(first) || isSilentH(nextWord) {
// preserve original casing style
if words[i] == "A" {
words[i] = "An"
} else {
words[i] = "an"
}
}
}
}

return strings.Join(words, " ")
}

import (
"fmt"
"simd/archsimd/_gen/simdgen"
"strings"
)
fmt.Println("\nOutput:")
fmt.Println(output)
}
*/

//Write a function that fixes spacing inside single quotes // "' awesome '" -> "'awesome'" // "' hello world '" -> "'hello world'"

/*
package main

import (
"fmt"
"strings"
)

func fixSingleQuotes(text string) string {
var result strings.Builder
inQuotes := false
var buffer strings.Builder

for _, r := range text {
if r == '\'' {
if inQuotes {
// closing quote → process buffer
content := strings.TrimSpace(buffer.String())
result.WriteRune('\'')
result.WriteString(content)
result.WriteRune('\'')
buffer.Reset()
inQuotes = false
} else {
// opening quote
inQuotes = true
buffer.Reset()
}
continue
}

if inQuotes {
buffer.WriteRune(r)
} else {
result.WriteRune(r)
}
}

// in case of unmatched quote, flush safely
if buffer.Len() > 0 {
result.WriteString(buffer.String())
}

return result.String()
}

func main() {
tests := []string{
"' awesome '",
"' hello world '",
"nothing here",
"mix ' spaced text ' outside",
"'multiple   spaces   here'",
}

for _, t := range tests {
fmt.Println(fixSingleQuotes(t))
}

}
*/
