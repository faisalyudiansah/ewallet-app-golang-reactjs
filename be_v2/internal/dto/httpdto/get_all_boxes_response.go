package httpdto

import "ewallet-server-v2/internal/model"

type GetAllBoxesResponse struct {
	BoxNumbers []int `json:"box_numbers"`
}

func ConvertToGetAllBoxesResponse(boxes []model.GameBox) GetAllBoxesResponse {
	boxNumbers := []int{}

	for _, box := range boxes {
		boxNumbers = append(boxNumbers, int(box.GameBoxId))
	}

	return GetAllBoxesResponse{
		BoxNumbers: boxNumbers,
	}
}
