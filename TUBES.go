package main

import "fmt"

// ============================================================
// STRUCT / TIPE DATA
// ============================================================

type Keahlian struct {
	Nama string
}

type User struct {
	ID       string
	Nama     string
	Minat    []string
	Keahlian []Keahlian
}

type SyaratKeahlian struct {
	Nama string
}

type Karier struct {
	ID          string
	Nama        string
	Industri    string
	GajiRata    int // dalam juta IDR
	ReqMinat    []string
	ReqKeahlian []SyaratKeahlian
	SkorCocok   float64 // dihitung dinamis
}

// ============================================================
// DATA AWAL (seed)
// ============================================================

func dataKarierAwal() []Karier {
	return []Karier{
		{
			ID:       "K001",
			Nama:     "Data Scientist",
			Industri: "Teknologi",
			GajiRata: 15,
			ReqMinat: []string{"Teknologi", "Matematika"},
			ReqKeahlian: []SyaratKeahlian{
				{"Python"},
				{"Statistik"},
				{"Machine Learning"},
			},
		},
		{
			ID:       "K002",
			Nama:     "UI/UX Designer",
			Industri: "Teknologi",
			GajiRata: 10,
			ReqMinat: []string{"Desain", "Teknologi"},
			ReqKeahlian: []SyaratKeahlian{
				{"Figma"},
				{"Riset Pengguna"},
			},
		},
		{
			ID:       "K003",
			Nama:     "Dokter Umum",
			Industri: "Kesehatan",
			GajiRata: 12,
			ReqMinat: []string{"Kesehatan", "Sains"},
			ReqKeahlian: []SyaratKeahlian{
				{"Anatomi"},
				{"Farmakologi"},
			},
		},
		{
			ID:       "K004",
			Nama:     "Guru Matematika",
			Industri: "Pendidikan",
			GajiRata: 6,
			ReqMinat: []string{"Pendidikan", "Matematika"},
			ReqKeahlian: []SyaratKeahlian{
				{"Matematika"},
				{"Komunikasi"},
			},
		},
		{
			ID:       "K005",
			Nama:     "Software Engineer",
			Industri: "Teknologi",
			GajiRata: 18,
			ReqMinat: []string{"Teknologi", "Logika"},
			ReqKeahlian: []SyaratKeahlian{
				{"Golang"},
				{"Algoritma"},
				{"Git"},
			},
		},
		{
			ID:       "K006",
			Nama:     "Akuntan",
			Industri: "Keuangan",
			GajiRata: 9,
			ReqMinat: []string{"Keuangan", "Matematika"},
			ReqKeahlian: []SyaratKeahlian{
				{"Akuntansi"},
				{"Excel"},
			},
		},
		{
			ID:       "K007",
			Nama:     "Jurnalis",
			Industri: "Media",
			GajiRata: 7,
			ReqMinat: []string{"Menulis", "Komunikasi"},
			ReqKeahlian: []SyaratKeahlian{
				{"Penulisan"},
				{"Riset"},
			},
		},
		{
			ID:       "K008",
			Nama:     "Arsitek",
			Industri: "Konstruksi",
			GajiRata: 13,
			ReqMinat: []string{"Desain", "Sains"},
			ReqKeahlian: []SyaratKeahlian{
				{"AutoCAD"},
				{"Desain Struktural"},
			},
		},
	}
}

// ============================================================
// MODUL MANAJEMEN DATA (Spesifikasi a)
// ============================================================

func tambahMinat(user *User, minat string) {
	for _, m := range user.Minat {
		if m == minat {
			fmt.Println("  [!] Minat '" + minat + "' sudah ada.")
			return
		}
	}
	user.Minat = append(user.Minat, minat)
	fmt.Println("  [✓] Minat '" + minat + "' berhasil ditambahkan.")
}

func hapusMinat(user *User, minat string) {
	baru := []string{}
	ditemukan := false
	for _, m := range user.Minat {
		if m == minat {
			ditemukan = true
		} else {
			baru = append(baru, m)
		}
	}
	if !ditemukan {
		fmt.Println("  [!] Minat '" + minat + "' tidak ditemukan.")
		return
	}
	user.Minat = baru
	fmt.Println("  [✓] Minat '" + minat + "' berhasil dihapus.")
}

