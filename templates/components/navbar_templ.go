// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Navbar(swapOob bool) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<header class=\"navbar\" data-overflow-nav><!-- rest of navbar... --><nav id=\"navbarContainer\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if swapOob {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" hx-swap-oob=\"true\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" aria-label=\"Site sections\" class=\"!flex-grow:1\"><ul role=\"list\"><li><a href=\"/\"><span>head<b>count</b></span></a></li><li><a href=\"/about\">About</a></li></ul></nav><nav aria-label=\"Social media links\"><ul role=\"list\"><li><a href=\"https://github.com/lloydlobo/go-headcount\">GitHub</a></li></ul></nav></header><span hx-get=\"/contacts\" hx-target=\"#contact-list\" hx-indicator=\"#loader\" hx-trigger=\"load\"><span class=\"htmx-indicator indicator\" id=\"loader\" alt=\"Loading...\" aria-busy=\"true\"></span></span>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
