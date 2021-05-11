package lsp

// code generated by helper. DO NOT EDIT.

import (
	"context"

	"github.com/nokia/ntt/internal/lsp/protocol"
)

func (s *Server) CodeAction(context.Context, *protocol.CodeActionParams) ([]protocol.CodeAction, error) {
	return nil, notImplemented("CodeAction")
}

func (s *Server) CodeLens(ctx context.Context, params *protocol.CodeLensParams) ([]protocol.CodeLens, error) {
	return s.codeLens(ctx, params)
}

func (s *Server) ColorPresentation(context.Context, *protocol.ColorPresentationParams) ([]protocol.ColorPresentation, error) {
	return nil, notImplemented("ColorPresentation")
}

func (s *Server) Completion(ctx context.Context, params *protocol.CompletionParams) (*protocol.CompletionList, error) {
	return s.completion(ctx, params)
}

func (s *Server) Declaration(context.Context, *protocol.DeclarationParams) (protocol.Declaration, error) {
	return nil, notImplemented("Declaration")
}

func (s *Server) Definition(ctx context.Context, params *protocol.DefinitionParams) (protocol.Definition, error) {
	return s.definition(ctx, params)
}

func (s *Server) DidChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) error {
	return s.didChange(ctx, params)
}

func (s *Server) DidChangeConfiguration(context.Context, *protocol.DidChangeConfigurationParams) error {
	return notImplemented("DidChangeConfiguration")
}

func (s *Server) DidChangeWatchedFiles(context.Context, *protocol.DidChangeWatchedFilesParams) error {
	return notImplemented("DidChangeWatchedFiles")
}

func (s *Server) DidChangeWorkspaceFolders(context.Context, *protocol.DidChangeWorkspaceFoldersParams) error {
	return notImplemented("DidChangeWorkspaceFolders")
}

func (s *Server) DidClose(ctx context.Context, params *protocol.DidCloseTextDocumentParams) error {
	return s.didClose(ctx, params)
}

func (s *Server) DidOpen(ctx context.Context, params *protocol.DidOpenTextDocumentParams) error {
	return s.didOpen(ctx, params)
}

func (s *Server) DidSave(ctx context.Context, params *protocol.DidSaveTextDocumentParams) error {
	return s.didSave(ctx, params)
}

func (s *Server) DocumentColor(context.Context, *protocol.DocumentColorParams) ([]protocol.ColorInformation, error) {
	return nil, notImplemented("DocumentColor")
}

func (s *Server) DocumentHighlight(context.Context, *protocol.DocumentHighlightParams) ([]protocol.DocumentHighlight, error) {
	return nil, notImplemented("DocumentHighlight")
}

func (s *Server) DocumentLink(ctx context.Context, params *protocol.DocumentLinkParams) ([]protocol.DocumentLink, error) {
	return s.documentLink(ctx, params)
}

func (s *Server) DocumentSymbol(context.Context, *protocol.DocumentSymbolParams) ([]interface{}, error) {
	return nil, notImplemented("DocumentSymbol")
}

func (s *Server) ExecuteCommand(ctx context.Context, params *protocol.ExecuteCommandParams) (interface{}, error) {
	return s.executeCommand(ctx, params)
}

func (s *Server) Exit(ctx context.Context) error {
	return s.exit(ctx)
}

func (s *Server) FoldingRange(context.Context, *protocol.FoldingRangeParams) ([]protocol.FoldingRange, error) {
	return nil, notImplemented("FoldingRange")
}

func (s *Server) Formatting(context.Context, *protocol.DocumentFormattingParams) ([]protocol.TextEdit, error) {
	return nil, notImplemented("Formatting")
}

func (s *Server) Hover(context.Context, *protocol.HoverParams) (*protocol.Hover, error) {
	return nil, notImplemented("Hover")
}

func (s *Server) Implementation(context.Context, *protocol.ImplementationParams) (protocol.Definition, error) {
	return nil, notImplemented("Implementation")
}

func (s *Server) IncomingCalls(context.Context, *protocol.CallHierarchyIncomingCallsParams) ([]protocol.CallHierarchyIncomingCall, error) {
	return nil, notImplemented("IncomingCalls")
}

func (s *Server) Initialize(ctx context.Context, params *protocol.ParamInitialize) (*protocol.InitializeResult, error) {
	return s.initialize(ctx, params)
}

