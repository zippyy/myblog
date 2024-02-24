+++
author = 'Zippy'
comments = false
date = '1942-04-02T10:20:49Z'
description = ''
slug = 'services'
title = 'Services'
+++


**Under Construction**

Here is a List of services I offer, You can reach out to me by heading to the [Contact](https://techrelay.xyz/contact) for all my contact methods or by clicking {{< calendly calendar="techrelay" />}}. For most services my rate is $75 for the first hour, $50 for Subsequent Hours in Increments of 30 Minutes. You can also check out my [CV/Résumé Here](https://nbennett.pro)


## Digital Nomad

1. Digital Nomad VPN Setup and Support
2. Digital Nomad RDP Setup and Support
3. Web Site, Blog, Resume, Etc... Built in Hugo, Code stored on Github and Hosted on Netlify Pages or Cloudflare Pages (Both Code storage on Github and The Static Hosting on Netlify or Cloudflare is free for most use cases.)


## Web

1. Web Application Installation/Configuration 
2. Hosting Management
3. Web Sites in Hugo
4. Web Sites in WordPress
5. Email Stuff (Setup and Configuration of hosted email like MS365 or Gsuite)
6. Azure, MS365, Intune, Auto-Pilot and On-Prem AD, Exchange. (Management, Setup, Configuration, Troubleshooting, etc. Compensation depends on scope of projects, Not included in the rate mentioned above.)

## Consulting

I am available for a vast range of consulting niches including Windows/Windows Server Administration, Cloud (Azure Ad, MS365, Intune, Autopilot), Virtualization (vmWare, Proxmox, Hyper-V) including VDI like Horizon, Linux Server Administration, Mobile Device Management (MDM) and Remote Management & Monitoring (RMM). You name it and I can either already do it or quickly learn how. 

I have affordable rates, quality work ethics and great communication skills, You will not be let down by choosing to work with me. 

## More to come...
<br>
<br>
Use this form to request a quote for services.
<br>
<br>

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
    if (!isNaN(budgetValue) || budgetValue === '') {
      budgetField.value = '$' + parseFloat(budgetValue).toLocaleString();
      errorSpan.classList.remove('show');
    } else {
      errorSpan.classList.add('show');
    }
  });

budgetField.addEventListener('keydown', function(event) {
  if (event.key === 'Backspace' || event.key === 'Delete') {
    let budgetValue = budgetField.value.replace('$', '');
    if (!isNaN(budgetValue)) {
      budgetField.value = '$' + parseFloat(budgetValue).toLocaleString();
      errorSpan.classList.remove('show');
    } else {
      budgetField.value = '';
    }
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
            padding: 30px;
           
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

        .btn-submit {
            background-color: #F78A2C;
            color: #fff;
            font-size: 18px;
            font-weight: 600;
            border: none;
            border-radius: 5px;
            padding: 12px 20px;
            transition: all 0.3s ease;
            margin-top: 20px; /* increased margin */
}

        .btn-submit:hover {
            background-color: #E2761B;
}
    </style>
