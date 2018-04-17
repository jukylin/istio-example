package main

import(
	"net/http"
	"fmt"
	"io/ioutil"
	"strconv"
)

func sleep(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	sleep := r.Form.Get("sleep")
	resp, err := http.Get("http://sleep:8080?sleep=" + sleep)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println("resp.StatusCode", resp.StatusCode)

	if err != nil {
		fmt.Println(err.Error())
	}

	w.Write([]byte("sleep:" + string(body)))
	w.Write([]byte("resp.StatusCode" + strconv.Itoa(resp.StatusCode)))
}


func deploy(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	resp, err := http.Get("http://deploy:8080")

	if err != nil {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err.Error())
	}

	w.Write([]byte("deploy:" + string(body)))
	w.Write([]byte("resp.StatusCode" + strconv.Itoa(resp.StatusCode)))
}



func main()  {
	mux := http.NewServeMux()

	mux.HandleFunc("/sleep", sleep)
	mux.HandleFunc("/deploy", deploy)
	http.ListenAndServe(":8000", mux)
}
