package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"github.com/fclairamb/ftpserver/server"
)

func main() {
	options := &server.ServerOpts{
		Factory: &myFSFactory{},
		Port:    2121, // Port yang digunakan oleh server FTP
	}

	ftpServer := server.NewServer(options)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", options.Port))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	defer ftpServer.Shutdown()

	fmt.Println("FTP server berjalan di port 2121")
	if err := ftpServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// Definisikan myFSFactory sebagai custom filesystem factory
type myFSFactory struct{}

// Implementasikan metode CreateDriver untuk membuat driver FTP
func (f *myFSFactory) CreateDriver() (server.Driver, error) {
	return &myDriver{}, nil
}

// Definisikan myDriver sebagai custom driver FTP
type myDriver struct{}

// Implementasikan metode Init untuk inisialisasi driver
func (d *myDriver) Init(conn *server.Conn) {
}

// Implementasikan metode Stat untuk mendapatkan informasi tentang file atau folder
func (d *myDriver) Stat(path string) (server.FileInfo, error) {
	return nil, fmt.Errorf("File tidak ditemukan")
}

// Implementasikan metode ListDir untuk mendapatkan daftar isi folder
func (d *myDriver) ListDir(path string, callback func(server.FileInfo) error) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		callback(file)
	}

	return nil
}

// Implementasikan metode ChangeDir untuk mengubah direktori
func (d *myDriver) ChangeDir(path string) error {
	return nil
}

// Implementasikan metode DeleteDir untuk menghapus direktori
func (d *myDriver) DeleteDir(path string) error {
	return fmt.Errorf("Tidak diizinkan menghapus direktori")
}

// Implementasikan metode DeleteFile untuk menghapus file
func (d *myDriver) DeleteFile(path string) error {
	return fmt.Errorf("Tidak diizinkan menghapus file")
}

// Implementasikan metode Rename untuk mengganti nama file atau direktori
func (d *myDriver) Rename(oldPath, newPath string) error {
	return fmt.Errorf("Tidak diizinkan mengganti nama")
}

// Implementasikan metode MakeDir untuk membuat direktori baru
func (d *myDriver) MakeDir(path string) error {
	return fmt.Errorf("Tidak diizinkan membuat direktori baru")
}

// Implementasikan metode GetFile untuk mendapatkan file dari server
func (d *myDriver) GetFile(path string, offset int64) (int64, io.ReadCloser, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, nil, err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		file.Close()
		return 0, nil, err
	}

	fileSize := fileInfo.Size()
	file.Seek(offset, 0)
	return fileSize - offset, file, nil
}

// Implementasikan metode PutFile untuk mengirim file ke server
func (d *myDriver) PutFile(destPath string, data io.Reader, appendData bool) (int64, error) {
	file, err := os.Create(destPath)
	if err != nil {
		return 0, err
	}

	defer file.Close()
	return io.Copy(file, data)
}
