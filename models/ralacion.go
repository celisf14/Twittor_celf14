package models

/*Relacion modelo para grabar la relacion con otro usuario*/
type Relacion struct {
	UsuarioID         string `bson:"usuarioid" json:"usuarioID"`
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"usuarioRelacionID"`
}
