package lengconv

import "fmt"

type Meter float64
type Foot float64

func (m Meter) String() string { return fmt.Sprintf("%.2gm", m) }
func (f Foot) String() string  { return fmt.Sprintf("%.2gf", f) }
