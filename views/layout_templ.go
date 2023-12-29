// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.501
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func head(page string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<meta charset=\"utf-8\"><meta name=\"viewport\" content=\"width=device-width,initial-scale=1\"><title>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var2 := `Liber - `
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var2)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(page)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/layout.templ`, Line: 5, Col: 22}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><link rel=\"apple-touch-icon\" sizes=\"180x180\" href=\"/assets/apple-touch-icon.png\"><link rel=\"icon\" type=\"image/x-icon\" href=\"/assets/favicon.ico\"><link rel=\"icon\" type=\"image/png\" sizes=\"32x32\" href=\"/assets/favicon-32x32.png\"><link rel=\"icon\" type=\"image/png\" sizes=\"16x16\" href=\"/assets/favicon-16x16.png\"><link rel=\"manifest\" href=\"/assets/site.webmanifest\"><!--")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var4 := ` TODO: properly bundle these `
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var4)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("--><script src=\"https://unpkg.com/htmx.org@1.9.10\" integrity=\"sha384-D1Kt99CQMDuVetoL1lrYwg5t+9QdHe7NLX/SoJYkXDFfX37iInKRy5xLSi8nO7UC\" crossorigin=\"anonymous\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var5 := ``
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var5)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</script><link href=\"https://cdn.jsdelivr.net/npm/daisyui@4.4.24/dist/full.min.css\" rel=\"stylesheet\" type=\"text/css\"><script src=\"https://cdn.tailwindcss.com\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var6 := ``
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var6)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func nav() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var7 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var7 == nil {
			templ_7745c5c3_Var7 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"navbar bg-base-100\"><div class=\"flex-none\"><label for=\"sidenav\" class=\"btn btn-square btn-ghost drawer-button lg:hidden\"><svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" class=\"inline-block w-5 h-5 stroke-current\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 6h16M4 12h16M4 18h16\"></path></svg></label></div><div class=\"flex flex-1 md:gap-1 lg:gap-2\"><a href=\"/\" class=\"btn btn-ghost text-xl\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var8 := `Liber`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var8)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a><div class=\"form-control\"><input type=\"text\" placeholder=\"Search\" class=\"input input-bordered md:w-auto lg:w-96\"></div></div><div class=\"flex-none gap-2\"><button class=\"btn btn-ghost btn-circle\"><svg xmlns=\"http://www.w3.org/2000/svg\" class=\"h-5 w-5\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\"><g id=\"SVGRepo_bgCarrier\" stroke-width=\"2\"></g><g id=\"SVGRepo_tracerCarrier\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></g><g id=\"SVGRepo_iconCarrier\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M12.0002 8C9.79111 8 8.00024 9.79086 8.00024 12C8.00024 14.2091 9.79111 16 12.0002 16C14.2094 16 16.0002 14.2091 16.0002 12C16.0002 9.79086 14.2094 8 12.0002 8ZM10.0002 12C10.0002 10.8954 10.8957 10 12.0002 10C13.1048 10 14.0002 10.8954 14.0002 12C14.0002 13.1046 13.1048 14 12.0002 14C10.8957 14 10.0002 13.1046 10.0002 12Z\" fill=\"currentColor\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M11.2867 0.5C9.88583 0.5 8.6461 1.46745 8.37171 2.85605L8.29264 3.25622C8.10489 4.20638 7.06195 4.83059 6.04511 4.48813L5.64825 4.35447C4.32246 3.90796 2.83873 4.42968 2.11836 5.63933L1.40492 6.83735C0.67773 8.05846 0.954349 9.60487 2.03927 10.5142L2.35714 10.7806C3.12939 11.4279 3.12939 12.5721 2.35714 13.2194L2.03927 13.4858C0.954349 14.3951 0.67773 15.9415 1.40492 17.1626L2.11833 18.3606C2.83872 19.5703 4.3225 20.092 5.64831 19.6455L6.04506 19.5118C7.06191 19.1693 8.1049 19.7935 8.29264 20.7437L8.37172 21.1439C8.6461 22.5325 9.88584 23.5 11.2867 23.5H12.7136C14.1146 23.5 15.3543 22.5325 15.6287 21.1438L15.7077 20.7438C15.8954 19.7936 16.9384 19.1693 17.9553 19.5118L18.3521 19.6455C19.6779 20.092 21.1617 19.5703 21.8821 18.3606L22.5955 17.1627C23.3227 15.9416 23.046 14.3951 21.9611 13.4858L21.6432 13.2194C20.8709 12.5722 20.8709 11.4278 21.6432 10.7806L21.9611 10.5142C23.046 9.60489 23.3227 8.05845 22.5955 6.83732L21.8821 5.63932C21.1617 4.42968 19.678 3.90795 18.3522 4.35444L17.9552 4.48814C16.9384 4.83059 15.8954 4.20634 15.7077 3.25617L15.6287 2.85616C15.3543 1.46751 14.1146 0.5 12.7136 0.5H11.2867ZM10.3338 3.24375C10.4149 2.83334 10.7983 2.5 11.2867 2.5H12.7136C13.2021 2.5 13.5855 2.83336 13.6666 3.24378L13.7456 3.64379C14.1791 5.83811 16.4909 7.09167 18.5935 6.38353L18.9905 6.24984C19.4495 6.09527 19.9394 6.28595 20.1637 6.66264L20.8771 7.86064C21.0946 8.22587 21.0208 8.69271 20.6764 8.98135L20.3586 9.24773C18.6325 10.6943 18.6325 13.3057 20.3586 14.7523L20.6764 15.0186C21.0208 15.3073 21.0946 15.7741 20.8771 16.1394L20.1637 17.3373C19.9394 17.714 19.4495 17.9047 18.9905 17.7501L18.5936 17.6164C16.4909 16.9082 14.1791 18.1618 13.7456 20.3562L13.6666 20.7562C13.5855 21.1666 13.2021 21.5 12.7136 21.5H11.2867C10.7983 21.5 10.4149 21.1667 10.3338 20.7562L10.2547 20.356C9.82113 18.1617 7.50931 16.9082 5.40665 17.6165L5.0099 17.7501C4.55092 17.9047 4.06104 17.714 3.83671 17.3373L3.1233 16.1393C2.9058 15.7741 2.97959 15.3073 3.32398 15.0186L3.64185 14.7522C5.36782 13.3056 5.36781 10.6944 3.64185 9.24779L3.32398 8.98137C2.97959 8.69273 2.9058 8.2259 3.1233 7.86067L3.83674 6.66266C4.06106 6.28596 4.55093 6.09528 5.0099 6.24986L5.40676 6.38352C7.50938 7.09166 9.82112 5.83819 10.2547 3.64392L10.3338 3.24375Z\" fill=\"currentColor\"></path></g></svg></button><div class=\"dropdown dropdown-end\"><button class=\"btn btn-ghost btn-circle\"><svg xmlns=\"http://www.w3.org/2000/svg\" class=\"h-5 w-5\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\"><g id=\"SVGRepo_bgCarrier\" stroke-width=\"2\"></g><g id=\"SVGRepo_tracerCarrier\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></g><g id=\"SVGRepo_iconCarrier\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M12 1C8.96243 1 6.5 3.46243 6.5 6.5C6.5 9.53757 8.96243 12 12 12C15.0376 12 17.5 9.53757 17.5 6.5C17.5 3.46243 15.0376 1 12 1ZM8.5 6.5C8.5 4.567 10.067 3 12 3C13.933 3 15.5 4.567 15.5 6.5C15.5 8.433 13.933 10 12 10C10.067 10 8.5 8.433 8.5 6.5Z\" fill=\"currentColor\"></path> <path d=\"M8 14C4.68629 14 2 16.6863 2 20V22C2 22.5523 2.44772 23 3 23C3.55228 23 4 22.5523 4 22V20C4 17.7909 5.79086 16 8 16H16C18.2091 16 20 17.7909 20 20V22C20 22.5523 20.4477 23 21 23C21.5523 23 22 22.5523 22 22V20C22 16.6863 19.3137 14 16 14H8Z\" fill=\"currentColor\"></path></g></svg></button><ul tabindex=\"0\" class=\"mt-3 z-[1] p-2 shadow menu menu-sm dropdown-content bg-base-100 rounded-box w-52\"><li><a class=\"justify-between\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var9 := `Profile`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var9)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <span class=\"badge\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var10 := `New`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var10)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></a></li><li><a>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var11 := `Settings`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var11)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></li><li><a>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var12 := `Logout`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var12)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></li></ul></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func sidebar() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var13 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var13 == nil {
			templ_7745c5c3_Var13 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul class=\"menu p-4 w-80 min-h-full bg-base-200 text-base-content\"><li><a>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var14 := `Sidebar Item 1`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var14)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></li><li><a>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var15 := `Sidebar Item 2`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var15)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></li></ul>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func layout(page string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var16 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var16 == nil {
			templ_7745c5c3_Var16 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en-US\"><head>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = head(page).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</head><body hx-boost=\"true\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = nav().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"fixed drawer lg:drawer-open\"><input id=\"sidenav\" type=\"checkbox\" class=\"drawer-toggle\"><div class=\"drawer-content flex flex-col items-center justify-center\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var16.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"drawer-side\"><label for=\"sidenav\" aria-label=\"close sidebar\" class=\"drawer-overlay\"></label>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = sidebar().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
