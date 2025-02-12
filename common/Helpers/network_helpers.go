package Helpers

import (
	"bytes"
	"crypto/tls"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/gorilla/websocket"

	pb "govenv/api/proto/safeSocketMsg"
	"govenv/pkg/common/PyHelpers"

	"google.golang.org/protobuf/proto"
)

// ####################################################################
// ################# Check if service is started ######################
func IsServiceListen(server string, port string, timeout int, ws bool) bool {
	_timeout := time.Second * time.Duration(timeout)
	if ws {
		uri := fmt.Sprintf("ws://%s:%s", server, port)
		dialer := websocket.Dialer{HandshakeTimeout: _timeout}
		conn, _, err := dialer.Dial(uri, nil)
		if err != nil {
			return false
		}
		defer conn.Close()
		return true
	} else {
		address := fmt.Sprintf("%s:%s", server, port)
		conn, err := net.DialTimeout("tcp", address, _timeout)
		if err != nil {
			return false
		}
		defer conn.Close()
		return true
	}
}

// ################# Check if service is started ######################
// ####################################################################
// ########### Insure correct exchange between sockets  ###############
type SafeSocket struct {
	Name         string
	Conn         net.Conn
	SelfHost     string
	SelfPort     string
	SelfSockName string
	SelfPublicIp string
}

func NewSafeSocket(name string, conn net.Conn) *SafeSocket {
	if name == "" {
		name = "safeSocket"
	}
	selfHost, selfPort, _ := net.SplitHostPort(conn.LocalAddr().String())
	selfSockName, _ := os.Hostname()
	selfPublicIp := os.Getenv("MY_PUBLIC_IP")
	return &SafeSocket{
		Name:         name,
		Conn:         conn,
		SelfHost:     selfHost,
		SelfPort:     selfPort,
		SelfSockName: selfSockName,
		SelfPublicIp: selfPublicIp,
	}
}

func (safeSocket *SafeSocket) SayHelloId() {
	// must not panic ! no error should occurs here !!
	// clients should send well formatted datas
	helloMsg, _ := proto.Marshal(
		&pb.HelloMsg{
			Name:         safeSocket.Name,
			SelfHost:     safeSocket.SelfHost,
			SelfPort:     safeSocket.SelfPort,
			SelfSockName: safeSocket.SelfSockName,
			SelfPublicIp: safeSocket.SelfPublicIp,
		})
	safeSocket.SendData(helloMsg)
}

func (safeSocket *SafeSocket) WaitHelloId() (string, string, string, string, string) {
	// must not panic ! no error should occurs here !!
	// clients should send well formatted datas
	dataSer := safeSocket.ReceiveData()
	data := &pb.HelloMsg{}
	_ = proto.Unmarshal(dataSer, data)
	// return infos from sender
	return data.Name, data.SelfHost, data.SelfPort, data.SelfSockName, data.SelfPublicIp
}

func (safeSocket *SafeSocket) SendData(serializedData []byte) {
	defer safeSocket.Conn.Close()
	length := uint32(len(serializedData))
	lenBuffer := new(bytes.Buffer)
	err := binary.Write(lenBuffer, binary.BigEndian, length)
	if err != nil {
		safeSocket.Conn.Close()
	}
	_, err = safeSocket.Conn.Write(lenBuffer.Bytes())
	if err != nil {
		safeSocket.Conn.Close()
	}
	_, err = safeSocket.Conn.Write(serializedData)
	if err != nil {
		safeSocket.Conn.Close()
	}
}

func (safeSocket *SafeSocket) ReceiveData() []byte {
	defer safeSocket.Conn.Close()
	chunk := make([]byte, 4)
	_, err := safeSocket.Conn.Read(chunk)
	if err != nil {
		safeSocket.Conn.Close()
	}
	slen := int(binary.BigEndian.Uint32(chunk))
	chunk = make([]byte, 0, slen)
	for len(chunk) < slen {
		buffer := make([]byte, slen-len(chunk))
		n, err := safeSocket.Conn.Read(buffer)
		if err != nil {
			safeSocket.Conn.Close()
		}
		chunk = append(chunk, buffer[:n]...)
	}
	return chunk
}

