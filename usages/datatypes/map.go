package main

import (
	"log"
	"sort"
	"time"
)

/**
When using the first way, if use codes below
var ranks map[string]int

ranks["a"] = 2
log.Println(ranks)
It will make an error, such as "panic: assignment to entry in nil map".
Because it is just declare a map, but don't assign a value to it. At that time, its value is nil.
That is the reason why that error come out.
*/
func mapDefineFirstWay() {
	var nilMap map[string]int // Declare a nip map.
	log.Println("nilMap == nil is", nilMap == nil, nilMap)

	// ranks["a"] = 2
	// log.Println(ranks)
	// panic: assignment to entry in nil map

	// Create an empty map
	ranks := make(map[string]int)

	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3
	log.Printf("bronze is %v, gold is %v, ranks content is %v\n", ranks["bronze"], ranks["gold"], ranks)
	// bronze is 3, gold is 1, ranks content is map[bronze:3 gold:1 silver:2]

	colors := make(map[int]string)
	colors[1] = "green"
	colors[2] = "red"
	colors[4] = "yellow"
	log.Println(colors) // map[1:green 2:red 4:yellow]
}

/**
The map variable is map[string]float64{"a":1.2, "b":5.6}.
string is the map key type.
float64 is the map value type.
*/
func mapDefineSecWay() {
	// Create a map and declare a variable to hold it, and set the content for map..
	myMap := map[string]float64{"a": 1.2, "b": 5.6}

	log.Printf("myMap content is %v\n", myMap) // map[a:1.2 b:5.6]
}

func createEmptyMap() {
	// As with slice literals, leaving the curly braces empty creates a map that starts empty.
	emptyFloatValueMap := map[string]float64{}

	// As with arrays and slices, if you access a map key it hasn’t been assigned to, you’ll get a zero value back.
	log.Printf("emptyFloatValueMap content is %v, a undefined element a value is %v\n",
		emptyFloatValueMap, emptyFloatValueMap["a"]) // emptyFloatValueMap content is map[], a undefined element value is 0

	emptyIntValueMap := map[string]int{}
	log.Printf("emptyIntValueMap content is %v, a undefined element a value is %v\n",
		emptyIntValueMap, emptyIntValueMap["c"]) // emptyIntValueMap content is map[], a undefined element value is 0

	emptyStringValueMap := map[string]string{}
	log.Printf("emptyStringValueMap content is %v, a undefined element a value is %v\n",
		emptyStringValueMap, emptyStringValueMap["d"]) // emptyStringValueMap content is map[], a undefined element a value is ""
}

/**
Accessing a map key optionally returns a second, Boolean value. It will be true if the returned value has
actually been assigned to the map, or false if the returned value just represents the default zero value.
Most Go developers assign this Boolean value to a variable named ok (because the name is nice and short).

delete() method will make a key and its corresponding value will be removed from the map.
*/
func checkElementExistAndDeleteElement() {
	counters := map[string]int{"a": 3, "b": 0, "d": 9}
	var value int
	var ok bool
	value, ok = counters["a"]
	log.Println(value, ok) // 3 true

	// Delete the "a" key and its corresponding value.
	delete(counters, "a")
	value, ok = counters["a"]
	// After delete "a" key, the value is 0, not 3. The ok is the false.
	log.Println(value, ok) // 0 false

	value, ok = counters["c"]
	log.Println(value, ok) // 0 false

	value, ok = counters["b"]
	log.Println(value, ok) // 0 true
}

// The for...range loop handles maps in random order!
// When using loop values or loop keys several times, it will show this issue easily.
func mapForRange() {
	grades := map[string]float64{"Alma": 74.2, "Rohit": 87.1, "Carl": 67.2}

	execTime0 := time.Now()
	// Loop key and value
	for name, grade := range grades {
		// If it wants to output % in log.Printf(), it should have two % as "%%", otherwise it won't have % output.
		log.Printf("%s has a grade of %0.1f%%\n", name, grade)
	}
	log.Println(time.Since(execTime0))

	execTime1 := time.Now()
	// Loop only keys
	for name := range grades {
		log.Printf("key name is %v\n", name)
	}
	log.Println(time.Since(execTime1))

	execTime2 := time.Now()
	// Loop only values
	for _, value := range grades {
		log.Printf("the map value is %v\n", value)
	}
	log.Println(time.Since(execTime2))
}

func mapSortLoop() {
	grades := map[string]float64{"Alma": 74.2, "Rohit": 87.1, "Carl": 67.2}

	// Build a slice which is named names with all the map keys.
	execTime := time.Now()
	var names []string
	for name := range grades {
		names = append(names, name)
	}
	log.Printf("make a slice using []string runtime is: %v\n", time.Since(execTime))

	execTime1 := time.Now()
	names1 := make([]string, 0, len(grades))
	for name1 := range grades {
		names1 = append(names1, name1)
	}
	log.Printf("make a slice using make([]string, 0, len(str) runtime is: %v\n", time.Since(execTime1))

	// Sort the slice alphabetically.
	sort.Strings(names)
	sort.Strings(names1)

	for _, name := range names {
		log.Printf("%s has a grade of %0.1f%%\n", name, grades[name])
	}
}

// The map is similar with PHP key=>value array.
func main() {
	mapDefineFirstWay()
	mapDefineSecWay()
	createEmptyMap()
	checkElementExistAndDeleteElement()
	mapForRange()
	mapSortLoop()
}
