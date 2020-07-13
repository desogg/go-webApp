package main


import (
"fmt"
"net/http" //Librería usada para crear un servidor Web
)

func main() {
	startHttpServer()
}

func login(w http.ResponseWriter, r *http.Request)  {
	if err := r.ParseForm(); err!=nil{
		fmt.Printf("ParseForm() err: %v", err)
		return
	}
	name := r.FormValue("nombre")
	pass := r.FormValue("pass")
	fmt.Printf("Nombre: " + name )
	fmt.Printf("Password: " + pass)
	http.Redirect(w,r,"/index.html",http.StatusSeeOther)
}

func startHttpServer() *http.Server {
	flServer := http.FileServer(http.Dir("./Views")) //Permite ver en el sistema de archivos una carpeta que servirá de referencia
	//--------Manejo de rutas ------------------------------------------
	http.Handle("/",flServer) 
	http.HandleFunc("/form",login)
	if err := http.ListenAndServe(":8888", nil); err != nil { //Se inicializa el servidor
		fmt.Printf("ERROR")
    }
	
	return nil
}