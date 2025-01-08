package enums

type Role int

const (
  Siswa Role = iota
  Guru
  Admin
)

func (r Role) String() string {
  return [...]string{"Siswa", "Guru", "Admin"}[r]
}