func tambahKeahlian(user *User, nama string) {
	for _, k := range user.Keahlian {
		if k.Nama == nama {
			fmt.Println("  [!] Keahlian '" + nama + "' sudah ada.")
			return
		}
	}
	user.Keahlian = append(user.Keahlian, Keahlian{nama})
	fmt.Println("  [✓] Keahlian '" + nama + "' berhasil ditambahkan.")
}

func hapusKeahlian(user *User, nama string) {
	baru := []Keahlian{}
	ditemukan := false
	for _, k := range user.Keahlian {
		if k.Nama == nama {
			ditemukan = true
		} else {
			baru = append(baru, k)
		}
	}
	if !ditemukan {
		fmt.Println("  [!] Keahlian '" + nama + "' tidak ditemukan.")
		return
	}
	user.Keahlian = baru
	fmt.Println("  [✓] Keahlian '" + nama + "' berhasil dihapus.")
}

func tampilProfil(user *User) {
	fmt.Println("  ┌─────────────────────────────────────")
	fmt.Printf("  │ Nama : %s (ID: %s)\n", user.Nama, user.ID)
	fmt.Print("  │ Minat: ")
	if len(user.Minat) == 0 {
		fmt.Println("(belum ada)")
	} else {
		for i, m := range user.Minat {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(m)
		}
		fmt.Println()
	}
	fmt.Println("  │ Keahlian:")
	if len(user.Keahlian) == 0 {
		fmt.Println("  │   (belum ada)")
	} else {
		for i, k := range user.Keahlian {
			fmt.Printf("  │   %d. %s\n", i+1, k.Nama)
		}
	}
	fmt.Println("  └─────────────────────────────────────")
}

// ============================================================
// MODUL REKOMENDASI & STATISTIK (Spesifikasi b & e)
// ============================================================

func hitungSkor(user *User, karier *Karier) float64 {
	if len(karier.ReqMinat) == 0 && len(karier.ReqKeahlian) == 0 {
		return 0
	}

	totalPoin := 0.0
	maksPoin := 0.0

	// Komponen minat: bobot 40%
	if len(karier.ReqMinat) > 0 {
		bobotPerMinat := 40.0 / float64(len(karier.ReqMinat))
		for _, rm := range karier.ReqMinat {
			maksPoin += bobotPerMinat
			for _, um := range user.Minat {
				if um == rm {
					totalPoin += bobotPerMinat
					break
				}
			}
		}
	}

	// Komponen keahlian: bobot 60%
	if len(karier.ReqKeahlian) > 0 {
		bobotPerKeahlian := 60.0 / float64(len(karier.ReqKeahlian))
		for _, rk := range karier.ReqKeahlian {
			maksPoin += bobotPerKeahlian
			for _, uk := range user.Keahlian {
				if uk.Nama == rk.Nama {
					totalPoin += bobotPerKeahlian
					break
				}
			}
		}
	}

	if maksPoin == 0 {
		return 0
	}
	skor := (totalPoin / maksPoin) * 100
	karier.SkorCocok = skor
	return skor
}

func hitungSemuaRekomendasi(user *User, daftarKarier []Karier) []Karier {
	hasil := []Karier{}
	for i := range daftarKarier {
		hitungSkor(user, &daftarKarier[i])
		if daftarKarier[i].SkorCocok > 0 {
			hasil = append(hasil, daftarKarier[i])
		}
	}
	return hasil
}

func tampilRekomendasi(daftar []Karier) {
	if len(daftar) == 0 {
		fmt.Println("  [!] Tidak ada rekomendasi yang ditemukan.")
		return
	}
	fmt.Println()
	fmt.Println("  No  Nama Karier           Industri       Gaji(jt)  Kecocokan")
	fmt.Println("  ─────────────────────────────────────────────────────────────")
	for i, k := range daftar {
		bar := buatBar(k.SkorCocok, 10)
		fmt.Printf("  %-3d %-22s %-15s %-9d %s %.1f%%\n",
			i+1, k.Nama, k.Industri, k.GajiRata, bar, k.SkorCocok)
	}
}

func buatBar(persen float64, panjang int) string {
	isi := int(persen / 100.0 * float64(panjang))
	bar := "["
	for i := 0; i < panjang; i++ {
		if i < isi {
			bar += "█"
		} else {
			bar += "░"
		}
	}
	bar += "]"
	return bar
}

