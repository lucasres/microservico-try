package main

import (
	"encoding/json"
	"log"
	"fmt"
	"net/http"
	"strings"
	"strconv"
)

var MAGICK_NUMBER = int64(21)

type Result struct {
	Status string
}

type RequestHome struct {
	CcNumber string
	CcName string
	Coupon string
}

func Home(w http.ResponseWriter, r *http.Request)  {
	data := RequestHome{
		CcNumber: r.FormValue("ccNumber"),
		CcName: r.FormValue("ccName"),
		Coupon: r.FormValue("coupon"),
	}
	sum := int64(0)
	//calcula a soma dos numeros
	for _, current := range strings.Split(data.CcNumber, "") {
		number, err := strconv.ParseInt(current, 0, 64)
		if(err != nil){
			log.Fatal("Parse string to in fail", err)
		}
		sum = sum + number
	}
	//por padrao nega a venda
	result := Result{
		Status: "denied",
	}
	//processamento fake eh a soma do cc ser igual a 21
	if (sum == MAGICK_NUMBER) {
		result.Status = "aproveed"
	}
	//serializa a resposta
	j, err := json.Marshal(result)
	if (err != nil) {
		log.Fatal("Error parse response", err)
	}
	//responde
	//imprime alguns logs
	fmt.Printf("COMPRADOR: %s\nCCNUMBER: %s\nCOUPON: %s\n", data.CcName, data.CcNumber, data.Coupon)
	fmt.Printf("APROVADO: %d == %d\n", sum, MAGICK_NUMBER)
	fmt.Fprintf(w, string(j))
}

func main()  {
	http.HandleFunc("/", Home)
	http.ListenAndServe(":9001", nil)
}