package main

import "fmt"

// ============================================================
// KONSTANTA & TIPE DATA
// ============================================================

const MAX_KARIER int = 8
const MAX_SYARAT int = 5
const MAX_PROFIL int = 15
const MAX_PILIHAN_MINAT int = 10
const MAX_PILIHAN_KEAHLIAN int = 18

type Karier struct {
	ID             string
	Nama           string
	Industri       string
	GajiRata       int
	ReqMinat       [MAX_SYARAT]string
	JmlReqMinat    int
	ReqKeahlian    [MAX_SYARAT]string
	JmlReqKeahlian int
	SkorCocok      float64
}

type User struct {
	Nama           string
	Minat          [MAX_PROFIL]string
	JumlahMinat    int
	Keahlian       [MAX_PROFIL]string
	JumlahKeahlian int
}

// ============================================================
// VARIABEL GLOBAL
// ============================================================

var daftarKarier [MAX_KARIER]Karier
var jumlahKarier int

var pilihanMinat [MAX_PILIHAN_MINAT]string
var jumlahPilihanMinat int
var pilihanKeahlian [MAX_PILIHAN_KEAHLIAN]string
var jumlahPilihanKeahlian int

// ============================================================
// INISIALISASI DATA
// ============================================================

func initData() {
	daftarKarier[0].ID = "K001"
	daftarKarier[0].Nama = "Data_Scientist"
	daftarKarier[0].Industri = "Teknologi"
	daftarKarier[0].GajiRata = 15
	daftarKarier[0].ReqMinat[0] = "Teknologi"
	daftarKarier[0].ReqMinat[1] = "Matematika"
	daftarKarier[0].JmlReqMinat = 2
	daftarKarier[0].ReqKeahlian[0] = "Python"
	daftarKarier[0].ReqKeahlian[1] = "Statistik"
	daftarKarier[0].ReqKeahlian[2] = "Machine_Learning"
	daftarKarier[0].JmlReqKeahlian = 3

	daftarKarier[1].ID = "K002"
	daftarKarier[1].Nama = "UI/UX_Designer"
	daftarKarier[1].Industri = "Teknologi"
	daftarKarier[1].GajiRata = 10
	daftarKarier[1].ReqMinat[0] = "Desain"
	daftarKarier[1].ReqMinat[1] = "Teknologi"
	daftarKarier[1].JmlReqMinat = 2
	daftarKarier[1].ReqKeahlian[0] = "Figma"
	daftarKarier[1].ReqKeahlian[1] = "Riset_Pengguna"
	daftarKarier[1].JmlReqKeahlian = 2

	daftarKarier[2].ID = "K003"
	daftarKarier[2].Nama = "Dokter_Umum"
	daftarKarier[2].Industri = "Kesehatan"
	daftarKarier[2].GajiRata = 12
	daftarKarier[2].ReqMinat[0] = "Kesehatan"
	daftarKarier[2].ReqMinat[1] = "Sains"
	daftarKarier[2].JmlReqMinat = 2
	daftarKarier[2].ReqKeahlian[0] = "Anatomi"
	daftarKarier[2].ReqKeahlian[1] = "Farmakologi"
	daftarKarier[2].JmlReqKeahlian = 2

	daftarKarier[3].ID = "K004"
	daftarKarier[3].Nama = "Guru_Matematika"
	daftarKarier[3].Industri = "Pendidikan"
	daftarKarier[3].GajiRata = 6
	daftarKarier[3].ReqMinat[0] = "Pendidikan"
	daftarKarier[3].ReqMinat[1] = "Matematika"
	daftarKarier[3].JmlReqMinat = 2
	daftarKarier[3].ReqKeahlian[0] = "Matematika"
	daftarKarier[3].ReqKeahlian[1] = "Komunikasi"
	daftarKarier[3].JmlReqKeahlian = 2

	daftarKarier[4].ID = "K005"
	daftarKarier[4].Nama = "Software_Engineer"
	daftarKarier[4].Industri = "Teknologi"
	daftarKarier[4].GajiRata = 18
	daftarKarier[4].ReqMinat[0] = "Teknologi"
	daftarKarier[4].ReqMinat[1] = "Logika"
	daftarKarier[4].JmlReqMinat = 2
	daftarKarier[4].ReqKeahlian[0] = "Golang"
	daftarKarier[4].ReqKeahlian[1] = "Algoritma"
	daftarKarier[4].ReqKeahlian[2] = "Git"
	daftarKarier[4].JmlReqKeahlian = 3

	daftarKarier[5].ID = "K006"
	daftarKarier[5].Nama = "Akuntan"
	daftarKarier[5].Industri = "Keuangan"
	daftarKarier[5].GajiRata = 9
	daftarKarier[5].ReqMinat[0] = "Keuangan"
	daftarKarier[5].ReqMinat[1] = "Matematika"
	daftarKarier[5].JmlReqMinat = 2
	daftarKarier[5].ReqKeahlian[0] = "Akuntansi"
	daftarKarier[5].ReqKeahlian[1] = "Excel"
	daftarKarier[5].JmlReqKeahlian = 2

	daftarKarier[6].ID = "K007"
	daftarKarier[6].Nama = "Jurnalis"
	daftarKarier[6].Industri = "Media"
	daftarKarier[6].GajiRata = 7
	daftarKarier[6].ReqMinat[0] = "Menulis"
	daftarKarier[6].ReqMinat[1] = "Komunikasi"
	daftarKarier[6].JmlReqMinat = 2
	daftarKarier[6].ReqKeahlian[0] = "Penulisan"
	daftarKarier[6].ReqKeahlian[1] = "Riset"
	daftarKarier[6].JmlReqKeahlian = 2

	daftarKarier[7].ID = "K008"
	daftarKarier[7].Nama = "Arsitek"
	daftarKarier[7].Industri = "Konstruksi"
	daftarKarier[7].GajiRata = 13
	daftarKarier[7].ReqMinat[0] = "Desain"
	daftarKarier[7].ReqMinat[1] = "Sains"
	daftarKarier[7].JmlReqMinat = 2
	daftarKarier[7].ReqKeahlian[0] = "AutoCAD"
	daftarKarier[7].ReqKeahlian[1] = "Desain_Struktural"
	daftarKarier[7].JmlReqKeahlian = 2

	jumlahKarier = 8

	pilihanMinat[0] = "Teknologi"
	pilihanMinat[1] = "Matematika"
	pilihanMinat[2] = "Desain"
	pilihanMinat[3] = "Kesehatan"
	pilihanMinat[4] = "Sains"
	pilihanMinat[5] = "Pendidikan"
	pilihanMinat[6] = "Logika"
	pilihanMinat[7] = "Keuangan"
	pilihanMinat[8] = "Menulis"
	pilihanMinat[9] = "Komunikasi"
	jumlahPilihanMinat = 10

	pilihanKeahlian[0] = "Python"
	pilihanKeahlian[1] = "Statistik"
	pilihanKeahlian[2] = "Machine_Learning"
	pilihanKeahlian[3] = "Figma"
	pilihanKeahlian[4] = "Riset_Pengguna"
	pilihanKeahlian[5] = "Anatomi"
	pilihanKeahlian[6] = "Farmakologi"
	pilihanKeahlian[7] = "Matematika"
	pilihanKeahlian[8] = "Komunikasi"
	pilihanKeahlian[9] = "Golang"
	pilihanKeahlian[10] = "Algoritma"
	pilihanKeahlian[11] = "Git"
	pilihanKeahlian[12] = "Akuntansi"
	pilihanKeahlian[13] = "Excel"
	pilihanKeahlian[14] = "Penulisan"
	pilihanKeahlian[15] = "Riset"
	pilihanKeahlian[16] = "AutoCAD"
	pilihanKeahlian[17] = "Desain_Struktural"
	jumlahPilihanKeahlian = 18
}

