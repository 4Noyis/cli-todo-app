package main

func main() {
	todos := Todos{}

	todos.Add("Buy sim card")
	todos.Add("Send email to prof.")

	todos.Toggle(0)

	todos.Print()
}
