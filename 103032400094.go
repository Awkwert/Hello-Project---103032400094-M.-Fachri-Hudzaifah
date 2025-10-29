package main

import (
	"fmt"
	"strings"
)

// In-memory "database" for demonstration purposes
var dataStore = make(map[string]User) // Map to store user data
var loggedInUser string
var borrowedBooksGlobal []string // Tambahkan variabel global untuk menyimpan daftar buku yang dipinjam

// Struktur untuk menyimpan data pengguna
type User struct {
	Name     string
	NIM      string
	Password string
}

// Daftar buku yang tersedia di perpustakaan
var books = []string{
	"Pemrograman Go",
	"Dasar-dasar Algoritma",
	"Database Modern",
	"Jaringan Komputer",
	"Kecerdasan Buatan",
}

func hashPassword(password string) string {
	// Sederhana hash password menggunakan panjang string (hanya untuk ilustrasi)
	return fmt.Sprintf("%x", len(password))
}

func registerUser(name, nim, password string) string {
	// Validate NIM: Must be numeric and start with "1030324"
	if !isValidNIM(nim) {
		return "Invalid NIM. It must be numeric and start with 1030324."
	}

	// Register a new user
	if _, exists := dataStore[name]; exists {
		return "Name already exists."
	}
	dataStore[name] = User{Name: name, NIM: nim, Password: hashPassword(password)}
	return "User registered successfully."
}

func loginUser(name, password string) string {
	// Login a user by checking the name and password
	if user, exists := dataStore[name]; !exists {
		return "Name not found."
	} else if user.Password != hashPassword(password) {
		return "Incorrect password."
	}
	loggedInUser = name
	return "Login successful."
}

func borrowBook() {
	if loggedInUser == "" {
		fmt.Println("Please login first.")
		return
	}

	var borrowedBooks []string
	for {
		// Tampilkan daftar buku yang tersedia
		if len(books) == 0 {
			fmt.Println("Tidak ada buku yang tersedia untuk dipinjam.")
			break
		}
		fmt.Println("\nDaftar Buku Tersedia:")
		for i, b := range books {
			fmt.Printf("%d. %s\n", i+1, b)
		}

		// Input nomor buku
		fmt.Print("Masukkan nomor buku yang ingin dipinjam: ")
		var choice int
		fmt.Scanln(&choice)

		if choice > 0 && choice <= len(books) {
			selectedBook := books[choice-1]
			fmt.Printf("Anda ingin meminjam buku '%s'? (y/n): ", selectedBook)
			var confirm string
			fmt.Scanln(&confirm)

			if strings.ToLower(confirm) == "y" {
				borrowedBooks = append(borrowedBooks, selectedBook)
				books = append(books[:choice-1], books[choice:]...) // Hapus buku yang dipinjam
				fmt.Printf("Buku '%s' berhasil dipinjam!\n", selectedBook)
				fmt.Print("Apakah Anda ingin meminjam buku lagi? (y/n): ")
				fmt.Scanln(&confirm)

				if strings.ToLower(confirm) != "y" {
					break
				}
			} else {
				fmt.Println("Kembali ke daftar buku.")
				continue
			}
		} else {
			fmt.Println("Nomor buku tidak valid.")
		}
	}
	// Simpan daftar buku yang dipinjam ke variabel global
	borrowedBooksGlobal = borrowedBooks
}

func listBooks() {
	if loggedInUser == "" {
		fmt.Println("Please login first.")
		return
	}
	fmt.Println("\nDaftar Buku Tersedia:")
	for i, b := range books {
		fmt.Printf("%d. %s\n", i+1, b)
	}
}

func logoutUser() {
	if loggedInUser == "" {
		fmt.Println("You are not logged in.")
		return
	}

	if len(borrowedBooksGlobal) > 0 {
		fmt.Printf("Anda memiliki buku yang dipinjam: %s. Apakah Anda yakin ingin logout? (buku yang Anda pinjam akan dibatalkan) (y/n): ", strings.Join(borrowedBooksGlobal, ", "))
		var confirm string
		fmt.Scanln(&confirm)
		if strings.ToLower(confirm) != "y" {
			fmt.Println("Logout dibatalkan.")
			return
		}
	}

	loggedInUser = ""
	borrowedBooksGlobal = nil
	fmt.Println("Logged out successfully.")
}

func main() {
	fmt.Println("Selamat datang di perpustakaan OPENLIB!")

	for {
		if loggedInUser == "" {
			fmt.Println("\nOptions:\n1. Register\n2. Login\n3. Exit")
			fmt.Print("Choose an option: ")
			var choice string
			fmt.Scanln(&choice)

			switch choice {
			case "1":
				fmt.Print("Enter name: ")
				var name string
				fmt.Scanln(&name)

				fmt.Print("Enter NIM: ")
				var nim string
				fmt.Scanln(&nim)

				fmt.Print("Enter password: ")
				var password string
				fmt.Scanln(&password)

				fmt.Println(registerUser(name, nim, password))

			case "2":
				fmt.Print("Enter name: ")
				var name string
				fmt.Scanln(&name)

				fmt.Print("Enter password: ")
				var password string
				fmt.Scanln(&password)

				fmt.Println(loginUser(name, password))

			case "3":
				fmt.Println("Goodbye!")
				return

			default:
				fmt.Println("Invalid option. Please try again.")
			}
		} else {
			fmt.Println("\nOptions:\n1. Borrow a Book\n2. List Available Books\n3. Logout\n4. Exit")
			fmt.Print("Choose an option: ")
			var choice string
			fmt.Scanln(&choice)

			switch choice {
			case "1":
				borrowBook()

			case "2":
				listBooks()

			case "3":
				logoutUser()

			case "4":
				fmt.Printf("Nama: %s\nNIM: %s\nNama Buku: %s\nSelamat membaca!\n", dataStore[loggedInUser].Name, dataStore[loggedInUser].NIM, strings.Join(borrowedBooksGlobal, ", "))
				return

			default:
				fmt.Println("Invalid option. Please try again.")
			}
		}
	}
}

// Fungsi untuk memeriksa ketersediaan buku
func isBookAvailable(book string) bool {
	for _, b := range books {
		if strings.EqualFold(b, book) {
			return true
		}
	}
	return false
}

// Fungsi untuk memvalidasi NIM
func isValidNIM(nim string) bool {
	return len(nim) > 7 && strings.HasPrefix(nim, "1030324")
}
