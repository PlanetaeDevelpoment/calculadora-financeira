const tags = {
    'app-header': 'header.html'
}
const apiUrl = 'http://0.0.0.0:8000/';

function loadPage() {
    for (const tag in tags) {
        inputTemplate(tag);        
    }
}

async function inputTemplate(tag) {
    document.querySelectorAll(tag).forEach(async (element) => {
        element.innerHTML = await fetchTemplate(tags[tag])
    });
} 

async function fetchTemplate(template) {
    return await fetch(`${apiUrl}templates/${template}`).then(response => response.text()).then(html => {
        console.log(html);
        return html;
    }).catch(error => {
        console.log('Error loading template', error);
    });
}


async function onStartup() {
    await loadPage();
}

window.addEventListener('load', onStartup);
