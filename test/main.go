package main

func Suma(x, y int) int {
	return x + y
}

func GetMax(x, y int) int {
	/*
		funcion sin ser testeada llamada
		code covarage
		go test -cover # Muestra cuanto porcentaje del codigo no esta testeado
		go test -coverprofile=salida.out # Genera un archivo
		go tool cover -func=salida.out # Muestra cuales funciones no estan en testing
		go tool cover -html=salida.out # Salida en navegador

	*/
	if x > y {
		return x
	}
	return y
}
