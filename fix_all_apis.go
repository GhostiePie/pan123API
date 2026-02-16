package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func processFile(path string) error {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("解析文件 %s 失败: %v", path, err)
	}

	// 查找需要删除的函数和需要保留的函数
	var withConfigFuncs []*ast.FuncDecl
	var simpleFuncs []*ast.FuncDecl

	for _, decl := range node.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			fnName := fn.Name.Name
			if strings.Contains(fnName, "WithConfig") {
				withConfigFuncs = append(withConfigFuncs, fn)
			} else if fn.Recv != nil {
				// 检查是否是APIClient的方法
				for _, field := range fn.Recv.List {
					if starExpr, ok := field.Type.(*ast.StarExpr); ok {
						if ident, ok := starExpr.X.(*ast.Ident); ok && ident.Name == "APIClient" {
							// 检查是否有参数，排除空函数
							if fn.Type.Params != nil && len(fn.Type.Params.List) > 0 {
								simpleFuncs = append(simpleFuncs, fn)
							}
						}
					}
				}
			}
		}
	}

	// 如果没有WithConfig函数，跳过
	if len(withConfigFuncs) == 0 {
		return nil
	}

	// 创建新文件内容
	var output strings.Builder

	// 写入包声明
	output.WriteString("package APIs\n\n")

	// 收集导入（简化处理，假设所有文件都有相同的导入）
	// 在实际中应该处理导入声明

	// 收集类型声明
	for _, decl := range node.Decls {
		switch decl.(type) {
		case *ast.GenDecl:
			// 类型声明
			printer.Fprint(&output, fset, decl)
			output.WriteString("\n")
		}
	}

	// 为每个WithConfig函数，找到对应的简单函数并替换
	for _, wcFunc := range withConfigFuncs {
		baseName := strings.Replace(wcFunc.Name.Name, "WithConfig", "", 1)

		// 查找对应的简单函数
		var simpleFunc *ast.FuncDecl
		for _, sf := range simpleFuncs {
			if sf.Name.Name == baseName {
				simpleFunc = sf
				break
			}
		}

		if simpleFunc != nil {
			// 用WithConfig函数体替换简单函数体
			simpleFunc.Body = wcFunc.Body

			// 更新函数签名（移除config参数）
			// 这里简化处理，直接使用WithConfig函数但重命名
			wcFunc.Name.Name = baseName

			// 写入函数
			printer.Fprint(&output, fset, wcFunc)
			output.WriteString("\n")
		}
	}

	// 写入文件
	return os.WriteFile(path, []byte(output.String()), 0644)
}

func main() {
	dir := "./APIs"

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("读取目录失败: %v\n", err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		name := file.Name()
		if !strings.HasSuffix(name, ".go") {
			continue
		}

		if name == "Client.go" || name == "Utils.go" {
			continue
		}

		path := filepath.Join(dir, name)
		fmt.Printf("处理文件: %s\n", path)

		if err := processFile(path); err != nil {
			fmt.Printf("错误处理文件 %s: %v\n", path, err)
		}
	}

	fmt.Println("完成!")
}
