package model

import "errors"

// Imovel representa um imovel
type Imovel struct {
	Nome  string `json:"nome"`
	X     int    `json:"coodenada_X"`
	Y     int    `json:"coodenada_Y"`
	valor int
}

//ErrValorDeImovelInvalido valor invalido
var ErrValorDeImovelInvalido = errors.New("Valor informado inválido")

//ErrValorMuitoAlto valor muito alto
var ErrValorMuitoAlto = errors.New("Valor informado muito Alto")

//SetValor definine o valor do Imovel
func (i *Imovel) SetValor(valor int) (err error) {

	err = nil
	if valor < 50000 {
		err = ErrValorDeImovelInvalido
		return
	} else if valor > 1000000 {
		err = ErrValorMuitoAlto
		return
	}

	i.valor = valor
	return
}

//GetValor retorna o valor do Imovel
func (i *Imovel) GetValor() int {
	return i.valor
}
