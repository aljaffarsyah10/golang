// Pakai go run *.go karena run semua file
// Kalau dari cmd pakai go run . jika go run *.go tidak bisa di cmd
// karena * tidak dapat dibaca di cmd
// untuk testing pakai
// curl -X GET --user batman:secret http://localhost:9000/student
// curl -X GET --user batman:secret http://localhost:9000/student?id=s001

package main

import "net/http"
import "fmt"
import "encoding/json"

func main() {
    mux := new(CustomMux)

    mux.HandleFunc("/student", ActionStudent)
    
    mux.RegisterMiddleware(MiddlewareAuth)
    mux.RegisterMiddleware(MiddlewareAllowOnlyGet)
    
    server := new(http.Server)
    server.Addr = ":9000"
    server.Handler = mux
    
    fmt.Println("server started at localhost:9000")
    server.ListenAndServe()
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
    if id := r.URL.Query().Get("id"); id != "" {
        OutputJSON(w, SelectStudent(id))
        return
    }

    OutputJSON(w, GetStudents())
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
    res, err := json.Marshal(o)
    if err != nil {
        w.Write([]byte(err.Error()))
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(res)
}