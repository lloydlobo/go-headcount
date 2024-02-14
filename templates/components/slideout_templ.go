// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/lloydlobo/go-headcount/templates"

// Requires: <script defer src="https://unpkg.com/alpinejs@latest/dist/cdn.min.js"></script>
func Slideout(contents templ.Component, triggerText string, isOpen bool) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-cloak x-data=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString("{ slideOut: " + templates.BoolToStrJS(isOpen) + " }"))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" @keydown.window.escape=\"slideOut = false\"><!-- trigger open -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !isOpen {
			if triggerText == "" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button @click=\"slideOut = !slideOut\" class=\"!big\">Toggle</button>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button @click=\"slideOut = !slideOut\" class=\"big\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var2 string
				templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(templ.EscapeString(triggerText))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates\components\slideout.templ`, Line: 22, Col: 38}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- overlay --><div x-cloak x-transition.opacity x-show=\"slideOut\" class=\"fixed\" style=\"inset: 0; background: #00000050;\"></div><!-- content --><aside x-cloak x-show=\"slideOut\" x-transition:enter=\"transition ease-out duration-300\" x-transition:enter-start=\"translate-x-full\" x-transition:enter-end=\"translate-x-0\" x-transition:leave=\"transition ease-in duration-300\" x-transition:leave-start=\"translate-x-0\" x-transition:leave-end=\"translate-x-full\" @click.away=\"slideOut = false\" class=\"fixed box slideout-aside\"><!-- trigger close --><button @click=\"slideOut = !slideOut\" class=\"iconbutton\" title=\"Close\" type=\"button\" role=\"button\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = XIcon().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button><!-- children content -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = contents.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</aside></div><style type=\"text/css\">\n        aside.slideout-aside {\n            position: fixed; \n            top: 0; \n            right: 0; \n            bottom: 0; \n            min-width: 300px; \n            margin-block: 0;\n            z-index: 50;\n        }\n    </style><style type=\"text/css\">\n        .transition {\n            transition: ease-out 300ms;\n        }\n        .ease-out {\n            transition-timing-function: ease-out;\n        }\n        .duration-300 {\n            transition-duration: 300ms;\n        }\n        .translate-x-full {\n            transform: translateX(100%);\n        }\n        .translate-x-0 {\n            transform: translateX(0);\n        }\n        .ease-in {\n            transition-timing-function: ease-in;\n        }\n    </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
