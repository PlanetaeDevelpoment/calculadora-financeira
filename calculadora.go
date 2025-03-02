package main

import "time"

const (
	baseAviso   = 30
	baseSalario = 30
)

func Date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func SaldoSalario(salario float64, dataDemissao time.Time) (saldo float64) {
	dataAposAviso := dataDemissao.AddDate(0, 0, baseAviso+3)
	return salario * (float64(dataAposAviso.Day()) / baseSalario)
}
