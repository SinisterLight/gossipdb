package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/SinisterLight/gossipdb"
	"io/ioutil"
	"net/http"
)

var (
	members   = flag.String("members", "", "comma seperated list of members")
	rpc_port  = flag.Int("rpc_port", 0, "RPC port")
	http_port = flag.Int("http_port", 0, "http port")
	key       string
	value     string
)

type keyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {
	flag.Parse()

	gossipDb, _ := gossipdb.NewGossipDb(*members, *rpc_port, &keyParser{})

	http.HandleFunc("/value", getValueHandler(gossipDb))
	http.HandleFunc("/", clusterHealthHandler(gossipDb))
	http.HandleFunc("/add_key", addKeyHandler(gossipDb))

	fmt.Printf("Listening on :%d\n", *http_port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *http_port), nil); err != nil {
		fmt.Println(err)
	}
}

func addKeyHandler(g *gossipdb.GossipDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		g.Set(body)
	}
}

func clusterHealthHandler(g *gossipdb.GossipDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(g.Members())
	}
}

func getValueHandler(g *gossipdb.GossipDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data keyValue
		val, found := g.Get(r.FormValue("key"))
		if found {
			err := json.Unmarshal(val, &data)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(data.Value)
			json.NewEncoder(w).Encode(data)
		}
	}
}

type keyParser struct{}

func (p *keyParser) ToKey(b []byte) string {
	var data keyValue
	if err := json.NewDecoder(bytes.NewReader(b)).Decode(&data); err != nil {
		fmt.Println(err)
		return ""
	}
	return data.Key
}
