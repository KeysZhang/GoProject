package service

import(
	"net/http"
	"github.com/axgle/mahonia"
	"fmt"
)

func midwareGbk(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	r.ParseForm()
	if len(r.Form) != 0{
		for k, _ := range r.Form {
			enc := mahonia.NewEncoder("utf-8")
			r.Form[k][0] = enc.ConvertString(r.Form[k][0])
		}
		fmt.Println(r.Form)
	}
	next(rw, r)
	// do some stuff after
  }