// ============================================================
// HELPER: INPUT & TAMPILAN
// ============================================================

// inputInt membaca integer dari user.
// Jika bukan angka, minta ulang sampai benar.
// Implementasi: baca per karakter, cek apakah semua digit 0-9.
// inputInt membaca integer dari user.
// Jika bukan angka, minta ulang sampai benar.
func inputInt(prompt string) int {
	var s string
	var n int
	var ok bool
	var i int
	var panjang int

	ok = false
	for !ok {
		fmt.Print(prompt)
		fmt.Scan(&s)

		// Cek apakah s adalah angka valid (boleh negatif)
		panjang = len(s)
		ok = true
		n = 0
		i = 0

		// Tangani tanda minus di depan
		if panjang > 0 && s[0] == '-' {
			i = 1
		}
		// Minimal harus ada 1 digit setelah tanda minus
		if i >= panjang {
			ok = false
		}
		// Periksa satu per satu apakah semua karakter adalah digit
		for i < panjang && ok {
			if s[i] >= '0' && s[i] <= '9' {
				n = n*10 + int(s[i]-'0')
				i++
			} else {
				ok = false
			}
		}
		// Terapkan tanda negatif jika ada
		if ok && panjang > 0 && s[0] == '-' {
			n = -n
		}
		if !ok {
			fmt.Println("  ⚠️  Masukkan angka yang valid.")
		}
	}
	return n
}

