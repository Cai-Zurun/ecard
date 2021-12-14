package utils

/*
将int转成两字节的[]byte，信息序列号只占两字节
*/
func IntToSeqBytes(num int) (seq []byte) {
    seq = append(seq, byte(num/256))
    seq = append(seq, byte(num%256))
    return
}
