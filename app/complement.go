package app

import "grest.dev/grest/cache"

// pake struct biar bisa ditambah sewaktu-waktu tanpa harus error
type Complement struct {
	Ctx       *Ctx        
	EndPoint  string      
	Method    string      
	Model     interface{} 
	Old       interface{} 
	New       interface{} 
	DataIDKey string       // dipake untuk dapetin New pake db.First(tx, model, url.Value{}.Add(DataIDKey, DataID))
	DataID    string      
	DataCode  string      
}

func (f *Complement) Run() {
	f.DeleteCache()
	go f.AddHistory()
}

func (f *Complement) DeleteCache() {
	cache.Delete(f.Ctx.Company.Slug + "." + f.EndPoint + "." + f.DataID)
	cache.Delete(f.Ctx.Company.Slug + "." + f.EndPoint + "." + f.DataCode)
	go cache.DeleteWithPrefix(f.Ctx.Company.Slug + "." + f.EndPoint + "?") // ini agak lama jadi jalan di goroutine aj biar proses yg lain nda perlu nungguin ini
}

func (f *Complement) AddHistory() {
	// lengkapi isi field history nya sesuai context
	// apabila method nya POST, PUT, PATCH, isi New nya diisi dari GetByID nya Model
	// sebelum di proses lebih lanjut Old dan New nya di override, dikonversi menjadi structured
	// simpan ke tabel history
	f.SendWebhook()
}

func (f *Complement) SendWebhook() {
	// cek apakah ada data webhook atas ctx terkait
	// apabila ada, maka kirim ke webhook sesuai dengan Method, Old & New yg sudah dilengkapi (dan dikonversi menjadi structured) dari proses AddHistory
	// jika pada api v2 yg dikirim hanya new atau old nya saja (salah satu, tidak bisa dua-dua nya), pada api v3 untuk PUT & PATCH yg dikirim dua-dua nya (kirim juga old nya)
}
