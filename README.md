# test-sv-fullstack
## _Aplikasi Web Fullstack Manajemen Artikel Menggunakan Go dan Vuejs_

Aplikasi manajemen post artikel berbasis web.
- Frontend: VueJS
- Backend (RESTful API): Go
- DBMS MySQL
- Dilengkapi dengan Postman Collection

## Fitur

- Membuat artikel baru. Form artikel terdiri dari Title, Content, dan Category
- Mengupdate artikel
- Artikel terdiri dari status "Published", "Draft", dan "Trashed"
- Setiap tabel artikel pada tab, menggunakan pagination
- Melihat preview blog (dengan pagination)


## Instalasi

### 1. Clone Project 
```sh
https://github.com/ahmadfaisalridwan94/test-sv-fullstack.git
```

### 2. Backend
###### Install dependencies
#
untuk mendownload paket-paket yang diperlukan dan menginstalnya, jalankan
```sh
cd test-sv-fullstack/backend
go get
```

###### Konfigurasi database
#
```sh
.env
```
Ubah menyesuikan dengan konfigurasi db
```sh
DB_HOST=127.0.0.1
DB_DRIVER=mysql 
DB_USER=root
DB_PASSWORD=
DB_NAME=article
DB_PORT=3306 #Default mysql port
```

Membuat database dengan menjalankan query berikut
```sh
CREATE DATABASE article
```

Membuat tabel `post` dengan menjalankan query berikut
```sh
CREATE TABLE `posts` (
  `Id` INT unsigned NOT NULL AUTO_INCREMENT,
  `Title` VARCHAR(200) NOT NULL,
  `Content` TEXT NOT NULL,
  `Category` VARCHAR(100) NOT NULL,
  `CreatedAt` TIMESTAMP NOT NULL DEFAULT current_timestamp(),
  `UpdatedAt` TIMESTAMP NOT NULL DEFAULT current_timestamp(),
  `Status` VARCHAR(100) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB
```
Atau dapat juga membuat tabel `post` dengan menggunakan migrate
```sh
migrate -database "mysql://root@tcp(localhost:3306)/article" -path database/migrations up
```
Menjalankan server backend
```sh
go run main.go
```

### 3. Frontend
###### Install dependencies
#
```sh
cd test-sv-fullstack/frontend
npm install
```
###### Konfigurasi base url
#
```sh
frontend/src/services/api/axios.js
```
###### Menjalankan aplikasi frontend
#
```sh
npm run dev
```