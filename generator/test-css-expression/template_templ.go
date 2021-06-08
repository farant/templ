// Code generated by templ DO NOT EDIT.

package testcssexpression

import "github.com/a-h/templ"
import "strings"

func className() templ.CSS {
	var templCSSBuilder strings.Builder
	templCSSBuilder.WriteString(`background-color:#ffffff;`)
	templCSSBuilder.WriteString(templ.SanitizeCSS(`color`, red))
	templCSSID := templ.CSSID(`className`, templCSSBuilder.String())
	return templ.CSS{
		ID: templCSSID,
		Class: templ.SafeCSS(`.` + templCSSID + `{` + templCSSBuilder.String() + `}`),
	}
}
