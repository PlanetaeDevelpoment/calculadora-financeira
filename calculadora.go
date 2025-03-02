package main

import (
	"time"
)

const (
	baseAviso   = 30
	baseSalario = 30
)

func Date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func SaldoSalario(salario float64, dataContratação, dataDemissao time.Time) (saldo float64) {
	porcentagemSaldo := getPorcentagemSaldo(dataContratação, dataDemissao)
	return salario * porcentagemSaldo
}

func getPorcentagemSaldo(dataContratação, dataDemissao time.Time) float64 {
	tempoDeAviso := getTempoDeAviso(dataContratação, dataDemissao)
	dataAposAviso := dataDemissao.AddDate(0, 0, tempoDeAviso)
	return (float64(dataAposAviso.Day()) / baseSalario)
}

func getTempoDeAviso(dataContratação, dataDemissao time.Time) int {
	anosContribuição := getAnosContribuição(dataContratação, dataDemissao)
	return baseAviso + (3 * anosContribuição)
}

func getAnosContribuição(inicio, fim time.Time) int {
	return int(fim.Sub(inicio).Hours() / (24 * 365))
}
