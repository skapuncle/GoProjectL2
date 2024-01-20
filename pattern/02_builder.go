package main

/*Паттерн "Строитель" (Builder) - это порождающий паттерн проектирования,
который позволяет создавать сложные объекты шаг за шагом. Он отделяет процесс
конструирования сложного объекта от его представления, так что в результате одного
и того же процесса конструирования могут получаться разные представления.

Преимущества паттерна "Строитель" включают:

Разделение процесса создания объекта и его представления.
Возможность создания сложных объектов шаг за шагом.
Возможность создания разных представлений одного и того же объекта.
Минусы паттерна "Строитель" включают:

Увеличение числа классов в коде.
Сложность в поддержке кода из-за добавления новых типов строителей.
*/

import "fmt"

// Car представляет автомобиль, который мы хотим построить.
type Car struct {
	Wheels int    // Количество колес.
	Seats  int    // Количество мест.
	Color  string // Цвет автомобиля.
}

// CarBuilder представляет интерфейс для построения автомобиля.
type CarBuilder interface {
	SetWheels(int) CarBuilder   // Устанавливает количество колес.
	SetSeats(int) CarBuilder    // Устанавливает количество мест.
	SetColor(string) CarBuilder // Устанавливает цвет автомобиля.
	Build() *Car                // Собирает автомобиль.
}

// CarBuilderImpl реализует интерфейс CarBuilder.
type CarBuilderImpl struct {
	car Car // Атрибут для хранения автомобиля.
}

// SetWheels устанавливает количество колес.
func (cb *CarBuilderImpl) SetWheels(wheels int) CarBuilder {
	cb.car.Wheels = wheels
	return cb
}

// SetSeats устанавливает количество мест.
func (cb *CarBuilderImpl) SetSeats(seats int) CarBuilder {
	cb.car.Seats = seats
	return cb
}

// SetColor устанавливает цвет автомобиля.
func (cb *CarBuilderImpl) SetColor(color string) CarBuilder {
	cb.car.Color = color
	return cb
}

// Build возвращает собранный автомобиль.
func (cb *CarBuilderImpl) Build() *Car {
	return &cb.car
}

func main() {
	// Создаем новый строитель автомобилей и устанавливаем параметры.
	car := NewCarBuilder().
		SetWheels(4).
		SetSeats(5).
		SetColor("Red").
		Build()

	// Выводим информацию о созданном автомобиле.
	fmt.Printf("%+v\n", car)
}

// NewCarBuilder создает новый строитель автомобилей.
func NewCarBuilder() CarBuilder {
	return &CarBuilderImpl{}
}
