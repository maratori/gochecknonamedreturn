package gochecknonamedreturn

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/maratori/pt"
	"github.com/stretchr/testify/assert"
)

func TestNamedReturnPosNoPos(t *testing.T) {
	pt.PackageParallel(t,
		pt.Group("should return NoPos",
			pt.Test("nil", func(t *testing.T) {
				pos := namedReturnPos(nil)
				assert.Exactly(t, token.NoPos, pos)
			}),
			pt.Test("nil list", func(t *testing.T) {
				pos := namedReturnPos(&ast.FieldList{
					List: nil,
				})
				assert.Exactly(t, token.NoPos, pos)
			}),
			pt.Test("empty list", func(t *testing.T) {
				pos := namedReturnPos(&ast.FieldList{
					List: []*ast.Field{},
				})
				assert.Exactly(t, token.NoPos, pos)
			}),
			pt.Group("single type",
				pt.Test("nil", func(t *testing.T) {
					pos := namedReturnPos(&ast.FieldList{
						List: []*ast.Field{
							nil,
						},
					})
					assert.Exactly(t, token.NoPos, pos)
				}),
				pt.Test("nil names", func(t *testing.T) {
					pos := namedReturnPos(&ast.FieldList{
						List: []*ast.Field{{
							Names: nil,
						}},
					})
					assert.Exactly(t, token.NoPos, pos)
				}),
				pt.Test("empty names", func(t *testing.T) {
					pos := namedReturnPos(&ast.FieldList{
						List: []*ast.Field{{
							Names: []*ast.Ident{},
						}},
					})
					assert.Exactly(t, token.NoPos, pos)
				}),
				pt.Test("single empty name", func(t *testing.T) {
					pos := namedReturnPos(&ast.FieldList{
						List: []*ast.Field{{
							Names: []*ast.Ident{{
								NamePos: 10,
								Name:    "",
							}},
						}},
					})
					assert.Exactly(t, token.NoPos, pos)
				}),
				pt.Test("two empty names", func(t *testing.T) {
					pos := namedReturnPos(&ast.FieldList{
						List: []*ast.Field{{
							Names: []*ast.Ident{
								{
									NamePos: 10,
									Name:    "",
								},
								{
									NamePos: 20,
									Name:    "",
								},
							},
						}},
					})
					assert.Exactly(t, token.NoPos, pos)
				}),
			),
			pt.Group("two types",
				pt.Test("nil", func(t *testing.T) {
					pos := namedReturnPos(&ast.FieldList{
						List: []*ast.Field{
							nil,
							nil,
						},
					})
					assert.Exactly(t, token.NoPos, pos)
				}),
				pt.Test("nil names", func(t *testing.T) {
					pos := namedReturnPos(&ast.FieldList{
						List: []*ast.Field{
							{
								Names: nil,
							},
							nil,
						},
					})
					assert.Exactly(t, token.NoPos, pos)
				}),
				pt.Test("empty names", func(t *testing.T) {
					pos := namedReturnPos(&ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{},
							},
							{
								Names: nil,
							},
						},
					})
					assert.Exactly(t, token.NoPos, pos)
				}),
				pt.Test("single empty name", func(t *testing.T) {
					pos := namedReturnPos(&ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{{
									NamePos: 10,
									Name:    "",
								}},
							},
							{
								Names: []*ast.Ident{},
							},
						},
					})
					assert.Exactly(t, token.NoPos, pos)
				}),
				pt.Test("two empty names", func(t *testing.T) {
					pos := namedReturnPos(&ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									{
										NamePos: 10,
										Name:    "",
									},
									{
										NamePos: 20,
										Name:    "",
									},
								},
							},
							{
								Names: []*ast.Ident{{
									NamePos: 30,
									Name:    "",
								}},
							},
						},
					})
					assert.Exactly(t, token.NoPos, pos)
				}),
			),
			pt.Test("different combinations", func(t *testing.T) {
				pos := namedReturnPos(&ast.FieldList{
					List: []*ast.Field{
						nil,
						{
							Names: nil,
						},
						{
							Names: []*ast.Ident{},
						},
						{
							Names: []*ast.Ident{
								nil,
							},
						},
						{
							Names: []*ast.Ident{{
								NamePos: 30,
								Name:    "",
							}},
						},
						{
							Names: []*ast.Ident{
								{
									NamePos: 20,
									Name:    "",
								},
								{
									NamePos: 10,
									Name:    "",
								},
							},
						},
					},
				})
				assert.Exactly(t, token.NoPos, pos)
			}),
		),
	)
}

