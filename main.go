package main

func main() {
	a := App{}
	a.Initialize("DB_USERNAME", "DB_PASSWORD", "rest_api_demo")
	a.Run(":8080")
}