func cetakHeader(judul string) {
	fmt.Println()
	fmt.Println("  +==========================================+")
	fmt.Printf("  |  %-42s|\n", "🗂  "+judul)
	fmt.Println("  +==========================================+")
}

func cetakGaris() {
	fmt.Println("  ..........................................")
}

// buatBar membuat visualisasi bar berdasarkan persentase (0-100)
func buatBar(persen float64, panjang int) string {
	var isi int
	var bar string
	var i int

	isi = int(persen / 100.0 * float64(panjang))
	bar = "["
	for i = 0; i < panjang; i++ {
		if i < isi {
			bar = bar + "="
		} else {
			bar = bar + "."
		}
	}
	bar = bar + "]"
	return bar
}

// ada memeriksa apakah 'kata' ada di dalam 'kalimat' (case-sensitive)
// menggunakan perbandingan karakter satu per satu, tanpa library tambahan
func ada(kalimat string, kata string) bool {
	var pKalimat int
	var pKata int
	var i int
	var j int
	var cocok bool

	pKalimat = len(kalimat)
	pKata = len(kata)

	if pKata == 0 {
		return true
	}
	if pKata > pKalimat {
		return false
	}

	// Cek setiap posisi kemungkinan awal kata di dalam kalimat
	for i = 0; i <= pKalimat-pKata; i++ {
		cocok = true
		for j = 0; j < pKata && cocok; j++ {
			if kalimat[i+j] != kata[j] {
				cocok = false
			}
		}
		if cocok {
			return true
		}
	}
	return false
}

// ============================================================
// MODUL 1: KELOLA PROFIL (Spesifikasi a)
// ============================================================

func tampilProfil(user *User) {
	var i int
	fmt.Println()
	fmt.Println("  +=====================================+")
	fmt.Println("  |         👤  PROFIL SAYA              |")
	fmt.Println("  +=====================================+")
	fmt.Printf("  |  Nama     : %-23s|\n", user.Nama)
	fmt.Println("  |  ❤  Minat    :                       |")
	if user.JumlahMinat == 0 {
		fmt.Println("  |    (belum ada)                       |")
	} else {
		for i = 0; i < user.JumlahMinat; i++ {
			fmt.Printf("  |    %d. %-31s|\n", i+1, user.Minat[i])
		}
	}
	fmt.Println("  |  🛠  Keahlian :                       |")
	if user.JumlahKeahlian == 0 {
		fmt.Println("  |    (belum ada)                       |")
	} else {
		for i = 0; i < user.JumlahKeahlian; i++ {
			fmt.Printf("  |    %d. %-31s|\n", i+1, user.Keahlian[i])
		}
	}
	fmt.Println("  +=====================================+")
}

