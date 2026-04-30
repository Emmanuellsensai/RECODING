# Go-Reloaded RECODING Practice ⟲

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
10. [Count Alphabetic Characters](#10-count-alphabetic-characters)
11. [Repeat String N Times](#11-repeat-string-n-times)
12. [Largest Integer in Slice](#12-largest-integer-in-slice)
13. [Remove Spaces](#13-remove-spaces)
14. [Is Palindrome (Case-Insensitive)](#14-is-palindrome-case-insensitive)
15. [Word Frequency](#15-word-frequency)
16. [Remove Vowels](#16-remove-vowels)
17. [Caesar Cipher](#17-caesar-cipher)
18. [Trim Spaces Inside Quotes](#18-trim-spaces-inside-quotes)

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
	return strings.ContainsRune("!?:;.,", rune(word[0]))
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

Then the loop ranges over each punctuation character in the string. On each pass,  finds every occurrence of space before a punctuation and removes the space then writes the string back without the space.

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

## 10. Count Alphabetic Characters

**Task:** Count only alphabetic characters in a string, ignoring digits, spaces, and symbols.

```
"Hello 123!" -> 5
```

```go
package main

import "fmt"

func CountAlpha(str string) int {
	count := 0
	for _, v := range str {
		if (v >= 'A' && v <= 'Z') || (v >= 'a' && v <= 'z') {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountAlpha("Hello 123"))
}
```

**How it works:**

The loop ranges over every character in the string as a `rune`. The `if` condition checks if the character falls within the ASCII range of uppercase (`'A'` to `'Z'`) or lowercase (`'a'` to `'z'`) letters. If yes, `count` is incremented. Digits, spaces, and symbols all fall outside these ranges, so they're skipped automatically.

For `"Hello 123"`: `H`, `e`, `l`, `l`, `o` = 5 letters. The space and `1`, `2`, `3` are ignored → returns `5`.

---

## 11. Repeat String N Times

**Task:** Return a string repeated n times. If n is 0 or negative, return an empty string.

```
Repeat("ha", 3) -> "hahaha"
Repeat("go", 0) -> ""
```

```go
package main

import (
	"fmt"
	"strings"
)

func Repeatstri(str string, num int) string {
	if num <= 0 {
		return ""
	}
	return strings.Repeat(str, num)
}

func main() {
	fmt.Println(Repeatstri("ha", 5))
}
```

**How it works:**

The guard `if num <= 0` handles the edge case first — returning an empty string immediately without calling anything else. For valid values, `strings.Repeat` does the heavy lifting: it takes a string and a count and returns the string concatenated that many times.

The guard matters because `strings.Repeat` panics if you pass a negative number — so the check protects you from that.

---

## 12. Largest Integer in Slice

**Task:** Return the largest integer in a slice. Return an error if the slice is empty.

```
{3, 1, 4, 1, 5, 9} -> 9, nil
```

```go
package main

import (
	"errors"
	"fmt"
	"slices"
)

func largestint(integer1 []int) (int, error) {
	if len(integer1) == 0 {
		return 0, errors.New("Error")
	}
	maxi := slices.Max(integer1)
	return maxi, nil
}

func main() {
	fmt.Println(largestint([]int{3, 1, 4, 11, 5, 9}))
}
```

**How it works:**

The empty check comes first — if the slice has no elements, `slices.Max` would panic, so we return `0` and an error before it ever gets called. `errors.New("Error")` creates a simple error value carrying that message.

For a non-empty slice, `slices.Max` walks the slice and returns the highest value. The function then returns that alongside `nil` (meaning no error).

Go's multiple return values make this pattern clean — the caller gets both the result and the error in one call and can check them separately.

---

## 13. Remove Spaces

**Task:** Remove all extra whitespace from a string, collapsing it to no spaces at all.

```
"   Hello      world   " -> "Helloworld"
```

```go
package main

import (
	"fmt"
	"strings"
)

func removespace(text string) string {
	var sb strings.Builder
	words := strings.Fields(text)

	for _, word := range words {
		sb.WriteString(word)
	}
	return strings.Join(words, "")
}

func main() {
	fmt.Println(removespace("                  Hello               world        "))
	fmt.Println(len(removespace("                  Hello               world        ")))
}
```

**How it works:**

`strings.Fields` splits the text on any whitespace (spaces, tabs, newlines) and returns only the non-empty chunks — so no matter how many spaces are between words, you just get the words themselves.

`strings.Join(words, "")` glues them back together with an empty string as separator, meaning no spaces at all between them.

The `sb.WriteString` loop and the `Join` are both doing the same thing here — the `Join` at the end is what actually gets returned.

---

## 14. Is Palindrome (Case-Insensitive)

**Task:** Check if a string is a palindrome, ignoring case.

```
isPalindrome("Racecar") -> true
isPalindrome("Hello")   -> false
```

```go
package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(word string) bool {
	word = strings.ToLower(word)
	r := []rune(word)

	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	words := []string{"racecar", "hello", "madam", "world", "level"}

	for _, word := range words {
		fmt.Println(IsPalindrome(word))
	}
}
```

**How it works:**

First, `strings.ToLower` normalizes the case so `"Racecar"` and `"racecar"` are treated the same.

Then the string is converted to `[]rune` — a slice of Unicode characters. This matters for correctness with non-ASCII text, but for plain English it's equivalent to `[]byte`.

The loop only goes halfway (`len(r)/2`) — there's no need to go further because you're comparing from both ends simultaneously. `r[i]` is the character from the left, `r[len(r)-1-i]` is its mirror from the right. If any pair mismatches, it's not a palindrome. If the loop finishes without returning false, it is.

---

## 15. Word Frequency

**Task:** Count how many times each word appears in a string.

```
"go is fun and go is easy" -> map[go:2 is:2 fun:1 and:1 easy:1]
```

```go
package main

import (
	"fmt"
	"strings"
)

func wordFreq(str string) map[string]int {
	result := make(map[string]int)
	words := strings.Fields(str)

	for _, word := range words {
		result[word]++
	}
	return result
}

func main() {
	fmt.Println(wordFreq("go is fun and go easy"))
}
```

**How it works:**

`make(map[string]int)` creates an empty map where keys are strings (words) and values are integers (counts).

`strings.Fields` splits the sentence into words. The loop walks each word — if the key already exists in the map, `++` increments it. If the key doesn't exist yet, Go automatically initializes it to `0` before incrementing, so you get `1` on the first occurrence. No need to check if the key exists first.

For `"go is fun and go easy"`:
- First `"go"` → `map[go:1]`
- `"is"` → `map[go:1 is:1]`
- `"fun"` → `map[go:1 is:1 fun:1]`
- Second `"go"` → `map[go:2 is:1 fun:1]`
- and so on.

---

## 16. Remove Vowels

**Task:** Remove all vowels from a string.

```
"hello world" -> "hll wrld"
```

```go
package main

import (
	"fmt"
	"strings"
)

func removeVowels(str string) string {
	var sb strings.Builder
	for _, char := range str {
		if strings.ContainsRune("aeiouAEIOU", char) {
			continue
		}
		sb.WriteRune(char)
	}
	return sb.String()
}

func main() {
	fmt.Println(removeVowels("A boy like micheal is dumb ASF...."))
}
```

**How it works:**

`strings.Builder` is an efficient way to build up a string one character at a time without creating lots of intermediate string copies.

The loop ranges over every character. `strings.ContainsRune("aeiouAEIOU", char)` checks if the character is a vowel — both upper and lowercase covered. If it is a vowel, `continue` skips the rest of the loop body (skipping `WriteRune`), so it never gets added. If it's not a vowel, `sb.WriteRune(char)` adds it to the result.

`sb.String()` at the end returns everything that was written.

---

## 17. Caesar Cipher

**Task:** Encrypt a string by shifting each letter n positions forward in the alphabet. Wraps around at the end.

```
caesarCipher("abc", 1) -> "bcd"
caesarCipher("xyz", 2) -> "zab"
```

```go
package main

import (
	"fmt"
	"strings"
)

func ceaserCipher(str string, num int) string {
	var sb strings.Builder
	for _, ch := range str {
		shifted := ch + rune(num)
		if (ch >= 'a' && ch <= 'z' && shifted > 'z') || (ch >= 'A' && ch <= 'Z' && shifted > 'Z') {
			shifted -= 26
		}
		sb.WriteRune(shifted)
	}
	return sb.String()
}

func main() {
	fmt.Println(ceaserCipher("abc", 1))
}
```

**How it works:**

For each character, `shifted := ch + rune(num)` moves it forward by `num` positions in the ASCII table. Since letters are contiguous in ASCII, adding 1 to `'a'` gives `'b'`, adding 1 to `'z'` would give `'{'` — which is wrong.

The `if` condition catches that overflow: if the character was lowercase and the shift went past `'z'`, subtract 26 to wrap back into the lowercase range. Same logic for uppercase letters wrapping past `'Z'`.

For `"xyz"` with shift 2:
- `'x'` + 2 = `'z'` — no overflow ✓
- `'y'` + 2 = `'{'` — overflow! `'{' - 26 = 'a'` ✓
- `'z'` + 2 = `'|'` — overflow! `'|' - 26 = 'b'` ✓

Result: `"zab"` ✓

---

## 18. Trim Spaces Inside Quotes

**Task:** Remove all spaces inside single quotes, including internal ones.

```
"'  hello  '" -> "'hello'"
```

```go
package main

import (
	"fmt"
	"strings"
)

func trimwquote(str string) string {
	word := strings.Fields(str)
	return strings.Join(word, "")
}

func main() {
	result := "'  hello      '"
	fmt.Printf("%q\n", trimwquote(result))
}
```

**How it works:**

`strings.Fields` splits the string on all whitespace and returns only the non-empty pieces. For `"'  hello  '"` it gives `["'", "hello", "'"]`.

`strings.Join(word, "")` glues them back with no separator — giving `"'hello'"`.

The `%q` format verb in `fmt.Printf` prints the string with surrounding double quotes and escaped characters, making it easy to see exactly what's in the output including any hidden spaces.

---

## Running Any Function

Each function is in its own commented-out block. Uncomment the block you want to test and run:

```bash
go run main.go
```

---

# Go-Reloaded Practice Functions

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
10. [Count Alphabetic Characters](#10-count-alphabetic-characters)
11. [Repeat String N Times](#11-repeat-string-n-times)
12. [Largest Integer in Slice](#12-largest-integer-in-slice)
13. [Remove Spaces](#13-remove-spaces)
14. [Is Palindrome (Case-Insensitive)](#14-is-palindrome-case-insensitive)
15. [Word Frequency](#15-word-frequency)
16. [Remove Vowels](#16-remove-vowels)
17. [Caesar Cipher](#17-caesar-cipher)
18. [Trim Spaces Inside Quotes](#18-trim-spaces-inside-quotes)

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
	return strings.ContainsRune("!?:;.,", rune(word[0]))
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
	result := strings.Join(word, " ")
	for _, p := range ".,!?:;" {
		result = strings.ReplaceAll(result, " "+string(p), string(p))
	}
	return result
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

Then the loop ranges over each punctuation character in `".,!?:;"`. On each pass, `strings.ReplaceAll` finds every occurrence of space + that punctuation and removes the space.

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

## 10. Count Alphabetic Characters

**Task:** Count only alphabetic characters in a string, ignoring digits, spaces, and symbols.

```
"Hello 123!" -> 5
```

```go
package main

import "fmt"

func CountAlpha(str string) int {
	count := 0
	for _, v := range str {
		if (v >= 'A' && v <= 'Z') || (v >= 'a' && v <= 'z') {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountAlpha("Hello 123"))
}
```

**How it works:**

The loop ranges over every character in the string as a `rune`. The `if` condition checks if the character falls within the ASCII range of uppercase (`'A'` to `'Z'`) or lowercase (`'a'` to `'z'`) letters. If yes, `count` is incremented. Digits, spaces, and symbols all fall outside these ranges, so they're skipped automatically.

For `"Hello 123"`: `H`, `e`, `l`, `l`, `o` = 5 letters. The space and `1`, `2`, `3` are ignored → returns `5`.

---

## 11. Repeat String N Times

**Task:** Return a string repeated n times. If n is 0 or negative, return an empty string.

```
Repeat("ha", 3) -> "hahaha"
Repeat("go", 0) -> ""
```

```go
package main

import (
	"fmt"
	"strings"
)

func Repeatstri(str string, num int) string {
	if num <= 0 {
		return ""
	}
	return strings.Repeat(str, num)
}

func main() {
	fmt.Println(Repeatstri("ha", 5))
}
```

**How it works:**

The guard `if num <= 0` handles the edge case first — returning an empty string immediately without calling anything else. For valid values, `strings.Repeat` does the heavy lifting: it takes a string and a count and returns the string concatenated that many times.

The guard matters because `strings.Repeat` panics if you pass a negative number — so the check protects you from that.

---

## 12. Largest Integer in Slice

**Task:** Return the largest integer in a slice. Return an error if the slice is empty.

```
{3, 1, 4, 1, 5, 9} -> 9, nil
```

```go
package main

import (
	"errors"
	"fmt"
	"slices"
)

func largestint(integer1 []int) (int, error) {
	if len(integer1) == 0 {
		return 0, errors.New("Error")
	}
	maxi := slices.Max(integer1)
	return maxi, nil
}

func main() {
	fmt.Println(largestint([]int{3, 1, 4, 11, 5, 9}))
}
```

**How it works:**

The empty check comes first — if the slice has no elements, `slices.Max` would panic, so we return `0` and an error before it ever gets called. `errors.New("Error")` creates a simple error value carrying that message.

For a non-empty slice, `slices.Max` walks the slice and returns the highest value. The function then returns that alongside `nil` (meaning no error).

Go's multiple return values make this pattern clean — the caller gets both the result and the error in one call and can check them separately.

---

## 13. Remove Spaces

**Task:** Remove all extra whitespace from a string, collapsing it to no spaces at all.

```
"   Hello      world   " -> "Helloworld"
```

```go
package main

import (
	"fmt"
	"strings"
)

func removespace(text string) string {
	var sb strings.Builder
	words := strings.Fields(text)

	for _, word := range words {
		sb.WriteString(word)
	}
	return strings.Join(words, "")
}

func main() {
	fmt.Println(removespace("                  Hello               world        "))
	fmt.Println(len(removespace("                  Hello               world        ")))
}
```

**How it works:**

`strings.Fields` splits the text on any whitespace (spaces, tabs, newlines) and returns only the non-empty chunks — so no matter how many spaces are between words, you just get the words themselves.

`strings.Join(words, "")` glues them back together with an empty string as separator, meaning no spaces at all between them.

The `sb.WriteString` loop and the `Join` are both doing the same thing here — the `Join` at the end is what actually gets returned.

---

## 14. Is Palindrome (Case-Insensitive)

**Task:** Check if a string is a palindrome, ignoring case.

```
isPalindrome("Racecar") -> true
isPalindrome("Hello")   -> false
```

```go
package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(word string) bool {
	word = strings.ToLower(word)
	r := []rune(word)

	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	words := []string{"racecar", "hello", "madam", "world", "level"}

	for _, word := range words {
		fmt.Println(IsPalindrome(word))
	}
}
```

**How it works:**

First, `strings.ToLower` normalizes the case so `"Racecar"` and `"racecar"` are treated the same.

Then the string is converted to `[]rune` — a slice of Unicode characters. This matters for correctness with non-ASCII text, but for plain English it's equivalent to `[]byte`.

The loop only goes halfway (`len(r)/2`) — there's no need to go further because you're comparing from both ends simultaneously. `r[i]` is the character from the left, `r[len(r)-1-i]` is its mirror from the right. If any pair mismatches, it's not a palindrome. If the loop finishes without returning false, it is.

---

## 15. Word Frequency

**Task:** Count how many times each word appears in a string.

```
"go is fun and go is easy" -> map[go:2 is:2 fun:1 and:1 easy:1]
```

```go
package main

import (
	"fmt"
	"strings"
)

func wordFreq(str string) map[string]int {
	result := make(map[string]int)
	words := strings.Fields(str)

	for _, word := range words {
		result[word]++
	}
	return result
}

func main() {
	fmt.Println(wordFreq("go is fun and go easy"))
}
```

**How it works:**

`make(map[string]int)` creates an empty map where keys are strings (words) and values are integers (counts).

`strings.Fields` splits the sentence into words. The loop walks each word — if the key already exists in the map, `++` increments it. If the key doesn't exist yet, Go automatically initializes it to `0` before incrementing, so you get `1` on the first occurrence. No need to check if the key exists first.

For `"go is fun and go easy"`:
- First `"go"` → `map[go:1]`
- `"is"` → `map[go:1 is:1]`
- `"fun"` → `map[go:1 is:1 fun:1]`
- Second `"go"` → `map[go:2 is:1 fun:1]`
- and so on.

---

## 16. Remove Vowels

**Task:** Remove all vowels from a string.

```
"hello world" -> "hll wrld"
```

```go
package main

import (
	"fmt"
	"strings"
)

func removeVowels(str string) string {
	var sb strings.Builder
	for _, char := range str {
		if strings.ContainsRune("aeiouAEIOU", char) {
			continue
		}
		sb.WriteRune(char)
	}
	return sb.String()
}

func main() {
	fmt.Println(removeVowels("A boy like micheal is dumb ASF...."))
}
```

**How it works:**

`strings.Builder` is an efficient way to build up a string one character at a time without creating lots of intermediate string copies.

The loop ranges over every character. `strings.ContainsRune("aeiouAEIOU", char)` checks if the character is a vowel — both upper and lowercase covered. If it is a vowel, `continue` skips the rest of the loop body (skipping `WriteRune`), so it never gets added. If it's not a vowel, `sb.WriteRune(char)` adds it to the result.

`sb.String()` at the end returns everything that was written.

---

## 17. Caesar Cipher

**Task:** Encrypt a string by shifting each letter n positions forward in the alphabet. Wraps around at the end.

```
caesarCipher("abc", 1) -> "bcd"
caesarCipher("xyz", 2) -> "zab"
```

```go
package main

import (
	"fmt"
	"strings"
)

func ceaserCipher(str string, num int) string {
	var sb strings.Builder
	for _, ch := range str {
		shifted := ch + rune(num)
		if (ch >= 'a' && ch <= 'z' && shifted > 'z') || (ch >= 'A' && ch <= 'Z' && shifted > 'Z') {
			shifted -= 26
		}
		sb.WriteRune(shifted)
	}
	return sb.String()
}

func main() {
	fmt.Println(ceaserCipher("abc", 1))
}
```

**How it works:**

For each character, `shifted := ch + rune(num)` moves it forward by `num` positions in the ASCII table. Since letters are contiguous in ASCII, adding 1 to `'a'` gives `'b'`, adding 1 to `'z'` would give `'{'` — which is wrong.

The `if` condition catches that overflow: if the character was lowercase and the shift went past `'z'`, subtract 26 to wrap back into the lowercase range. Same logic for uppercase letters wrapping past `'Z'`.

For `"xyz"` with shift 2:
- `'x'` + 2 = `'z'` — no overflow ✓
- `'y'` + 2 = `'{'` — overflow! `'{' - 26 = 'a'` ✓
- `'z'` + 2 = `'|'` — overflow! `'|' - 26 = 'b'` ✓

Result: `"zab"` ✓

---

## 18. Trim Spaces Inside Quotes

**Task:** Remove all spaces inside single quotes, including internal ones.

```
"'  hello  '" -> "'hello'"
```

```go
package main

import (
	"fmt"
	"strings"
)

func trimwquote(str string) string {
	word := strings.Fields(str)
	return strings.Join(word, "")
}

func main() {
	result := "'  hello      '"
	fmt.Printf("%q\n", trimwquote(result))
}
```

**How it works:**

`strings.Fields` splits the string on all whitespace and returns only the non-empty pieces. For `"'  hello  '"` it gives `["'", "hello", "'"]`.

`strings.Join(word, "")` glues them back with no separator — giving `"'hello'"`.

The `%q` format verb in `fmt.Printf` prints the string with surrounding double quotes and escaped characters, making it easy to see exactly what's in the output including any hidden spaces.

---

## Running Any Function

Each function is in its own commented-out block. Uncomment the block you want to test and run:

```bash
go run main.go
```

---

*Built as part of the go-reloaded project at Learn2Earn.*

## Author

**Emmanuel Usang** — [GitHub](https://github.com/Emmanuellsensai) [Portfolio](https://emmanuellsensai.github.io)
