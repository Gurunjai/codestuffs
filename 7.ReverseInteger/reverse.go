package reverse

func ReverseInteger(x int) int {
    const MaxInt = int32(^uint32(0) >> 1)
    const MinInt = -MaxInt - 1
    
    out := 0
    for x != 0 {
        out = (out * 10) + (x % 10)
        
        if (out <= int(MinInt) || out > int(MaxInt)) {
            return 0
        }
        x /= 10
    }
    
    return out
}