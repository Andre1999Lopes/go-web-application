package controllers

import (
	"go-web-application/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("./templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco, errPreco := strconv.ParseFloat(r.FormValue("preco"), 64)
		quantidade, errQtde := strconv.Atoi(r.FormValue("quantidade"))

		if errPreco != nil {
			log.Println("Erro na conversão do preço:", errPreco)
		}

		if errQtde != nil {
			log.Println("Erro na conversão da quantidade:", errQtde)

		}

		models.CriarNovoProduto(nome, descricao, preco, quantidade)
	}
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	produto := models.BuscaProdutoPorId(id)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, errId := strconv.Atoi(r.FormValue("id"))
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco, errPreco := strconv.ParseFloat(r.FormValue("preco"), 64)
		quantidade, errQtde := strconv.Atoi(r.FormValue("quantidade"))

		if errId != nil {
			log.Println("Erro na conversão do id:", errPreco)
		}

		if errPreco != nil {
			log.Println("Erro na conversão do preço:", errPreco)
		}

		if errQtde != nil {
			log.Println("Erro na conversão da quantidade:", errQtde)
		}

		models.AtualizaProduto(id, quantidade, nome, descricao, preco)
		http.Redirect(w, r, "/", 301)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.ExcluirProduto(id)
	http.Redirect(w, r, "/", 301)
}
