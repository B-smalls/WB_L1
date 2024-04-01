package main

import "fmt"

// AuthService - интерфейс, используемый в нашей системе.
type AuthService interface {
	ParseJSON()
}

// ThirdPartyAuthService - некоторая реализация (ParseXML), написанная другими разработчиками и не может быть изменена.
type ThirdPartyAuthService struct {
}

// ParseXml - метод, который нам нужно адаптировать к нашему интерфейсу AuthService.
func (s *ThirdPartyAuthService) ParseXml() {
	fmt.Println("Parsing XML auth data")
}

// AdapterAuthService - адаптер, который адаптирует ThirdPartyAuthService к нашему интерфейсу AuthService.
type AdapterAuthService struct {
	authService ThirdPartyAuthService
}

// ParseJSON - метод, который преобразует данные JSON в XML и вызывает метод ParseXml вашего ThirdPartyAuthService.
func (s AdapterAuthService) ParseJSON() {
	fmt.Println("Converting JSON data to XML...")
	s.authService.ParseXml()
}

func main() {
	// Создание экземпляра адаптера и вызов метода ParseJSON для демонстрации его работы.
	a := AdapterAuthService{}
	a.ParseJSON()
}
