package main
import "fmt"

type mahasiswa struct {
	nama string
	nim  int
}

type nilai struct {
	kuis  float64
	tugas  float64
	praktikum  float64
	uas float64
	uts float64	
}
func main() {
    var mhs mahasiswa
	var n nilai
	var akhir float64
	fmt.Scanf("%s %d", &mhs.nama, &mhs.nim)
	fmt.Scanf("%f %f %f %f %f", &n.kuis, &n.tugas, &n.praktikum, &n.uas, &n.uts)
	akhir = (0.15*n.kuis + 0.15*n.tugas + 0.20*n.praktikum + 0.25*n.uas + 0.25*n.uts)
	fmt.Printf("NIM: %d\nNama: %s\n", mhs.nim, mhs.nama)
	fmt.Printf("Nilai tugas: %.2f\nNilai kuis: %.2f\nNilai praktikum: %.2f\nNilai UAS: %.2f\nNilai UTS: %.2f\nNilai akhir: %.2f\n", n.tugas, n.kuis, n.praktikum, n.uas, n.uts, akhir)
}

