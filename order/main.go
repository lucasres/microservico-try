package main

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"log"
	"net/http"
	"text/template"
)

var URI_MAKE_PAYMENT = "http://localhost:9001"

//estrutura de dados para armazenar o retorna de payment
type Result struct {
	Status string
}

type RequestPaymentService struct {
	CcNumber string
	CcName string
	Coupon string
}

func MakeRequestPaymentService(r RequestPaymentService) Result {
	//adiciona os dados para irem na request
	values := url.Values{}
	values.Add("ccNumber", r.CcNumber)
	values.Add("ccName", r.CcName)
	values.Add("coupon", r.Coupon)
	//faz a request
	response, err := http.PostForm(URI_MAKE_PAYMENT, values)
	if (err != nil) {
		log.Fatal("Erro in request", err)
	}
	//fecha a conexao
	defer response.Body.Close()
	//ler o retorno para dentro de data
	data, err := ioutil.ReadAll(response.Body)
	if (err != nil) {
		log.Fatal("Erro in read response", err)
	} 
	//faz a deserializacao dos dados
	result := Result{}
	json.Unmarshal(data, &result)
	return result
}

func Home(w http.ResponseWriter, r *http.Request) {
	//cria um template
	t := template.Must(template.ParseFiles("templates/home.html"))
	//executa a criação do template
	err := t.Execute(w, Result{})
	if(err != nil) {
		log.Fatal("Erro in parse html", err)
	}
}

func Finish(w http.ResponseWriter, r *http.Request) {
	//pega os dados da request
	request := RequestPaymentService{
		CcName: r.FormValue("ccName"),
		CcNumber: r.FormValue("ccNumber"),
		Coupon: r.FormValue("coupon"),
	}
	result := MakeRequestPaymentService(request)
	//cria um template
	t := template.Must(template.ParseFiles("templates/home.html"))
	//executa a criação do template
	err := t.Execute(w, result)
	if(err != nil) {
		log.Fatal("Erro in parse html", err)
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/finish", Finish)
	http.ListenAndServe(":9000", nil)
}