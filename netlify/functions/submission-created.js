require("dotenv").config()

import fetch from 'node-fetch';
const { ELASTICEMAIL_API_KEY } = process.env

exports.handler = async event => {
  const payload = JSON.parse(event.body).payload
  console.log(`Recieved a submission: ${payload.email}`)

  return fetch("https://api.elasticemail.com/v2/contact/add", {
    method: "POST",
    headers: {
      Authorization: `Token ${ELASTICEMAIL_API_KEY}`,
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ email: payload.email, notes: payload.name }),
  })
    .then(response => response.json())
    .then(data => {
      console.log(`Submitted to ElasticEmail:\n ${data}`)
    })
    .catch(error => ({ statusCode: 422, body: String(error) }))
} 
