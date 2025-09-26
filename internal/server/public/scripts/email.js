function selectEmail(id) {
    const emailNode = document.getElementById(id);
    console.log(emailNode)
    emailNode.classList.add('active-email')
    emailNode.classList.add('bg-gray-600');
    emailNode.classList.add('underline');
}

function showEmail(numberOfEmails, id, subject, body, from) {
    console.log(`Showing email ${id} with subject: ${subject}, body: ${body}, from: ${from}`);
    for (let i = 0; i < numberOfEmails; i++) {
        document.getElementById(`email-${i}`).classList.remove('active-email');
        document.getElementById(`email-${i}`).classList.remove('bg-gray-600');
        document.getElementById(`email-${i}`).classList.remove('underline');
    }
    selectEmail(id);

    document.getElementById('email-subject').textContent = subject;
    document.getElementById('email-from').textContent = from;
    fillIframe(body);
}

function fillIframe(content) {
    if (!content) {
        return;
    }
    const viewer = document.getElementById('email-body');
    const htmlContent = `<p>${content}</p>`;

    viewer.contentDocument.open();
    viewer.contentDocument.write(htmlContent);
    viewer.contentDocument.close();
}
