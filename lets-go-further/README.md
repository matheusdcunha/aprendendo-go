# Let's Go Further

Essa é a sessão que vai conter os códigos do projeto do livro: 

- **[Let's Go Further](https://lets-go-further.alexedwards.net/)**

---

## O projeto

Esse projeto é uma  aplicação chamada Greenlight — uma API JSON para recuperar e gerenciar informações sobre filmes. Você pode pensar na funcionalidade principal como sendo parecida com a Open Movie Database API.

No final, a API Greenlight suportará os seguintes endpoints e ações:

| Método |	Padrão de URL |	Ação |
| :------: | :--------------: | :-----: |
| GET	| `/v1/healthcheck`	| Exibir informações de saúde e versão da aplicação |
| GET	|`/v1/movies` |	Exibir detalhes de todos os filmes |
| POST	| `/v1/movies` |	Criar um novo filme |
| GET	| `/v1/movies/:id`	| Exibir detalhes de um filme específico |
| PATCH	| `/v1/movies/:id`|	Atualizar detalhes de um filme específico |
| DELETE	| `/v1/movies/:id` |	Deletar um filme específico |
| POST	| `/v1/users` |	Registrar um novo usuário |
| PUT	| `/v1/users/activated` |	Ativar um usuário específico |
| PUT	| `/v1/users/password` | Atualizar a senha de um usuário específico |
| POST	| `/v1/tokens/authentication` |	Gerar um novo token de autenticação |
| POST	| `/v1/tokens/password-reset` |	Gerar um novo token de redefinição de senha |
| GET |	`/debug/vars` |	Exibir métricas da aplicação |

Para dar uma ideia de como a API ficará do ponto de vista do cliente, ao final deste livro o endpoint `GET /v1/movies/:id` retornará uma resposta HTTP com um corpo JSON parecido com este:

```sh
$ curl -H "Authorization: Bearer RIDBIAE3AMMK57T6IAEBUGA7ZQ" localhost:4000/v1/movies/1
{
    "movie": {
        "id": 1,
        "title": "Moana",
        "year": 2016,
        "runtime": "107 mins",
        "genres": [
            "animation",
            "adventure"
        ],
        "version": 1
    }
}
```