package main

import "fmt"

type Optioner interface {
	Name() string
	IsValid() bool
}

type OptionCommon struct {
	ShortName string "short option name"
	LongName  string "long option name"
}

type IntOption struct {
	OptionCommon
	Value, Min, Max int
}

func (option IntOption) Name() string {
	return name(option.ShortName, option.LongName)
}

func (option IntOption) IsValid() bool {
	return option.Min <= option.Value && option.Value <= option.Max
}

type StringOption struct {
	OptionCommon
	Value string
}

func (option StringOption) Name() string {
	return name(option.ShortName, option.LongName)
}

func (option StringOption) IsValid() bool {
	return true
}

func name(shoetName, longName string) string {
	if longName == "" {
		return shoetName
	}
	return longName
}

type FloatOption struct {
	Optioner
	Value float64
}

type GenericOption struct {
	OptionCommon
}

func (option GenericOption) Name() string {
	return name(option.ShortName, option.LongName)
}

func (option GenericOption) IsValid() bool {
	return true
}

func main() {
	fileOption := StringOption{
		OptionCommon{"f", "file"}, "idex.html"}
	topOption := IntOption{
		OptionCommon: OptionCommon{"t", "top"},
		Max:          100,
	}
	sizeOption := FloatOption{
		GenericOption{OptionCommon{"s", "size"}}, 19.5}

	var a []Optioner
	a = append(a, topOption, fileOption, sizeOption)
	for _, option := range a {
		fmt.Print("name=", option.Name(), "  valid=", option.IsValid())
		fmt.Print(" value=")
		switch option := option.(type) {
		case IntOption:
			fmt.Print(option.Value, " min =", option.Min, " max =", option.Max, "\n")
		case StringOption:
			fmt.Println(option.Value)
		case FloatOption:
			fmt.Println(option.Value)
		}
	}
	// sizeOption := FloatOption{
	// 	GenericOption{OptionCommon{"s", "size"}}, 19.5}

}

// type Count int

// type Tasks struct {
// 	slice []string
// 	Count
// }

// func (count *Count) Increment()  { *count++ }
// func (count *Count) Decrement()  { *count-- }
// func (count Count) IsZero() bool { return count == 0 }

// func (tasks *Tasks) add(task string) {
// 	tasks.slice = append(tasks.slice, task)
// 	tasks.Increment()
// }

// func (tasks *Tasks) Tally() int {
// 	return int(tasks.Count)
// }

// func main() {
// 	tasks := Tasks{}
// 	fmt.Println(tasks.IsZero(), tasks.Tally(), tasks)
// 	tasks.add("One")
// 	tasks.add("Two")
// 	fmt.Println(tasks.IsZero(), tasks.Tally(), tasks)
// }

//-------------------------------------------------------------------------------------------
// type Person struct {
// 	Title     string
// 	Forenames []string
// 	Surname   string
// }

// type Author1 struct {
// 	Names    Person
// 	Title    []string
// 	YearBorn int
// }
// type Author2 struct {
// 	Person
// 	Title    []string
// 	YearBorn int
// }

// func main() {
// 	author1 := Author1{Person{"Mr", []string{"Robert", "Luise", "Balfour"}, "Stevenson"}, []string{"Kidnapped", "Treasure island"}, 1850}
// 	fmt.Println(author1)
// 	author1.Names.Title = ""
// 	author1.Names.Forenames = []string{"Oscar", "Fingal", "O'Flahertie", "Wills"}
// 	author1.Names.Surname = "Wilde"
// 	author1.Title = []string{"The Picture of Dorian Gray"}
// 	author1.YearBorn += 4
// 	fmt.Println(author1)

// 	author2 := Author2{Person{"Mr", []string{"Robert", "Luise", "Balfour"}, "Stevenson"}, []string{"Kidnapped", "Treasure island"}, 1850}
// 	fmt.Println(author2)
// 	author2.Title = []string{"The picture of Dorian Gray"}
// 	author2.Person.Title = ""
// 	author2.Forenames = []string{"Oscar", "Fingal", "O'Flahertie", "Wills"}
// 	author2.Surname = "Wilde"
// 	author2.YearBorn += 5
// 	fmt.Println(author2)
// }
//-------------------------------------------------------------------------------------------
// func main() {
// 	points := [][2]int{{4, 6}, {}, {-7, 11}, {15, 17}, {14, -8}}
// 	for _, point := range points {
// 		fmt.Printf("%d, %d ", point[0], point[1])
// 	}
// 	fmt.Println()
// 	pointsS := []struct{ x, y int }{{4, 6}, {}, {-7, 11}, {15, 17}, {14, -8}}
// 	for _, point := range pointsS {
// 		fmt.Printf("(%d, %d) ", point.x, point.y)
// 	}

// }
