package main

import "fmt"

func main() {
	var room = "пещера"

	if room == "пещера" {
		fmt.Println("Вы находитесь в тускло освещенной пещере.")
	} else if room == "вход" {
		fmt.Println("Здесь есть вход в пещеру и путь на восток.")
	} else if room == "гора" {
		fmt.Println("Здесь крутой утес. Тропа ведет к подножью горы.")
	} else {
		fmt.Println("Здесь ничего нет.")
	}

}