func tambahMinat(user *User) {
	var i int
	var nomor int
	var minat string
	var sudahAda bool

	if user.JumlahMinat >= MAX_PROFIL {
		fmt.Println("  ⚠️  Daftar minat sudah penuh.")
		return
	}

	fmt.Println()
	fmt.Println("  Pilihan minat yang tersedia:")
	cetakGaris()
	for i = 0; i < jumlahPilihanMinat; i++ {
		fmt.Printf("  %2d. %s\n", i+1, pilihanMinat[i])
	}
	cetakGaris()

	nomor = inputInt("  Pilih nomor minat: ")
	if nomor < 1 || nomor > jumlahPilihanMinat {
		fmt.Println("  ⚠️  Nomor tidak valid.")
		return
	}
	minat = pilihanMinat[nomor-1]

	// Cek apakah minat sudah ada di profil
	sudahAda = false
	for i = 0; i < user.JumlahMinat; i++ {
		if user.Minat[i] == minat {
			sudahAda = true
		}
	}
	if sudahAda {
		fmt.Println("  ⚠️  Minat '" + minat + "' sudah ada di profil.")
		return
	}

	user.Minat[user.JumlahMinat] = minat
	user.JumlahMinat++
	fmt.Println("  ✅ Minat '" + minat + "' berhasil ditambahkan!")
}

func hapusMinat(user *User) {
	var i int
	var nomor int
	var hapus string

	if user.JumlahMinat == 0 {
		fmt.Println("  ⚠️  Belum ada minat di profil.")
		return
	}

	fmt.Println()
	fmt.Println("  Minat Anda saat ini:")
	cetakGaris()
	for i = 0; i < user.JumlahMinat; i++ {
		fmt.Printf("  %2d. %s\n", i+1, user.Minat[i])
	}
	cetakGaris()

	nomor = inputInt("  Pilih nomor minat yang dihapus: ")
	if nomor < 1 || nomor > user.JumlahMinat {
		fmt.Println("  ⚠️  Nomor tidak valid.")
		return
	}

	hapus = user.Minat[nomor-1]
	// Geser semua elemen di belakangnya ke kiri untuk menutup celah
	for i = nomor - 1; i < user.JumlahMinat-1; i++ {
		user.Minat[i] = user.Minat[i+1]
	}
	user.Minat[user.JumlahMinat-1] = ""
	user.JumlahMinat--
	fmt.Println("  🗑  Minat '" + hapus + "' berhasil dihapus.")
}

func tambahKeahlian(user *User) {
	var i int
	var nomor int
	var keahlian string
	var sudahAda bool

	if user.JumlahKeahlian >= MAX_PROFIL {
		fmt.Println("  ⚠️  Daftar keahlian sudah penuh.")
		return
	}

	fmt.Println()
	fmt.Println("  Pilihan keahlian yang tersedia:")
	cetakGaris()
	for i = 0; i < jumlahPilihanKeahlian; i++ {
		fmt.Printf("  %2d. %s\n", i+1, pilihanKeahlian[i])
	}
	cetakGaris()

	nomor = inputInt("  Pilih nomor keahlian: ")
	if nomor < 1 || nomor > jumlahPilihanKeahlian {
		fmt.Println("  ⚠️  Nomor tidak valid.")
		return
	}
	keahlian = pilihanKeahlian[nomor-1]

	sudahAda = false
	for i = 0; i < user.JumlahKeahlian; i++ {
		if user.Keahlian[i] == keahlian {
			sudahAda = true
		}
	}
	if sudahAda {
		fmt.Println("  ⚠️  Keahlian '" + keahlian + "' sudah ada di profil.")
		return
	}

	user.Keahlian[user.JumlahKeahlian] = keahlian
	user.JumlahKeahlian++
	fmt.Println("  ✅ Keahlian '" + keahlian + "' berhasil ditambahkan!")
}