func (s *Server) Initialized(ctx context.Context, params *protocol.InitializedParams) error {
	return s.initialized(ctx, params)
}

func (s *Server) LogTrace(context.Context, *protocol.LogTraceParams) error {
	return notImplemented("LogTrace")
}

func (s *Server) NonstandardRequest(ctx context.Context, method string, params interface{}) (interface{}, error) {
	return s.nonstandardRequest(ctx, method, params)
}

func (s *Server) OnTypeFormatting(context.Context, *protocol.DocumentOnTypeFormattingParams) ([]protocol.TextEdit, error) {
	return nil, notImplemented("OnTypeFormatting")
}

func (s *Server) OutgoingCalls(context.Context, *protocol.CallHierarchyOutgoingCallsParams) ([]protocol.CallHierarchyOutgoingCall, error) {
	return nil, notImplemented("OutgoingCalls")
}

func (s *Server) PrepareCallHierarchy(context.Context, *protocol.CallHierarchyPrepareParams) ([]protocol.CallHierarchyItem, error) {
	return nil, notImplemented("PrepareCallHierarchy")
}

func (s *Server) PrepareRename(context.Context, *protocol.PrepareRenameParams) (*protocol.Range, error) {
	return nil, notImplemented("PrepareRename")
}

func (s *Server) RangeFormatting(context.Context, *protocol.DocumentRangeFormattingParams) ([]protocol.TextEdit, error) {
	return nil, notImplemented("RangeFormatting")
}

func (s *Server) References(ctx context.Context, params *protocol.ReferenceParams) ([]protocol.Location, error) {
	return s.references(ctx, params)
}

func (s *Server) Rename(context.Context, *protocol.RenameParams) (*protocol.WorkspaceEdit, error) {
	return nil, notImplemented("Rename")
}

func (s *Server) Resolve(context.Context, *protocol.CompletionItem) (*protocol.CompletionItem, error) {
	return nil, notImplemented("Resolve")
}

func (s *Server) ResolveCodeLens(context.Context, *protocol.CodeLens) (*protocol.CodeLens, error) {
	return nil, notImplemented("ResolveCodeLens")
}

func (s *Server) ResolveDocumentLink(context.Context, *protocol.DocumentLink) (*protocol.DocumentLink, error) {
	return nil, notImplemented("ResolveDocumentLink")
}

func (s *Server) SelectionRange(context.Context, *protocol.SelectionRangeParams) ([]protocol.SelectionRange, error) {
	return nil, notImplemented("SelectionRange")
}

func (s *Server) SemanticTokensFull(context.Context, *protocol.SemanticTokensParams) (*protocol.SemanticTokens, error) {
	return nil, notImplemented("SemanticTokensFull")
}

func (s *Server) SemanticTokensFullDelta(context.Context, *protocol.SemanticTokensDeltaParams) (interface{}, error) {
	return nil, notImplemented("SemanticTokensFullDelta")
}

func (s *Server) SemanticTokensRange(context.Context, *protocol.SemanticTokensRangeParams) (*protocol.SemanticTokens, error) {
	return nil, notImplemented("SemanticTokensRange")
}

func (s *Server) SetTrace(ctx context.Context, params *protocol.SetTraceParams) error {
	return s.setTrace(ctx, params)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.shutdown(ctx)
}

func (s *Server) SignatureHelp(context.Context, *protocol.SignatureHelpParams) (*protocol.SignatureHelp, error) {
	return nil, notImplemented("SignatureHelp")
}

func (s *Server) Symbol(context.Context, *protocol.WorkspaceSymbolParams) ([]protocol.SymbolInformation, error) {
	return nil, notImplemented("Symbol")
}

func (s *Server) TypeDefinition(context.Context, *protocol.TypeDefinitionParams) (protocol.Definition, error) {
	return nil, notImplemented("TypeDefinition")
}

func (s *Server) WillSave(context.Context, *protocol.WillSaveTextDocumentParams) error {
	return notImplemented("WillSave")
}

func (s *Server) WillSaveWaitUntil(context.Context, *protocol.WillSaveTextDocumentParams) ([]protocol.TextEdit, error) {
	return nil, notImplemented("WillSaveWaitUntil")
}

func (s *Server) WorkDoneProgressCancel(context.Context, *protocol.WorkDoneProgressCancelParams) error {
	return notImplemented("WorkDoneProgressCancel")
}
