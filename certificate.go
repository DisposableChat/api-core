package core

import "os"

func GetCert(certFilePath, keyFilePath string) (cert []byte, key []byte, err error) {
	cert, err = os.ReadFile(certFilePath)
	if err != nil {
		return
	}

	key, err = os.ReadFile(keyFilePath)
	if err != nil {
		return
	}

	return
}

func GetCertPath() (certFilePath, keyFilePath string) {
	certFilePath = os.Getenv("SSL_CERTIFICATE_PATH")
	keyFilePath = os.Getenv("SSL_CERTIFICATE_KEY_PATH")

	return
}

func (s *Server) SetCertificates(certFile string, keyFile string) (string, string) {
	s.certFile = certFile
	s.keyFile = keyFile

	return certFile, keyFile
}