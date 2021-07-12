package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Machine struct {
	Milk int
	Water int
	Sugar int
	Coffee int
	Money int
}

type Coffee struct {
	Coffee_type string
	price int
	need_milk int
	need_water int
	need_coffee int
}

func (u *Machine) BuyCoffee(s Coffee, count int){
	if (u.Milk - s.need_milk < 0) || (u.Water - s.need_water < 0) || (u.Coffee - s.need_coffee < 0 || (u.Sugar - count * 7 < 0)){
		fmt.Println("Not enough resources")
	} else {
		colorReset := "\033[0m"
		colorRed := "\033[31m"
		for i := 5; i > 0; i-- {
				fmt.Printf("\rCoffee is being prepared there is still " + strconv.Itoa(i) + " seconds")
			time.Sleep(time.Second)
		}
		fmt.Println()
		fmt.Println()
		fmt.Println(string(colorRed), "Please take you " + s.Coffee_type + "\n", string(colorReset))
		u.Milk -= s.need_milk
		u.Water -= s.need_water
		u.Coffee -= s.need_coffee
		u.Sugar -= count * 7
		u.Money += s.price

		byte := bytes.Buffer{}
		today := time.Now()
		byte.WriteString(s.Coffee_type + " - " + today.String() + "\n")
		f, err := os.OpenFile("Purchase history", os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}
		/*f, err = os.OpenFile("Purchase history", os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			panic(err)
		}*/
		io.WriteString(f, byte.String())
		f.Close()

		time.Sleep(3*time.Second)
	}
}

func Purch_History(){
	f, err := os.Open("Purchase history")
	if err != nil {
		panic(err)
	}

	c, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(c))

	f.Close()

	fmt.Println("For exit press Enter")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func (u *Machine) OutInfo(){
	fmt.Println("In machine: ")
	fmt.Println("Milk = ", u.Milk)
	fmt.Println("Water = ", u.Water)
	fmt.Println("Sugar = ", u.Sugar)
	fmt.Println("Coffee = ", u.Coffee)
	fmt.Println("Money = ", u.Money, "\n")
}

func (u *Machine) AddMilk(quantity int){
	u.Milk += quantity
}

func (u *Machine) AddWater(quantity int){
	u.Water += quantity
}

func (u *Machine) AddSugar(quantity int){
	u.Sugar += quantity
}

func (u *Machine) AddCoffee(quantity int){
	u.Coffee += quantity
}

func (u *Machine) WithDraw_money(quantity int){
	if (u.Money - quantity) < 0 {
		fmt.Println("Out of money")
	} else {
		u.Money -= quantity
	}
}

