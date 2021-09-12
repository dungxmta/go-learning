package helper

import "net"

// generate the next ip

func NextIPStr(ip string) string {
	i := net.ParseIP(ip)
	if i == nil {
		return ""
	}
	ni := NextIP(i, 1)
	return ni.String()
}

func NextIP(ip net.IP, inc uint) net.IP {
	i := ip.To4()
	v := uint(i[0])<<24 + uint(i[1])<<16 + uint(i[2])<<8 + uint(i[3])
	v += inc
	v3 := byte(v & 0xFF)
	v2 := byte((v >> 8) & 0xFF)
	v1 := byte((v >> 16) & 0xFF)
	v0 := byte((v >> 24) & 0xFF)
	return net.IPv4(v0, v1, v2, v3)
}