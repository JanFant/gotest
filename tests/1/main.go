package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type statistics struct {
	numbers []float64
	mean    float64
	median  float64
	stdDev  float64
	mode    []float64
}

func getStats(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	sort.Float64s(stats.numbers)
	stats.mean = sum(numbers) / float64(len(numbers))
	stats.median = median(numbers)
	stdDev(&stats)
	return stats
}

func sum(numbers []float64) (total float64) {
	for _, x := range numbers {
		total += x
	}
	return total
}

func median(number []float64) float64 {
	middle := len(number) / 2
	result := number[middle]
	if len(number)%2 == 0 {
		result = (result + number[middle-1]) / 2
	}
	return result
}

func stdDev(stat *statistics) {
	var sum float64
	for _, num := range stat.numbers {
		sum += math.Pow((num - stat.mean), 2)
	}
	stat.stdDev = math.Sqrt(sum / float64(len(stat.numbers)-1))
}

func mode(stat *statistics) {
	var (
		freq []float64
	)
	k := 0
	m := 0
	for _, numj := range stat.numbers {
		freq[k*2] = numj
		for m < len(stat.numbers) {
			if freq[k] == stat.numbers[m] {
				freq[k*2+1]++
				m++
			} else {
				break
			}
		}

	}

}

func main() {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
	}

}

const (
	form = `<form action="/" method="POST">
			<label for="numbers">Numbers (comma or space-separated):</label><br /> 
			<input type="text" name="numbers" size="30"><br /> 
			<input type="submit" value="Calculate"> 
			</form>`
	pageTop = `<!DOCTYPE HTML><html><head>
			   <style>.error{color:#FF0000;}</style></head><title>Statistick</title>
			   <body><h3>Statistis</h3>
			   <p>Computes basic statistics for a given list of numbers</p>`
	pageBottom = `</body></html>`
	anError    = `<p class="error">%s</p>`
)

func homePage(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	fmt.Fprint(writer, pageTop, form)
	if err != nil {
		fmt.Fprintf(writer, anError, err)
	} else {
		if numbers, message, ok := processRequest(request); ok {
			stats := getStats(numbers)
			fmt.Fprintf(writer, formatStats(stats))
		} else if message != "" {
			fmt.Fprintf(writer, anError, message)
		}
	}
	fmt.Fprint(writer, pageBottom)
}

func processRequest(request *http.Request) ([]float64, string, bool) {
	var numbers []float64
	if slice, found := request.Form["numbers"]; found && len(slice) > 0 {
		text := strings.Replace(slice[0], ",", " ", -1)
		for _, field := range strings.Fields(text) {
			if x, err := strconv.ParseFloat(field, 64); err != nil {
				return numbers, "'" + field + "' is invalid", false
			} else {
				numbers = append(numbers, x)
			}
		}
	}
	if len(numbers) == 0 {
		return numbers, "", false
	}
	return numbers, "", true
}

func formatStats(stats statistics) string {
	return fmt.Sprintf(`<table border="1">
	<tr><th colspan="2">Results</th></tr>
	<tr><td>Numbers</td><td>%v</td></tr>
	<tr><td>Count</td><td>%d</td></tr>
	<tr><td>Mean</td><td>%f</td></tr>
	<tr><td>Median</td><td>%f</td></tr>
	<tr><td>Std. Dev.</td><td>%f</td></tr>
	</table>`, stats.numbers, len(stats.numbers), stats.mean, stats.median, stats.stdDev)
}
