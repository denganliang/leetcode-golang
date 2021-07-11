

type MaxQueue struct {
    q []int  //queue
    m []int  //max
}


func Constructor() MaxQueue {
    mq := MaxQueue{}
    return mq
}


func (this *MaxQueue) Max_value() int {
    if len(this.m) == 0  {
        return -1
    }
    return this.m[0]
}


func (this *MaxQueue) Push_back(value int)  {
    this.q = append(this.q,value)
    if len(this.m) == 0 {
        this.m = append(this.m, value)
    } else {
        i := len(this.m)-1 ; 
        for i>=0 {
            if this.m[i] > value {
                break;
            }
            i--
        }
        this.m = this.m[:i+1]
        this.m = append(this.m, value)
    }
}


func (this *MaxQueue) Pop_front() int {
    if len(this.q) == 0 {
        return -1
    }
    v := this.q[0]
    this.q = this.q[1:]
    if v == this.m[0] {
        this.m = this.m[1:]
    }
    return v
}


/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
