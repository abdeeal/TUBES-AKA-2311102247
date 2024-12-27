package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Komputer represents a computer in the warnet
type Komputer struct {
	ID         int
	Tipe       string // "Biasa", "VIP", atau "Battle Arena"
	Status     string // "Tersedia" atau "Digunakan"
	Pengguna   string
	WaktuMulai time.Time
	Durasi     time.Duration
}

// Warnet represents the internet cafe
type Warnet struct {
	Komputer []Komputer
}

// TambahKomputer menambahkan komputer baru ke warnet
func (w *Warnet) TambahKomputer(id int, tipe string) {
	w.Komputer = append(w.Komputer, Komputer{
		ID:     id,
		Tipe:   tipe,
		Status: "Tersedia",
	})
}

// TampilkanSemuaKomputerIteratif menampilkan semua komputer (tersedia dan digunakan) secara iteratif
func (w *Warnet) TampilkanSemuaKomputerIteratif() {
	garis()
	fmt.Println("| Biasa                     | Battle Arena                   | VIP                     |")
	garis()

	biasa, battleArena, vip := kelompokkanKomputerIteratif(w.Komputer)
	maxRows := max(len(biasa), len(battleArena), len(vip))

	for i := 0; i < maxRows; i++ {
		fmt.Printf("| %-25s | %-30s | %-23s |\n",
			formatCell(biasa, i),
			formatCell(battleArena, i),
			formatCell(vip, i))
	}

	garis()
}

// TampilkanKomputerTersediaIteratif menampilkan komputer yang tersedia secara iteratif
func (w *Warnet) TampilkanKomputerTersediaIteratif() {
	garis()
	fmt.Println("| Biasa                     | Battle Arena                   | VIP                     |")
	garis()

	biasa, battleArena, vip := kelompokkanKomputerTersediaIteratif(w.Komputer)
	maxRows := max(len(biasa), len(battleArena), len(vip))

	for i := 0; i < maxRows; i++ {
		fmt.Printf("| %-25s | %-30s | %-23s |\n",
			formatCell(biasa, i),
			formatCell(battleArena, i),
			formatCell(vip, i))
	}

	garis()
}

// kelompokkanKomputerIteratif mengelompokkan semua komputer secara iteratif
func kelompokkanKomputerIteratif(komputers []Komputer) ([]Komputer, []Komputer, []Komputer) {
	var biasa, battleArena, vip []Komputer

	for _, komputer := range komputers {
		switch komputer.Tipe {
		case "Biasa":
			biasa = append(biasa, komputer)
		case "Battle Arena":
			battleArena = append(battleArena, komputer)
		case "VIP":
			vip = append(vip, komputer)
		}
	}

	return biasa, battleArena, vip
}

// kelompokkanKomputerTersediaIteratif mengelompokkan komputer yang tersedia secara iteratif
func kelompokkanKomputerTersediaIteratif(komputers []Komputer) ([]Komputer, []Komputer, []Komputer) {
	var biasa, battleArena, vip []Komputer

	for _, komputer := range komputers {
		if komputer.Status == "Tersedia" {
			switch komputer.Tipe {
			case "Biasa":
				biasa = append(biasa, komputer)
			case "Battle Arena":
				battleArena = append(battleArena, komputer)
			case "VIP":
				vip = append(vip, komputer)
			}
		}
	}

	return biasa, battleArena, vip
}

// TampilkanSemuaKomputerRekursif menampilkan semua komputer (tersedia dan digunakan) secara rekursif
func (w *Warnet) TampilkanSemuaKomputerRekursif() {
	garis()
	fmt.Println("| Biasa                     | Battle Arena                   | VIP                     |")
	garis()

	biasa, battleArena, vip := kelompokkanKomputerRekursif(w.Komputer, 0, []Komputer{}, []Komputer{}, []Komputer{})
	cetakTabelRekursif(biasa, battleArena, vip, 0, max(len(biasa), len(battleArena), len(vip)))
	garis()
}

// TampilkanKomputerTersediaRekursif menampilkan komputer yang tersedia secara rekursif
func (w *Warnet) TampilkanKomputerTersediaRekursif() {
	garis()
	fmt.Println("| Biasa                     | Battle Arena                   | VIP                     |")
	garis()

	biasa, battleArena, vip := kelompokkanKomputerTersediaRekursif(w.Komputer, 0, []Komputer{}, []Komputer{}, []Komputer{})
	cetakTabelRekursif(biasa, battleArena, vip, 0, max(len(biasa), len(battleArena), len(vip)))
	garis()
}

