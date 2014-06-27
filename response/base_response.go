package response

import ( 
	"net/http"
	"encoding/json"
	)

type BaseResponse struct {

}

func (resp BaseResponse) GetJSONResponseData() ([]byte, int) {
	// return nil, http.StatusInternalServerError
	jsonData, _ := json.Marshal(resp)
	return jsonData, http.StatusOK
}
