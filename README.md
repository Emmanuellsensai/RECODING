# RECODING: Go-Reloaded Practice Functions

A collection of helper functions built while learning text processing in Go.
Each function solves one specific problem from the go-reloaded project.

---

## Table of Contents

1. [Hex to Decimal](#1-hex-to-decimal)
2. [Binary to Decimal](#2-binary-to-decimal)
3. [Capitalize Word](#3-capitalize-word)
4. [Uppercase Last N Words](#4-uppercase-last-n-words)
5. [Is Punctuation](#5-is-punctuation)
6. [Fix Punctuation Spacing](#6-fix-punctuation-spacing)
7. [A or An](#7-a-or-an)
8. [Fix Articles in a Sentence](#8-fix-articles-in-a-sentence)
9. [Fix Single Quote Spacing](#9-fix-single-quote-spacing)

---

## 1. Hex to Decimal

**Task:** Convert a hexadecimal string to its decimal integer value.

```
"1E" -> 30
"FF" -> 255
```

```go
package main

import (
	"fmt"
	"strconv"
)

func hexToDecimal(hexStr string) (int64, error) {
	val, err := strconv.ParseInt(hexStr, 16, 64)

	if err != nil {
		fmt.Println(err)
	}
	return val, nil
}

func main() {
	val1, _ := hexToDecimal("1E")
	val2, _ := hexToDecimal("FF")

	fmt.Println(val1) // 30
	fmt.Println(val2) // 255
}
```

**How it works:**

`strconv.ParseInt` takes three arguments:
- The string to parse (`hexStr`)
- The **base** — `16` means hexadecimal (digits are `0-9` and `a-f`)
- The **bitSize** — `64` means the result fits into a 64-bit integer (range up to ~9.2 quintillion)

For `"1E"`: `1×16 + 14 = 30` ✓  
For `"FF"`: `15×16 + 15 = 255` ✓

The function returns two values — the result and an error — because Go expects you to handle the case where the input isn't valid hex.

---

## 2. Binary to Decimal

**Task:** Convert a binary string to its decimal integer value.

```
"10"       -> 2
"1010"     -> 10
"11111111" -> 255
```

```go
package main

import (
	"fmt"
	"strconv"
)

func binToDecimal(binStr string) (int64, error) {
	result, err := strconv.ParseInt(binStr, 2, 64)

	if err != nil {
		fmt.Println(err)
	}
	return result, nil
}

func main() {
	val1, _ := binToDecimal("10")
	val2, _ := binToDecimal("1010")
	val3, _ := binToDecimal("11111111")

	fmt.Println(val1)
	fmt.Println(val2)
	fmt.Println(val3)
}
```

**How it works:**

Same as `hexToDecimal` but base `2` instead of `16`. In base 2, the only valid digits are `0` and `1`. Each position represents a power of 2.

For `"1010"`: `1×8 + 0×4 + 1×2 + 0×1 = 10` ✓  
For `"11111111"`: all 8 bits set = 255 ✓

---

## 3. Capitalize Word

**Task:** Capitalize only the first letter of a word, lowercasing the rest.

```
"hELLO" -> "Hello"
"WORLD" -> "World"
```

```go
package main

import (
	"fmt"
	"strings"
)

func capitalizeWord(word string) string {
	return strings.Title(strings.ToLower(word))
}

func main() {
	test1 := "hELLO" //"Hello"
	test2 := "WORLD" //"World"
	fmt.Println(capitalizeWord(test1) + "\n" + capitalizeWord(test2))
}
```

**How it works:**

Two steps happen inside one return:
1. `strings.ToLower(word)` — converts the entire word to lowercase first, so `"hELLO"` becomes `"hello"`
2. `strings.Title(...)` — capitalizes the first letter, turning `"hello"` into `"Hello"`

The two functions are nested, so the output of `ToLower` feeds directly into `Title` without needing a separate variable.

**Task 2:** write a function that takes a string as input and returns the string with the first letter of each word capitalized. For example, if the input is "hello world! my name is james bond.", the output should be "Hello World! My Name Is James Bond.".

```go
package main

import (
	"fmt"
	"strings"
)

// write a function that takes a string as input and returns the string with the first letter of each word capitalized. For example, if the input is "hello world", the output should be "Hello World".
func Capitalize(word string) string {
	words := strings.Fields(word)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + word[1:]
		}
	}
	return strings.Join(words, " ")
}

func main() {
	input := "hello world! my name is james bond."
	output := Capitalize(input)
	fmt.Println(output) // Output: "Hello World! My Name Is James Bond."
}
```

---

## 4. Uppercase Last N Words

**Task:** Apply uppercase to the last N words in a slice.

```
words = ["this", "is", "so", "exciting"], n = 2
-> ["this", "is", "SO", "EXCITING"]
```

```go
package main

import (
	"fmt"
	"strings"
)

func uppercaseLastN(words []string, n int) []string {
	start := len(words) - n
	if start < 0 {
		start = 0
	}
	for i := start; i < len(words); i++ {
		words[i] = strings.ToUpper(words[i])
	}
	return words
}

func main() {
	n := 2
	vart := []string{"this", "is", "so", "exciting"}
	fmt.Println(uppercaseLastN(vart, n))
}
```

**How it works:**

`start` is calculated as `len(words) - n`, giving you the index of the first word to uppercase. The `if start < 0` guard handles the case where `n` is larger than the slice — instead of crashing, it just starts from index 0 and uppercases everything.

For `n=2` and 4 words: `start = 4 - 2 = 2`, so it uppercases from index 2 onwards — `"so"` and `"exciting"`.

---

## 5. Is Punctuation

**Task:** Check if a string is a punctuation mark from the project's list.

```
"," -> true
"!" -> true
"x" -> false
```

```go
package main

import (
	"fmt"
	"strings"
)

func isPunctuation(word string) bool {
	if strings.ContainsRune("!?:;.,", rune(word[0])) {
		return true
	}
	return false
}

func main() {
	tests := []string{",", "!", "x"}

	for _, t := range tests {
		fmt.Printf("%q -> %v\n", t, isPunctuation(t))
	}
}
```

**How it works:**

`word[0]` gets the first byte of the string. `rune(word[0])` converts it to a Unicode character. `strings.ContainsRune` checks if that character exists anywhere in the punctuation string `"!?:;.,"`.

Returns `true` if it's in the set, `false` if not. One line, no loop needed.

---

## 6. Fix Punctuation Spacing

**Task:** Given a slice of tokens, remove the space before punctuation marks.

```
["hello", ",", "world", "!"] -> "hello, world!"
```

```go
package main

import (
	"fmt"
	"strings"
)

func fixPunctuation(word []string) string {
	var sb strings.Builder
	result := strings.Join(word, " ")

	for i := 0; i < len(result); i++ {
		if result[i] == ' ' && i+1 < len(result) && strings.ContainsRune("!?,.:;", rune(result[i+1])) {
			continue
		}
		sb.WriteByte(result[i])
	}
	return sb.String()
}

func main() {
	token1 := []string{"hello", ",", "world", "!"}
	fmt.Println(fixPunctuation(token1)) // hello, world!

	test2 := []string{"I", "am", "angry", "!"}
	fmt.Println(fixPunctuation(test2)) // I am angry!

	test3 := []string{"I", "just", "needed", "to", "remember", "slice", "syntax", "that", "is", "why", "?"}
	fmt.Println(fixPunctuation(test3))
}
```

**How it works:**

First, `strings.Join` glues all tokens together with spaces — giving you `"hello , world !"`.

Then the loop ranges over each punctuation character in `".,!?:;"`. tuation and removes the space.

Ranging directly over a string in Go gives you each character as a `rune`, so no need for a separate slice.

- After handling `","`: `"hello, world !"`
- After handling `"!"`: `"hello, world!"`

---

## 7. A or An

**Task:** Determine whether to use `"a"` or `"an"` before a given word.

```
"apple"  -> "an"
"horse"  -> "an"  (h is treated as needing "an")
"book"   -> "a"
"honest" -> "an"
```

```go
package main

import (
	"fmt"
	"strings"
)

func aOrAn(word string) string {
	if strings.ContainsRune("aeiouAEIOUhH", rune(word[0])) {
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
```

**How it works:**

Checks the first character of the word against the set `"aeiouAEIOUhH"`. If it's a vowel or `h`/`H`, return `"an"`. Otherwise return `"a"`.

The project treats `h` as requiring `"an"` — so `"a horse"` becomes `"an horse"`, `"a honest man"` becomes `"an honest man"`.

---

## 8. Fix Articles in a Sentence

**Task:** Walk a full sentence and fix all `"a"` → `"an"` corrections where needed.

```
"There it was. A amazing rock. A honest man. A book."
-> "There it was. An amazing rock. An honest man. A book."
```

```go
package main

import (
	"fmt"
	"strings"
)

func article(word string) string {
	if strings.ContainsRune("aeiouAEIOUhH", rune(word[0])) {
		return "An"
	}
	return "A"
}

func fixArticles(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" || words[i] == "A" {
			words[i] = article(words[i+1])
		}
	}
	return strings.Join(words, " ")
}

func main() {
	input := "There it was. A amazing rock. A honest man. A book."
	output := fixArticles(input)
	fmt.Println("\nOutput:")
	fmt.Println(output)
}
```

**How it works:**

`strings.Fields` splits the sentence into individual words. The loop walks through every word except the last one — `len(words)-1` stops one short so `words[i+1]` never goes out of bounds.

When it finds `"a"` or `"A"`, it passes the **next word** to `article()`, which checks what that next word starts with and returns the correct article. That result is written back onto the current position.

For `"A amazing"`: `words[i]` is `"A"`, `words[i+1]` is `"amazing"`, `article("amazing")` returns `"An"`, so `words[i]` becomes `"An"`. The word `"amazing"` itself never changes.

`"A book"` is left unchanged because `'b'` is not in the vowel/h set.

---

## 9. Fix Single Quote Spacing

**Task:** Remove spaces immediately inside single quotes.

```
"' awesome '"     -> "'awesome'"
"' hello world '" -> "'hello world'"
```

```go
package main

import (
	"fmt"
	"strings"
)

func fixSingleQuotes(text string) string {
	result := strings.ReplaceAll(text, "' ", "'")
	result = strings.ReplaceAll(result, " '", "'")
	return result
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
```

**How it works:**

Two replacements handle both sides of the quote:
1. `"' "` → `"'"` — removes the space after an opening quote
2. `" '"` → `"'"` — removes the space before a closing quote

For `"' awesome '"`:
- After step 1: `"'awesome '"`
- After step 2: `"'awesome'"`

Internal spaces between words inside the quotes are untouched — only the spaces directly touching the quote characters are removed. `"nothing here"` passes through completely unchanged since it has no quotes at all.

---

## Running Any Function

Each function is in its own commented-out block. Uncomment the block you want to test and run:

```bash
go run main.go
```

---

*Built as part of the go-reloaded project at Zone01.*
