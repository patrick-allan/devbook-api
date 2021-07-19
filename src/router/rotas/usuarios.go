package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:        "/usuarios",
		Metodo:     http.MethodPost,
		Funcao:     controllers.CriarUsuario,
		RequerAuth: false,
	},
	{
		URI:        "/usuarios",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarUsuarios,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{usuarioID}",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarUsuario,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{usuarioID}",
		Metodo:     http.MethodPut,
		Funcao:     controllers.AtualizarUsuario,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{usuarioID}",
		Metodo:     http.MethodDelete,
		Funcao:     controllers.DeletarUsuario,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{usuarioID}/seguir",
		Metodo:     http.MethodPost,
		Funcao:     controllers.SeguiUsuario,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{usuarioID}/parar-de-seguir",
		Metodo:     http.MethodPost,
		Funcao:     controllers.PararDeSeguirUsuario,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{usuarioID}/seguidores",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarSeguidores,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{usuarioID}/seguindo",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarSeguindo,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{usuarioID}/atualizar-senha",
		Metodo:     http.MethodPost,
		Funcao:     controllers.AtualizarSenha,
		RequerAuth: true,
	},
}
