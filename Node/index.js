const express = require('express');
const bodyParser = require('body-parser');

const app = express();

app.get('/api', (req, res) => {
    res.status(200).json({ message: 'Welcome to the API!!' });
})

app.listen(5000);