func tampilStatistik(user *User, daftarKarier []Karier) {
	fmt.Println()
	fmt.Println("  ╔══════════════════════════════════════════╗")
	fmt.Println("  ║         STATISTIK KECOCOKAN KARIER       ║")
	fmt.Println("  ╠══════════════════════════════════════════╣")

	total := 0.0
	maks := 0.0
	namaMaks := ""
	for i := range daftarKarier {
		hitungSkor(user, &daftarKarier[i])
		total += daftarKarier[i].SkorCocok
		if daftarKarier[i].SkorCocok > maks {
			maks = daftarKarier[i].SkorCocok
			namaMaks = daftarKarier[i].Nama
		}
	}

	rataRata := 0.0
	if len(daftarKarier) > 0 {
		rataRata = total / float64(len(daftarKarier))
	}

	for _, k := range daftarKarier {
		bar := buatBar(k.SkorCocok, 15)
		fmt.Printf("  ║  %-22s %s %.1f%%\n", k.Nama, bar, k.SkorCocok)
	}

	fmt.Println("  ╠══════════════════════════════════════════╣")
	fmt.Printf("  ║  Rata-rata kecocokan  : %.1f%%\n", rataRata)
	fmt.Printf("  ║  Karier terbaik       : %s (%.1f%%)\n", namaMaks, maks)
	fmt.Printf("  ║  Total karier dianalisis: %d\n", len(daftarKarier))
	fmt.Println("  ╚══════════════════════════════════════════╝")
}

// ============================================================
// MODUL PENCARIAN (Spesifikasi c)
// ============================================================

// toLower manual tanpa library strings
func toLower(s string) string {
	hasil := []byte(s)
	for i, c := range hasil {
		if c >= 'A' && c <= 'Z' {
			hasil[i] = c + 32
		}
	}
	return string(hasil)
}

// contains manual tanpa library strings
func contains(s, sub string) bool {
	if len(sub) == 0 {
		return true
	}
	if len(sub) > len(s) {
		return false
	}
	for i := 0; i <= len(s)-len(sub); i++ {
		cocok := true
		for j := 0; j < len(sub); j++ {
			if s[i+j] != sub[j] {
				cocok = false
				break
			}
		}
		if cocok {
			return true
		}
	}
	return false
}

// Sequential Search: cari berdasarkan nama atau industri
func sequentialSearch(daftar []Karier, query string) []Karier {
	hasil := []Karier{}
	q := toLower(query)
	for _, k := range daftar {
		if contains(toLower(k.Nama), q) || contains(toLower(k.Industri), q) {
			hasil = append(hasil, k)
		}
	}
	return hasil
}

// Urutkan alfabet untuk keperluan Binary Search
func urutAlfabet(daftar []Karier) []Karier {
	salinan := make([]Karier, len(daftar))
	copy(salinan, daftar)
	// Insertion sort untuk mengurutkan A–Z berdasarkan nama
	for i := 1; i < len(salinan); i++ {
		kunci := salinan[i]
		j := i - 1
		for j >= 0 && toLower(salinan[j].Nama) > toLower(kunci.Nama) {
			salinan[j+1] = salinan[j]
			j--
		}
		salinan[j+1] = kunci
	}
	return salinan
}

