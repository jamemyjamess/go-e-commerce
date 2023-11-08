package main

import (
	"fmt"

	"github.com/jamemyjamess/go-e-commerce/config"
	"github.com/jamemyjamess/go-e-commerce/pkg/databases/postgres"
	"github.com/jamemyjamess/go-e-commerce/servers"
)

// func main() {
// 	// แปลง string เป็น int
// 	stringInt := "13"
// 	numberCoverted, _ := strconv.Atoi(stringInt)
// 	fmt.Printf("numberCoverted type: %T value: %d\n", numberCoverted, numberCoverted)
// 	//numberCoverted type: int value: 13

// 	// แปลง string เป็น float
// 	stringFloat := "13.3"
// 	numberFloatCoverted, _ := strconv.ParseFloat(stringFloat, 64)
// 	fmt.Printf("numberFloatCoverted type: %T value: %.2f\n", numberFloatCoverted, numberFloatCoverted)
// 	// numberFloatCoverted type: float64 value: 13.30

// 	// แปลง int เป็น string
// 	numberInt := 13
// 	stringIntCoverted := strconv.Itoa(numberInt)
// 	fmt.Printf("stringIntCoverted type: %T value: %s\n", stringIntCoverted, stringIntCoverted)
// 	// stringIntCoverted type: string value: 13

// 	// แปลง float หรือ int เป็น string
// 	numberFloat := 13.3
// 	numberIntString := fmt.Sprintf("%.2f", numberFloat)
// 	numberFloatString := fmt.Sprintf("%d", numberInt)
// 	fmt.Printf("numberIntString type: %T value: %s\n", numberIntString, numberIntString)
// 	// numberIntString type: string value: 13.30
// 	fmt.Printf("numberFloatString type: %T value: %s\n", numberFloatString, numberFloatString)
// 	// numberFloatString type: string value: 13

// 	// แปลง float เป็น int
// 	floatFromInt := float64(numberInt)
// 	fmt.Printf("numberFloatString type: %T value: %.2f\n", floatFromInt, floatFromInt)
// 	// numberFloatString type: float64 value: 13.00

// 	// แปลง int เป็น float
// 	intFromFloat := int(numberFloat)
// 	fmt.Printf("numberFloatString type: %T value: %d\n", intFromFloat, intFromFloat)
// 	// numberFloatString type: int value: 13
// }

type Customer struct {
	FirstName string
	LastName  string
}

// func main() {
// 	user := [2]string{"jame", "cat"}
// 	fmt.Println("user[0]: ", user[0]) // jame
// 	fmt.Println("user[1]: ", user[1]) // cat

// 	customer := Customer{
// 		FirstName: "Suttipong",
// 		LastName:  "Saksittikarn",
// 	}

// 	fmt.Printf("customer: %#v\n", customer)
// 	// customer: main.Customer{FirstName:"Suttipong", LastName:"Saksittikarn"}
// 	fmt.Printf("firstname: %v\n", customer.FirstName)
// 	// firstname: Suttipong
// 	fmt.Printf("lastname: %v\n", customer.LastName)
// 	// lastname: Saksittikarn

// 	sum := 0
// 	for i := 0; i < 10; i++ {
// 		sum = sum + i
// 	}
// 	fmt.Println("sum: ", sum) // sum: 45

// 	strings := []string{"hello", "world"}
// 	for i, s := range strings {
// 		fmt.Println("index:", i, "string:", s)
// 	}
// 	// index:  0  string:  hello
// 	// index:  1  string:  world

// 	arr := [6]int{1, 2, 3, 4, 5, 6}

// 	fmt.Println("arr[:3] : ", arr[:3])

// 	numbers := []int{1, 2, 3, 4, 5, 6}
// 	fmt.Printf("numbers: %#v", numbers)
// }

type customer struct {
	firstname string // use in local package only
	lastname  string
	code      int
	Address   string // can use in external package
}

func print(x int) {
	fmt.Println("result: ", x)
}

func add(x, y int) int {
	return x + y
}

// func main() {
// 	print(3)
// 	print(add(3, 4))
// }

func main() {
	// config
	cf := config.InitConfig()

	// connect db
	postgresDB := postgres.ConnectPostgresDB(cf.Db())

	server := servers.NewServer(&cf, postgresDB)
	server.Start()
}

// output:
// 	hello
//	world
//	ilove
//	you
