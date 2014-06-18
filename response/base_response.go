package response

import ( 
	"net/http"
	)

type BaseResponse struct {

}

func (resp BaseResponse) GetJSONResponseData() ([]byte, int) {
	return nil, http.StatusInternalServerError
}
