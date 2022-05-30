<script>
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import { loading, User } from '../store';

	let email;
	let username;
	let password;
	let confirmPassword;
	let privacy = false;

	let passedEmail = false;
	let passedUsername = false;
	let passedPassword = false;
	let passedConfirm = false;
	let passedPrivacy = false;

	//checking email
	function checkEmail() {
		const error = document.createElement('p');
		error.classList.add('error');
		const sample =
			/^(([^<>()[\]\.,;:\s@\"]+(\.[^<>()[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$/i;
		const el = document.getElementById('email');
		error.innerHTML = 'Email is not vaild';
		if (!email.match(sample)) {
			if (el.parentNode.querySelector('.error')) return;
			el.parentNode.insertBefore(error, el.previousElementSibling);
			el.classList.add('border-error');
			passedEmail = false;
		} else {
			el.classList.remove('border-error');
			if (el.parentNode.querySelector('.error')) {
				el.parentNode.removeChild(el.parentNode.querySelector('.error'));
			}
			passedEmail = true;
		}
	}
	// checking username
	function checkUsername() {
		const error = document.createElement('p');
		error.classList.add('error');
		const el = document.getElementById('username');
		error.innerHTML = 'Username must be at least 3 characters';
		if (username.length <= 3) {
			el.classList.add('border-error');
			if (el.parentNode.querySelector('.error')) return;
			el.parentNode.insertBefore(error, el.previousElementSibling);
			passedUsername = false;
		} else {
			if (el.parentNode.querySelector('.error')) {
				el.parentNode.removeChild(el.parentNode.querySelector('.error'));
			}
			el.classList.remove('border-error');
			passedUsername = true;
		}
	}
	/**
	 *todo make check password great agian
	 */
	// checking Password
	function checkPassword() {
		const error = document.createElement('p');
		error.classList.add('error');
		const el = document.getElementById('password');
		const uppercase = document.getElementById('uppercase');
		const lowercase = document.getElementById('lowecase');
		const number = document.getElementById('number');
		const length = document.getElementById('length');
		if (!/[A-Z]/.test(password)) {
			el.classList.add('border-error');
			uppercase.classList.add('text-error');
		} else if (!/[a-z]/.test(password)) {
			uppercase.classList.remove('text-error');
			el.classList.add('border-error');
			lowercase.classList.add('text-error');
		} else if (!/[0-9]/.test(password)) {
			lowercase.classList.remove('text-error');
			el.classList.add('border-error');
			number.classList.add('text-error');
		} else if (password.length <= 7) {
			number.classList.remove('text-error');
			el.classList.add('border-error');
			length.classList.add('text-error');
		} else {
			el.classList.remove('border-error');
			length.classList.remove('text-error');
			passedPassword = true;
		}
	}
	// checking confirm password
	function checkConfrimPassword() {
		const error = document.createElement('p');
		error.classList.add('error');
		const el = document.getElementById('confirmPassword');
		if (password !== confirmPassword) {
			el.classList.add('border-error');
			error.innerHTML = 'Password and confirm password are not same';
			if (el.parentNode.querySelector('.error')) return;
			el.parentNode.insertBefore(error, el.previousElementSibling);
			passedConfirm = false;
		} else {
			el.classList.remove('border-error');
			if (el.parentNode.querySelector('.error')) {
				el.parentNode.removeChild(el.parentNode.querySelector('.text-error'));
			}
			passedConfirm = true;
		}
	}
	//privacy
	function checkPrivacy() {
		if (!privacy) {
			privacy = true;
			passedPrivacy = true;
		} else {
			privacy = false;
			passedPrivacy = false;
		}
	}
	// submit
	async function submit() {
		if (passedEmail && passedUsername && passedPassword && passedConfirm && passedPrivacy) {
			$loading = true;
			const Response = await fetch('/api/signup', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					email,
					username,
					password,
					confirmPassword
				})
			});
			const data = await Response.json();
			if (Response.status == 201) {
				$loading = false;
				$User = data;
				console.log($User);
				goto('/set-profile');
			}
			// handleing duplicate email
			if (Response.status == 401 && data.emailErr) {
				$loading = false;
				const error = document.createElement('p');
				const el = document.getElementById('email');
				error.innerHTML = 'This Email has been used before';
				error.classList.add('error');
				el.classList.add('border-error');
				el.parentNode.insertBefore(error, el.previousElementSibling);
			}
			// handleing duplicate username
			if (Response.status == 403 && data.UsernameErr) {
				$loading = false;
				const error = document.createElement('p');
				const el = document.getElementById('username');
				error.innerHTML = 'This Username has been used before';
				error.classList.add('error');
				el.classList.add('border-error');
				el.parentNode.insertBefore(error, el.previousElementSibling);
			}
		}
	}
	onMount(async () => {
		if ($User) {
			goto('/dashboard');
		}
	});
</script>

<svelte:head>
	<title>Sign Up</title>
</svelte:head>
<main class="md:flex md:justify-center md:m-auto md:py-8  md:items-center">
	<form
		on:submit|preventDefault={submit}
		action="/signup"
		method="post"
		class="flex flex-col justify-between md:mr-56 p-4 gap-3 items-center"
	>
		<section class="w-80 ">
			<label class="text-base text-text" for="email">Email</label>
			<input
				bind:value={email}
				on:change={checkEmail}
				type="email"
				name="email"
				id="email"
				class="outline-none border-2 border-border border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3 text-text"
			/>
		</section>
		<section class="w-80 ">
			<label class="text-base text-text" for="username">Username</label>
			<input
				bind:value={username}
				on:change={checkUsername}
				type="text"
				name="username"
				id="username"
				class="outline-none border-2 border-border border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3 text-text "
			/>
		</section>
		<section class="w-80 ">
			<label class="text-base text-text" for="password">Password</label>
			<input
				bind:value={password}
				on:change={checkPassword}
				type="password"
				name="password"
				id="password"
				class="outline-none border-2 border-border border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3 text-text "
			/>
			<ul class=" px-4">
				<li class="text-xs my-1 text-text" id="uppercase"><p>At least one uppercase letter</p></li>
				<li class="text-xs my-1 text-text" id="lowercase"><p>At least one lowercase letter</p></li>
				<li class="text-xs my-1 text-text" id="number"><p>contain at least one number</p></li>
				<li class="text-xs my-1 text-text" id="length"><p>should be more than 8 character</p></li>
			</ul>
		</section>
		<section class="w-80 ">
			<label class="text-base text-text" for="confirmPassword">Confirm Password</label>
			<input
				bind:value={confirmPassword}
				on:change={checkConfrimPassword}
				type="password"
				name="confirmPassword"
				id="confirmPassword"
				class="outline-none border-2 border-border border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3 text-text"
			/>
		</section>
		<section class="m-2 w-80 px-4 text-text">
			<label for="privacy" on:click={checkPrivacy}>
				<input
					bind:checked={privacy}
					type="checkbox"
					name="privacy"
					id=""
					class="border-border bg-main-bg"
				/>
				Privacy Bullshit
			</label>
		</section>
		<section class="w-80">
			<input
				type="submit"
				value="Sign Up"
				class="main-btn   py-2 px-20 mx-auto block w-11/12 hover:shadow-xl text-text text-lg"
			/>
		</section>
		<h5 class="text-base text-center text-text">
			Already a member?<a href="/login" class="text-lg hover:text-main mx-1 text-white">Login</a>
		</h5>
	</form>
	<h1 class="md:ml-56 md:p-4 hidden">There will be pic here</h1>
</main>
