package main

func main() {

}

func squareIsWhite(coordinates string) bool {
	return (coordinates[0]%2 == 1 && coordinates[1]%2 == 0) || (coordinates[0]%2 == 0 && coordinates[1]%2 == 1)
}
