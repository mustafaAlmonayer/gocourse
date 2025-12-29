package basics

import (
	"fmt"
	"maps"
)

func mapsEx() {

	// Declare a map
	// var m map{keyType}valueType

	// Declare map using make
	// mapVar := make(map[keyType]valueType)

	// Using Map Literals
	// mapVar := map[keyType]valueType{
	// 	key1: value1,
	// 	key2: value2,
	// 	...
	// }

	myMap := make(map[string]int)
	fmt.Println(myMap)

	myMap["key1"] = 10
	myMap["code"] = 20
	fmt.Println(myMap)
	fmt.Println("Value for 'key1':", myMap["key1"])

	// Zero Value for non-existing keys
	fmt.Println("Value for 'nonExistingKey':", myMap["nonExistingKey"])
	myMap["key1"] = 30
	fmt.Println("Updated Value for 'key1':", myMap["key1"])

	delete(myMap, "key1")
	fmt.Println(myMap)

	// See if key exists
	{
		value, ok := myMap["code"]
		if ok {
			fmt.Println("Key 'code' exists with value:", value)
		} else {
			fmt.Println("Key 'code' does not exist")
		}
	}
	{
		value, ok := myMap["key1"]
		if ok {
			fmt.Println("Key 'key1' exists with value:", value)
		} else {
			fmt.Println("Key 'key1' does not exist")
		}
	}

	clear(myMap)
	fmt.Println("Map after clearing:", myMap)

	myMap2 := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println("Map initialized using literals:", myMap2)

	if maps.Equal(myMap, myMap2) {
		fmt.Println("myMap and myMap2 are equal")
	} else {
		fmt.Println("myMap and myMap2 are not equal")
	}

	for k, v := range myMap2 {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}

	// Nil
	var myNilMap map[string]int
	fmt.Println("Nil map:", myNilMap)
	// myNilMap["newKey"] = 100 // This will cause a runtime panic

	// Map of Map
	mapOfMap := make(map[string]map[string]string)

	mapOfMap["map1"] = make(map[string]string)
	mapOfMap["map1"]["key1"] = "value1"
	mapOfMap["map1"]["key2"] = "value2"

	fmt.Println("Map of Map:", mapOfMap)

}
