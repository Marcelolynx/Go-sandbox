package entity

import "fmt"

type Fulano struct {
	PrimeiroNome string
	Sobrenome    string
}

func (v *Fulano) Validate() error {
	if v.PrimeiroNome == "" {
		return fmt.Errorf("Primeiro nome é obrigatório")
	}
	if v.Sobrenome == "" {
		return fmt.Errorf("Sobrenome é obrigatório")
	}
	return nil
}

func NewFulano(primeiroNome string, sobreNome string) (*Fulano, error) {
	fulano := &Fulano{
		PrimeiroNome: primeiroNome,
		Sobrenome:    sobreNome,
	}
	err := fulano.Validate()
	if err != nil {
		return nil, err
	}
	return fulano, nil
}
