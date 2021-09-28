package doer

//go:generate mockgen -destination=../mocks/mock_doer.go -package=mocks mockgen-example/doer Doer

type Doer interface {
	DoSomething(int, string) error
}
