package lsp_test

import (
	"fmt"
	"testing"

	"github.com/nokia/ntt/internal/loc"
	"github.com/nokia/ntt/internal/lsp"
	"github.com/nokia/ntt/internal/lsp/protocol"
	"github.com/nokia/ntt/internal/ntt"
	"github.com/stretchr/testify/assert"
)

func generateSymbols(t *testing.T, suite *ntt.Suite) (*ntt.ParseInfo, []protocol.DocumentSymbol) {

	name := fmt.Sprintf("%s_Module_0.ttcn3", t.Name())
	syntax := suite.ParseWithAllErrors(name)
	list := lsp.NewAllDefinitionSymbolsFromCurrentModule(syntax)
	ret := make([]protocol.DocumentSymbol, 0, len(list))
	for _, l := range list {
		if l, ok := l.(protocol.DocumentSymbol); ok {
			ret = append(ret, l)
		}
	}

	return syntax, ret
}

func setRange(syntax *ntt.ParseInfo, begin loc.Pos, end loc.Pos) protocol.Range {
	b := syntax.Position(begin)
	e := syntax.Position(end)
	ret := protocol.Range{
		Start: protocol.Position{Line: float64(b.Line - 1), Character: float64(b.Column)},
		End:   protocol.Position{Line: float64(e.Line - 1), Character: float64(e.Column)}}

	return ret
}
func TestFunctionDefWithModuleDotRunsOn(t *testing.T) {
	suite := buildSuite(t, `module Test
    {
        type component B0 {
			var integer i := 1;
			timer t1 := 2.0;
			port P p;
		}
		function f() runs on TestFunctionDefWithModuleDotRunsOn_Module_1.C0 system B0 {}
	  }`, `module TestFunctionDefWithModuleDotRunsOn_Module_1
      {
		  type component C0 {}
	  }`)

	syntax, list := generateSymbols(t, suite)

	assert.Equal(t, []protocol.DocumentSymbol{
		{Name: "B0", Kind: protocol.Class, Detail: "component type",
			Range:          setRange(syntax, 26, 105),
			SelectionRange: setRange(syntax, 26, 105),
			Children: []protocol.DocumentSymbol{
				{Name: "i", Detail: "var integer", Kind: protocol.Variable,
					Range:          setRange(syntax, 49, 67),
					SelectionRange: setRange(syntax, 49, 67)},
				{Name: "t1", Detail: "timer", Kind: protocol.Event,
					Range:          setRange(syntax, 72, 87),
					SelectionRange: setRange(syntax, 72, 87)},
				{Name: "p", Detail: "port P", Kind: protocol.Interface,
					Range:          setRange(syntax, 92, 100),
					SelectionRange: setRange(syntax, 92, 100)}}},
		{Name: "f", Kind: protocol.Method, Detail: "function definition",
			Range:          setRange(syntax, 108, 188),
			SelectionRange: setRange(syntax, 108, 188),
			Children: []protocol.DocumentSymbol{
				{Name: "runs on", Detail: "TestFunctionDefWithModuleDotRunsOn_Module_1.C0", Kind: protocol.Class,
					Range:          setRange(syntax, 129, 175),
					SelectionRange: setRange(syntax, 129, 175)},
				{Name: "system", Detail: "B0", Kind: protocol.Class,
					Range:          setRange(syntax, 183, 185),
					SelectionRange: setRange(syntax, 183, 185)}}}}, list)
}

func TestRecordOfTypeDefWithTypeRef(t *testing.T) {
	suite := buildSuite(t, `module Test
    {
        type integer Byte(0..255)
		type record of Byte Octets
	  }`)

	syntax, list := generateSymbols(t, suite)

	assert.Equal(t, []protocol.DocumentSymbol{
		{Name: "Byte", Kind: protocol.Struct, Detail: "subtype",
			Range:          setRange(syntax, 26, 51),
			SelectionRange: setRange(syntax, 26, 51),
			Children:       nil},
		{Name: "Octets", Kind: protocol.Array, Detail: "record of type",
			Range:          setRange(syntax, 54, 80),
			SelectionRange: setRange(syntax, 54, 80),
			Children: []protocol.DocumentSymbol{
				{Name: "Byte", Detail: "element type", Kind: protocol.Struct,
					Range:          setRange(syntax, 69, 73),
					SelectionRange: setRange(syntax, 69, 73)}}}}, list)
}
