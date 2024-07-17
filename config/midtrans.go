package config

import (
	"os"

	"github.com/veritrans/go-midtrans"
)

type MidClient struct {
	Client *midtrans.Client
}

func NewMidtransClient() *MidClient {
	var envTypeMidtrans midtrans.EnvironmentType

	envType := os.Getenv("MIDTRANS_ENVIRONMENT_TYPE")
	if envType == "production" {
		envTypeMidtrans = midtrans.Production
	} else {
		envTypeMidtrans = midtrans.Sandbox
	}

	midclient := midtrans.NewClient()
	midclient.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")
	midclient.ClientKey = os.Getenv("MIDTRANS_CLIENT_KEY")
	midclient.APIEnvType = envTypeMidtrans

	return &MidClient{
		Client: &midclient,
	}
}
