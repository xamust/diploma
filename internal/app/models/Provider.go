package models

import "fmt"

type Provider struct {
	ProviderList []string
}

func getProvider() *Provider {
	return &Provider{ProviderList: []string{
		"Topolo",
		"Rond",
		"Kildy",
	}}
}

func getProviderVoiceCall() *Provider {
	return &Provider{ProviderList: []string{
		"TransparentCalls",
		"E-Voice",
		"JustPhone",
	}}
}

func getProviderEmail() *Provider {
	return &Provider{ProviderList: []string{
		"Gmail",
		"Yahoo",
		"Hotmail",
		"MSN",
		"Orange",
		"Comcast",
		"AOL",
		"Live",
		"RediffMail",
		"GMX",
		"Protonmail",
		"Yandex",
		"Mail.ru",
	}}
}

func SearchProvider(provider string) error {
	for _, v := range getProvider().ProviderList {
		if provider == v {
			return nil
		}
	}
	return fmt.Errorf("данный провайдер отсутствует в системе")
}

func SearchProviderVoiceCall(provider string) error {
	for _, v := range getProviderVoiceCall().ProviderList {
		if provider == v {
			return nil
		}
	}
	return fmt.Errorf("данный провайдер отсутствует в системе")
}

func SearchProviderEmail(provider string) error {
	for _, v := range getProviderEmail().ProviderList {
		if provider == v {
			return nil
		}
	}
	return fmt.Errorf("данный провайдер отсутствует в системе")
}
