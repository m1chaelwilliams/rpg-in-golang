package animations

type Animation struct {
	First        int
	Last         int
	Step         int     // how many indices do we move per frame
	SpeedInTps   float32 // how many ticks before next frame
	frameCounter float32
	frame        int
}

func (a *Animation) Update() {
	a.frameCounter -= 1.0
	if a.frameCounter < 0.0 {
		a.frameCounter = a.SpeedInTps
		a.frame += a.Step
		if a.frame > a.Last {
			// loop back to the beginning
			a.frame = a.First
		}
	}
}

func (a *Animation) Frame() int {
	return a.frame
}

func NewAnimation(first, last, step int, speed float32) *Animation {
	return &Animation{
		first,
		last,
		step,
		speed,
		speed,
		first,
	}
}
