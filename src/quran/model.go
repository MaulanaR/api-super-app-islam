package quran

import (
	"grest.dev/grest/db"
)

const EndPoint = "surah"

type Surah struct {
	db.Model
	Nomor      db.NullInt64  `json:"nomor" db:"s.nomor" gorm:"primaryKey;column:nomor;index:surah_nomor"`
	Arti       db.NullString `json:"arti" db:"s.arti" gorm:"column:arti;index:surah_arti"`
	Asma       db.NullString `json:"asma" db:"s.asma" gorm:"column:asma;index:surah_asma"`
	Ayat       db.NullInt64  `json:"ayat" db:"s.ayat" gorm:"column:ayat;index:surah_ayat"`
	Nama       db.NullString `json:"nama" db:"s.nama" gorm:"column:nama;index:surah_nama"`
	Jenis      db.NullString `json:"type" db:"s.jenis" gorm:"column:jenis"`
	Urut       db.NullString `json:"urut" db:"s.urut" gorm:"column:urut"`
	Audio      db.NullString `json:"audio" db:"s.audio" gorm:"column:audio"`
	Keterangan db.NullText   `json:"keterangan" db:"s.keterangan" gorm:"column:keterangan"`
}

func (Surah) TableVersion() string {
	return "22.02.090949"
}

func (Surah) TableName() string {
	return "surah"
}

func (Surah) TableAliasName() string {
	return "s"
}

func (s *Surah) SetRelation() {
}

func (s *Surah) SetFilter() {
}

func (s *Surah) SetSort() {
}

type Ayat struct {
	db.Model
	NomorSurah db.NullInt64  `json:"surah" db:"s.nomor_surah" gorm:"column:nomor_surah;index:"ayat_nomor_surah""`
	IDAyat     db.NullString `json:"key" db:"s.key" gorm:"column:key;primaryKey;index:"ayat_key""`
	Nomor      db.NullInt64  `json:"nomor" db:"s.ayat" gorm:"column:ayat;index:"ayat_ayat""`
	Id         db.NullText   `json:"id" db:"s.indonesia" gorm:"column:indonesia"`
	Ar         db.NullText   `json:"ar" db:"s.arabic" gorm:"column:arabic"`
	Latih      db.NullText   `json:"tr" db:"s.latin" gorm:"column:latin"`
}

func (Ayat) TableVersion() string {
	return "22.02.090949"
}

func (Ayat) TableName() string {
	return "ayat"
}

func (Ayat) TableAliasName() string {
	return "s"
}

func (s *Ayat) SetRelation() {
}

func (s *Ayat) SetFilter() {
}

func (s *Ayat) SetSort() {
}

type PerawiHadits struct {
	db.Model
	Kode   db.NullString `json:"kode" db:"s.kode" gorm:"column:kode;primaryKey;index:perawi_kode"`
	Perawi db.NullString `json:"perawi" db:"s.perawi" gorm:"column:perawi;index:perawi_name"`
}

func (PerawiHadits) TableVersion() string {
	return "22.02.090949"
}

func (PerawiHadits) TableName() string {
	return "perawi_hadits"
}

func (PerawiHadits) TableAliasName() string {
	return "s"
}

func (s *PerawiHadits) SetRelation() {
}

func (s *PerawiHadits) SetFilter() {
}

func (s *PerawiHadits) SetSort() {
}

type Hadits struct {
	db.Model
	Kode       db.NullString `json:"kode" db:"s.kode gorm:"column:kode;primaryKey;index:hadits_kode"`
	KodePerawi db.NullString `json:"kode_perawi" db:"s.kode_perawi" gorm:"column:kode_perawi;index:hadits_perawi_kode"`
	Perawi     db.NullString `json:"perawi" db:"i.perawi" gorm:"-"`
	Nomor      db.NullInt64  `json:"number" db:"s.nomor" gorm:"column:nomor"`
	Id         db.NullText   `json:"id" db:"s.indonesia" gorm:"column:indonesia"`
	Ar         db.NullText   `json:"arab" db:"s.arabic" gorm:"column:arabic"`
}

func (Hadits) TableVersion() string {
	return "22.02.090949"
}

func (Hadits) TableName() string {
	return "hadits"
}

func (Hadits) TableAliasName() string {
	return "s"
}

func (s *Hadits) SetRelation() {
	s.Relation = append(s.Relation, db.NewRelation("left", "perawi_hadits", "i", []db.Filter{{Column: "i.kode", Column2: "s.kode_perawi"}}))
}

func (s *Hadits) SetFilter() {
}

func (s *Hadits) SetSort() {
}

type Dzikir struct {
	db.Model
	Id        db.NullInt64  `json:"id" db:"s.id" gorm:"column:id;primaryKey;index:dzikir_id"`
	Arabic    db.NullText   `json:"arabic" db:"s.arabic" gorm:"column:arabic"`
	Latin     db.NullText   `json:"latin" db:"s.latin" gorm:"column:latin"`
	Faedah    db.NullText   `json:"faedah" db:"s.faedah" gorm:"column:faedah"`
	Narrator  db.NullText   `json:"narrator" db:"s.narrator" gorm:"column:narrator"`
	Note      db.NullText   `json:"note" db:"s.note" gorm:"column:note"`
	Title     db.NullText   `json:"title" db:"s.title" gorm:"column:title"`
	Indonesia db.NullText   `json:"indonesia" db:"s.indonesia" gorm:"column:indonesia"`
	Time      db.NullString `json:"time" db:"s.time" gorm:"column:time"`
}

func (Dzikir) TableVersion() string {
	return "22.02.090949"
}

func (Dzikir) TableName() string {
	return "dzikir"
}

func (Dzikir) TableAliasName() string {
	return "s"
}

func (s *Dzikir) SetRelation() {
}

func (s *Dzikir) SetFilter() {
}

func (s *Dzikir) SetSort() {
}

type Doa struct {
	db.Model
	Id        db.NullInt64 `json:"id" db:"s.id" gorm:"column:id;primaryKey;index:doa_id"`
	Doa       db.NullText  `json:"doa" db:"s.doa" gorm:"column:doa;index:doa_doa"`
	Arabic    db.NullText  `json:"arabic" db:"s.arabic" gorm:"column:arabic"`
	Latin     db.NullText  `json:"latin" db:"s.latin" gorm:"column:latin"`
	Indonesia db.NullText  `json:"indonesia" db:"s.indonesia" gorm:"column:indonesia"`
}

func (Doa) TableVersion() string {
	return "22.02.090949"
}

func (Doa) TableName() string {
	return "doa"
}

func (Doa) TableAliasName() string {
	return "s"
}

func (s *Doa) SetRelation() {
}

func (s *Doa) SetFilter() {
}

func (s *Doa) SetSort() {
}
