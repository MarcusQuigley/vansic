package models

type User struct {
	Id        int
	Firstname string
	Lastname  string
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.Id = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

// func chap4() {
// 	var i int = 42
// 	i := 42
// 	fmt.Printf("%d\n", i)
// 	var firstname *string = new(string) //pointer operator
// 	*firstname = "Marco" //deref operator
// 	fmt.Println(*firstname, firstname)
// 	lastname := "Quigs"
// 	ptrl := &lastname
// 	fmt.Println(ptrl, *ptrl)
// 	lastname = "jell"
// 	fmt.Println(ptrl, *ptrl)
// }
// func chap5() {
// 	var arr [3]int
// 	arr[0] = 1
// 	fmt.Printf("%v\n", arr)
// 	arr2 := [3]int{4, 5, 6}
// 	fmt.Printf("%v\n", arr2)
// 	myslice := arr2[1:2]
// 	fmt.Printf("%v\n", myslice)
// 	myslice2 := []int{1, 2, 3}
// 	fmt.Printf("%v\n", myslice2)
// 	myslice2 = append(myslice2, 22)
// 	fmt.Printf("%v\n", myslice2)
// 	m := map[int]string{1: "foo", 2: "boo", 3: "doo"}
// 	fmt.Printf("%v\n", m)
// 	m[3] = "loo"
// 	fmt.Printf("%v\n", m)
// 	delete(m, 2)
// 	fmt.Printf("%v\n", m)
// 	type user struct {
// 		id   int
// 		name string
// 		age  int
// 	}
// 	// var u user = (3,"er",2)
// 	//  user uu := (3,"er",2)
// 	var u user
// 	u.age = 3
// 	u.id = 33
// 	u.name = "def"
// 	fmt.Printf("%v\n", u)
// 	u2 := user{id: 3, age: 44,
// 		name: "bawbag",
// 	}
// 	fmt.Printf("%v\n", u2)
// }
