# API Super Islamic App (Indonesia & Arabic Only)

This api was created to handle super Islamic apps which include :

- [x] Al-Qur'an
- [x] Hadits 9 Imam (Muslim Bukhari, Tirmidzi, Nasai, Abu Daud, Ibnu Majah, Ahmad, Darimi, Malik)
- [x] Dzikir Pagi/Petang
- [x] Doa Harian

## Installation

make sure golang is installed

```bash
go run main.go
```

## Usage
1. Call Migration
```bash
GET {{localhost}}/api/migration
```
2. Call Seeder Data
```bash
GET {{localhost}}/api/seed/surah
```
```bash
GET {{localhost}}/api/seed/ayat
```
```bash
GET {{localhost}}/api/seed/hadits
```
```bash
GET {{localhost}}/api/seed/dzikir
```
```bash
GET {{localhost}}/api/seed/doa
```
## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License


## Data Source 
1. https://github.com/muslim-dev/dzikrr
2. https://doa-doa-api-ahmadramadhan.fly.dev/
3. https://github.com/gadingnst/hadith-api
4. https://github.com/bachors/Al-Quran-ID-API