// Binary Search: cari berdasarkan nama tepat (data harus terurut A–Z)
func binarySearch(daftarTerurut []Karier, target string) int {
	kiri := 0
	kanan := len(daftarTerurut) - 1
	targetLower := toLower(target)
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		val := toLower(daftarTerurut[tengah].Nama)
		if val == targetLower {
			return tengah
		} else if val < targetLower {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}

// ============================================================
// MODUL PENGURUTAN (Spesifikasi d)
// ============================================================

// Selection Sort: urutkan berdasarkan gaji rata-rata (tertinggi ke terendah)
func selectionSortGaji(daftar []Karier) []Karier {
	salinan := make([]Karier, len(daftar))
	copy(salinan, daftar)
	n := len(salinan)
	for i := 0; i < n-1; i++ {
		idxMaks := i
		for j := i + 1; j < n; j++ {
			if salinan[j].GajiRata > salinan[idxMaks].GajiRata {
				idxMaks = j
			}
		}
		salinan[i], salinan[idxMaks] = salinan[idxMaks], salinan[i]
	}
	return salinan
}

// Insertion Sort: urutkan berdasarkan skor kecocokan (tertinggi ke terendah)
func insertionSortSkor(daftar []Karier) []Karier {
	salinan := make([]Karier, len(daftar))
	copy(salinan, daftar)
	for i := 1; i < len(salinan); i++ {
		kunci := salinan[i]
		j := i - 1
		for j >= 0 && salinan[j].SkorCocok < kunci.SkorCocok {
			salinan[j+1] = salinan[j]
			j--
		}
		salinan[j+1] = kunci
	}
	return salinan
}

// ============================================================
// TAMPILAN HELPER
// ============================================================

func cetakHeader(judul string) {
	fmt.Println()
	fmt.Println("  ┌──────────────────────────────────────────┐")
	fmt.Printf("  │  %-42s│\n", judul)
	fmt.Println("  └──────────────────────────────────────────┘")
}

func cetakGaris() {
	fmt.Println("  ──────────────────────────────────────────────")
}

func inputStr(prompt string) string {
	fmt.Print(prompt)
	var s string
	fmt.Scan(&s)
	return s
}

func inputInt(prompt string) int {
	fmt.Print(prompt)
	var n int
	fmt.Scan(&n)
	return n
}

// ============================================================
// MENU MANAJEMEN DATA
// ============================================================

func menuManajemenData(user *User) {
	for {
		cetakHeader("KELOLA MINAT & KEAHLIAN")
		fmt.Println("  1. Lihat profil")
		fmt.Println("  2. Tambah minat")
		fmt.Println("  3. Hapus minat")
		fmt.Println("  4. Tambah keahlian")
		fmt.Println("  5. Hapus keahlian")
		fmt.Println("  0. Kembali")
		cetakGaris()
		pilihan := inputInt("  Pilihan: ")

		switch pilihan {
		case 1:
			tampilProfil(user)
		case 2:
			minat := inputStr("  Masukkan minat baru: ")
			tambahMinat(user, minat)
		case 3:
			tampilProfil(user)
			minat := inputStr("  Minat yang dihapus: ")
			hapusMinat(user, minat)
		case 4:
			nama := inputStr("  Nama keahlian: ")
			tambahKeahlian(user, nama)
		case 5:
			tampilProfil(user)
			nama := inputStr("  Keahlian yang dihapus: ")
			hapusKeahlian(user, nama)
		case 0:
			return
		default:
			fmt.Println("  [!] Pilihan tidak valid.")
		}
	}
}

// ============================================================
// MENU PENCARIAN
// ============================================================

func menuPencarian(daftarKarier []Karier) {
	for {
		cetakHeader("CARI KARIER")
		fmt.Println("  1. Sequential Search (nama/industri)")
		fmt.Println("  2. Binary Search (nama tepat)")
		fmt.Println("  0. Kembali")
		cetakGaris()
		pilihan := inputInt("  Pilihan: ")

		switch pilihan {
		case 1:
			query := inputStr("  Masukkan kata kunci: ")
			hasil := sequentialSearch(daftarKarier, query)
			fmt.Printf("\n  Ditemukan %d hasil untuk '%s':\n", len(hasil), query)
			tampilRekomendasi(hasil)

		case 2:
			terurut := urutAlfabet(daftarKarier)
			fmt.Println("\n  Daftar karier (A–Z):")
			for i, k := range terurut {
				fmt.Printf("  %2d. %s\n", i+1, k.Nama)
			}
			target := inputStr("\n  Masukkan nama karier (tepat): ")
			idx := binarySearch(terurut, target)
			if idx == -1 {
				fmt.Println("  [!] Karier '" + target + "' tidak ditemukan.")
			} else {
				k := terurut[idx]
				fmt.Println()
				fmt.Println("  [✓] Karier ditemukan!")
				fmt.Printf("  Nama      : %s\n", k.Nama)
				fmt.Printf("  Industri  : %s\n", k.Industri)
				fmt.Printf("  Gaji rata : %d juta/bulan\n", k.GajiRata)
				fmt.Println("  Syarat keahlian:")
				for _, rk := range k.ReqKeahlian {
					fmt.Printf("    - %s\n", rk.Nama)
				}
			}

		case 0:
			return
		default:
			fmt.Println("  [!] Pilihan tidak valid.")
		}
	}
}

// ============================================================
// MENU PENGURUTAN
// ============================================================

func menuPengurutan(user *User, daftarKarier []Karier) {
	for {
		cetakHeader("URUTKAN REKOMENDASI")
		fmt.Println("  1. Selection Sort — berdasarkan gaji (tertinggi)")
		fmt.Println("  2. Insertion Sort — berdasarkan kecocokan (tertinggi)")
		fmt.Println("  0. Kembali")
		cetakGaris()
		pilihan := inputInt("  Pilihan: ")

		switch pilihan {
		case 1:
			hasil := hitungSemuaRekomendasi(user, daftarKarier)
			terurut := selectionSortGaji(hasil)
			fmt.Println("\n  Karier diurutkan berdasarkan gaji (Selection Sort):")
			fmt.Println("\n  No  Nama Karier           Industri       Gaji(jt)  Kecocokan")
			fmt.Println("  ─────────────────────────────────────────────────────────────")
			for i, k := range terurut {
				bar := buatBar(k.SkorCocok, 10)
				fmt.Printf("  %-3d %-22s %-15s %-9d %s %.1f%%\n",
					i+1, k.Nama, k.Industri, k.GajiRata, bar, k.SkorCocok)
			}

		case 2:
			hasil := hitungSemuaRekomendasi(user, daftarKarier)
			terurut := insertionSortSkor(hasil)
			fmt.Println("\n  Karier diurutkan berdasarkan kecocokan (Insertion Sort):")
			fmt.Println("\n  No  Nama Karier           Industri       Gaji(jt)  Kecocokan")
			fmt.Println("  ─────────────────────────────────────────────────────────────")
			for i, k := range terurut {
				bar := buatBar(k.SkorCocok, 10)
				fmt.Printf("  %-3d %-22s %-15s %-9d %s %.1f%%\n",
					i+1, k.Nama, k.Industri, k.GajiRata, bar, k.SkorCocok)
			}

		case 0:
			return
		default:
			fmt.Println("  [!] Pilihan tidak valid.")
		}
	}
}

// ============================================================
// MAIN
// ============================================================

func main() {
	daftarKarier := dataKarierAwal()

	// Setup pengguna awal
	fmt.Println("  ╔══════════════════════════════════════════╗")
	fmt.Println("  ║   APLIKASI REKOMENDASI KARIER            ║")
	fmt.Println("  ║   Algoritma Pemrograman                  ║")
	fmt.Println("  ╚══════════════════════════════════════════╝")
	fmt.Println()
	fmt.Print("  Masukkan nama Anda: ")
	var namaPengguna string
	fmt.Scan(&namaPengguna)

	user := &User{
		ID:       "U001",
		Nama:     namaPengguna,
		Minat:    []string{},
		Keahlian: []Keahlian{},
	}

	fmt.Println()
	fmt.Println("  Selamat datang, " + user.Nama + "!")
	fmt.Println("  Silakan lengkapi profil Anda di menu 'Kelola Data'.")

	for {
		cetakHeader("MENU UTAMA")
		fmt.Println("  1. Kelola Minat & Keahlian")
		fmt.Println("  2. Lihat Rekomendasi Karier")
		fmt.Println("  3. Cari Karier")
		fmt.Println("  4. Urutkan Rekomendasi")
		fmt.Println("  5. Statistik Kecocokan")
		fmt.Println("  0. Keluar")
		cetakGaris()
		pilihan := inputInt("  Pilihan: ")

		switch pilihan {
		case 1:
			menuManajemenData(user)

		case 2:
			cetakHeader("REKOMENDASI KARIER")
			hasil := hitungSemuaRekomendasi(user, daftarKarier)
			if len(hasil) == 0 {
				fmt.Println("  [!] Lengkapi minat dan keahlian dulu di menu 1.")
			} else {
				// Default: tampilkan urut skor tertinggi
				terurut := insertionSortSkor(hasil)
				tampilRekomendasi(terurut)
			}

		case 3:
			menuPencarian(daftarKarier)

		case 4:
			menuPengurutan(user, daftarKarier)

		case 5:
			cetakHeader("STATISTIK KECOCOKAN")
			tampilStatistik(user, daftarKarier)

		case 0:
			fmt.Println()
			fmt.Println("  Terima kasih! Semoga sukses dalam perjalanan karier Anda.")
			fmt.Println()
			return

		default:
			fmt.Println("  [!] Pilihan tidak valid.")
		}
	}
}