func (u Machine)WriteToFile(){
	byte := bytes.Buffer{}
	byte.WriteString(strconv.Itoa(u.Milk) + " ")
	byte.WriteString(strconv.Itoa(u.Water) + " ")
	byte.WriteString(strconv.Itoa(u.Sugar) + " ")
	byte.WriteString(strconv.Itoa(u.Coffee) + " ")
	byte.WriteString(strconv.Itoa(u.Money) + " ")

	f, err := os.Open("Coffee_Machine")
	if err != nil {
		panic(err)
	}
	f, err = os.OpenFile("Coffee_Machine", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	io.WriteString(f, byte.String())
	io.WriteString(f, "              ")
	f.Close()
}

func (u *Machine)LoadData()(Machine){
	f, err := os.Open("Coffee_Machine")
	if err != nil {
		panic(err)
	}

	c, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	arr := strings.Split(string(c)," ")
	milk, err := strconv.Atoi(arr[0])
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	water, err := strconv.Atoi(arr[1])
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	sugar, err := strconv.Atoi(arr[2])
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	coffee, err := strconv.Atoi(arr[3])
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	money, err := strconv.Atoi(arr[4])
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	var m Machine
	m = Machine{milk,water,sugar,coffee,money}
	f.Close()
	return m
}

func MainMenu(){
	fmt.Println("If you User - Enter 1\nIf you Admin - Enter 2\nFor exit - any character")
}

func UserMenu(){
	fmt.Println("	     ░░░░░░░░▄▄▄▀▀▀▀▀▀▀▀▀▄▄▄░░░░░░░░░\n" +
					"	     ░░░░░░▄▀░░▄▄▄▄▄▄▄▄▄▄▄░░▀▄░░░░░░░\n" +
					"	     ░░░░░█░▄███████████████▄░█░░░░░░\n" +
					"	     ░░░░░█░████████████████▀░█░░░░░░\n" +
					"	     ░░░░░▀█▄▄▀▀▀▀▀▀▀▀▀▀▀▀▀▄████▄░░░░\n" +
					"	     ░░▄▄▄▄███████████████████░░██▄░░\n" +
					"	     ▄▀░░░░░██████████████████▄██▀░▀▄\n" +
					"	     █▄░░░░░░▀█████████████▀▀▀▀░░░░▄█\n" +
					"	     ▀█▄░░░░░░░▀█████████▀░░░░░░░░▄█▀\n" +
					"	     ░▀██▄▄░░░░░░░░░░░░░░░░░░░░▄▄██▀░\n" +
					"	     ░░░░▀███▄▄▄▄░░░░░░░░▄▄▄▄███▀░░░░\n" +
					"	     ░░░░░░░░▀▀▀██████████▀▀▀░░░░░░░░\n" +
					"	     ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\n" +
					"	     ░░░███████████████████████████░░\n" +
					"	     ░░░█────█────█───█───█───█───█░░\n" +
					"	     ░░░█─██─█─██─█─███─███─███─███░░\n" +
					"	     ░░░█─████─██─█───█───█───█───█░░\n" +
					"	     ░░░█─██─█─██─█─███─███─███─███░░\n" +
					"	     ░░░█────█────█─███─███───█───█░░\n" +
					"	     ░░░███████████████████████████░░\n" +
					"	     ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░")
	colorCyan := "\033[36m"
	colorReset := "\033[0m"
	colorYellow := "\033[33m"
	fmt.Println(string(colorCyan), "1 - ░░░░░░░░░░░░░ 2 - ░░░░░░░░░░░░░ 3 - ░░░░░░░░░░░░░\n" +
		                           "     ░ Ristretto ░     ░ Kon-panna ░     ░ Espresso  ░\n" +
		       string(colorYellow),"    ░  50 rub   ░     ░  40 rub   ░     ░  60 rub   ░\n" +
		       string(colorCyan),  "    ░░░░░░░░░░░░░     ░░░░░░░░░░░░░     ░░░░░░░░░░░░░\n")

	fmt.Println(string(colorCyan), "4 - ░░░░░░░░░░░░░ 5 - ░░░░░░░░░░░░░ 6 - ░░░░░░░░░░░░░\n" +
		                           "     ░ Americano ░     ░  Makiato  ░     ░ Kapuchino ░\n" +
		       string(colorYellow),"    ░  110 rub  ░     ░  80 rub   ░     ░  70 rub   ░\n" +
		       string(colorCyan),  "    ░░░░░░░░░░░░░     ░░░░░░░░░░░░░     ░░░░░░░░░░░░░\n")

	fmt.Println(string(colorCyan), "                  7 - ░░░░░░░░░░░░░\n" +
		                           "                       ░   Latte   ░\n" +
		       string(colorYellow),"                      ░  90 rub   ░\n" +
		       string(colorCyan),  "                      ░░░░░░░░░░░░░", string(colorReset))

	fmt.Println(string(colorCyan), "Exit to main menu - 8", string(colorReset))

	/*fmt.Println(string(colorCyan), "1 - Ristretto" + string(colorYellow),"50 rub\n" +
		string(colorCyan), "2 - Kon-panna" + string(colorYellow),"40 rub\n"+
		string(colorCyan), "3 - Espresso" + string(colorYellow),"60 rub\n"+
		string(colorCyan), "4 - Amerikano" + string(colorYellow),"110 rub\n"+
		string(colorCyan), "5 - Makiato" + string(colorYellow),"80 rub\n"+
		string(colorCyan), "6 - Kapuchino" + string(colorYellow),"70 rub\n"+
		string(colorCyan), "7 - Latte" + string(colorYellow),"90 rub", string(colorReset))*/
}

func AdminMenu(){
	colorPurple := "\033[35m"
	colorReset := "\033[0m"
	fmt.Println(string(colorPurple),"Chose action", string(colorReset))
	fmt.Println(string(colorPurple),"1 for added milk", string(colorReset))
	fmt.Println(string(colorPurple),"2 for added water", string(colorReset))
	fmt.Println(string(colorPurple),"3 for added sugar", string(colorReset))
	fmt.Println(string(colorPurple),"4 for added coffee", string(colorReset))
	fmt.Println(string(colorPurple),"5 for withdraw money", string(colorReset))
	fmt.Println(string(colorPurple),"6 for save data", string(colorReset))
	fmt.Println(string(colorPurple),"7 for load data", string(colorReset))
	fmt.Println(string(colorPurple),"8 for view purchase history", string(colorReset))
	fmt.Println(string(colorPurple),"9 for delete purchase history", string(colorReset))
	fmt.Println(string(colorPurple),"For out enter any key", string(colorReset))

}

func ClearPurchHistory(){
	f, err := os.Open("Purchase history")
	if err != nil {
		panic(err)
	}
	f, err = os.OpenFile("Purchase history", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	io.WriteString(f, " ")
	f.Close()
	fmt.Println("Deleting complete")
}

func main() {
	var mascoffee [7]Coffee
	mascoffee[0] = Coffee{"Ristretto", 50, 10, 20, 15}
	mascoffee[1] = Coffee{"Kon-panna", 40, 15, 10, 10}
	mascoffee[2] = Coffee{"Espresso", 60, 18, 10, 12}
	mascoffee[3] = Coffee{"Amerikano", 110, 5, 200, 25}
	mascoffee[4] = Coffee{"Makiato", 80, 80, 10, 20}
	mascoffee[5] = Coffee{"Kapuchino", 70, 60, 0, 30}
	mascoffee[6] = Coffee{"Latte", 90, 70, 15, 5}
	MyMachine := Machine{1000, 1000, 500, 400, 0}
	for j := true;j;{
		clear := exec.Command("cmd", "/c", "cls")
		clear.Stdout = os.Stdout
		clear.Run()
		var flag int
		MainMenu()
		fmt.Scanf("%d\n", &flag)
	switch flag {
		case 1:
			fmt.Println("Chose coffee")
			UserMenu()
			var flag3 int
			fmt.Scanf("%d\n", &flag3)
			count := 0
			colorBlue := "\033[34m"
			colorReset := "\033[0m"
			switch flag3 {
				case 1:
					fmt.Println("Enter count of spoons of sugar")
					fmt.Scanf("%d\n", &count)
					fmt.Println(string(colorBlue),"\nПорция буквально на один глоток (до 20 мл). Вкусовые рецепторы будут приятно удивлены силой вкуса этого концентрированного напитка.\nПричем количество кофеина в расчете на 100 мл здесь нижу за счет меньшего времени приготовления.\n", string(colorReset))
					MyMachine.BuyCoffee(mascoffee[0], count)

				case 2:
					fmt.Println("Enter count of spoons of sugar")
					fmt.Scanf("%d\n", &count)
					fmt.Println(string(colorBlue),"\nПод взбитыми сливками, посыпанными коричным порошком, прячется эспрессо.\nМягкий напиток служит завершением обеда. Подается с ложечкой, ведь сначала надо съесть сливки и лишь затем пить.\n", string(colorReset))
					MyMachine.BuyCoffee(mascoffee[1], count)

				case 3:
					fmt.Println("Enter count of spoons of sugar")
					fmt.Scanf("%d\n", &count)
					fmt.Println(string(colorBlue),"\nОбразно эсперссо называют кофейным соком. Небольшая (30г) порция исключительно ароматного напитка с насыщенным вкусом не нуждается в дополнениях.\nГотовится из кофе тонкого помола в кофемашинах или рожковых кофеварках под давлением, поэтомуимеет характерную плотную пенку (крема).\nДля подачи используются специальные чашки (демитассе) - небольшие, но толстостенные.\n", string(colorReset))
					MyMachine.BuyCoffee(mascoffee[2], count)

				case 4:
					fmt.Println("Enter count of spoons of sugar")
					fmt.Scanf("%d\n", &count)
					fmt.Println(string(colorBlue),"\nРазница между эспрессо и американо в размере порции и концентрации кофе. Это 200-400 мл разбавленного кипятком эспрессо.\nВкус напитка ненасыщенный, поэтому его часто дополняют сахаром, молоком. Подают с десертом.\n", string(colorReset))
					MyMachine.BuyCoffee(mascoffee[3], count)

				case 5:
					fmt.Println("Enter count of spoons of sugar")
					fmt.Scanf("%d\n", &count)
					fmt.Println(string(colorBlue),"\nВ небольшое количество вспененного молока вливают эспрессо, и на поверхности образуются кофейные пятна.\nЗа внешний вид напиток получил название:macchiato в переводе с итальянского означает пятнистый.\nВ прозрачном стакане хорошо видно три слоя с плавными переходами: молоко, кофе и пенку.\n", string(colorReset))
					MyMachine.BuyCoffee(mascoffee[4], count)

				case 6:
					fmt.Println("Enter count of spoons of sugar")
					fmt.Scanf("%d\n", &count)
					fmt.Println(string(colorBlue),"\nСамый популярный вид кофейно-сливочного коктейля. Молочная пенка покрывает эспрессо, причем молока в 2 раза больше, чем кофе.\nПлотная пенка позволяет создавать рисунки в технике латте-арт.\n", string(colorReset))
					MyMachine.BuyCoffee(mascoffee[5], count)

				case 7:
					fmt.Println("Enter count of spoons of sugar")
					fmt.Scanf("%d\n", &count)
					fmt.Println(string(colorBlue),"\nОтличия кофе латте от капучино в пропорциях: здесь молоко преобладает.\nСледовательно, содержание кофеина невысоко, и этот вид кофейного коктейля годится для вечера. Часто дополняется сладкими ароматизированными сиропами.\n", string(colorReset))
					MyMachine.BuyCoffee(mascoffee[6], count)

				case 8:
					break

			default:
				fmt.Println("incorrect input")
			}

		case 2:
			var password string
			fmt.Println("Enter password")
			fmt.Scanf("%s\n", &password)
			if password == "123" {
				for i := true;i; {
					clear := exec.Command("cmd", "/c", "cls")
					clear.Stdout = os.Stdout
					clear.Run()
					MyMachine.OutInfo()
					AdminMenu()
					var flag1 int
					var quantity int
					fmt.Scanf("%d\n", &flag1)
					switch flag1 {
						case 1:
							fmt.Println("Enter quantity of milk")
							fmt.Scanf("%d\n", &quantity)
							MyMachine.AddMilk(quantity)

						case 2:
							fmt.Println("Enter quantity of water")
							fmt.Scanf("%d\n", &quantity)
							MyMachine.AddWater(quantity)

						case 3:
							fmt.Println("Enter quantity of sugar")
							fmt.Scanf("%d\n", &quantity)
							MyMachine.AddSugar(quantity)

						case 4:
							fmt.Println("Enter quantity of coffee")
							fmt.Scanf("%d\n", &quantity)
							MyMachine.AddCoffee(quantity)

						case 5:
							fmt.Println("Enter how much money do you want to withdraw")
							fmt.Scanf("%d\n", &quantity)
							MyMachine.WithDraw_money(quantity)

						case 6:
							MyMachine.WriteToFile()

						case 7:
							MyMachine = MyMachine.LoadData()
							MyMachine.OutInfo()

						case 8:
							Purch_History()

						case 9:
							ClearPurchHistory()

						default:
							i = false
						}
					}
			} else {
				fmt.Println("Incorrect password")
			}
	default:
		MyMachine.WriteToFile()
		j = false
		}
	}
}
