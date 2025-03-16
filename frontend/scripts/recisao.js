const inputMapping = {
    salário: 'salário',
    contratação: 'contratação',
    demissão: 'demissão',
    avisoPrévio: 'aviso',
    motivoEncerramento: 'motivo',
    fériasVencidas: 'férias-vencidas'
}
const resultMapping = {
    SaldoSalário: "saldo-salário",
    SaldoFérias: "saldo-férias",
    SaldoAviso: "saldo-aviso",
    SaldoDécimoTerceiro: "saldo-13",
    SaldoFGTS: "saldo-fgts",
    SaldoTotal: "saldo-total"
}

const demissãoPorJustaCausa = 5;
const exemploResultado = {
    SaldoFGTS: 1102.4825599268518,
    SaldoAviso: 0,
    SaldoSalário: 92.5,
    SaldoFérias: 1444.4444444444443,
    SaldoDécimoTerceiro: 0,
    SaldoTotal: 2639.427004371296
}
const backendUrl = 'https://tsowo4cg2uqbjf6rc6palclojm0phyrg.lambda-url.us-east-1.on.aws/';

async function interceptSubmit(event) {
    event.preventDefault();
    const requisição = {
        Salário: getValorInput(inputMapping.salário),
        DataContratação: getValorInput(inputMapping.contratação),
        DataDemissão: getValorInput(inputMapping.demissão),
        AvisoPrevio: getValorInput(inputMapping.avisoPrévio),
        Motivo: getValorInput(inputMapping.motivoEncerramento),
        fériasVencidas: getValorInput(inputMapping.fériasVencidas)
    }
    const error = validate(requisição);
    if (error) {
        alert(error);
        return;
    }
    const resposta = await enviarRequisição(requisição)
    preencherResultado(resposta);
}

function getValorInput(inputId) {
    const input = document.getElementById(inputId);
    if (input.type === 'checkbox') {
        return input.checked;
    }
    return input.value;
}

function validate(requisição) {
    if (requisição.Salário <= 0) {
        return 'Salário deve ser maior que zero';
    }
    if (requisição.DataContratação === "") {
        return 'Data de contratação deve ser informada';
    }
    if (requisição.DataDemissão === "") {
        return 'Data de demissão deve ser informada';
    }
    if (requisição.AvisoPrevio === "") {
        return 'Aviso Prévio deve ser informado';
    }
    if (requisição.Motivo === "") {
        return 'Motivo deve ser informado';
    }
    requisição.AvisoPrevio = +requisição.AvisoPrevio;
    requisição.Motivo = +requisição.Motivo;
    requisição.JustaCausa = requisição.Motivo === demissãoPorJustaCausa;
    return null;
}

async function enviarRequisição(requisição) {
    console.log(requisição);
    // return await fetch(`${backendUrl}`, {
    //     method: 'POST',
    //     headers: {
    //         'Content-Type': 'application/json'
    //     },
    //     body: JSON.stringify(requisição)
    // }).then(response => response.json())
    //     .then(data => {
    //         return data;
    //     })
    //     .catch(error => {
    //         console.log('Error:', error);
    //     });
    return exemploResultado;
}

function preencherResultado(resposta) {
    const resultado = document.getElementById('resultados');
    for (const key in resultMapping) {
        const valor = resposta[key];
        const elemento = document.getElementById(resultMapping[key]);
        if (valor >= 0) {
            elemento.innerHTML = "R$: " + valor.toFixed(2);
        } else {
            elemento.innerHTML = "- R$: " + Math.abs(valor).toFixed(2);
            elemento.style.color = 'red';
        }
    }
}

window.addEventListener('submit', interceptSubmit);