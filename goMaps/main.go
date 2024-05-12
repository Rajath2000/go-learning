package main

import "fmt"

func main() {

	lang := make(map[string]string)

	lang["js"] = "Java script"
	lang["c"] = "c"
	lang["py"] = "python"

	fmt.Println(lang)
	fmt.Println(lang["py"])

	delete(lang, "py")

	fmt.Println(lang)

	//loops

	for _, value := range lang {
		fmt.Println(value)
	}

}
