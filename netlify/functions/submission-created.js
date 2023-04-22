require("dotenv").config();

import fetch from 'node-fetch';

exports.handler = async event => {
  const payload = JSON.parse(event.body).payload;
  const apiUrl = `https://api.elasticemail.com/v2/contact/add?publicAccountID=61458a7b-ca00-4bb5-9e21-5db82d4846a0&email=${payload.email}`;
  
  console.log(`Received a submission: ${payload.email}`);
  console.log(`API URL: ${apiUrl}`);

  return fetch(apiUrl, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ notes: payload.name }),
  })
    .then(response => response.json())
    .then(data => {
      console.log(`Submitted to ElasticEmail:\n ${JSON.stringify(data)}`);
    })
    .catch(error => ({ statusCode: 422, body: String(error) }));
}
