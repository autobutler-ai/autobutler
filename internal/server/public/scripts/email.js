function showEmail(numberOfEmails, id, subject, body, from) {
    for (let i = 0; i < numberOfEmails; i++) {
        document.getElementById(`email-${i}`).classList.remove('bg-gray-600');
        document.getElementById(`email-${i}`).classList.remove('underline');
    }
    const emailNode = document.getElementById(id);
    emailNode.classList.add('bg-gray-600');
    emailNode.classList.add('underline');

    document.getElementById('email-subject').textContent = subject;
    document.getElementById('email-from').textContent = from;
    document.getElementById('email-body').innerHTML = `<p>${body}</p>`;
}
