package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	models "userAPI/models/user"
	utilModels "userAPI/services/util/models"
)

// urlViaCep url para via CEP
const urlViaCep = "https://viacep.com.br/ws/"

// IsCPFValido valida se e um CPF valido
func IsCPFValido(cpf string) bool {

	cpf = strings.Replace(cpf, ".", "", -1)
	cpf = strings.Replace(cpf, "-", "", -1)
	if len(cpf) != 11 {
		return false
	}
	var eq bool
	var dig string
	for _, val := range cpf {
		if len(dig) == 0 {
			dig = string(val)
		}
		if string(val) == dig {
			eq = true
			continue
		}
		eq = false
		break
	}
	if eq {
		return false
	}

	i := 10
	sum := 0
	for index := 0; index < len(cpf)-2; index++ {
		pos, _ := strconv.Atoi(string(cpf[index]))
		sum += pos * i
		i--
	}

	prod := sum * 10
	mod := prod % 11
	if mod == 10 {
		mod = 0
	}
	digit1, _ := strconv.Atoi(string(cpf[9]))
	if mod != digit1 {
		return false
	}
	i = 11
	sum = 0
	for index := 0; index < len(cpf)-1; index++ {
		pos, _ := strconv.Atoi(string(cpf[index]))
		sum += pos * i
		i--
	}
	prod = sum * 10
	mod = prod % 11
	if mod == 10 {
		mod = 0
	}
	digit2, _ := strconv.Atoi(string(cpf[10]))
	if mod != digit2 {
		return false
	}

	return true
}

// IsTelefoneValido verifica se e um telefone valido
func IsTelefoneValido(telefone models.Telefone) (bool, error) {
	ddds := []int{61, 62, 64, 65, 66, 67, 82, 71, 73, 74, 75, 77, 85, 88, 98, 99, 83, 81, 87, 86, 89, 84, 79, 68, 96, 92, 97, 91, 93, 94, 69, 95, 63, 27, 28, 31, 32, 33, 34, 35, 37, 38, 21, 22, 24, 11, 12, 13, 14, 15, 16, 17, 18, 19, 41, 42, 43, 44, 45, 46, 51, 53, 54, 55, 47, 48, 49}

	if telefone.DDD == 0 || !ArrayContains(ddds, telefone.DDD) {

		return false, errors.New("DDD não existe na lista ! ")
	}

	sizeTelefone := len(telefone.Numero)
	isSizeTelefoneValido := sizeTelefone == 9 || sizeTelefone == 8

	if !isSizeTelefoneValido {

		return false, errors.New("Tamanho do numero telefone Invalido ! ")
	}

	return true, nil
}

// ArrayContains verifica se contains o item no array
func ArrayContains(arrayValue interface{}, itemComparacao interface{}) bool {
	arr := reflect.ValueOf(arrayValue)
	if arr.Kind() != reflect.Slice {
		fmt.Print("Tipo de array invalido para a comparação !")
		return false
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == itemComparacao {
			return true
		}
	}

	return false
}

// IsEmailValido verifica se e um e-mail valido
func IsEmailValido(email string) bool {
	regexEmailValido := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if regexEmailValido.MatchString(email) {
		return false
	}
	return true
}

//IsValidaEndereco faz a validação do endereço e a comparação com o valida CEP para ver se esta trazendo o endereço correto
func IsValidaEndereco(endereco *models.Endereco) (bool, error) {
	if strings.TrimSpace(endereco.CEP) == "" || len(strings.TrimSpace(endereco.CEP)) != 8 {
		return false, errors.New("CEP invalido ! ")
	}

	response, erroViaCep := http.Get(urlViaCep + endereco.CEP + "/json")
	if erroViaCep != nil {
		return false, erroViaCep
	}

	data, _ := ioutil.ReadAll(response.Body)

	var viaCep utilModels.ViaCep
	json.Unmarshal(data, &viaCep)

	fmt.Print(string(data))

	if !CompareerStringIgnoreCase(endereco.Bairro, viaCep.Bairro) && viaCep.Bairro != "" {
		return false, errors.New("O bairro informado : " + endereco.Bairro + " Diferente do cadastro no correios : " + viaCep.Bairro)
	}

	if !CompareerStringIgnoreCase(endereco.UF, viaCep.Uf) && viaCep.Uf != "" {
		return false, errors.New("O Estado informado : " + endereco.UF + " Diferente do cadastro no correios : " + viaCep.Uf)
	}

	if !CompareerStringIgnoreCase(endereco.Logradouro, viaCep.Logradouro) && viaCep.Logradouro != "" {
		return false, errors.New("O Logradouro informado : " + endereco.Logradouro + " Diferente do cadastro no correios : " + viaCep.Logradouro)
	}

	if !CompareerStringIgnoreCase(endereco.Cidade, viaCep.Cidade) && viaCep.Cidade != "" {
		return false, errors.New("A cidade  informado : " + endereco.Cidade + " Diferente do cadastro no correios : " + viaCep.Cidade)
	}

	return true, nil
}

// CompareerStringIgnoreCase compara as strings ignorando o upper case
func CompareerStringIgnoreCase(value, compaererValue string) bool {
	if strings.ToUpper(strings.TrimSpace(value)) != strings.ToUpper(strings.TrimSpace(compaererValue)) {
		return false
	}
	return true
}
