package commands

import "errors"

var errorMessages = map[string]string{
	"ErrProductNotFound":      "Ürün bulunamadı.",
	"ErrProductQuery":         "Ürün sorgulanırken bir hata oluştu.",
	"ErrActiveCampaignExists": "Bu ürün için zaten aktif bir kampanya bulunmaktadır.",
	"ErrCampaignNotFound":     "Kampanya bulunamadı.",
	"ErrCampaignUpdate":       "Kampanya durumu güncellenirken bir hata oluştu.",
	"ErrCampaignCreate":       "Kampanya oluşturulurken bir hata oluştu.",
	"ErrCampaignValidation":   "Kampanya doğrulama hatası.",
}

func GetErrorMessage(err string) error {
	if errMsg, ok := errorMessages[err]; ok {
		return errors.New(errMsg)
	}
	return errors.New("Beklenmeyen bir hata oluştu.")
}
