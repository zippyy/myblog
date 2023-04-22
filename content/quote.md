+++
author = 'Zippy'
comments = false
date = "2020-04-09T22:53:20-06:00"
description = ''
slug = 'quote'
title = 'Quote'
+++

<!DOCTYPE HTML>
<body>
<form name="quote-request" class="quote-request-form" action="/thank-you/" method="POST" data-netlify="true">
    <input type="hidden" name="form-name" value="quote-request" />
    <div class="form-group">
        <label class="form-label" for="Name"><i class="fas fa-user"></i>Name</label>
        <input id="quote-request-name" name="Name" type="text" placeholder="Enter your name" class="form-input" required="" autocomplete="off">
    </div>
    <div class="form-group">
        <label class="form-label" for="Email"><i class="fas fa-envelope"></i>Email Address</label>
        <input id="quote-request-email" name="Email" type="email" placeholder="Enter your email address" class="form-input" required="" autocomplete="off">
    </div>
    <div class="form-group">
        <label class="form-label" for="Phone"><i class="fas fa-phone"></i>Phone Number</label>
        <input id="quote-request-phone" name="Phone" type="text" placeholder="Enter your phone number" class="form-input" required="" autocomplete="off">
    </div>
    <div class="form-group">
        <label class="form-label" for="ProjectType"><i class="fas fa-clipboard"></i>Project Type</label>
        <select id="quote-request-project-type" name="ProjectType" class="form-input" required="">
            <option value="" disabled="" selected="">Select project type</option>
            <option value="Digital Nomad">Digital Nomad</option>
            <option value="Hugo Websites">Hugo Websites</option>
            <option value="Cloud">Cloud</option>
            <option value="Consulting">Consulting</option>
            <option value="Other">Other</option>
        </select>
    </div>
    <div class="form-group">
    <label class="form-label" for="Budget"><i class="fas fa-money-bill-wave"></i>Budget</label>
    <input id="quote-request-budget" name="Budget" type="text" placeholder="Enter your budget" class="form-input" required="" autocomplete="off" oninput="if(!this.value.startsWith('$')){this.value = '$' + this.value}">
    <span class="error">Please enter a valid number.</span>
</div>
<div class="form-group">
        <label class="form-label" for="Message"><i class="fas fa-pencil-alt"></i>Message</label>
        <textarea class="form-input" id="quote-request-message" name="Message" placeholder="Enter your message" rows="6"></textarea>
    </div>
    <div class="form-group">
        <button type="submit" value="Submit" id="Form-submit" class="btn-submit"><i class="fas fa-paper-plane"></i>Request Quote</button>
    </div>
</form>

<script>
  document.addEventListener('DOMContentLoaded', function() {
  const budgetField = document.getElementById('quote-request-budget');
  const errorSpan = document.querySelector('.error');

  budgetField.addEventListener('input', function() {
    let budgetValue = budgetField.value.trim().replace('$', '').replace(',', '');
    if (!isNaN(budgetValue)) {
      budgetField.value = '$' + parseFloat(budgetValue).toLocaleString();
      errorSpan.classList.remove('show');
    } else {
      errorSpan.classList.add('show');
    }
  });

  budgetField.addEventListener('keydown', function(event) {
    if (event.key === 'Backspace') {
      budgetField.value = budgetField.value.replace('$', '');
    }
  });

  const form = document.querySelector('.quote-request-form');
  form.addEventListener('submit', function(e) {
    const budgetValue = budgetField.value.trim().replace('$', '').replace(',', '');
    if (!isNaN(budgetValue)) {
      budgetField.value = budgetValue;
      budgetField.dispatchEvent(new Event('input'));
    } else {
      errorSpan.classList.add('show');
      e.preventDefault();
    }
  });
});

  </script>

</body>
 <style>
        .quote-request-form {
            background-color: #002538;
            border: none;
            padding: 30px;
            border-radius: 10px;
        }

        .form-label {
            display: flex;
            align-items: center;
            margin-bottom: 10px;
            font-size: 18px;
            font-weight: 600;
            color: #fff;
        }

        .form-label i {
            font-size: 20px;
            margin-right: 10px;
        }

        .form-input {
            display: block;
            width: 100%;
            padding: 12px;
            background-color: #333;
            color: #fff;
        }

        .error {
            display: none;
            color: red;
            font-size: 14px;
            margin-top: 5px;
        }

        .error.show {
            display: block;
        }
    </style>
