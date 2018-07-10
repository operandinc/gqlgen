// Code generated by github.com/vektah/gqlgen, DO NOT EDIT.

package scalars

import (
	"bytes"
	context "context"
	external "external"
	strconv "strconv"
	time "time"

	model "github.com/vektah/gqlgen/example/scalars/model"
	graphql "github.com/vektah/gqlgen/graphql"
	introspection "github.com/vektah/gqlgen/neelance/introspection"
	query "github.com/vektah/gqlgen/neelance/query"
	schema "github.com/vektah/gqlgen/neelance/schema"
)

// MakeExecutableSchema creates an ExecutableSchema from the Resolvers interface.
func MakeExecutableSchema(resolvers Resolvers) graphql.ExecutableSchema {
	return &executableSchema{resolvers: resolvers}
}

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(resolvers ResolverRoot) graphql.ExecutableSchema {
	return MakeExecutableSchema(shortMapper{r: resolvers})
}

type Resolvers interface {
	Query_user(ctx context.Context, id external.ObjectID) (*model.User, error)
	Query_search(ctx context.Context, input model.SearchArgs) ([]model.User, error)

	User_primitiveResolver(ctx context.Context, obj *model.User) (string, error)
	User_customResolver(ctx context.Context, obj *model.User) (model.Point, error)
}

type ResolverRoot interface {
	Query() QueryResolver
	User() UserResolver
}
type QueryResolver interface {
	User(ctx context.Context, id external.ObjectID) (*model.User, error)
	Search(ctx context.Context, input model.SearchArgs) ([]model.User, error)
}
type UserResolver interface {
	PrimitiveResolver(ctx context.Context, obj *model.User) (string, error)
	CustomResolver(ctx context.Context, obj *model.User) (model.Point, error)
}

type shortMapper struct {
	r ResolverRoot
}

func (s shortMapper) Query_user(ctx context.Context, id external.ObjectID) (*model.User, error) {
	return s.r.Query().User(ctx, id)
}

func (s shortMapper) Query_search(ctx context.Context, input model.SearchArgs) ([]model.User, error) {
	return s.r.Query().Search(ctx, input)
}

func (s shortMapper) User_primitiveResolver(ctx context.Context, obj *model.User) (string, error) {
	return s.r.User().PrimitiveResolver(ctx, obj)
}

func (s shortMapper) User_customResolver(ctx context.Context, obj *model.User) (model.Point, error) {
	return s.r.User().CustomResolver(ctx, obj)
}

type executableSchema struct {
	resolvers Resolvers
}

func (e *executableSchema) Schema() *schema.Schema {
	return parsedSchema
}

func (e *executableSchema) Query(ctx context.Context, op *query.Operation) *graphql.Response {
	ec := executionContext{graphql.GetRequestContext(ctx), e.resolvers}

	buf := ec.RequestMiddleware(ctx, func(ctx context.Context) []byte {
		data := ec._Query(ctx, op.Selections)
		var buf bytes.Buffer
		data.MarshalGQL(&buf)
		return buf.Bytes()
	})

	return &graphql.Response{
		Data:   buf,
		Errors: ec.Errors,
	}
}

func (e *executableSchema) Mutation(ctx context.Context, op *query.Operation) *graphql.Response {
	return graphql.ErrorResponse(ctx, "mutations are not supported")
}

func (e *executableSchema) Subscription(ctx context.Context, op *query.Operation) func() *graphql.Response {
	return graphql.OneShot(graphql.ErrorResponse(ctx, "subscriptions are not supported"))
}

type executionContext struct {
	*graphql.RequestContext

	resolvers Resolvers
}

var addressImplementors = []string{"Address"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Address(ctx context.Context, sel []query.Selection, obj *model.Address) graphql.Marshaler {
	fields := graphql.CollectFields(ec.Doc, sel, addressImplementors, ec.Variables)

	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Address")
		case "id":
			out.Values[i] = ec._Address_id(ctx, field, obj)
		case "location":
			out.Values[i] = ec._Address_location(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) _Address_id(ctx context.Context, field graphql.CollectedField, obj *model.Address) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "Address"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.ID
	return model.MarshalID(res)
}

func (ec *executionContext) _Address_location(ctx context.Context, field graphql.CollectedField, obj *model.Address) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "Address"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.Location
	if res == nil {
		return graphql.Null
	}
	return *res
}

