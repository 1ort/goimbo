package service

import (
	"io"

	"github.com/1ort/goimbo/model"
	"github.com/dchest/captcha"
)

const (
	defaultWidth  = 240
	defaultHeight = 80
)

type ImageCaptcha struct {
	width  int
	height int
}

type ImageCaptchaConfig struct {
	Width  int
	Height int
}

func NewImageCaptcha(cfg *ImageCaptchaConfig) model.Captcha {
	ic := &ImageCaptcha{
		width:  cfg.Width,
		height: cfg.Height,
	}
	if ic.width == 0 {
		ic.width = defaultWidth
	}
	if ic.height == 0 {
		ic.height = defaultHeight
	}
	return ic
}

func (ic *ImageCaptcha) New() (string, error) {
	return captcha.New(), nil
}

func (ic *ImageCaptcha) Verify(id, solution string) (bool, error) {
	return captcha.VerifyString(id, solution), nil
}

func (ic *ImageCaptcha) Write(w io.Writer, id string) error {
	return captcha.WriteImage(w, id, ic.width, ic.height)
}
