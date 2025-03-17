## Calculate Sum of Squares with Concurrency

Buatlah fungsi di Golang yang menghitung jumlah kuadrat dari angka 1 hingga n secara concurrent menggunakan goroutine, channel, dan select statement.

### Contoh:
Jika n = 5, maka hasilnya harus:
1² + 2² + 3² + 4² + 5² = 55

### Solusi:
1. Fungsi sumOfSquares:

- Fungsi ini menerima nilai n, sebuah channel squares untuk mengirim hasil kuadrat, dan channel quit untuk sinyal berhenti.
- Menggunakan loop for dengan select statement, fungsi mengirim nilai i * i ke channel squares setiap iterasi.

2. Penggunaan Select:

- Pada setiap iterasi, select memilih kasus yang siap.
- Setelah mengirim kuadrat, nilai i dinaikkan. Jika i melebihi n, fungsi mengirim sinyal ke channel quit dan mengakhiri eksekusi.

3. Fungsi main:

- Mendefinisikan n dan membuat channel squares dan quit.
- Memulai goroutine sumOfSquares.
- Melalui loop, fungsi utama menerima tepat n nilai dari squares dan menjumlahkannya.
- Terakhir, main menunggu sinyal dari channel quit sebelum mencetak hasil.
