package main

//Crie uma struct chamada Playlist com os campos "nome"
//e "músicas". O campo "músicas" deve ser um
//slice de structs, cada uma representando uma música
//com os campos "título", "artista" e "duração".
//Escreva uma função que receba uma Playlist como
//parâmetro e imprima o nome da playlist,
//o tempo total da playlist e as informações de cada
//música.
//Usando as mesmas structs do exercício anterior, escreva
//uma função que receba um título de música
//como parâmetro e retorne uma lista das Playlists
//que possuem
//esse título.

type Musica struct {
	titulo  string
	artista string
	duracao int
}

func encontrarMusicas(playlists []Playlist, nomeMusica string) []Playlist {
	var resultado []Playlist
	for _, playlist := range playlists {
		for _, musica := range playlist.musicas {
			if musica.titulo == nomeMusica {
				resultado = append(resultado, playlist)
			}
		}
	}
	return resultado
}

type Playlist struct {
	nome    string
	musicas []Musica
}

func main() {

}
