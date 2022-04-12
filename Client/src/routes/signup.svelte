<script>
	import { onMount } from 'svelte';

	let email;
	let username;
	let password;
	let confirmPassword;
	let privacy = false;
	let error;
	let passedEmail = false;
	let passedUsername = false;
	let passedPassword = false;
	let passedConfirm = false;
	let passedPrivacy = false;
	onMount(() => {
		// create error golbaly
		error = document.createElement('p');
		error.classList.add('text-error', 'text-center', 'text-base', 'my-2');
	});
	//checking email
	function checkEmail() {
		const sample =
			/^(([^<>()[\]\.,;:\s@\"]+(\.[^<>()[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$/i;
		const el = document.getElementById('email');
		error.innerHTML = 'Email is not vaild';
		if (!email.match(sample)) {
			if (el.parentNode.querySelector('.text-error')) return;
			el.parentNode.insertBefore(error, el.previousElementSibling);
			el.classList.add('border-error');
			passedEmail = false;
		} else {
			el.classList.remove('border-error');
			if (el.parentNode.querySelector('.text-error')) {
				el.parentNode.removeChild(el.parentNode.querySelector('.text-error'));
			}
			passedEmail = true;
		}
	}
	// checking username
	function checkUsername() {
		const el = document.getElementById('username');
		error.innerHTML = 'Username must be at least 3 characters';
		if (username.length <= 3) {
			el.classList.add('border-error');
			if (el.parentNode.querySelector('.text-error')) return;
			el.parentNode.insertBefore(error, el.previousElementSibling);
			passedUsername = false;
		} else {
			if (el.parentNode.querySelector('.text-error')) {
				el.parentNode.removeChild(el.parentNode.querySelector('.text-error'));
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
		const el = document.getElementById('confirmPassword');
		if (password !== confirmPassword) {
			el.classList.add('border-error');
			error.innerHTML = 'Password and confirm password are not same';
			if (el.parentNode.querySelector('.text-error')) return;
			el.parentNode.insertBefore(error, el.previousElementSibling);
			passedConfirm = false;
		} else {
			el.classList.remove('border-error');
			if (el.parentNode.querySelector('.text-error')) {
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
	function submit() {
		if (passedEmail && passedUsername && passedPassword && passedConfirm && passedPrivacy) {
			fetch('http://localhost:8585/signup', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					email,
					password,
					confirmPassword
				})
			}).then((res) => {
				if (res.status == 301) {
					location.href = '/set-profile';
				}
			});
		}
	}
</script>

<main class="md:flex md:justify-center md:m-auto md:py-8  md:items-center">
	<form
		on:submit|preventDefault={submit}
		action="/signup"
		method="post"
		class="flex flex-col justify-between md:mr-56 p-4 gap-3 items-center"
	>
		<section class="w-80 ">
			<label class="text-base" for="email">Email</label>
			<input
				bind:value={email}
				on:change={checkEmail}
				type="email"
				name="email"
				id="email"
				class="outline-none border-2 border-main border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3"
			/>
		</section>
		<section class="w-80 ">
			<label class="text-base" for="username">Username</label>
			<input
				bind:value={username}
				on:change={checkUsername}
				type="text"
				name="username"
				id="username"
				class="outline-none border-2 border-main border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3"
			/>
		</section>
		<section class="w-80 ">
			<label class="text-base" for="password">Password</label>
			<input
				bind:value={password}
				on:change={checkPassword}
				type="password"
				name="password"
				id="password"
				class="outline-none border-2 border-main border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3"
			/>
			<ul class=" px-4">
				<li class="text-sm my-1" id="uppercase"><p>At least one uppercase letter</p></li>
				<li class="text-sm my-1" id="lowercase"><p>At least one lowercase letter</p></li>
				<li class="text-sm my-1" id="number"><p>contain at least one number</p></li>
				<li class="text-sm my-1" id="length"><p>should be more than 8 character</p></li>
			</ul>
		</section>
		<section class="w-80 ">
			<label class="text-base" for="confirmPassword">Confirm Password</label>
			<input
				bind:value={confirmPassword}
				on:change={checkConfrimPassword}
				type="password"
				name="confirmPassword"
				id="confirmPassword"
				class="outline-none border-2 border-main border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3"
			/>
		</section>
		<section class="m-2 w-80 px-4">
			<label for="privacy" on:click={checkPrivacy}>
				<input bind:checked={privacy} type="checkbox" name="privacy" id="" />
				Privacy Bullshit
			</label>
		</section>
		<section class="w-80">
			<input
				type="submit"
				value="Sign Up"
				class="cursor-pointer text-lg rounded-lg text-white bg-main py-2 px-20 mx-auto block w-11/12 hover:shadow-xl hover:bg-main-darker  "
			/>
		</section>
		<h5 class="text-base text-center">
			Already a member?<a href="/login" class="text-lg text-main mx-1">Login</a>
		</h5>
	</form>
	<h1 class="md:ml-56 md:p-4 hidden">There will be pic here</h1>
</main>
