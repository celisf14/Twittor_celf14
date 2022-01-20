package models

/*Tweet captura del body, elmensaje que nos llega*/
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje,omitempty"`
}
