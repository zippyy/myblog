+++
author = 'Zippy'
comments = false
date = "2020-04-09T22:53:20-06:00"
description = ''
slug = 'contact'
title = 'Contact'
+++

{{< calendly calendar="techrelay" />}}
<br>
<br>
Personal Email: {{< cloakemail "nick@techrelay.xyz" >}}

Personal Telegram: @zippyy

Discord: Zippyy#8551

Battle.net: Zippy420#1359

Steam: Zippyy

Matrix: {{< cloakemail address="@nick:matrix.techrelay.xyz" protocol="matrix" >}}

IRC: LiberaChat IRC Network on Channel #TechRelay

You may also use the form below to reach out to me directly.

<!DOCTYPE HTML>
<form name="contact" class="contact-form width-normal" action="/thank-you/" method="POST" data-netlify="true">
    <input type="hidden" name="form-name" value="contact" />
    <div class="form-group">
        <label class="form-label" for="Name"><i class="fas fa-user"></i>Name</label>
        <input id="contact-form-name" name="Name" type="text" placeholder="Enter your name" class="form-input" required="" autocomplete="off">
    </div>
    <div class="form-group">
        <label class="form-label" for="Email"><i class="fas fa-envelope"></i>Email Address</label>
        <input id="contact-form-email" name="Email" type="email" placeholder="Enter your email address" class="form-input" required="" autocomplete="off">
    </div>
    <div class="form-group">
        <label class="form-label" for="Subject"><i class="fas fa-comment"></i>Subject</label>
        <input id="contact-form-subject" name="Subject" type="text" placeholder="Enter the subject of your message" class="form-input" required="" autocomplete="off">
    </div>
    <div class="form-group">
        <label class="form-label" for="Message"><i class="fas fa-pencil-alt"></i>Message</label>
        <textarea class="form-input" id="contact-form-message" name="Message" placeholder="Enter your message" rows="6"></textarea>
    </div>
    <div class="form-group">
        <button type="submit" value="Submit" id="Form-submit" class="btn-submit"><i class="fas fa-paper-plane"></i>Send Message</button>
    </div>
</form>

<style>
    .contact-form {
        background-color: #1a1a1a;
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
        border: none;
        border-radius: 5px;
        font-size: 16px;
        background-color: #333;
        color: #fff;
        transition: background-color 0.3s ease-in-out;
    }

    .form-input:focus {
        outline: none;
        background-color: #444;
    }

    .form-input::placeholder {
        color: #888;
    }
    
    .btn-submit {
        background-color: #007bff;
        color: #fff;
        border: none;
        padding: 10px 20px;
        border-radius: 5px;
        font-size: 18px;
        cursor: pointer;
        transition: background-color 0.3s ease-in-out;
        display: flex;
        align-items: center;
    }

    .btn-submit i {
        font-size: 18px;
        margin-right: 10px;
    }
    
    .btn-submit:hover {
        background-color: #0069d9;
    }
</style>

<!-- Add Font Awesome icons -->
<script src="https://kit.fontawesome
