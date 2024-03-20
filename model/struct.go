package model

type Buku struct {
	IDBuku      int    `gorm:"primaryKey;column:id_buku" json:"-"`
	Judul       string `gorm:"column:judul" json:"judul"`
	Pengarang   string `gorm:"column:pengarang" json:"pengarang"`
	TahunTerbit string `gorm:"column:tahun_terbit" json:"tahun_terbit"`
	ISBN        string `gorm:"column:isbn" json:"isbn"`
}

type Anggota struct {
	IDAnggota   int    `gorm:"primaryKey;column:id_anggota" json:"-"`
	Nama        string `gorm:"column:nama" json:"nama"`
	Alamat      string `gorm:"column:alamat" json:"alamat"`
	NoTelepon   string `gorm:"column:no_telepon" json:"no_telepon"`
	NoIdentitas string `gorm:"column:no_identitas" json:"no_identitas"`
}

type Peminjaman struct {
	IDPeminjaman     int    `gorm:"primaryKey;column:id_peminjaman" json:"-"`
	IDAnggota        int    `gorm:"column:id_anggota" json:"id_anggota"`
	TglPeminjaman    string `gorm:"column:tgl_peminjaman" json:"tgl_peminjaman"`
	TglPengembalian  string `gorm:"column:tgl_pengembalian" json:"tgl_pengembalian"`
	StatusPeminjaman string `gorm:"column:status_peminjaman" json:"status_peminjaman"`
}

type DetailPeminjaman struct {
	IDDetailPeminjaman int `gorm:"primaryKey;column:id_detail_peminjaman" json:"-"`
	IDPeminjaman       int `gorm:"column:id_peminjaman" json:"id_peminjaman"`
	IDBuku             int `gorm:"column:id_buku" json:"id_buku"`
}

type StokBuku struct {
	IDStok         int `gorm:"primaryKey;column:id_stok" json:"-"`
	IDBuku         int `gorm:"column:id_buku" json:"id_buku"`
	JumlahTersedia int `gorm:"column:jumlah_tersedia" json:"jumlah_tersedia"`
}
