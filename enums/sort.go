package enums

type SORT int

const (
  ASC SORT = iota
  DESC
)

func (s SORT) String() string {
  return [...]string{"ASC", "DESC"}[s]
}
