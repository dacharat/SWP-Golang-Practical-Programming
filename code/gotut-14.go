package main

import "fmt"

func main() {
	grades := make(map[string]float32)

	grades["Timmy"] = 42
	grades["Jess"] = 92
	grades["Sam"] = 71

	fmt.Println(grades)

	tims_grade := grades["Timmy"]
	fmt.Println(tims_grade)

	delete(grades, "Timmy")
	fmt.Println(grades)

	for k, v := range grades {
		fmt.Println(k, ":", v)
	}

}
