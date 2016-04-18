package forms

type EncodeRequest struct {
    Url string `form:"url" json:"url" binding:"required"`
}