var queryImplementors = []string{"Query"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Query(ctx context.Context, sel []query.Selection) graphql.Marshaler {
	fields := graphql.CollectFields(ec.Doc, sel, queryImplementors, ec.Variables)

	ctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{
		Object: "Query",
	})

	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Query")
		case "user":
			out.Values[i] = ec._Query_user(ctx, field)
		case "search":
			out.Values[i] = ec._Query_search(ctx, field)
		case "__schema":
			out.Values[i] = ec._Query___schema(ctx, field)
		case "__type":
			out.Values[i] = ec._Query___type(ctx, field)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) _Query_user(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	args := map[string]interface{}{}
	var arg0 external.ObjectID
	if tmp, ok := field.Args["id"]; ok {
		var err error
		arg0, err = model.UnmarshalID(tmp)
		if err != nil {
			ec.Error(ctx, err)
			return graphql.Null
		}
	}
	args["id"] = arg0
	ctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{
		Object: "Query",
		Args:   args,
		Field:  field,
	})
	return graphql.Defer(func() (ret graphql.Marshaler) {
		defer func() {
			if r := recover(); r != nil {
				userErr := ec.Recover(ctx, r)
				ec.Error(ctx, userErr)
				ret = graphql.Null
			}
		}()

		resTmp, err := ec.ResolverMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
			return ec.resolvers.Query_user(ctx, args["id"].(external.ObjectID))
		})
		if err != nil {
			ec.Error(ctx, err)
			return graphql.Null
		}
		if resTmp == nil {
			return graphql.Null
		}
		res := resTmp.(*model.User)
		if res == nil {
			return graphql.Null
		}
		return ec._User(ctx, field.Selections, res)
	})
}

func (ec *executionContext) _Query_search(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	args := map[string]interface{}{}
	var arg0 model.SearchArgs
	if tmp, ok := field.Args["input"]; ok {
		var err error
		arg0, err = UnmarshalSearchArgs(tmp)
		if err != nil {
			ec.Error(ctx, err)
			return graphql.Null
		}
	} else {
		var tmp interface{} = map[string]interface{}{"isBanned": false, "location": "37,144"}
		var err error
		arg0, err = UnmarshalSearchArgs(tmp)
		if err != nil {
			ec.Error(ctx, err)
			return graphql.Null
		}
	}

	args["input"] = arg0
	ctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{
		Object: "Query",
		Args:   args,
		Field:  field,
	})
	return graphql.Defer(func() (ret graphql.Marshaler) {
		defer func() {
			if r := recover(); r != nil {
				userErr := ec.Recover(ctx, r)
				ec.Error(ctx, userErr)
				ret = graphql.Null
			}
		}()

		resTmp, err := ec.ResolverMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
			return ec.resolvers.Query_search(ctx, args["input"].(model.SearchArgs))
		})
		if err != nil {
			ec.Error(ctx, err)
			return graphql.Null
		}
		if resTmp == nil {
			return graphql.Null
		}
		res := resTmp.([]model.User)
		arr1 := graphql.Array{}
		for idx1 := range res {
			arr1 = append(arr1, func() graphql.Marshaler {
				rctx := graphql.GetResolverContext(ctx)
				rctx.PushIndex(idx1)
				defer rctx.Pop()
				return ec._User(ctx, field.Selections, &res[idx1])
			}())
		}
		return arr1
	})
}

func (ec *executionContext) _Query___schema(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "Query"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := ec.introspectSchema()
	if res == nil {
		return graphql.Null
	}
	return ec.___Schema(ctx, field.Selections, res)
}