func hapusKeahlian(user *User) {
	var i int
	var nomor int
	var hapus string

	if user.JumlahKeahlian == 0 {
		fmt.Println("  ⚠️  Belum ada keahlian di profil.")
		return
	}

	fmt.Println()
	fmt.Println("  Keahlian Anda saat ini:")
	cetakGaris()
	for i = 0; i < user.JumlahKeahlian; i++ {
		fmt.Printf("  %2d. %s\n", i+1, user.Keahlian[i])
	}
	cetakGaris()

	nomor = inputInt("  Pilih nomor keahlian yang dihapus: ")
	if nomor < 1 || nomor > user.JumlahKeahlian {
		fmt.Println("  ⚠️  Nomor tidak valid.")
		return
	}

	hapus = user.Keahlian[nomor-1]
	for i = nomor - 1; i < user.JumlahKeahlian-1; i++ {
		user.Keahlian[i] = user.Keahlian[i+1]
	}
	user.Keahlian[user.JumlahKeahlian-1] = ""
	user.JumlahKeahlian--
	fmt.Println("  🗑  Keahlian '" + hapus + "' berhasil dihapus.")
}

func menuKelolaProfil(user *User) {
	var pilihan int
	for {
		cetakHeader("KELOLA MINAT & KEAHLIAN")
		fmt.Println("  1. 👁  Lihat profil")
		fmt.Println("  2. ➕ Tambah minat")
		fmt.Println("  3. ➖ Hapus minat")
		fmt.Println("  4. ➕ Tambah keahlian")
		fmt.Println("  5. ➖ Hapus keahlian")
		fmt.Println("  0. 🔙 Kembali")
		cetakGaris()
		pilihan = inputInt("  Pilihan: ")
		switch pilihan {
		case 1:
			tampilProfil(user)
		case 2:
			tambahMinat(user)
		case 3:
			hapusMinat(user)
		case 4:
			tambahKeahlian(user)
		case 5:
			hapusKeahlian(user)
		case 0:
			return
		default:
			fmt.Println("  ⚠️  Pilihan tidak valid.")
		}
	}
}

// ============================================================
// MODUL 2: REKOMENDASI & STATISTIK (Spesifikasi b & e)
// ============================================================

// hitungSkor menghitung persentase kecocokan user dengan satu karier.
// Bobot: minat 40% + keahlian 60%
func hitungSkor(user *User, k *Karier) float64 {
	var i int
	var j int
	var totalPoin float64
	var maksPoin float64
	var bobotMinat float64
	var bobotKeahlian float64

	if k.JmlReqMinat == 0 && k.JmlReqKeahlian == 0 {
		return 0
	}

	totalPoin = 0
	maksPoin = 0

	// Komponen minat (bobot 40%)
	bobotMinat = 40.0 / float64(k.JmlReqMinat)
	for i = 0; i < k.JmlReqMinat; i++ {
		maksPoin = maksPoin + bobotMinat
		for j = 0; j < user.JumlahMinat; j++ {
			if user.Minat[j] == k.ReqMinat[i] {
				totalPoin = totalPoin + bobotMinat
			}
		}
	}

	// Komponen keahlian (bobot 60%)
	bobotKeahlian = 60.0 / float64(k.JmlReqKeahlian)
	for i = 0; i < k.JmlReqKeahlian; i++ {
		maksPoin = maksPoin + bobotKeahlian
		for j = 0; j < user.JumlahKeahlian; j++ {
			if user.Keahlian[j] == k.ReqKeahlian[i] {
				totalPoin = totalPoin + bobotKeahlian
			}
		}
	}

	if maksPoin == 0 {
		return 0
	}
	k.SkorCocok = (totalPoin / maksPoin) * 100
	return k.SkorCocok
}

func hitungSemuaSkor(user *User) {
	var i int
	for i = 0; i < jumlahKarier; i++ {
		hitungSkor(user, &daftarKarier[i])
	}
}

