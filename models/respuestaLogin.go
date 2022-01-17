package models

/*RespuestaLogin tiene el token que devuelve el loginn*/
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}