func (safeSocket *SafeSocket) Getsockname() (string, string) {
	return safeSocket.SelfHost, safeSocket.SelfPort
}

// ########### Insure correct exchange between sockets  ###############
// ####################################################################
// ############# create custom generic socket connection ##############
//
//	// add memory map param ?
func MySocket(name string, server string, port string, timeout int) *SafeSocket {
	address := fmt.Sprintf("%s:%s", server, port)
	_timeout := time.Duration(timeout) * time.Second
	if server != "127.0.0.1" && strings.ToLower(server) != "localhost" {
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         server,
		}
		conn, err := tls.Dial("tcp", address, tlsConfig)
		if err != nil {
			conn.Close()
			return nil
		}
		return NewSafeSocket(name, conn)
	} else {
		conn, err := net.DialTimeout("tcp", address, _timeout)
		if err != nil {
			conn.Close()
			return nil
		}
		return NewSafeSocket(name, conn)
	}
}

// ############# create custom generic socket connection ##############
// ####################################################################
// ######################## get My Public IP  #########################
// # https://unix.stackexchange.com/questions/22615/how-can-i-get-my-external-ip-address-in-a-shell-script/81699#81699
// # OpenDNS (since 2009)
// #$ dig @resolver3.opendns.com myip.opendns.com +short
// #$ dig @resolver4.opendns.com myip.opendns.com +short
// ## Akamai (since 2009)
// #$ dig @ns1-1.akamaitech.net ANY whoami.akamai.net +short
// ## Akamai approximate
// ## NOTE: This returns only an approximate IP from your block,
// ## but has the benefit of working with private DNS proxies.
// #$ dig +short TXT whoami.ds.akahelp.net
// ## Google (since 2010)
// ## Supports IPv6 + IPv4, use -4 or -6 to force one.
// #$ dig @ns1.google.com TXT o-o.myaddr.l.google.com +short
func getMyPublicIp() string {
	myPublicIP := ""
	var DnsServersCmd []string
	if runtime.GOOS == "windows" {
		DnsServersCmd = []string{
			"nslookup myip.opendns.com resolver1.opendns.com",
			"nslookup myip.opendns.com resolver2.opendns.com",
			"nslookup myip.opendns.com resolver3.opendns.com",
			"nslookup myip.opendns.com resolver4.opendns.com",
			"nslookup whoami.akamai.net ns1-1.akamaitech.net",
		}
	} else {
		DnsServersCmd = []string{
			"dig +time=1 +tries=1 @ns1.google.com TXT o-o.myaddr.l.google.com +short",
			"dig +time=1 +tries=1 @1.0.0.1 TXT whoami.cloudflare +short",
			"dig +time=1 +tries=1 @1.1.1.1 TXT whoami.cloudflare +short",
			"dig +time=1 +tries=1 @ns1-1.akamaitech.net ANY whoami.akamai.net +short",
		}
	}
	IpPattern, _ := regexp.Compile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`)
	cpt := 0
	for {
		if cpt > 20 {
			break
		}
		cmdArgs := strings.Fields(DnsServersCmd[cpt%len(DnsServersCmd)])
		cmd := &exec.Cmd{}
		if runtime.GOOS != "windows" {
			cmd = exec.Command(cmdArgs[0], cmdArgs[1], cmdArgs[2])
		} else {
			cmd = exec.Command(cmdArgs[0], cmdArgs[1], cmdArgs[2], cmdArgs[3], cmdArgs[4], cmdArgs[5], cmdArgs[6])
		}
		output, err := cmd.CombinedOutput()
		strOutput := PyHelpers.Trim(string(output), false)
		if err != nil {
			continue
		}
		if IpPattern.MatchString(strOutput) {
			myPublicIP = IpPattern.FindString(strOutput)
			return myPublicIP
		}
		cpt += 1
	}
	return ""
}

// ######################## get My Public IP  #########################
//#####################################################################

// init func
func init() {
	if os.Getenv("MY_PUBLIC_IP") == "" {
		fmt.Printf("\nMY_PUBLIC_IP should be set manually in OS env...\n")
		os.Setenv("MY_PUBLIC_IP", getMyPublicIp())
	}
}
