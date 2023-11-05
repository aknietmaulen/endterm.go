package main

type PaymentMethod interface {
	Pay(amount int) bool
}

type CreditCard struct {
	CardNumber  string
	CVV         string
	MoneyAmount int
}

func NewCreditCard(cardNumber, cvv string, moneyAmount int) *CreditCard {
	return &CreditCard{
		CardNumber:  cardNumber,
		CVV:         cvv,
		MoneyAmount: moneyAmount,
	}
}

func (c *CreditCard) Pay(amount int) bool {
	if c.MoneyAmount >= amount {
		c.MoneyAmount -= amount
		return true
	}
	return false
}

type PayPal struct {
	Email       string
	MoneyAmount int
}

func NewPayPal(email string, moneyAmount int) *PayPal {
	return &PayPal{
		Email:       email,
		MoneyAmount: moneyAmount,
	}
}

func (p *PayPal) Pay(amount int) bool {
	if p.MoneyAmount >= amount {
		p.MoneyAmount -= amount
		return true
	}
	return false
}
