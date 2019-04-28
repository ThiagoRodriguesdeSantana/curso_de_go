package operadora

import (
	"strconv"

	"github.com/ThiagoRodriguesdeSantana/curso_de_go/pacotes/prefixo"
)

//NomeDaOperadora nome da operadora do estado
var NomeDaOperadora = "xpto telocon"

//PrefixoDaCapitalOperadora concatenação do prefico com a operadora
var PrefixoDaCapitalOperadora = strconv.Itoa(prefixo.Capital) + " " + NomeDaOperadora
