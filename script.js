document.getElementById('nidForm').addEventListener('submit', function(event) {
    event.preventDefault();
    const nidInput = document.getElementById('nidInput').value;
    const resultDiv = document.getElementById('result');
    resultDiv.textContent = 'Verifying NID: ' + nidInput;
    verifyNID(nidInput, resultDiv);
});

// Fetch API call to backend
async function verifyNID(nid, resultDiv) {
    try {
        const response = await fetch(`http://localhost:8080/verify-nid?nid=${nid}`);
        if (!response.ok) {
            if (response.status === 404) {
                resultDiv.textContent = 'NID not found.';
            } else {
                resultDiv.textContent = 'Server error.';
            }
            return;
        }
        const data = await response.json();
        resultDiv.innerHTML = `
            <div class="bg-white-100 p-4 rounded">
                <div><strong>NID:</strong> ${data.nid}</div>
                <div><strong>Name:</strong> ${data.name}</div>
                <div><strong>Date of Birth:</strong> ${data.date_of_birth}</div>
                <div><strong>Father's Name:</strong> ${data.father_name}</div>
                <div><strong>Mother's Name:</strong> ${data.mother_name}</div>
                <div><strong>Address:</strong> ${data.address}</div>
                <div><strong>Gender:</strong> ${data.gender}</div>
            </div>
        `;
    } catch (error) {
        resultDiv.textContent = 'Error verifying NID: ' + error.message;
    }
}