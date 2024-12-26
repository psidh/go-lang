const express = require('express');

const app = express();

const port = 3001;

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.get('/', (req, res) => {
  return res.status(200).send('Welcome to server');
});

app.get('/get', (req, res) => {
  return res.status(200).json({ message: 'hello from get method' });
});

app.post('/post', (req, res) => {
  let mjson  = req.body;
  console.log(mjson);

  return res.status(200).send(mjson);
});

app.post('/platform', (req, res) => {
  console.log('Platform');
  console.log(req.body);
  
  return res.status(200).send(JSON.stringify(req.body));
});

app.listen(port, () => {
  console.log('Server listening at PORT: ', port);
});
