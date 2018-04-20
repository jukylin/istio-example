package main

import(
	"net/http"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
	"net"
)

func sleep(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	sleep := r.Form.Get("sleep")

	t1 := time.Now()
	resp, err := http.Get("http://sleep:8080?sleep=" + sleep)
	elapsed := time.Since(t1)
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
	w.Write([]byte("\nresp.StatusCode " + strconv.Itoa(resp.StatusCode)))
	w.Write([]byte("\n间隔 " + elapsed.String()))
	w.Write([]byte("\n ip " + getPrivateIPIfAvailable().String()))
	w.Write([]byte("\n"))
}

func getPrivateIPIfAvailable() net.IP {
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		default:
			continue
		}
		if !ip.IsLoopback() {
			return ip
		}
	}
	return net.IPv4zero
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
	w.Write([]byte("\nresp.StatusCode" + strconv.Itoa(resp.StatusCode)))
	w.Write([]byte("\n ip " + getPrivateIPIfAvailable().String()))
	w.Write([]byte("\n"))
}


func retry(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	status := r.Form.Get("status")

	t1 := time.Now()
	resp, err := http.Get("http://sleep:8080/retry?status=" + status)
	elapsed := time.Since(t1)
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
	w.Write([]byte("\nresp.StatusCode " + strconv.Itoa(resp.StatusCode)))
	w.Write([]byte("\n间隔 " + elapsed.String()))
	w.Write([]byte("\n"))

}


func main()  {
	mux := http.NewServeMux()

	mux.HandleFunc("/sleep", sleep)
	mux.HandleFunc("/retry", retry)
	mux.HandleFunc("/deploy", deploy)
	http.ListenAndServe(":8000", mux)
}
