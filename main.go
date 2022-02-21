package main

import (
	"fmt"
	"os"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-1-bakyazi/kitaplik"
)

func printUsage() {
	fmt.Println("kitaplik uygulamasında kullanabileceğiniz komutlar :")
	fmt.Println("* search => arama işlemi için")
	fmt.Println("    Örnek kullanim: $ kitaplik search moby dick")
	fmt.Println("* list => listeleme işlemi için")
	fmt.Println("    Örnek kullanim: $ kitaplik list")

}

func printBooks(books []string) {
	for i, book := range books {
		fmt.Printf("%d. %s\n", i+1, book)
	}
}

func main() {
	args := os.Args

	if len(args) == 1 {
		printUsage()
		return
	}

	command := args[1]

	switch command {
	case "search":
		if len(args) < 3 {
			printUsage()
			return
		}
		books := kitaplik.Search(args[2:])
		if len(books) == 0 {
			fmt.Printf("Arama kriterlerine göre bir kitap bulunamadı.\n")
			return
		}
		printBooks(books)
	case "list":
		books := kitaplik.List()
		if len(books) == 0 {
			fmt.Printf("Kitaplıkta herhangi bir kitap yer almıyor.\n")
			return
		}
		printBooks(books)
	default:
		printUsage()
	}
}