func (ec *executionContext) _Query___type(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := field.Args["name"]; ok {
		var err error
		arg0, err = graphql.UnmarshalString(tmp)
		if err != nil {
			ec.Error(ctx, err)
			return graphql.Null
		}
	}
	args["name"] = arg0
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "Query"
	rctx.Args = args
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := ec.introspectType(args["name"].(string))
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

var userImplementors = []string{"User"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _User(ctx context.Context, sel []query.Selection, obj *model.User) graphql.Marshaler {
	fields := graphql.CollectFields(ec.Doc, sel, userImplementors, ec.Variables)

	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("User")
		case "id":
			out.Values[i] = ec._User_id(ctx, field, obj)
		case "name":
			out.Values[i] = ec._User_name(ctx, field, obj)
		case "created":
			out.Values[i] = ec._User_created(ctx, field, obj)
		case "isBanned":
			out.Values[i] = ec._User_isBanned(ctx, field, obj)
		case "primitiveResolver":
			out.Values[i] = ec._User_primitiveResolver(ctx, field, obj)
		case "customResolver":
			out.Values[i] = ec._User_customResolver(ctx, field, obj)
		case "address":
			out.Values[i] = ec._User_address(ctx, field, obj)
		case "tier":
			out.Values[i] = ec._User_tier(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) _User_id(ctx context.Context, field graphql.CollectedField, obj *model.User) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "User"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.ID
	return model.MarshalID(res)
}

func (ec *executionContext) _User_name(ctx context.Context, field graphql.CollectedField, obj *model.User) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "User"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.Name
	return graphql.MarshalString(res)
}

func (ec *executionContext) _User_created(ctx context.Context, field graphql.CollectedField, obj *model.User) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "User"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.Created
	return model.MarshalTimestamp(res)
}

func (ec *executionContext) _User_isBanned(ctx context.Context, field graphql.CollectedField, obj *model.User) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "User"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.IsBanned
	return graphql.MarshalBoolean(bool(res))
}

func (ec *executionContext) _User_primitiveResolver(ctx context.Context, field graphql.CollectedField, obj *model.User) graphql.Marshaler {
	ctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{
		Object: "User",
		Args:   nil,
		Field:  field,
	})
	return graphql.Defer(func() (ret graphql.Marshaler) {
		defer func() {
			if r := recover(); r != nil {
				userErr := ec.Recover(ctx, r)
				ec.Error(ctx, userErr)
				ret = graphql.Null
			}
		}()

		resTmp, err := ec.ResolverMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
			return ec.resolvers.User_primitiveResolver(ctx, obj)
		})
		if err != nil {
			ec.Error(ctx, err)
			return graphql.Null
		}
		if resTmp == nil {
			return graphql.Null
		}
		res := resTmp.(string)
		return graphql.MarshalString(res)
	})
}

func (ec *executionContext) _User_customResolver(ctx context.Context, field graphql.CollectedField, obj *model.User) graphql.Marshaler {
	ctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{
		Object: "User",
		Args:   nil,
		Field:  field,
	})
	return graphql.Defer(func() (ret graphql.Marshaler) {
		defer func() {
			if r := recover(); r != nil {
				userErr := ec.Recover(ctx, r)
				ec.Error(ctx, userErr)
				ret = graphql.Null
			}
		}()

		resTmp, err := ec.ResolverMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
			return ec.resolvers.User_customResolver(ctx, obj)
		})
		if err != nil {
			ec.Error(ctx, err)
			return graphql.Null
		}
		if resTmp == nil {
			return graphql.Null
		}
		res := resTmp.(model.Point)
		return res
	})
}

func (ec *executionContext) _User_address(ctx context.Context, field graphql.CollectedField, obj *model.User) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "User"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.Address
	return ec._Address(ctx, field.Selections, &res)
}

func (ec *executionContext) _User_tier(ctx context.Context, field graphql.CollectedField, obj *model.User) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "User"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.Tier
	return res
}

var __DirectiveImplementors = []string{"__Directive"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Directive(ctx context.Context, sel []query.Selection, obj *introspection.Directive) graphql.Marshaler {
	fields := graphql.CollectFields(ec.Doc, sel, __DirectiveImplementors, ec.Variables)

	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Directive")
		case "name":
			out.Values[i] = ec.___Directive_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___Directive_description(ctx, field, obj)
		case "locations":
			out.Values[i] = ec.___Directive_locations(ctx, field, obj)
		case "args":
			out.Values[i] = ec.___Directive_args(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___Directive_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Directive"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.Name()
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___Directive_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Directive"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.Description()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

func (ec *executionContext) ___Directive_locations(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Directive"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.Locations()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			return graphql.MarshalString(res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Directive_args(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Directive"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.Args()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___InputValue(ctx, field.Selections, res[idx1])
		}())
	}
	return arr1
}

