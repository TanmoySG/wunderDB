package txlModel

import "encoding/json"

func (j *TxlogSchemaJson) Marshal() ([]byte, error) {
	return json.Marshal(j)
}
