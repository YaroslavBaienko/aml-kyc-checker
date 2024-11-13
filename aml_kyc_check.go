package main

import (
	"fmt"
	"strings"
)

// Структура для хранения информации о клиенте
type Customer struct {
	Name          string
	Country       string
	Occupation    string
	SourceOfFunds string
	IsPEP         bool
}

// Черный список стран с высоким уровнем риска
var highRiskCountries = []string{"Afghanistan", "Iran", "North Korea", "Syria"}

// Функция для проверки, является ли клиент из страны с высоким уровнем риска
func isHighRiskCountry(country string) bool {
	for _, riskCountry := range highRiskCountries {
		if strings.EqualFold(riskCountry, country) {
			return true
		}
	}
	return false
}

// Функция для проверки источника средств клиента
func validateSourceOfFunds(source string) bool {
	allowedSources := []string{"salary", "investment", "inheritance"}
	for _, allowed := range allowedSources {
		if strings.EqualFold(allowed, source) {
			return true
		}
	}
	return false
}

// Функция для выполнения проверки клиента
func performKYCCustomerCheck(customer Customer) {
	fmt.Printf("Проверка клиента: %s\n", customer.Name)

	if isHighRiskCountry(customer.Country) {
		fmt.Println("Предупреждение: Клиент из страны с высоким уровнем риска!")
	} else {
		fmt.Println("Страна клиента - в допустимом списке.")
	}

	if customer.IsPEP {
		fmt.Println("Предупреждение: Клиент является политически значимым лицом (PEP)!")
	} else {
		fmt.Println("Клиент не является политически значимым лицом.")
	}

	if validateSourceOfFunds(customer.SourceOfFunds) {
		fmt.Println("Источник средств клиента проверен и допустим.")
	} else {
		fmt.Println("Источник средств клиента вызывает подозрения!")
	}

	fmt.Println("Проверка завершена.\n")
}

// Точка входа
func main() {
	// Пример клиентов для проверки
	customers := []Customer{
		{"John Doe", "United Kingdom", "Engineer", "salary", false},
		{"Jane Smith", "Iran", "Business Owner", "inheritance", false},
		{"Ahmed Khan", "United States", "Politician", "investment", true},
	}

	for _, customer := range customers {
		performKYCCustomerCheck(customer)
	}
}
