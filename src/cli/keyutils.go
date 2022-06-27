package cli

import (
	"errors"
	"strings"
)

// ecdsa-sha2-nistp521 AAAAE2VjZHNhLXNoYTItbmlzdHA1MjEAAAAIbmlzdHA1Mj
//EAAACFBACS5PzQRWQG9x+76VxFvmJkV3mfNwmERp0aoPIjJNcX7nUWFjUk2BLCQpp9l
//iCGrRtqPhLa0lqJl+84gn8DAIEPxwF41rf1gId8sgrhajMEmu64/zrKbwKYIUZATs1e
//9UBz1Ervjuguhhg6JUUjND06eU+D9mw60SGLmydSBlE+CQf62Q== pietro@go

func KeyVerify(key string) error {
	pb, err := split(key)
	if err != nil {
		return err
	}

	err = pb.algoVerify()
	if err != nil {
		return err
	}

	return nil
}

func split(key string) (PBKey, error) {
	ks := strings.Split(key, " ")
	pb := PBKey{}

	if len(ks) != 3 {
		return pb, errors.New("could not parse public key")
	}

	pb.Algo = ks[0]
	pb.Payload = ks[1]
	pb.Host = ks[2]

	return pb, nil
}

func (algo PBKey) algoVerify() error {
	/*
		dsa		ecdsa	ecdsa-sk	ed25519		ed25519-sk	rsa
		1024	256					Fixed		Fixed		1024
				384											3072
				521											4096
	*/

	types := []string{"ssh-rsa", "ecdsa-sha2-nistp256", "ecdsa-sha2-nistp384", "ecdsa-sha2-nistp521", "ssh-ed25519", "ssh-dss"}

	for i := range types {
		if algo.Algo == types[i] {
			return nil
		}
	}
	return errors.New("no algorithm matched")
}
