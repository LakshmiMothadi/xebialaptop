package main

import "fmt"

type Datastorer interface {
	Put(key string, value interface{}) error
	Get(key string) (interface{}, bool, error)
	Remove(key string) error
}

type InMemory struct {
	id      int
	name    string
	area    string
	pincode int
	state   string
}

func (data *InMemory) Put(key string, value interface{}) error {
	data.id = 1
	data.name = "mothadi"
	data.area = "kadapa"
	data.pincode = 516003
	data.state = "ap"
	if key == "student" {
		fmt.Println(*data)
	} else {
		fmt.Println("Not working")
	}

	return nil
}

func (data *InMemory) Get(key string) (interface{}, bool, error) {

	getDeatils := data.Put("student", " ")
	fmt.Println(getDeatils)

	return nil, false, nil
}

func (data *InMemory) Remove(key string) error {
	data.Get("get")
	//addData :=

	return nil
}

/*func user(b Datastorer) {
	fmt.Println(b.Get)
}
*/

/* func main() {
	//var details InMemory
	//details.Get("get")

	var details2 InMemory
	details2.Remove("remove")

	/* var detail InMemory
	detail.Put("put", " ")
	user(&detail)
}
*/
