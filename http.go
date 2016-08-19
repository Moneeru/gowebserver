package main

import (
    "io"
    "fmt"
    "log"
    "net/http"
    "strconv"
)

func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello world!")
}

func to_roman(n int)  string {
    if n == 2 {
        return "II"
    }
    return "I"
}

func to_roman(n int)  string {
    if n == 3 {
        return "III"
    }
    return "I"
    
    func to_roman(n int)  string {
    if n == 4 {
        return "IV"
    }
    return "I"
}

func to_roman(n int)  string {
    if n == 5 {
        return "V"
    }
    return "I"
}

func to_roman(n int)  string {
    if n == 6 {
        return "VI"
    }
    return "I"
}

func to_roman(n int)  string {
    if n == 7 {
        return "VII"
    }
    return "I"
}

func to_roman(n int)  string {
    if n == 8 {
        return "VIII"
    }
    return "I"
}

func to_roman(n int)  string {
    if n == 9 {
        return "IX"
    }
    return "I"
}
    
    func to_roman(n int)  string {
    if n == 10 {
        return "X"
    }
    return "I"
}
    
}
type romanGenerator int
func (n romanGenerator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    ascii_num := r.URL.Path[7:]
    i, err := strconv.Atoi(ascii_num)
    if err != nil {
        log.Print(err)
    }
    fmt.Fprintf(w, "Here's your number: %s\n", to_roman(i))
}



func main() {
    h := http.NewServeMux()

    h.Handle("/roman/", romanGenerator(1))
    h.HandleFunc("/", hello)

    err := http.ListenAndServe(":8000", h)
    log.Fatal(err)
}