func tampilRekomendasi(daftar [MAX_KARIER]Karier, jumlah int) {
	var i int
	var bar string

	if jumlah == 0 {
		fmt.Println("  [!] Tidak ada rekomendasi.")
		return
	}
	fmt.Println()
	fmt.Println("  No   Nama Karier         Industri      Gaji(jt)  Kecocokan")
	cetakGaris()
	for i = 0; i < jumlah; i++ {
		bar = buatBar(daftar[i].SkorCocok, 8)
		fmt.Printf("  %-4d %-19s %-13s %-9d %s %.1f%%\n",
			i+1, daftar[i].Nama, daftar[i].Industri,
			daftar[i].GajiRata, bar, daftar[i].SkorCocok)
	}
}

func tampilStatistik(user *User) {
	var i int
	var total float64
	var maks float64
	var namaMaks string
	var rataRata float64
	var bar string

	hitungSemuaSkor(user)

	total = 0
	maks = 0
	namaMaks = ""

	for i = 0; i < jumlahKarier; i++ {
		total = total + daftarKarier[i].SkorCocok
		if daftarKarier[i].SkorCocok > maks {
			maks = daftarKarier[i].SkorCocok
			namaMaks = daftarKarier[i].Nama
		}
	}

	rataRata = total / float64(jumlahKarier)

	fmt.Println()
	fmt.Println("  +==========================================+")
	fmt.Println("  |   📈 STATISTIK KECOCOKAN KARIER          |")
	fmt.Println("  +==========================================+")
	for i = 0; i < jumlahKarier; i++ {
		bar = buatBar(daftarKarier[i].SkorCocok, 12)
		fmt.Printf("  |  %-18s %s %.1f%%\n",
			daftarKarier[i].Nama, bar, daftarKarier[i].SkorCocok)
	}
	fmt.Println("  +==========================================+")
	fmt.Printf("  |  📊 Rata-rata kecocokan  : %.1f%%\n", rataRata)
	fmt.Printf("  |  🏆 Karier terbaik       : %s (%.1f%%)\n", namaMaks, maks)
	fmt.Printf("  |  🗂  Total karier dianalisis: %d\n", jumlahKarier)
	fmt.Println("  +==========================================+")
}

// ============================================================
// MODUL 3: PENCARIAN (Spesifikasi c)
// ============================================================

// sequentialSearch mencari karier berdasarkan nama atau industri.
// Memeriksa satu per satu dari indeks 0 sampai akhir (tidak perlu terurut).
func sequentialSearch(query string) {
	var i int
	var ditemukan bool

	fmt.Println()
	fmt.Printf("  Hasil pencarian untuk '%s':\n", query)
	cetakGaris()

	ditemukan = false
	for i = 0; i < jumlahKarier; i++ {
		if ada(daftarKarier[i].Nama, query) || ada(daftarKarier[i].Industri, query) {
			fmt.Printf("  - %-20s | Industri: %-12s | Gaji: %d jt\n",
				daftarKarier[i].Nama, daftarKarier[i].Industri, daftarKarier[i].GajiRata)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("  ❌ Tidak ada karier yang cocok.")
	}
}

// urutAlfabet membuat salinan daftarKarier yang sudah diurutkan A-Z by nama.
// Menggunakan Insertion Sort karena data harus terurut sebelum Binary Search.
func urutAlfabet() [MAX_KARIER]Karier {
	var salinan [MAX_KARIER]Karier
	var i int
	var j int
	var kunci Karier

	for i = 0; i < jumlahKarier; i++ {
		salinan[i] = daftarKarier[i]
	}
	// Insertion Sort berdasarkan nama (ascending A-Z)
	for i = 1; i < jumlahKarier; i++ {
		kunci = salinan[i]
		j = i - 1
		for j >= 0 && salinan[j].Nama > kunci.Nama {
			salinan[j+1] = salinan[j]
			j--
		}
		salinan[j+1] = kunci
	}
	return salinan
}

// binarySearch mencari karier berdasarkan nama tepat (exact match).
// Prasyarat: data sudah terurut A-Z. Cara kerja: bagi dua ruang pencarian setiap langkah.
func binarySearch(terurut [MAX_KARIER]Karier, target string) {
	var kiri int
	var kanan int
	var tengah int
	var hasil int
	var i int
	var k Karier

	kiri = 0
	kanan = jumlahKarier - 1
	hasil = -1

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2
		if terurut[tengah].Nama == target {
			hasil = tengah
			kiri = kanan + 1 // paksa keluar dari loop
		} else if terurut[tengah].Nama < target {
			kiri = tengah + 1 // target ada di sisi kanan
		} else {
			kanan = tengah - 1 // target ada di sisi kiri
		}
	}

	fmt.Println()
	if hasil != -1 {
		k = terurut[hasil]
		fmt.Println("  ✅ Karier ditemukan!")
		fmt.Println("  Nama      : " + k.Nama)
		fmt.Println("  Industri  : " + k.Industri)
		fmt.Printf("  Gaji rata : %d juta/bulan\n", k.GajiRata)
		fmt.Println("  Syarat keahlian:")
		for i = 0; i < k.JmlReqKeahlian; i++ {
			fmt.Println("    - " + k.ReqKeahlian[i])
		}
	} else {
		fmt.Println("  ❌ Karier '" + target + "' tidak ditemukan.")
	}
}

