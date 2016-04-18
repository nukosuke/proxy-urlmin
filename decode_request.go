package forms

type DecodeRequest struct {
    Url string `form:"url" json:"url" binding:"required"`
}