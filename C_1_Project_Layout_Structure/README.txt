Project Layout Structure
Ringkasang :
.
├── go.mod
|   # file go.mod dipergunakan oleh go module (jika go mod diaktifkan).
|
├── Makefile
|   # file Makefile dipergunakan oleh command `make`.
|
├── assets/
|   # folder assets berisi static assets, seperti gambar, logo, dll.
|
├── build/
|   # folder build isinya adalah files untuk keperluan build dan
|   # juga CI (continous integration). Contoh file yang dimaksud adalah
|   # seperti Dockerfile, file CI tool (.travis-ci.yml, .gitlab-ci.yml)
|   # dan file untuk keperluan build ke bentuk lain seperti file deb, rpm, pkg.
|   |
│   ├── ci/
|   |   # tempatkan file untuk CI dalam folder ini
|   |
│   └── package/
|       # tempatkan file untuk keperluan build dalam folder ini
|
├── cmd/
|   # folder cmd isinya adalah source code utama aplikasi.
|   #
|   # jika aplikasi merupakan sebuah app monolith, maka folder ini isinya
|   # adalah langsung source code utama.
|   # salah satu contoh, folder ini isinya adalah file-file bisnis logic utama,
|   # seperti services dan repositories.
|   #
|   # jika arsitektur microservices diadopsi, dengan layout monorepo,
|   # maka isi dari cmd adalah source code yang dibagi per service.
|   |
│   ├── your_app_1/
│   ├── your_app_2/
│   ├── your_app_3/
│   └── ...
|
├── configs/
|   # folder configs isinya adalah file konfigurasi.
|
├── deployments/
|   # folder deployments isinya adalah file yang berhubungan dengan orchestration,
|   # deployments, dan juga CD. Seperti docker-compose.yml, k8s file, dll.
|
├── docs/
|   # folder docs isinya adalah file design dan dokumentasi.
|
├── examples/
|   # folder examples isinya adalah file example.
|
├── init/
|   # folder init isinya adalah file-file system init (systemd, upstart, sysv)
|   # dan file konfigurasi process manager atau supervisor (runit, supervisord).
|
├── internal/
|   # folder internal isinya adalah file private aplikasi dan library.
|   # sebetulnya folder ini kegunaannya sama seperti `pkg`, perbedaannya adalah package
|   # dalam folder internal ini hanya bisa di-import dalam project ini, tidak bisa di-import
|   # ke project lain.
|
├── pkg/
|   # folder pkg isinya adalah file utility yg di-reuse dalam project yang sama,
|   # atau bisa juga di re-use oleh project lain.
|   |
│   ├── your_public_lib_1/
│   ├── your_public_lib_2/
│   ├── your_public_lib_3/
│   └── ...
|
├── test/
|   # folder test isinya adalah file testing. untuk struktur file-nya sendiri bebas,
|   # mau disusun seperti apa.
|   #
|   # khusus untuk unit test, baiknya tidak ditempatkan di sini,
|   # tapi ditempatkan di dalam package yang sama dengan file yang akan di-unit-test.
|
├── vendor/
|   # berisi clone dari 3rd party dependencies. Folder ini digunakan jika konfigurasi vendor diaktifkan
|
├── web/
|   # berisi aplikasi web. untuk microservices saya sarankan untuk menempatkan aplikasi web dalam folder `cmd/app`
|
└── ...
