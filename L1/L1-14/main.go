package main

import (
	"fmt"
	"reflect"
)

func determineTypeWithSwitch(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Printf("Type: int, Value: %d\n", v)
	case string:
		fmt.Printf("Type: string, Value: %s\n", v)
	case bool:
		fmt.Printf("Type: bool, Value: %t\n", v)
	case chan int:
		fmt.Printf("Type: chan int, Value: %v\n", v)
	case chan string:
		fmt.Printf("Type: chan string, Value: %v\n", v)
	case chan bool:
		fmt.Printf("Type: chan bool, Value: %v\n", v)
	default:
		fmt.Printf("Type: unknown (%T), Value: %v\n", v, v)
	}
}

func determineTypeWithReflect(value interface{}) {
	valueType := reflect.TypeOf(value)
	valueKind := valueType.Kind()
	
	fmt.Printf("Type: %v, Kind: %v, Value: %v\n", valueType, valueKind, value)
	
	if valueKind == reflect.Chan {
		fmt.Printf("Channel element type: %v\n", valueType.Elem())
		fmt.Printf("Channel direction: %v\n", valueType.ChanDir())
	}
}

func isChannelType(value interface{}) bool {
	return reflect.TypeOf(value).Kind() == reflect.Chan
}

func getDetailedTypeInfo(value interface{}) string {
	valueType := reflect.TypeOf(value)
	
	switch valueType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("integer type: %v", valueType)
	case reflect.String:
		return fmt.Sprintf("string type: %v", valueType)
	case reflect.Bool:
		return fmt.Sprintf("boolean type: %v", valueType)
	case reflect.Chan:
		return fmt.Sprintf("channel type: %v, element: %v", valueType, valueType.Elem())
	default:
		return fmt.Sprintf("other type: %v", valueType)
	}
}

func main() {
	testInteger := 42
	testString := "hello"
	testBoolean := true
	testChannelInt := make(chan int)
	testChannelString := make(chan string)
	testChannelBool := make(chan bool, 10)
	testFloat := 3.14
	testSlice := []int{1, 2, 3}
	
	fmt.Println("=== Type detection with switch ===")
	determineTypeWithSwitch(testInteger)
	determineTypeWithSwitch(testString)
	determineTypeWithSwitch(testBoolean)
	determineTypeWithSwitch(testChannelInt)
	determineTypeWithSwitch(testChannelString)
	determineTypeWithSwitch(testChannelBool)
	determineTypeWithSwitch(testFloat)
	determineTypeWithSwitch(testSlice)
	fmt.Println()
	
	fmt.Println("=== Type detection with reflection ===")
	determineTypeWithReflect(testInteger)
	determineTypeWithReflect(testString)
	determineTypeWithReflect(testBoolean)
	determineTypeWithReflect(testChannelInt)
	determineTypeWithReflect(testChannelString)
	determineTypeWithReflect(testChannelBool)
	fmt.Println()
	
	fmt.Println("=== Channel detection ===")
	fmt.Printf("Is %v a channel? %v\n", testInteger, isChannelType(testInteger))
	fmt.Printf("Is %v a channel? %v\n", testChannelInt, isChannelType(testChannelInt))
	fmt.Println()
	
	fmt.Println("=== Detailed type info ===")
	fmt.Println(getDetailedTypeInfo(testInteger))
	fmt.Println(getDetailedTypeInfo(testString))
	fmt.Println(getDetailedTypeInfo(testBoolean))
	fmt.Println(getDetailedTypeInfo(testChannelInt))
	fmt.Println(getDetailedTypeInfo(testFloat))
	fmt.Println()
	
	fmt.Println("=== Type assertion examples ===")
	var genericValue interface{} = "I am a string"
	
	if stringValue, ok := genericValue.(string); ok {
		fmt.Printf("Successfully asserted as string: %s\n", stringValue)
	}
	
	if _, ok := genericValue.(int); !ok {
		fmt.Println("Failed to assert as int (expected)")
	}
	
	close(testChannelInt)
	close(testChannelString)
	close(testChannelBool)
}
