package main

import "fmt"

type Item struct {
	id       string  // Именованное поле (агрегирование)
	price    float64 // Именованное поле (агрегирование)
	quantity int     // Именованное поле (агрегирование)
}

func (item *Item) Cost() float64 {
	return item.price * float64(item.quantity)
}

type SpecialItem struct {
	Item      // Анонимное поле (встраивание)
	catalogId int
}

func main() {
	special := SpecialItem{Item{"Green", 3, 5}, 207}
	fmt.Println(special.id, special.price, special.quantity, special.catalogId)
	fmt.Println(special.Cost())
}

//--------------------------------------------------------------------
// type Part struct {
// 	Id   int    // Именованное поле (агрегирование)
// 	Name string // Именованное поле (агрегирование)
// }

// func (part *Part) LowerCase() {
// 	part.Name = strings.ToLower(part.Name)
// }

// func (part *Part) UpperCase() {
// 	part.Name = strings.ToUpper(part.Name)
// }

// func (part Part) String() string {
// 	return fmt.Sprintf("%d %q", part.Id, part.Name)
// }

// func (part Part) HasPrefix(prefix string) bool {
// 	return strings.HasPrefix(part.Name, prefix)
// }

// func main() {
// 	part := Part{5, "wrench"}
// 	part.UpperCase()
// 	part.Id += 11
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
