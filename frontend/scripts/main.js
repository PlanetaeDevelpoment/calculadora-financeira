const tags = {
    'app-header': 'header.html',
    'app-footer': 'footer.html'
}

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
    return await fetch(`templates/${template}`).then(response => response.text()).then(html => {
        return html;
    }).catch(error => {
        console.log('Error loading template', error);
    });
}


async function onStartup() {
    await loadPage();
}

window.addEventListener('load', onStartup);