// kelompokkanKomputerRekursif mengelompokkan semua komputer secara rekursif
func kelompokkanKomputerRekursif(komputers []Komputer, index int, biasa, battleArena, vip []Komputer) ([]Komputer, []Komputer, []Komputer) {
	if index >= len(komputers) {
		return biasa, battleArena, vip
	}

	komputer := komputers[index]
	switch komputer.Tipe {
	case "Biasa":
		biasa = append(biasa, komputer)
	case "Battle Arena":
		battleArena = append(battleArena, komputer)
	case "VIP":
		vip = append(vip, komputer)
	}

	return kelompokkanKomputerRekursif(komputers, index+1, biasa, battleArena, vip)
}

// kelompokkanKomputerTersediaRekursif mengelompokkan komputer yang tersedia secara rekursif
func kelompokkanKomputerTersediaRekursif(komputers []Komputer, index int, biasa, battleArena, vip []Komputer) ([]Komputer, []Komputer, []Komputer) {
	if index >= len(komputers) {
		return biasa, battleArena, vip
	}

	komputer := komputers[index]
	if komputer.Status == "Tersedia" {
		switch komputer.Tipe {
		case "Biasa":
			biasa = append(biasa, komputer)
		case "Battle Arena":
			battleArena = append(battleArena, komputer)
		case "VIP":
			vip = append(vip, komputer)
		}
	}

	return kelompokkanKomputerTersediaRekursif(komputers, index+1, biasa, battleArena, vip)
}

// cetakTabelRekursif mencetak tabel secara rekursif
func cetakTabelRekursif(biasa, battleArena, vip []Komputer, currentRow, maxRows int) {
	if currentRow >= maxRows {
		return
	}

	fmt.Printf("| %-25s | %-30s | %-23s |\n",
		formatCellRekursif(biasa, currentRow),
		formatCellRekursif(battleArena, currentRow),
		formatCellRekursif(vip, currentRow))

	cetakTabelRekursif(biasa, battleArena, vip, currentRow+1, maxRows)
}

// formatCellRekursif mengembalikan string untuk sel tertentu secara rekursif
func formatCellRekursif(komputers []Komputer, index int) string {
	if index < len(komputers) {
		return fmt.Sprintf("K-%d %s (%s)", komputers[index].ID, komputers[index].Tipe, komputers[index].Status)
	}
	return ""
}

func garis() {
	fmt.Print("+")
	for i := 0; i < 27; i++ {
		fmt.Print("-")
	}
	fmt.Print("+")
	for i := 0; i < 32; i++ {
		fmt.Print("-")
	}
	fmt.Print("+")
	for i := 0; i < 25; i++ {
		fmt.Print("-")
	}
	fmt.Print("+")
	fmt.Println()
}

// formatCell mengembalikan string untuk sel tertentu
func formatCell(komputers []Komputer, index int) string {
	if index < len(komputers) {
		return fmt.Sprintf("K-%d %s (%s)", komputers[index].ID, komputers[index].Tipe, komputers[index].Status)
	}
	return ""
}

