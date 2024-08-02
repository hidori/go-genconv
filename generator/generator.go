package generator

import (
	"go/ast"
	"go/token"

	"github.com/hidori/go-astutil"
	"github.com/hidori/go-typeutil"
	"github.com/pkg/errors"
)

type GeneratorConfig struct {
	TagName        string
	Initialism     []string
	Validate       bool
	ValidationFunc string
	ValidationTag  string
}

type Generator struct {
	config *GeneratorConfig
}

func NewGenerator(config *GeneratorConfig) *Generator {
	return &Generator{
		config: config,
	}
}

func (g *Generator) Generate(fileSet *token.FileSet, file *ast.File) ([]ast.Decl, error) {
	var decls []ast.Decl

	for _, d := range file.Decls {
		genDecl := typeutil.AsOrEmpty[*ast.GenDecl](d)
		funcDecl := typeutil.AsOrEmpty[*ast.FuncDecl](d)

		switch {
		case genDecl != nil:
			_decls, err := g.fromGenDecl(genDecl)
			if err != nil {
				return nil, errors.WithStack(err)
			}

			decls = append(decls, _decls...)

		case funcDecl != nil:
			_decls, err := g.fromFuncDecl(funcDecl)
			if err != nil {
				return nil, errors.WithStack(err)
			}

			decls = append(decls, _decls...)

		default:
			// nothing to do
		}
	}

	return decls, nil
}

func (g *Generator) fromGenDecl(genDecl *ast.GenDecl) ([]ast.Decl, error) {
	switch genDecl.Tok {
	case token.IMPORT:
		return []ast.Decl{genDecl}, nil

	default:
		return []ast.Decl{}, nil
	}
}

func (g *Generator) fromFuncDecl(funcDecl *ast.FuncDecl) ([]ast.Decl, error) {
	name := astutil.NewIdent(
		"Convert" + funcDecl.Name.Name,
	)
	funcType := astutil.NewFuncType(
		nil,
		astutil.NewFieldList(
			[]*ast.Field{
				astutil.NewField([]*ast.Ident{ast.NewIdent("src")}, ast.NewIdent("SrcType")),
			},
		),
		astutil.NewFieldList(
			[]*ast.Field{
				astutil.NewField(nil, ast.NewIdent("DestType")),
				astutil.NewField(nil, ast.NewIdent("error")),
			},
		),
	)
	body := astutil.NewBlockStmt(
		[]ast.Stmt{
			astutil.NewReturnStmt(
				[]ast.Expr{
					astutil.NewSelectorExpr(astutil.NewIdent("t"), astutil.NewIdent("getter data")),
				},
			),
		},
	)

	return []ast.Decl{&ast.FuncDecl{
		Name: name,
		Type: funcType,
		Body: body,
	}}, nil
}