var __EnumValueImplementors = []string{"__EnumValue"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___EnumValue(ctx context.Context, sel []query.Selection, obj *introspection.EnumValue) graphql.Marshaler {
	fields := graphql.CollectFields(ec.Doc, sel, __EnumValueImplementors, ec.Variables)

	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__EnumValue")
		case "name":
			out.Values[i] = ec.___EnumValue_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___EnumValue_description(ctx, field, obj)
		case "isDeprecated":
			out.Values[i] = ec.___EnumValue_isDeprecated(ctx, field, obj)
		case "deprecationReason":
			out.Values[i] = ec.___EnumValue_deprecationReason(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___EnumValue_name(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__EnumValue"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.Name()
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___EnumValue_description(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__EnumValue"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.Description()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

func (ec *executionContext) ___EnumValue_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__EnumValue"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.IsDeprecated()
	return graphql.MarshalBoolean(res)
}

func (ec *executionContext) ___EnumValue_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__EnumValue"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.DeprecationReason()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

var __FieldImplementors = []string{"__Field"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Field(ctx context.Context, sel []query.Selection, obj *introspection.Field) graphql.Marshaler {
	fields := graphql.CollectFields(ec.Doc, sel, __FieldImplementors, ec.Variables)

	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Field")
		case "name":
			out.Values[i] = ec.___Field_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___Field_description(ctx, field, obj)
		case "args":
			out.Values[i] = ec.___Field_args(ctx, field, obj)
		case "type":
			out.Values[i] = ec.___Field_type(ctx, field, obj)
		case "isDeprecated":
			out.Values[i] = ec.___Field_isDeprecated(ctx, field, obj)
		case "deprecationReason":
			out.Values[i] = ec.___Field_deprecationReason(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___Field_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Field"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.Name()
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___Field_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Field"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.Description()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

func (ec *executionContext) ___Field_args(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Field"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.Args()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___InputValue(ctx, field.Selections, res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Field_type(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Field"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.Type()
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Field"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.IsDeprecated()
	return graphql.MarshalBoolean(res)
}

func (ec *executionContext) ___Field_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Field"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.DeprecationReason()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

var __InputValueImplementors = []string{"__InputValue"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___InputValue(ctx context.Context, sel []query.Selection, obj *introspection.InputValue) graphql.Marshaler {
	fields := graphql.CollectFields(ec.Doc, sel, __InputValueImplementors, ec.Variables)

	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__InputValue")
		case "name":
			out.Values[i] = ec.___InputValue_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___InputValue_description(ctx, field, obj)
		case "type":
			out.Values[i] = ec.___InputValue_type(ctx, field, obj)
		case "defaultValue":
			out.Values[i] = ec.___InputValue_defaultValue(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___InputValue_name(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__InputValue"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.Name()
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___InputValue_description(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__InputValue"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.Description()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

func (ec *executionContext) ___InputValue_type(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__InputValue"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.Type()
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func (ec *executionContext) ___InputValue_defaultValue(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__InputValue"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.DefaultValue()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

var __SchemaImplementors = []string{"__Schema"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Schema(ctx context.Context, sel []query.Selection, obj *introspection.Schema) graphql.Marshaler {
	fields := graphql.CollectFields(ec.Doc, sel, __SchemaImplementors, ec.Variables)

	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Schema")
		case "types":
			out.Values[i] = ec.___Schema_types(ctx, field, obj)
		case "queryType":
			out.Values[i] = ec.___Schema_queryType(ctx, field, obj)
		case "mutationType":
			out.Values[i] = ec.___Schema_mutationType(ctx, field, obj)
		case "subscriptionType":
			out.Values[i] = ec.___Schema_subscriptionType(ctx, field, obj)
		case "directives":
			out.Values[i] = ec.___Schema_directives(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___Schema_types(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Schema"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.Types()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___Type(ctx, field.Selections, res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Schema_queryType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Schema"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.QueryType()
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_mutationType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Schema"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.MutationType()
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_subscriptionType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Schema"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.SubscriptionType()
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_directives(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Schema"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.Directives()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___Directive(ctx, field.Selections, res[idx1])
		}())
	}
	return arr1
}

var __TypeImplementors = []string{"__Type"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Type(ctx context.Context, sel []query.Selection, obj *introspection.Type) graphql.Marshaler {
	fields := graphql.CollectFields(ec.Doc, sel, __TypeImplementors, ec.Variables)

	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Type")
		case "kind":
			out.Values[i] = ec.___Type_kind(ctx, field, obj)
		case "name":
			out.Values[i] = ec.___Type_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___Type_description(ctx, field, obj)
		case "fields":
			out.Values[i] = ec.___Type_fields(ctx, field, obj)
		case "interfaces":
			out.Values[i] = ec.___Type_interfaces(ctx, field, obj)
		case "possibleTypes":
			out.Values[i] = ec.___Type_possibleTypes(ctx, field, obj)
		case "enumValues":
			out.Values[i] = ec.___Type_enumValues(ctx, field, obj)
		case "inputFields":
			out.Values[i] = ec.___Type_inputFields(ctx, field, obj)
		case "ofType":
			out.Values[i] = ec.___Type_ofType(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___Type_kind(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Type"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.Kind()
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___Type_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Type"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.Name()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

func (ec *executionContext) ___Type_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Type"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.Description()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

func (ec *executionContext) ___Type_fields(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	args := map[string]interface{}{}
	var arg0 bool
	if tmp, ok := field.Args["includeDeprecated"]; ok {
		var err error
		arg0, err = graphql.UnmarshalBoolean(tmp)
		if err != nil {
			ec.Error(ctx, err)
			return graphql.Null
		}
	}
	args["includeDeprecated"] = arg0
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Type"
	rctx.Args = args
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.Fields(args["includeDeprecated"].(bool))
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___Field(ctx, field.Selections, res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_interfaces(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Type"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.Interfaces()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___Type(ctx, field.Selections, res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_possibleTypes(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Type"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.PossibleTypes()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___Type(ctx, field.Selections, res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_enumValues(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	args := map[string]interface{}{}
	var arg0 bool
	if tmp, ok := field.Args["includeDeprecated"]; ok {
		var err error
		arg0, err = graphql.UnmarshalBoolean(tmp)
		if err != nil {
			ec.Error(ctx, err)
			return graphql.Null
		}
	}
	args["includeDeprecated"] = arg0
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Type"
	rctx.Args = args
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.EnumValues(args["includeDeprecated"].(bool))
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___EnumValue(ctx, field.Selections, res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_inputFields(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Type"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.InputFields()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___InputValue(ctx, field.Selections, res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_ofType(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Type"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	res := obj.OfType()
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func UnmarshalSearchArgs(v interface{}) (model.SearchArgs, error) {
	var it model.SearchArgs
	var asMap = v.(map[string]interface{})

	for k, v := range asMap {
		switch k {
		case "location":
			var err error
			var ptr1 model.Point
			if v != nil {
				err = (&ptr1).UnmarshalGQL(v)
				it.Location = &ptr1
			}

			if err != nil {
				return it, err
			}
		case "createdAfter":
			var err error
			var ptr1 time.Time
			if v != nil {
				ptr1, err = model.UnmarshalTimestamp(v)
				it.CreatedAfter = &ptr1
			}

			if err != nil {
				return it, err
			}
		case "isBanned":
			var err error

			var castTmp bool
			castTmp, err = graphql.UnmarshalBoolean(v)
			it.IsBanned = model.Banned(castTmp)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func (ec *executionContext) introspectSchema() *introspection.Schema {
	return introspection.WrapSchema(parsedSchema)
}

func (ec *executionContext) introspectType(name string) *introspection.Type {
	t := parsedSchema.Resolve(name)
	if t == nil {
		return nil
	}
	return introspection.WrapType(t)
}

var parsedSchema = schema.MustParse(`type Query {
    user(id: ID!): User
    search(input: SearchArgs = {location: "37,144", isBanned: false}): [User!]!
}

type User {
    id: ID!
    name: String!
    created: Timestamp
    isBanned: Boolean!
    primitiveResolver: String!
    customResolver: Point!
    address: Address
    tier: Tier
}

type Address {
    id: ID!
    location: Point
}

input SearchArgs {
    location: Point
    createdAfter: Timestamp
    isBanned: Boolean
}

enum Tier {
    A
    B
    C
}

scalar Timestamp
scalar Point
`)