// max mengembalikan nilai maksimum dari beberapa angka
func max(values ...int) int {
	maxVal := values[0]
	for _, v := range values {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

// MasukkanPenggunaKomputer mengalokasikan komputer untuk pengguna
func (w *Warnet) MasukkanPenggunaKomputerIteratif(jumlahPengguna int) {
	tersedia := []int{}
	for i, komputer := range w.Komputer {
		if komputer.Status == "Tersedia" {
			tersedia = append(tersedia, i)
		}
	}

	if len(tersedia) < jumlahPengguna {
		fmt.Println("Komputer yang tersedia tidak mencukupi untuk semua pengguna.")
		return
	}

	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= jumlahPengguna; i++ {
		idx := rand.Intn(len(tersedia))
		w.Komputer[tersedia[idx]].Status = "Digunakan"
		w.Komputer[tersedia[idx]].Pengguna = fmt.Sprintf("Pengguna-%d", i)
		w.Komputer[tersedia[idx]].WaktuMulai = time.Now()
		tersedia = append(tersedia[:idx], tersedia[idx+1:]...)
	}

	fmt.Println("Pengguna berhasil dimasukkan ke komputer.")
}

func (w *Warnet) MasukkanPenggunaKomputerRekursif(jumlahPengguna int) {
	helperMasukkanPenggunaRekursif(w, jumlahPengguna, []int{}, 0)
}

func helperMasukkanPenggunaRekursif(w *Warnet, jumlahPengguna int, tersedia []int, index int) {
	if jumlahPengguna == 0 {
		// Semua pengguna sudah dialokasikan, keluar dari rekursi
		return
	}

	if index < len(w.Komputer) {
		if w.Komputer[index].Status == "Tersedia" {
			tersedia = append(tersedia, index)
		}
		helperMasukkanPenggunaRekursif(w, jumlahPengguna, tersedia, index+1)
		return
	}

	if len(tersedia) == 0 {
		fmt.Println("Tidak ada komputer yang tersedia.")
		return
	}

	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(tersedia))
	selectedIndex := tersedia[idx]
	w.Komputer[selectedIndex].Status = "Digunakan"
	w.Komputer[selectedIndex].Pengguna = fmt.Sprintf("Pengguna-%d", jumlahPengguna)
	w.Komputer[selectedIndex].WaktuMulai = time.Now()

	// Hapus komputer yang digunakan dari daftar tersedia
	tersedia = append(tersedia[:idx], tersedia[idx+1:]...)
	helperMasukkanPenggunaRekursif(w, jumlahPengguna-1, tersedia, index)
}

func main() {
	warnet := Warnet{}

	jenisKomputer := []string{"Biasa", "VIP", "Battle Arena"}
	jumlahKomputer := 800

	for i := 1; i <= jumlahKomputer; i++ {
		jenis := jenisKomputer[rand.Intn(len(jenisKomputer))]
		warnet.TambahKomputer(i, jenis)
	}

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tampilkan Semua Komputer")
		fmt.Println("2. Tampilkan Komputer yang Tersedia")
		fmt.Println("3. Masukkan Pengguna Komputer")
		fmt.Println("0. Keluar")
		fmt.Print("Pilihan Anda: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			startIteratif := time.Now()
			warnet.TampilkanSemuaKomputerIteratif()
			iteratifDuration := time.Since(startIteratif).Seconds()

			startRekursif := time.Now()
			warnet.TampilkanSemuaKomputerRekursif()
			rekursifDuration := time.Since(startRekursif).Seconds()

			fmt.Printf("Waktu eksekusi (Iteratif): %.6f detik\n", iteratifDuration)
			fmt.Printf("Waktu eksekusi (Rekursif): %.6f detik\n", rekursifDuration)

		case 2:
			startIteratif := time.Now()
			warnet.TampilkanKomputerTersediaIteratif()
			iteratifDuration := time.Since(startIteratif).Seconds()

			startRekursif := time.Now()
			warnet.TampilkanKomputerTersediaRekursif()
			rekursifDuration := time.Since(startRekursif).Seconds()

			fmt.Printf("Waktu eksekusi (Iteratif): %.6f detik\n", iteratifDuration)
			fmt.Printf("Waktu eksekusi (Rekursif): %.6f detik\n", rekursifDuration)

		case 3:
			fmt.Print("Masukkan jumlah pengguna: ")
			var jumlahPengguna int
			fmt.Scan(&jumlahPengguna)

			tersedia := []int{}
			for i, komputer := range warnet.Komputer {
				if komputer.Status == "Tersedia" {
					tersedia = append(tersedia, i)
				}
			}

			if len(tersedia) < jumlahPengguna {
				fmt.Println("Komputer yang tersedia tidak mencukupi untuk semua pengguna.")
				return
			}

			if jumlahPengguna%2 == 0 {
				a:= jumlahPengguna / 2
				startIteratif := time.Now()
				warnet.MasukkanPenggunaKomputerIteratif(a)
				iteratifDuration := time.Since(startIteratif).Seconds()

				startRekursif := time.Now()
				warnet.MasukkanPenggunaKomputerRekursif(a)
				rekursifDuration := time.Since(startRekursif).Seconds()

				fmt.Printf("Waktu eksekusi (Iteratif): %.6f detik\n", iteratifDuration)
				fmt.Printf("Waktu eksekusi (Rekursif): %.6f detik\n", rekursifDuration)
			} else {
				a:= jumlahPengguna / 2
				b:= a + 1
				startIteratif := time.Now()
				warnet.MasukkanPenggunaKomputerIteratif(b)
				iteratifDuration := time.Since(startIteratif).Seconds()

				startRekursif := time.Now()
				warnet.MasukkanPenggunaKomputerRekursif(a)
				rekursifDuration := time.Since(startRekursif).Seconds()

				fmt.Printf("Waktu eksekusi (Iteratif): %.6f detik\n", iteratifDuration)
				fmt.Printf("Waktu eksekusi (Rekursif): %.6f detik\n", rekursifDuration)
			}

		case 0:
			fmt.Println("Terima kasih telah menggunakan sistem warnet.")
			return

		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}