func menuPencarian() {
	var pilihan int
	var query string
	var target string
	var terurut [MAX_KARIER]Karier
	var i int

	for {
		cetakHeader("CARI KARIER")
		fmt.Println("  1. 🔎 Sequential Search (nama/industri)")
		fmt.Println("  2. ⚡ Binary Search (nama tepat)")
		fmt.Println("  0. 🔙 Kembali")
		cetakGaris()
		pilihan = inputInt("  Pilihan: ")
		switch pilihan {
		case 1:
			fmt.Print("  Kata kunci: ")
			fmt.Scan(&query)
			sequentialSearch(query)
		case 2:
			// Urutkan dulu, tampilkan daftar, BARU minta input
			terurut = urutAlfabet()
			fmt.Println()
			fmt.Println("  Daftar karier (A-Z):")
			cetakGaris()
			for i = 0; i < jumlahKarier; i++ {
				fmt.Printf("  %2d. %s\n", i+1, terurut[i].Nama)
			}
			cetakGaris()
			fmt.Println("  Contoh nama tepat: Data_Scientist, Akuntan, Jurnalis")
			fmt.Print("  Masukkan nama karier: ")
			fmt.Scan(&target)
			binarySearch(terurut, target)
		case 0:
			return
		default:
			fmt.Println("  ⚠️  Pilihan tidak valid.")
		}
	}
}

// ============================================================
// MODUL 4: PENGURUTAN (Spesifikasi d)
// ============================================================

// selectionSortGaji mengurutkan karier berdasarkan gaji tertinggi ke terendah.
// Cara kerja: cari nilai maksimum di sisa array, tukar ke posisi saat ini.
func selectionSortGaji(daftar [MAX_KARIER]Karier, n int) [MAX_KARIER]Karier {
	var i int
	var j int
	var idxMaks int

	for i = 0; i < n-1; i++ {
		idxMaks = i
		for j = i + 1; j < n; j++ {
			if daftar[j].GajiRata > daftar[idxMaks].GajiRata {
				idxMaks = j
			}
		}
		// Tukar posisi i dengan posisi maksimum
		daftar[i], daftar[idxMaks] = daftar[idxMaks], daftar[i]
	}
	return daftar
}

// insertionSortSkor mengurutkan karier berdasarkan skor kecocokan tertinggi ke terendah.
// Cara kerja: ambil satu elemen, sisipkan ke posisi yang benar di bagian kiri yang sudah terurut.
func insertionSortSkor(daftar [MAX_KARIER]Karier, n int) [MAX_KARIER]Karier {
	var i int
	var j int
	var kunci Karier

	for i = 1; i < n; i++ {
		kunci = daftar[i]
		j = i - 1
		// Geser elemen yang lebih kecil dari kunci ke kanan
		for j >= 0 && daftar[j].SkorCocok < kunci.SkorCocok {
			daftar[j+1] = daftar[j]
			j--
		}
		// Sisipkan kunci ke posisi yang tepat
		daftar[j+1] = kunci
	}
	return daftar
}

