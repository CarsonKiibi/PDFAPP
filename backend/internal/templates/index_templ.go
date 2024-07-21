// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
// index.templ

package main

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func index() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Code Compiler</title><script src=\"https://unpkg.com/htmx.org@1.9.10\"></script><script src=\"https://unpkg.com/monaco-editor@0.36.1/min/vs/loader.js\"></script><link href=\"https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css\" rel=\"stylesheet\"><style>\r\n\t\t\t\t#editor { height: 400px; }\r\n\t\t\t\t@media (min-width: 768px) { #editor { height: 600px; } }\r\n\t\t\t</style></head><body class=\"bg-gray-100\"><div class=\"container mx-auto p-4\"><h1 class=\"text-3xl font-bold mb-4\">Code Compiler</h1><div class=\"flex flex-col md:flex-row gap-4\"><div class=\"w-full md:w-1/2\"><div id=\"editor\" class=\"border border-gray-300 rounded\"></div><button class=\"mt-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded\" hx-post=\"/compile\" hx-trigger=\"click\" hx-target=\"#output\" hx-indicator=\"#loading\" hx-swap=\"innerHTML\">Compile</button><div id=\"loading\" class=\"htmx-indicator mt-2\">Compiling...</div></div><div class=\"w-full md:w-1/2\"><div id=\"output\" class=\"border border-gray-300 rounded p-4 h-[400px] md:h-[600px] overflow-auto\">Output will appear here</div></div></div></div><script>\r\n\t\t\t\trequire.config({ paths: { vs: 'https://unpkg.com/monaco-editor@0.36.1/min/vs' } });\r\n\t\t\t\trequire(['vs/editor/editor.main'], function() {\r\n\t\t\t\t\tvar editor = monaco.editor.create(document.getElementById('editor'), {\r\n\t\t\t\t\t\tvalue: '// Your code here',\r\n\t\t\t\t\t\tlanguage: 'plaintext',\r\n\t\t\t\t\t\ttheme: 'vs-light',\r\n\t\t\t\t\t\tminimap: { enabled: false }\r\n\t\t\t\t\t});\r\n\r\n\t\t\t\t\thtmx.on('htmx:beforeRequest', function(evt) {\r\n\t\t\t\t\t\tif (evt.detail.elt.getAttribute('hx-target') === '#output') {\r\n\t\t\t\t\t\t\tevt.detail.headers['Content-Type'] = 'application/x-www-form-urlencoded';\r\n\t\t\t\t\t\t\tevt.detail.parameters.code = editor.getValue();\r\n\t\t\t\t\t\t}\r\n\t\t\t\t\t});\r\n\r\n\t\t\t\t\thtmx.on('htmx:afterSwap', function(evt) {\r\n\t\t\t\t\t\tif (evt.detail.target.id === 'output') {\r\n\t\t\t\t\t\t\tvar response = evt.detail.xhr.response;\r\n\t\t\t\t\t\t\ttry {\r\n\t\t\t\t\t\t\t\tvar jsonResponse = JSON.parse(response);\r\n\t\t\t\t\t\t\t\tif (jsonResponse.line && jsonResponse.column) {\r\n\t\t\t\t\t\t\t\t\teditor.deltaDecorations([], [\r\n\t\t\t\t\t\t\t\t\t\t{\r\n\t\t\t\t\t\t\t\t\t\t\trange: new monaco.Range(jsonResponse.line, jsonResponse.column, jsonResponse.line, jsonResponse.column + 1),\r\n\t\t\t\t\t\t\t\t\t\t\toptions: {\r\n\t\t\t\t\t\t\t\t\t\t\t\tisWholeLine: true,\r\n\t\t\t\t\t\t\t\t\t\t\t\tclassName: 'bg-red-200',\r\n\t\t\t\t\t\t\t\t\t\t\t\tglyphMarginClassName: 'bg-red-200'\r\n\t\t\t\t\t\t\t\t\t\t\t}\r\n\t\t\t\t\t\t\t\t\t\t}\r\n\t\t\t\t\t\t\t\t\t]);\r\n\t\t\t\t\t\t\t\t\tevt.detail.target.innerHTML = `Error at line ${jsonResponse.line}, column ${jsonResponse.column}: ${jsonResponse.message}`;\r\n\t\t\t\t\t\t\t\t}\r\n\t\t\t\t\t\t\t} catch (e) {\r\n\t\t\t\t\t\t\t\t// If it's not JSON, it's probably the PDF\r\n\t\t\t\t\t\t\t\tvar blob = new Blob([response], { type: 'application/pdf' });\r\n\t\t\t\t\t\t\t\tvar url = URL.createObjectURL(blob);\r\n\t\t\t\t\t\t\t\tevt.detail.target.innerHTML = `<embed src=\"${url}\" type=\"application/pdf\" width=\"100%\" height=\"100%\">`;\r\n\t\t\t\t\t\t\t}\r\n\t\t\t\t\t\t}\r\n\t\t\t\t\t});\r\n\t\t\t\t});\r\n\t\t\t</script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
