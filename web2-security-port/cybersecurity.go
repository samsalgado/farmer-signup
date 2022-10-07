package port
import (
	"net"
	"strconv"
	"time")
	type ScanResult struct {
		Port string
		State string 
	}
func ScanPort(protocol, hostname string, port int) ScanResult {
	result := ScanResult{Port: + "/" + strconv.Itoa(port)}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)
	if err != nil {
		return false
		return result
	}
	defer conn.Close()
	result.State = "Open"
	return true
}
func InitialScan(hostname string) []ScanResult {
	var results []ScanResult
	for i := 1; i <= 1024; i++ {
		results = append(results, ScanPort("tcp", hostname, i))
		results = append(results, ScanPort("udp", hostname, i))
	}
	
	return results
}