func menuPengurutan(user *User) {
	var pilihan int
	var temp [MAX_KARIER]Karier
	var terurut [MAX_KARIER]Karier
	var i int

	for {
		cetakHeader("URUTKAN REKOMENDASI")
		fmt.Println("  1. 💰 Selection Sort  - berdasarkan gaji (tertinggi)")
		fmt.Println("  2. 🎯 Insertion Sort  - berdasarkan kecocokan (tertinggi)")
		fmt.Println("  0. 🔙 Kembali")
		cetakGaris()
		pilihan = inputInt("  Pilihan: ")
		switch pilihan {
		case 1:
			hitungSemuaSkor(user)
			for i = 0; i < jumlahKarier; i++ {
				temp[i] = daftarKarier[i]
			}
			terurut = selectionSortGaji(temp, jumlahKarier)
			fmt.Println()
			fmt.Println("  Diurutkan berdasarkan gaji (Selection Sort):")
			tampilRekomendasi(terurut, jumlahKarier)
		case 2:
			hitungSemuaSkor(user)
			for i = 0; i < jumlahKarier; i++ {
				temp[i] = daftarKarier[i]
			}
			terurut = insertionSortSkor(temp, jumlahKarier)
			fmt.Println()
			fmt.Println("  Diurutkan berdasarkan kecocokan (Insertion Sort):")
			tampilRekomendasi(terurut, jumlahKarier)
		case 0:
			return
		default:
			fmt.Println("  ⚠️  Pilihan tidak valid.")
		}
	}
}

// ============================================================
// MAIN
// ============================================================

func main() {
	var user User
	var pilihan int
	var temp [MAX_KARIER]Karier
	var terurut [MAX_KARIER]Karier
	var i int

	initData()

	fmt.Println()
	fmt.Println("  ╔══════════════════════════════════════════╗")
	fmt.Println("  ║   🎯  REKOMENDASI KARIER                 ║")
	fmt.Println("  ║       Temukan karier impianmu!           ║")
	fmt.Println("  ╚══════════════════════════════════════════╝")
	fmt.Println()
	fmt.Print("  👤 Masukkan nama Anda: ")
	fmt.Scan(&user.Nama)

	fmt.Println()
	fmt.Println("  ✨ Selamat datang, " + user.Nama + "!")
	fmt.Println("  📋 Silakan lengkapi profil Anda di menu 1.")

	for {
		cetakHeader("MENU UTAMA")
		fmt.Println("  1. 👤 Kelola Minat & Keahlian")
		fmt.Println("  2. 💼 Lihat Rekomendasi Karier")
		fmt.Println("  3. 🔍 Cari Karier")
		fmt.Println("  4. 📊 Urutkan Rekomendasi")
		fmt.Println("  5. 📈 Statistik Kecocokan")
		fmt.Println("  0. 🚪 Keluar")
		cetakGaris()
		pilihan = inputInt("  Pilihan: ")

		switch pilihan {
		case 1:
			menuKelolaProfil(&user)
		case 2:
			cetakHeader("REKOMENDASI KARIER")
			hitungSemuaSkor(&user)
			for i = 0; i < jumlahKarier; i++ {
				temp[i] = daftarKarier[i]
			}
			terurut = insertionSortSkor(temp, jumlahKarier)
			tampilRekomendasi(terurut, jumlahKarier)
		case 3:
			menuPencarian()
		case 4:
			menuPengurutan(&user)
		case 5:
			cetakHeader("STATISTIK KECOCOKAN")
			tampilStatistik(&user)
		case 0:
			fmt.Println()
			fmt.Println("  🌟 Terima kasih, semoga sukses dalam perjalanan karier Anda!")
			fmt.Println()
			return
		default:
			fmt.Println("  ⚠️  Pilihan tidak valid.")
		}
	}
}
	