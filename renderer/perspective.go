package renderer

type Perspective struct{
    Enabled bool
    FovY struct {
        Deg float64
        RadHalf float64
    }
    ZFar float64
    ZNear float64
}
