package modelos

import (

	arango "github.com/diegogub/aranGO"
)
type Usuario struct {
	arango.Document //Hay que incluirlo siempre en todas las estructuras porque añade campos nativos de arango
	Usuario   string `json:"usuario,omitempty"`
	Password  string `json:"password,omitempty"`
	Email 	  string `json:"email,omitempty"`
}
