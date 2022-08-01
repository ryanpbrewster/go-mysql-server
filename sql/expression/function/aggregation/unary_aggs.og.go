// Code generated by optgen; DO NOT EDIT.

package aggregation

import (
	"fmt"

	"github.com/dolthub/go-mysql-server/sql/transform"

	"github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/expression"
)

type Avg struct {
	unaryAggBase
}

var _ sql.FunctionExpression = (*Avg)(nil)
var _ sql.Aggregation = (*Avg)(nil)
var _ sql.WindowAdaptableExpression = (*Avg)(nil)

func NewAvg(e sql.Expression) *Avg {
	return &Avg{
		unaryAggBase{
			UnaryExpression: expression.UnaryExpression{Child: e},
			functionName:    "Avg",
			description:     "returns the average value of expr in all rows.",
		},
	}
}

func (a *Avg) Type() sql.Type {
	return sql.Float64
}

func (a *Avg) IsNullable() bool {
	return true
}

func (a *Avg) String() string {
	return fmt.Sprintf("AVG(%s)", a.Child)
}

func (a *Avg) WithWindow(window *sql.WindowDefinition) (sql.Aggregation, error) {
	res, err := a.unaryAggBase.WithWindow(window)
	return &Avg{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *Avg) WithChildren(children ...sql.Expression) (sql.Expression, error) {
	res, err := a.unaryAggBase.WithChildren(children...)
	return &Avg{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *Avg) NewBuffer() (sql.AggregationBuffer, error) {
	child, err := transform.Clone(a.UnaryExpression.Child)
	if err != nil {
		return nil, err
	}
	return NewAvgBuffer(child), nil
}

func (a *Avg) NewWindowFunction() (sql.WindowFunction, error) {
	child, err := transform.Clone(a.UnaryExpression.Child)
	if err != nil {
		return nil, err
	}
	return NewAvgAgg(child).WithWindow(a.Window())
}

type Count struct {
	unaryAggBase
}

var _ sql.FunctionExpression = (*Count)(nil)
var _ sql.Aggregation = (*Count)(nil)
var _ sql.WindowAdaptableExpression = (*Count)(nil)

func NewCount(e sql.Expression) *Count {
	return &Count{
		unaryAggBase{
			UnaryExpression: expression.UnaryExpression{Child: e},
			functionName:    "Count",
			description:     "returns a count of the number of non-NULL values of expr in the rows retrieved by a SELECT statement.",
		},
	}
}

func (a *Count) Type() sql.Type {
	return sql.Int64
}

func (a *Count) IsNullable() bool {
	return false
}

func (a *Count) String() string {
	return fmt.Sprintf("COUNT(%s)", a.Child)
}

func (a *Count) WithWindow(window *sql.WindowDefinition) (sql.Aggregation, error) {
	res, err := a.unaryAggBase.WithWindow(window)
	return &Count{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *Count) WithChildren(children ...sql.Expression) (sql.Expression, error) {
	res, err := a.unaryAggBase.WithChildren(children...)
	return &Count{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *Count) NewBuffer() (sql.AggregationBuffer, error) {
	child, err := transform.Clone(a.UnaryExpression.Child)
	if err != nil {
		return nil, err
	}
	return NewCountBuffer(child), nil
}

func (a *Count) NewWindowFunction() (sql.WindowFunction, error) {
	child, err := transform.Clone(a.UnaryExpression.Child)
	if err != nil {
		return nil, err
	}
	return NewCountAgg(child).WithWindow(a.Window())
}

type CountDistinct struct {
	naryAggBase
}

var _ sql.FunctionExpression = (*CountDistinct)(nil)
var _ sql.Aggregation = (*CountDistinct)(nil)
var _ sql.WindowAdaptableExpression = (*CountDistinct)(nil)

func NewCountDistinct(exprs []sql.Expression) *CountDistinct {
	return &CountDistinct{
		naryAggBase{
			NaryExpression: expression.NaryExpression{ChildExpressions: exprs},
			functionName:   "CountDistinct",
			description:    "returns the number of distinct values in a result set.",
		},
	}
}

func (a *CountDistinct) Type() sql.Type {
	return sql.Int64
}

func (a *CountDistinct) IsNullable() bool {
	return false
}

func (a *CountDistinct) String() string {
	return fmt.Sprintf("COUNTDISTINCT(%s)", a.ChildExpressions)
}

func (a *CountDistinct) WithWindow(window *sql.WindowDefinition) (sql.Aggregation, error) {
	res, err := a.naryAggBase.WithWindow(window)
	return &CountDistinct{naryAggBase: *res.(*naryAggBase)}, err
}

func (a *CountDistinct) WithChildren(children ...sql.Expression) (sql.Expression, error) {
	res, err := a.naryAggBase.WithChildren(children...)
	return &CountDistinct{naryAggBase: *res.(*naryAggBase)}, err
}

func (a *CountDistinct) NewBuffer() (sql.AggregationBuffer, error) {
	exprs := make([]sql.Expression, len(a.ChildExpressions))
	for i, expr := range a.ChildExpressions {
		child, err := transform.Clone(expr)
		if err != nil {
			return nil, err
		}
		exprs[i] = child
	}

	return NewCountDistinctBuffer(exprs), nil
}

func (a *CountDistinct) NewWindowFunction() (sql.WindowFunction, error) {
	child, err := transform.Clone(a.ChildExpressions[0])
	if err != nil {
		return nil, err
	}
	return NewCountDistinctAgg(child).WithWindow(a.Window())
}

type First struct {
	unaryAggBase
}

var _ sql.FunctionExpression = (*First)(nil)
var _ sql.Aggregation = (*First)(nil)
var _ sql.WindowAdaptableExpression = (*First)(nil)

func NewFirst(e sql.Expression) *First {
	return &First{
		unaryAggBase{
			UnaryExpression: expression.UnaryExpression{Child: e},
			functionName:    "First",
			description:     "returns the first value in a sequence of elements of an aggregation.",
		},
	}
}

func (a *First) Type() sql.Type {
	return a.Child.Type()
}

func (a *First) IsNullable() bool {
	return false
}

func (a *First) String() string {
	return fmt.Sprintf("FIRST(%s)", a.Child)
}

func (a *First) WithWindow(window *sql.WindowDefinition) (sql.Aggregation, error) {
	res, err := a.unaryAggBase.WithWindow(window)
	return &First{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *First) WithChildren(children ...sql.Expression) (sql.Expression, error) {
	res, err := a.unaryAggBase.WithChildren(children...)
	return &First{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *First) NewBuffer() (sql.AggregationBuffer, error) {
	child, err := transform.Clone(a.UnaryExpression.Child)
	if err != nil {
		return nil, err
	}
	return NewFirstBuffer(child), nil
}

func (a *First) NewWindowFunction() (sql.WindowFunction, error) {
	child, err := transform.Clone(a.UnaryExpression.Child)
	if err != nil {
		return nil, err
	}
	return NewFirstAgg(child).WithWindow(a.Window())
}

type Last struct {
	unaryAggBase
}

var _ sql.FunctionExpression = (*Last)(nil)
var _ sql.Aggregation = (*Last)(nil)
var _ sql.WindowAdaptableExpression = (*Last)(nil)

func NewLast(e sql.Expression) *Last {
	return &Last{
		unaryAggBase{
			UnaryExpression: expression.UnaryExpression{Child: e},
			functionName:    "Last",
			description:     "returns the last value in a sequence of elements of an aggregation.",
		},
	}
}

func (a *Last) Type() sql.Type {
	return a.Child.Type()
}

func (a *Last) IsNullable() bool {
	return false
}

func (a *Last) String() string {
	return fmt.Sprintf("LAST(%s)", a.Child)
}

func (a *Last) WithWindow(window *sql.WindowDefinition) (sql.Aggregation, error) {
	res, err := a.unaryAggBase.WithWindow(window)
	return &Last{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *Last) WithChildren(children ...sql.Expression) (sql.Expression, error) {
	res, err := a.unaryAggBase.WithChildren(children...)
	return &Last{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *Last) NewBuffer() (sql.AggregationBuffer, error) {
	child, err := transform.Clone(a.UnaryExpression.Child)
	if err != nil {
		return nil, err
	}
	return NewLastBuffer(child), nil
}

func (a *Last) NewWindowFunction() (sql.WindowFunction, error) {
	child, err := transform.Clone(a.UnaryExpression.Child)
	if err != nil {
		return nil, err
	}
	return NewLastAgg(child).WithWindow(a.Window())
}

type Max struct {
	unaryAggBase
}

var _ sql.FunctionExpression = (*Max)(nil)
var _ sql.Aggregation = (*Max)(nil)
var _ sql.WindowAdaptableExpression = (*Max)(nil)

func NewMax(e sql.Expression) *Max {
	return &Max{
		unaryAggBase{
			UnaryExpression: expression.UnaryExpression{Child: e},
			functionName:    "Max",
			description:     "returns the maximum value of expr in all rows.",
		},
	}
}

func (a *Max) Type() sql.Type {
	return a.Child.Type()
}

func (a *Max) IsNullable() bool {
	return false
}

func (a *Max) String() string {
	return fmt.Sprintf("MAX(%s)", a.Child)
}

func (a *Max) WithWindow(window *sql.WindowDefinition) (sql.Aggregation, error) {
	res, err := a.unaryAggBase.WithWindow(window)
	return &Max{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *Max) WithChildren(children ...sql.Expression) (sql.Expression, error) {
	res, err := a.unaryAggBase.WithChildren(children...)
	return &Max{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *Max) NewBuffer() (sql.AggregationBuffer, error) {
	child, err := transform.Clone(a.UnaryExpression.Child)
	if err != nil {
		return nil, err
	}
	return NewMaxBuffer(child), nil
}

func (a *Max) NewWindowFunction() (sql.WindowFunction, error) {
	child, err := transform.Clone(a.UnaryExpression.Child)
	if err != nil {
		return nil, err
	}
	return NewMaxAgg(child).WithWindow(a.Window())
}

type Min struct {
	unaryAggBase
}

var _ sql.FunctionExpression = (*Min)(nil)
var _ sql.Aggregation = (*Min)(nil)
var _ sql.WindowAdaptableExpression = (*Min)(nil)

func NewMin(e sql.Expression) *Min {
	return &Min{
		unaryAggBase{
			UnaryExpression: expression.UnaryExpression{Child: e},
			functionName:    "Min",
			description:     "returns the minimum value of expr in all rows.",
		},
	}
}

func (a *Min) Type() sql.Type {
	return a.Child.Type()
}

func (a *Min) IsNullable() bool {
	return false
}

func (a *Min) String() string {
	return fmt.Sprintf("MIN(%s)", a.Child)
}

func (a *Min) WithWindow(window *sql.WindowDefinition) (sql.Aggregation, error) {
	res, err := a.unaryAggBase.WithWindow(window)
	return &Min{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *Min) WithChildren(children ...sql.Expression) (sql.Expression, error) {
	res, err := a.unaryAggBase.WithChildren(children...)
	return &Min{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *Min) NewBuffer() (sql.AggregationBuffer, error) {
	child, err := transform.Clone(a.UnaryExpression.Child)
	if err != nil {
		return nil, err
	}
	return NewMinBuffer(child), nil
}

func (a *Min) NewWindowFunction() (sql.WindowFunction, error) {
	child, err := transform.Clone(a.UnaryExpression.Child)
	if err != nil {
		return nil, err
	}
	return NewMinAgg(child).WithWindow(a.Window())
}

type Sum struct {
	unaryAggBase
}

var _ sql.FunctionExpression = (*Sum)(nil)
var _ sql.Aggregation = (*Sum)(nil)
var _ sql.WindowAdaptableExpression = (*Sum)(nil)

func NewSum(e sql.Expression) *Sum {
	return &Sum{
		unaryAggBase{
			UnaryExpression: expression.UnaryExpression{Child: e},
			functionName:    "Sum",
			description:     "returns the sum of expr in all rows",
		},
	}
}

func (a *Sum) Type() sql.Type {
	return sql.Float64
}

func (a *Sum) IsNullable() bool {
	return false
}

func (a *Sum) String() string {
	return fmt.Sprintf("SUM(%s)", a.Child)
}

func (a *Sum) WithWindow(window *sql.WindowDefinition) (sql.Aggregation, error) {
	res, err := a.unaryAggBase.WithWindow(window)
	return &Sum{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *Sum) WithChildren(children ...sql.Expression) (sql.Expression, error) {
	res, err := a.unaryAggBase.WithChildren(children...)
	return &Sum{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *Sum) NewBuffer() (sql.AggregationBuffer, error) {
	child, err := transform.Clone(a.UnaryExpression.Child)
	if err != nil {
		return nil, err
	}
	return NewSumBuffer(child), nil
}

func (a *Sum) NewWindowFunction() (sql.WindowFunction, error) {
	child, err := transform.Clone(a.UnaryExpression.Child)
	if err != nil {
		return nil, err
	}
	return NewSumAgg(child).WithWindow(a.Window())
}

type JsonArray struct {
	unaryAggBase
}

var _ sql.FunctionExpression = (*JsonArray)(nil)
var _ sql.Aggregation = (*JsonArray)(nil)
var _ sql.WindowAdaptableExpression = (*JsonArray)(nil)

func NewJsonArray(e sql.Expression) *JsonArray {
	return &JsonArray{
		unaryAggBase{
			UnaryExpression: expression.UnaryExpression{Child: e},
			functionName:    "JsonArray",
			description:     "returns result set as a single JSON array.",
		},
	}
}

func (a *JsonArray) Type() sql.Type {
	return sql.JSON
}

func (a *JsonArray) IsNullable() bool {
	return false
}

func (a *JsonArray) String() string {
	return fmt.Sprintf("JSON_ARRAYAGG(%s)", a.Child)
}

func (a *JsonArray) WithWindow(window *sql.WindowDefinition) (sql.Aggregation, error) {
	res, err := a.unaryAggBase.WithWindow(window)
	return &JsonArray{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *JsonArray) WithChildren(children ...sql.Expression) (sql.Expression, error) {
	res, err := a.unaryAggBase.WithChildren(children...)
	return &JsonArray{unaryAggBase: *res.(*unaryAggBase)}, err
}

func (a *JsonArray) NewBuffer() (sql.AggregationBuffer, error) {
	child, err := transform.Clone(a.UnaryExpression.Child)
	if err != nil {
		return nil, err
	}
	return NewJsonArrayBuffer(child), nil
}

func (a *JsonArray) NewWindowFunction() (sql.WindowFunction, error) {
	child, err := transform.Clone(a.UnaryExpression.Child)
	if err != nil {
		return nil, err
	}
	return NewJsonArrayAgg(child).WithWindow(a.Window())
}
