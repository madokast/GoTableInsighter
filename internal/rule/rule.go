package rule

type Predicate struct {
}

type Rule struct{}

type RuleBatch struct {
	PredicateIndexer []Predicate
	CommonLhs        []int
	Rules            []Rule
}
