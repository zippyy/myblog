+++
author = 'Zippy'
date = '1942-04-02T22:20:47Z'
description = ''
slug = 'form'
title = 'form'
+++

<!DOCTYPE HTML>
<html>
<head>
	<meta charset="UTF-8">
	<title>Contact Us</title>
	<style>
		.contact-form {
		    background-color: #002538;
		    border: 1px solid #ddd;
		    padding: 20px;
		    border-radius: 10px;
		}
		    
		.form-label {
		    display: block;
		    margin-bottom: 10px;
		    font-size: 16px;
		    font-weight: 600;
		    color: #fff;
		}
		    
		.form-input {
		    display: block;
		    width: 100%;
		    padding: 12px;
		    border: 1px solid #ccc;
		    border-radius: 5px;
		    font-size: 16px;
		    transition: border-color 0.3s ease-in-out;
		    color: #000;
		    background-color: #fff;
		}

		.form-input:focus {
		    outline: none;
		    border-color: #007bff;
		}

		.form-input::placeholder {
		    color: #bbb;
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
		}

		.btn-submit:hover {
		    background-color: #0069d9;
		}

		.icon {
			font-size: 20px;
			margin-right: 5px;
		}

		.form-group {
			margin-bottom: 20px;
		}

	</style>
</head>
<body>
	<form name="contact" class="contact-form width-normal" action="/thank-you/" method="POST" data-netlify="true">
	    <input type="hidden" name="form-name" value="contact" />
	    <div class="form-group">
	        <label class="form-label" for="Name"><i class="icon fas fa-user"></i>Name</label>
	        <input id="contact-form-name" name="Name" type="text" placeholder="Enter your name" class="form-input" required="" autocomplete="off">
	    </div>
	    <div class="form-group">
	        <label class="form-label" for="Email"><i class="icon fas fa-envelope"></i>Email Address</label>
	        <input id="contact-form-email" name="Email" type="email" placeholder="Enter your email address" class="form-input" required="" autocomplete="off">
	    </div>
	    <div class="form-group">
	        <label class="form-label" for="Subject"><i class="icon fas fa-heading"></i>Subject</label>
	        <input id="contact-form-subject" name="Subject" type="text" placeholder="Enter the subject of your message" class="form-input" required="" autocomplete="off">
	    </div>
	    <div class="form-group">
	        <label class="form-label" for="Message"><i class="icon fas fa-pencil-alt"></i>Message</label>
	        <textarea class="form-input" id="contact-form-message" name="Message" placeholder="Enter your message" rows="6"></textarea>
	    </div>
	    <div class="form-group">
	        <button type="submit" value="Submit" id="Form-submit" class="btn-submit"><i class="icon fas fa-paper-plane"></i>Send Message</button>
	   