func TestNamedReturnPosFound(t *testing.T) {
	pt.PackageParallel(t,
		pt.Group("should return first name position",
			pt.Group("single type",
				pt.Group("single name",
					pt.Test("underscore", func(t *testing.T) {
						pos := namedReturnPos(&ast.FieldList{
							List: []*ast.Field{{
								Names: []*ast.Ident{{
									NamePos: 12,
									Name:    "_",
								}},
							}},
						})
						assert.EqualValues(t, 12, pos)
					}),
					pt.Test("err", func(t *testing.T) {
						pos := namedReturnPos(&ast.FieldList{
							List: []*ast.Field{{
								Names: []*ast.Ident{{
									NamePos: 23,
									Name:    "err",
								}},
							}},
						})
						assert.EqualValues(t, 23, pos)
					}),
					pt.Test("ignore", func(t *testing.T) {
						pos := namedReturnPos(&ast.FieldList{
							List: []*ast.Field{{
								Names: []*ast.Ident{{
									NamePos: 34,
									Name:    "ignore",
								}},
							}},
						})
						assert.EqualValues(t, 34, pos)
					}),
				),
				pt.Group("two names",
					pt.Group("underscore",
						pt.Test("nil first", func(t *testing.T) {
							pos := namedReturnPos(&ast.FieldList{
								List: []*ast.Field{{
									Names: []*ast.Ident{
										nil,
										{
											NamePos: 45,
											Name:    "_",
										},
									},
								}},
							})
							assert.EqualValues(t, 45, pos)
						}),
						pt.Test("nil second", func(t *testing.T) {
							pos := namedReturnPos(&ast.FieldList{
								List: []*ast.Field{{
									Names: []*ast.Ident{
										{
											NamePos: 56,
											Name:    "_",
										},
										nil,
									},
								}},
							})
							assert.EqualValues(t, 56, pos)
						}),
						pt.Test("empty first", func(t *testing.T) {
							pos := namedReturnPos(&ast.FieldList{
								List: []*ast.Field{{
									Names: []*ast.Ident{
										{
											NamePos: 43,
											Name:    "",
										},
										{
											NamePos: 34,
											Name:    "_",
										},
									},
								}},
							})
							assert.EqualValues(t, 34, pos)
						}),
						pt.Test("empty second", func(t *testing.T) {
							pos := namedReturnPos(&ast.FieldList{
								List: []*ast.Field{{
									Names: []*ast.Ident{
										{
											NamePos: 21,
											Name:    "_",
										},
										{
											NamePos: 12,
											Name:    "",
										},
									},
								}},
							})
							assert.EqualValues(t, 21, pos)
						}),
						pt.Test("underscore", func(t *testing.T) {
							pos := namedReturnPos(&ast.FieldList{
								List: []*ast.Field{{
									Names: []*ast.Ident{
										{
											NamePos: 78,
											Name:    "_",
										},
										{
											NamePos: 67,
											Name:    "_",
										},
									},
								}},
							})
							assert.EqualValues(t, 78, pos)
						}),
						pt.Test("err first", func(t *testing.T) {
							pos := namedReturnPos(&ast.FieldList{
								List: []*ast.Field{{
									Names: []*ast.Ident{
										{
											NamePos: 67,
											Name:    "err",
										},
										{
											NamePos: 76,
											Name:    "_",
										},
									},
								}},
							})
							assert.EqualValues(t, 67, pos)
						}),
						pt.Test("err second", func(t *testing.T) {
							pos := namedReturnPos(&ast.FieldList{
								List: []*ast.Field{{
									Names: []*ast.Ident{
										{
											NamePos: 54,
											Name:    "_",
										},
										{
											NamePos: 45,
											Name:    "err",
										},
									},
								}},
							})
							assert.EqualValues(t, 54, pos)
						}),
					),
					pt.Test("err, ignore", func(t *testing.T) {
						pos := namedReturnPos(&ast.FieldList{
							List: []*ast.Field{{
								Names: []*ast.Ident{
									{
										NamePos: 98,
										Name:    "err",
									},
									{
										NamePos: 89,
										Name:    "ignore",
									},
								},
							}},
						})
						assert.EqualValues(t, 98, pos)
					}),
				),
			),
			pt.Group("two types",
				pt.Test("nil first", func(t *testing.T) {
					pos := namedReturnPos(&ast.FieldList{
						List: []*ast.Field{
							nil,
							{
								Names: []*ast.Ident{{
									NamePos: 123,
									Name:    "_",
								}},
							},
						},
					})
					assert.EqualValues(t, 123, pos)
				}),
				pt.Test("nil second", func(t *testing.T) {
					pos := namedReturnPos(&ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{{
									NamePos: 234,
									Name:    "err",
								}},
							},
							nil,
						},
					})
					assert.EqualValues(t, 234, pos)
				}),
				pt.Test("nil names", func(t *testing.T) {
					pos := namedReturnPos(&ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{{
									NamePos: 345,
									Name:    "_",
								}},
							},
							{
								Names: nil,
							},
						},
					})
					assert.EqualValues(t, 345, pos)
				}),
				pt.Test("empty names list", func(t *testing.T) {
					pos := namedReturnPos(&ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{},
							},
							{
								Names: []*ast.Ident{{
									NamePos: 312,
									Name:    "abc",
								}},
							},
						},
					})
					assert.EqualValues(t, 312, pos)
				}),
				pt.Test("empty name", func(t *testing.T) {
					pos := namedReturnPos(&ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{{
									NamePos: 1,
									Name:    "",
								}},
							},
							{
								Names: []*ast.Ident{{
									NamePos: 132,
									Name:    "_",
								}},
							},
						},
					})
					assert.EqualValues(t, 132, pos)
				}),
				pt.Test("underscore", func(t *testing.T) {
					pos := namedReturnPos(&ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{{
									NamePos: 543,
									Name:    "_",
								}},
							},
							{
								Names: []*ast.Ident{{
									NamePos: 123,
									Name:    "_",
								}},
							},
						},
					})
					assert.EqualValues(t, 543, pos)
				}),
				pt.Test("many names", func(t *testing.T) {
					pos := namedReturnPos(&ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									{
										NamePos: 2,
										Name:    "",
									},
									{
										NamePos: 413,
										Name:    "_",
									},
								},
							},
							{
								Names: []*ast.Ident{
									{
										NamePos: 4,
										Name:    "err",
									},
									{
										NamePos: 3,
										Name:    "_",
									},
								},
							},
						},
					})
					assert.EqualValues(t, 413, pos)
				}),
			),
		),
	)
}
