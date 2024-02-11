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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<header class=\"navbar\" data-overflow-nav><!-- The navbar will still also remain horizontally scrollable. --><button class=\"iconbutton\" data-nav-expander aria-hidden>&#x2630; <!-- trigram for heaven --></button><!-- rest of navbar... --><nav id=\"navbarContainer\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if swapOob {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" hx-swap-oob=\"true\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" aria-label=\"Site sections\" class=\"contents\"><ul role=\"list\"><li><a href=\"/\" aria-label=\"Home\"><span>head<b>count</b></span></a></li><hr><li><a href=\"/about\">About</a></li><li><a href=\"https://github.com/lloydlobo/go-headcount\">GitHub</a></li><!-- <li><a href=\"/\"><img alt=\"\"/></a></li> --></ul></nav></header><!-- \n    Note: deprecated: using this directly at callsite i.e. in IndexPage\n    This should only appears during loading\n\t<span\n\t\thx-get=\"/contacts\"\n\t\thx-target=\"#contact-list\"\n\t\thx-indicator=\"#loader\"\n\t\thx-trigger=\"load\"\n\t>\n\t\t<span\n\t\t\tclass=\"htmx-indicator indicator !vh\"\n\t\t\tid=\"loader\"\n\t\t\talt=\"Loading...\"\n\t\t\taria-busy=\"true\"\n\t\t></span>\n\t</span>\n     -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
