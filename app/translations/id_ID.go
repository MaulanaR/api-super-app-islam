package translations

import "grest.dev/grest"

func AddIdID() {
	grest.AddTranslation("id-ID", map[string]string{
		"400_bad_request":                     "Permintaan tidak dapat dilakukan karena ada parameter yang salah atau tidak lengkap.",
		"401_unauthorized":                    "Token otentikasi tidak valid. Silakan logout dan login ulang",
		"403_forbidden":                       "Pengguna tidak memiliki izin untuk :action.",
		"404_not_found":                       "The resource you have specified cannot be found.",
		"500_internal_error":                  "Gagal terhubung ke server, silakan coba lagi nanti.",
		"deleted":                             "Data berhasil dihapus.",
		"not_found":                           "Data :entity dengan :key = :value tidak ditemukan.",
		"greater_than":                        "Nilai data :key harus lebih besar dari :value",
		"less_than":                           "Nilai data :key harus lebih kecil dari :value",
		"not_in":                              "Ekstensi file harus salah satu dari: (:value).",
		"storage_limit":                       "Tidak dapat mengunggah file karena batas penyimpanan Anda. Silakan tingkatkan batas penyimpanan.",
		"required_value":                      ":key wajib diisi.",
		"unique":                              ":attribute (:value) sudah ada dan tidak bisa digunakan lagi.",
		"success":                             "Sukses.",
		"attendance_is_exists":                "Kehadiran untuk hari yang dipilih sudah ada.",
		"invalid_operating_hour":              "Jam Kerja tidak valid",
		"invalid_operating_day":               "Hari Kerja tidak valid",
		"invalid_username_or_password":        "Username atau kata sandi tidak valid",
		"insufficient_timeoff_balance":        "Saldo cuti tidak mencukupi",
		"id":                                  "id",
		"code":                                "kode",
		"branches":                            "data cabang",
		"employees":                           "data karyawan",
		"data_stores.departments.create":      "membuat data departemen",
		"timeschedules_in_used_on_attendance": "Schedule ini sudah digunakan pada data absensi",
		"invalid_data_type":                   "Tipe data yang dikirimkan tidak valid atau karakter terlalu panjang. Silakan cek kembali parameter yang dikirimkan.",
	})
}
