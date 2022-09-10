<script>
	import { onMount } from 'svelte';
	import MetaMaskOnboarding from '@metamask/onboarding';
	import {Buffer} from 'buffer/'
	import store from 'store2'
	import { ClientJS } from 'clientjs';

	let rememberme = false

	const { isMetaMaskInstalled } = MetaMaskOnboarding;
	let accounts;
	const isMetaMaskConnected = () => accounts && accounts.length > 0;
	const publicAddress = () => accounts[0];

	if (isMetaMaskInstalled()) {
		ethereum.on('accountsChanged', (newAccounts) => { 
			handleNewAccounts(newAccounts)
		});
	}
	function handleNewAccounts(newAccounts) {
		accounts = newAccounts;
	}

	function handleSignup(publicAddress){
		fetch(`http://localho.st:8080/api/v1/user`, {
		body: JSON.stringify({ publicAddress }),
		headers: {
			'Content-Type': 'application/json'
		},
		method: 'POST'
		}).then(response => response.json());
	}

	function handleSignMessage({ publicAddress, nonce }) {
		return new Promise(async (resolve, reject) => {
			try {
				const msg = `0x${Buffer.from(nonce, 'utf8').toString('hex')}`;
				const signature = await ethereum.request({
					method: 'personal_sign',
					params: [msg, publicAddress],
				});
				return resolve({ publicAddress, signature });
			} catch (err) {
				reject(err);
			}
		});
	}

	function handleAuthenticate({ publicAddress, signature }) {
		var body = { publicAddress, signature };
		if (rememberme) {
			var client = new ClientJS(); 
			body.print = ""+client.getFingerprint();
		}
		fetch(`http://localho.st:8080/api/v1/auth`, {
		body: JSON.stringify(body),
		headers: {
			'Content-Type': 'application/json'
		},
		method: 'POST'
		}).then(response => response.json())
		.then(token => {
			store.setAll(token)
			location.reload()
		});
	}

	async function handleClick() {
		if (isMetaMaskInstalled()) {
			if (!isMetaMaskConnected()) {
				handleNewAccounts(await ethereum.request({
					method: 'eth_requestAccounts',
				}));
			}
			

			fetch(`http://localho.st:8080/api/v1/users?publicAddress=${publicAddress()}`)
			.then(response => response.json())
			.then(
				user => (user.nonce ? user : handleSignup(publicAddress()))
			).then(handleSignMessage).then(handleAuthenticate)
		} else {
			alert("no metamask")
		}
	};
</script>

<main>
	<h1>Hello!</h1>
	<p>Visit the <a href="https://svelte.dev/tutorial">Svelte tutorial</a> to learn how to build Svelte apps.</p>
	<input type=checkbox bind:checked={rememberme}> Remember me ?<br>
	<button on:click={handleClick}>Log ins</button>

</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>