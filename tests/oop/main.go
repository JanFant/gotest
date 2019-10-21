package main

import (
	"fmt"
	"strings"
	"unicode"
)

//--------------------------------------------------------------------
// type Exchanger interface {
// 	Exchange()
// }
// type StringPair struct {
// 	first, second string
// }

// func (pair *StringPair) Exchange() {
// 	pair.first, pair.second = pair.second, pair.first
// }

// type Point [2]int

// func (point *Point) Exchange() {
// 	point[0], point[1] = point[1], point[0]
// }

// func (pair StringPair) String() string {
// 	return fmt.Sprintf("%q+%q", pair.first, pair.second)
// }

// func exchangeThese(exchangers ...Exchanger) {
// 	for _, exchanger := range exchangers {
// 		exchanger.Exchange()
// 	}
// }

// func (pair *StringPair) Read(data []byte) (n int, err error) {
// 	if pair.first == "" && pair.second == "" {
// 		return 0, io.EOF
// 	}
// 	if pair.first != "" {
// 		n = copy(data, pair.first)
// 		pair.first = pair.first[n:]
// 	}
// 	if n < len(data) && pair.second != "" {
// 		m := copy(data[n:], pair.second)
// 		pair.second = pair.second[m:]
// 		n += m
// 	}
// 	return n, nil
// }

// func ToBytes(reader io.Reader, size int) ([]byte, error) {
// 	data := make([]byte, size)
// 	n, err := reader.Read(data)
// 	if err != nil {
// 		return data, err
// 	}
// 	return data[:n], nil // Отсечет все неиспользованные байты
// }

// func main() {
// 	// jekyll := StringPair{"Henry", "Jekyll"}
// 	// hyde := StringPair{"Edward", "Hyde"}
// 	// point := Point{5, -3}
// 	// fmt.Println("Before: ", jekyll, hyde, point)
// 	// jekyll.Exchange() // Интерпретируется как: (&jekyll).Exchange()
// 	// hyde.Exchange()   // Интерпретируется как: (&hyde).Exchange()
// 	// point.Exchange()  // Интерпретируется как: (&point).Exchange()
// 	// fmt.Println("After #1:", jekyll, hyde, point)
// 	// exchangeThese(&jekyll, &hyde, &point)
// 	// fmt.Println("After #2:", jekyll, hyde, point)

// 	const size = 16
// 	robert := &StringPair{"Robert L.", "Stevenson"}
// 	david := StringPair{"David", "Balfour"}
// 	for _, reader := range []io.Reader{robert, &david} {
// 		raw, err := ToBytes(reader, size)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Printf("%q\n", raw)
// 	}

// }

//--------------------------------------------------------------------
// type Item struct {
// 	id       string  // Именованное поле (агрегирование)
// 	price    float64 // Именованное поле (агрегирование)
// 	quantity int     // Именованное поле (агрегирование)
// }

// func (item *Item) Cost() float64 {
// 	return item.price * float64(item.quantity)
// }

// type SpecialItem struct {
// 	Item      // Анонимное поле (встраивание)
// 	catalogId int
// }

// type LuxuryItem struct {
// 	Item           // Анонимное поле(встраивание)
// 	markup float64 // Именованное поле (агрегирование)
// }

// func (item *LuxuryItem) Cost() float64 {
// 	return item.Cost() * item.markup
// }

// func main() {
// 	special := SpecialItem{Item{"Green", 3, 5}, 207}
// 	fmt.Println(special.id, special.price, special.quantity, special.catalogId)
// 	fmt.Println(special.Cost())
// }

// --------------------------------------------------------------------

type StringPair struct {
	first, second string
}

type Part struct {
	Id   int    // Именованное поле (агрегирование)
	Name string // Именованное поле (агрегирование)
}

type LowerCaser interface {
	LowerCase()
}

type UpperCaser interface {
	UpperCase()
}
type FixCaser interface {
	FixCase()
}

type LowerUpperCaser interface {
	LowerCaser
	UpperCaser
}

type ChangeCaser interface {
	LowerUpperCaser
	FixCaser
}

func (pair *StringPair) UpperCase() {
	pair.first = strings.ToUpper(pair.first)
	pair.second = strings.ToUpper(pair.second)
}

func (pair *StringPair) LowerCase() {
	pair.first = strings.ToLower(pair.first)
	pair.second = strings.ToLower(pair.second)
}

func (pair *StringPair) FixCase() {
	pair.first = fixCase(pair.first)
	pair.second = fixCase(pair.second)
}

func (part *Part) LowerCase() {
	part.Name = strings.ToLower(part.Name)
}

func (part *Part) UpperCase() {
	part.Name = strings.ToUpper(part.Name)
}

func (pair StringPair) String() string {
	return fmt.Sprintf("%s + %s", pair.first, pair.second)
}

func (part Part) String() string {
	return fmt.Sprintf("%d %q", part.Id, part.Name)
}

func (part Part) HasPrefix(prefix string) bool {
	return strings.HasPrefix(part.Name, prefix)
}

func (part *Part) FixCase() {
	part.Name = fixCase(part.Name)
}

func fixCase(s string) string {
	var chars []rune
	upper := true
	for _, char := range s {
		if upper {
			char = unicode.ToUpper(char)
		} else {
			char = unicode.ToLower(char)
		}
		chars = append(chars, char)
		upper = unicode.IsSpace(char) || unicode.Is(unicode.Hyphen, char)
	}
	return string(chars)
}

func main() {
	toastRack := Part{8427, "TOAST rack"}
	// toastRack.LowerCase()
	lobelia := StringPair{"LOBElia", "SACKVELLE-baggins"}
	// lobelia.LowerCase()

	for _, x := range []FixCaser{&toastRack, &lobelia} {
		x.FixCase()
	}

	fmt.Println(toastRack, lobelia)
}

// func main() {
// 	part := Part{5, "wrench"}
// 	part.UpperCase()
// 	part.Id += 11

// 	asStringV := Part.String
// 	sv := asStringV(part)
// 	HasPrefix := Part.HasPrefix
// 	asStringP := (*Part).String
// 	sp := asStringP(&part)

// 	lower := (*Part).LowerCase
// 	lower(&part)

// 	fmt.Println(sv, sp, HasPrefix(part, "w"), part)

// 	fmt.Println(part, part.HasPrefix("w"))
// }

//--------------------------------------------------------------------
// type Count int

// func (count *Count) Increment()  { *count++ }
// func (count *Count) Decrement()  { *count-- }
// func (count Count) IsZero() bool { return count == 0 }

// func main() {
// 	var count Count
// 	i := int(count)
// 	count.Increment()
// 	j := int(count)
// 	count.Decrement()
// 	k := int(count)
// 	fmt.Println(count, i, j, k, count.IsZero())

// }

// ------------------------------------------------------------
//RuneForRuneFunc aa
// type RuneForRuneFunc func(rune) rune

// var removePunctuation RuneForRuneFunc = func(char rune) rune {
// 	if unicode.Is(unicode.Terminal_Punctuation, char) {
// 		return -1
// 	}
// 	return char
// }

// func main() {
// 	phrases := []string{"Day; dusk, and night.", "All day long"}
// 	processPhrases(phrases, removePunctuation)
// }

// func processPhrases(phrases []string, function RuneForRuneFunc) {
// 	for _, phrase := range phrases {
// 		fmt.Println(strings.Map(function, phrase))
// 	}
// }
