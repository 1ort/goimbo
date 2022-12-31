package handler

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path"
	"time"

	"github.com/1ort/goimbo/model"
	"github.com/gin-gonic/gin"
)

type WebCaptchaWrapper struct {
	captcha model.Captcha
}

func NewWebCaptchaWrapper(c model.Captcha) *WebCaptchaWrapper {
	return &WebCaptchaWrapper{
		captcha: c,
	}
}

type CaptchaRequest struct {
	ID       string `form:"captchaId" binding:"required"`
	Solution string `form:"captchaSolution" binding:"required"`
}

// TODO: remove hardcoded template
var templ = `
<tr>
	<td class="postblock">Captcha</td>
	<td>
		<img id="captcha_image" src="/captcha/%s">
		<br>
		<input id="captcha_input" type="text" name="captchaSolution" size="48">
		<input type="hidden" name="captchaId" value="%s">
	</td>
</tr>
`

func (w *WebCaptchaWrapper) presentCaptcha() template.HTML {
	if w.captcha == nil {
		return template.HTML("")
	}
	capID, capIDErr := w.captcha.New()
	if capIDErr != nil {
		return template.HTML("")
	}
	return template.HTML(fmt.Sprintf(templ, capID, capID))
}

func (w *WebCaptchaWrapper) verify(c *gin.Context) error {
	if w.captcha == nil {
		return nil
	}
	var cr CaptchaRequest
	if err := c.ShouldBind(&cr); err != nil {
		return model.NewBadRequest("Captcha error")
	}
	verified, err := w.captcha.Verify(cr.ID, cr.Solution)
	if err != nil {
		return model.NewBadRequest("Captcha error")
	}
	if !verified {
		return model.NewBadRequest("Incorrect captcha")
	}
	return nil
}

func (ww *WebCaptchaWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if ww.captcha == nil {
		http.NotFound(w, r)
		return
	}
	_, file := path.Split(r.URL.Path)
	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]
	if id == "" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Type", "image/png")
	var content bytes.Buffer
	ww.captcha.Write(&content, id)
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
}
