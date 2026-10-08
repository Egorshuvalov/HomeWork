package main

import "fmt"

func main() {
	// Название блюда
	dishName := "Борщ"
	// Количество порций
	servings := 4
	// Время приготовления в минутах
	cookingTime := 45
	// Калорийность на порцию
	caloriesPerServing := 320.50

	// TODO: Выведи название блюда через fmt.Println(fmt.Sprintf("Блюдо: %s"))
	fmt.Println(fmt.Sprintf("Блюдо: %s", dishName))
	// TODO: Выведи количество порций через fmt.Println(fmt.Sprintf("Порции: %d", ))
	fmt.Println(fmt.Sprintf("Порции: %d", servings))
	// TODO: Выведи время приготовления через fmt.Println(fmt.Sprintf("Время приготовления: %d минут"))
	fmt.Println(fmt.Sprintf("Время приготовления: %d минут", cookingTime))
	// TODO: Выведи калорийность через fmt.Println(fmt.Sprintf("Калорийность на порцию: %.2f ккал"))
	fmt.Println(fmt.Sprintf("Калорийность на порцию: %.2f ккал", caloriesPerServing))
}
