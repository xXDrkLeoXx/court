func Redirect(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, World!\n")
  fmt.Fprintf(w, "%s", r.URL.Path)
}
