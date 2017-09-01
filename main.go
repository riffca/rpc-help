package main

import "net/http"
import "log"
import "rpc-util/service"
import user "rpc-util/service/user/proto"
import "rpc-util/utils/rpc"

var client = user.NewUserClient(service.UserConnection)

func main() {
	service.ConnectUserService()
	http.HandleFunc("/user", handle)
	http.Handle("/", http.FileServer(http.Dir("./")))
	log.Println("Server started on port: 9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := rpc.DefaultContext()
	defer cancel()
	n, err := client.GetName(ctx, &user.NameRequest{Name: "stas"})
	if err != nil {
		panic(err)
	}
	log.Println(n)
	w.Write([]byte("new school"))
}
