package tool

import "fmt"

// Aesthetic adds inserts spaces in between each character and pre/post pends with ~
func Aesthetic(unaesthetic string) (string, error) {
	if len(unaesthetic) == 0 {
		return "", fmt.Errorf("need statement to make aesthetic")
	}

	aesthetic := []rune{'~', ' '}
	for _, letter := range unaesthetic {
		switch letter {
		case ' ':
			continue
		default:
			aesthetic = append(aesthetic, letter, ' ')
		}
	}

	aesthetic = append(aesthetic, '~')

	return string(aesthetic), nil
}

// Voweless removes all vowels in string
func Voweless(vowelful string) (string, error) {
	if len(vowelful) == 0 {
		return "", fmt.Errorf("need statement to make voweless")
	}

	var voweless []rune
	for _, letter := range vowelful {
		switch letter {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U', ' ':
			continue
		default:
			voweless = append(voweless, letter)
		}
	}

	return string(voweless), nil
}

// DeBruijn give the De Bruijn Sequence given
func DeBruijn(alphabet string, subSequenceSize int) (string, error) {
	characters := []rune(alphabet)
	charactersCount := len(characters)
	a := make([]int, charactersCount*subSequenceSize)

	var deBruijnRunner func(t, p int) string
	deBruijnRunner = func(t, p int) string {
		if t > subSequenceSize {

		} else {
			a[t] = a[t-p]
			deBruijnRunner(t+1, p)
			for j := a[t-p] + 1; j <= charactersCount; j++ {
				a[t] = j
				deBruijnRunner(t+1, t)
			}
		}
		return "trash"
	}

	fmt.Println(alphabet, characters, subSequenceSize)

	return deBruijnRunner(1, 1), nil
}
