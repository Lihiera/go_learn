package lengconv

func MtoF(m Meter) Foot { return Foot(m * 3.28084) }
func FtoM(f Foot) Meter { return Meter(f * 0.3048) }
