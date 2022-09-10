import App from './App.svelte';
import Auth from './Auth.svelte';
import Web3 from 'web3';
import { ClientJS } from 'clientjs';

import store from 'store2'

const accessToken = store('access_token')
let app

if (accessToken) {
	var client = new ClientJS();
	var fingerprint = ""+client.getFingerprint();
	fetch(`http://localho.st:8080/api/v1/auth/validate`, {
		body: JSON.stringify({ 'print': fingerprint }),
		headers: {
			'Content-Type': 'application/json',
			'Authorization': `Bearer ${accessToken}`
		},
		method: 'POST'
		}).then(response => {
			if (response.status > 299)
				throw `${response.statusText}`
			app = new Auth({
				target: document.body,
			})
		}).catch(err => {
			app = new App({
				target: document.body,
			})
		}) ;
} else {
	app = new App({
		target: document.body,
	})
}

export default app;