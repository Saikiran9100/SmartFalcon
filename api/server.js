const express = require('express');
const app = express();
const port = 3000;
const { getContract } = require('./connection/gateway');
app.use(express.json());

app.post('/records', async (req, res) => {
    try {
        const {
            recordId, dealerCode, phoneNumber, pinCode,
            amount, state, transactionValue, transactionMode, note
        } = req.body;

        const contract = await getContract();
        await contract.submitTransaction('RegisterRecord', recordId, dealerCode, phoneNumber, pinCode, amount, state, transactionValue, transactionMode, note);

        res.json({ message: 'Record registered successfully' });
    } catch (error) {
        console.error(`Failed to register record: ${error}`);
        res.status(500).json({ error: error.message });
    }
});

app.get('/records/:recordId', async (req, res) => {
    try {
        const contract = await getContract();
        const result = await contract.evaluateTransaction('FetchRecord', req.params.recordId);
        res.json(JSON.parse(result.toString()));
    } catch (error) {
        console.error(`Failed to fetch record: ${error}`);
        res.status(500).json({ error: error.message });
    }
});

app.get('/records/:recordId/timeline', async (req, res) => {
    try {
        const contract = await getContract();
        const result = await contract.evaluateTransaction('RetrieveRecordTimeline', req.params.recordId);
        res.json(JSON.parse(result.toString()));
    } catch (error) {
        console.error(`Failed to retrieve record timeline: ${error}`);
        res.status(500).json({ error: error.message });
    }
});

app.listen(port, () => {
    console.log(`Server running at http://localhost:${port}`);
});
