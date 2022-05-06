<script>
	import { goto } from '$app/navigation';
	import { loading, User } from '../store';

	let email;
	let password;
	let passedEmail = false;
	let passedPassword = false;

	// check Email
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
			if (el.parentNode.querySelector('.text-error')) {
				el.parentNode.removeChild(el.parentNode.querySelector('.text-error'));
			}
			passedEmail = true;
		}
	}

	// check password
	function checkPassword() {
		const error = document.createElement('p');
		error.classList.add('error');
		const el = document.getElementById('password');
		if (!/[A-Z]/.test(password)) {
			el.classList.add('border-error');
		} else if (!/[a-z]/.test(password)) {
			el.classList.add('border-error');
		} else if (!/[0-9]/.test(password)) {
			el.classList.add('border-error');
		} else if (password.length <= 7) {
			el.classList.add('border-error');
		} else {
			el.classList.remove('border-error');
			passedPassword = true;
		}
	}
	async function sumbit() {
		$loading = true;
		if (passedEmail && passedPassword) {
			const response = await fetch('/api/login', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					email,
					password
				})
			});
			const data = await response.json();
			if (response.status == 200) {
				User.set(data);
				$loading = false;
				goto('/dashboard');
			}
			// handleing invalid email
			if (response.status == 403 && data.emailErr) {
				const error = document.createElement('p');
				error.classList.add('error');
				const el = document.getElementById('email');
				error.innerHTML = 'Email is not vaild';
				el.classList.add('border-error');
				el.parentNode.insertBefore(error, el.previousElementSibling);
				$loading = false;
			}
			// handleing invalid password
			if (response.status == 403 && data.passwordErr) {
				const error = document.createElement('p');
				error.classList.add('error');
				const el = document.getElementById('password');
				error.innerHTML = 'Password is not vaild';
				el.classList.add('border-error');
				el.parentNode.insertBefore(error, el.previousElementSibling);
				$loading = false;
			}
		}
	}
</script>

<svelte:head>
	<title>Login</title>
</svelte:head>
<main class=" flex justify-center m-auto py-8 md:my-16  items-center">
	<form
		action="/login"
		method="post"
		on:submit|preventDefault={sumbit}
		class="flex flex-col justify-between items-center  md:mr-56  md:p-4 gap-4"
	>
		<section class="w-80">
			<label class="text-base" for="email">Email</label>
			<input
				bind:value={email}
				on:change={checkEmail}
				type="email"
				name="email"
				id="email"
				class="outline-none border-2 border-main-bg border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3"
				placeholder="Email"
			/>
		</section>
		<section class="w-80">
			<label class="text-lg" for="password">Password</label>
			<input
				bind:value={password}
				on:change={checkPassword}
				type="password"
				name="password"
				id="password"
				class="outline-none border-2 border-main-bg border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3"
				placeholder="Password"
			/>
		</section>
		<section class="my-3">
			<h4 class="text-base">
				Forgot your Password? <a
					href="/reset-password"
					class="text-lg text-main  mx-1 hover:text-main-bg">Reset Password</a
				>
			</h4>
		</section>
		<section class="w-80">
			<input
				type="submit"
				value="Login"
				class="cursor-pointer text-lg rounded-lg bg-main-bg py-2 hover:text-main px-20 hover:shadow-xl mx-auto block w-11/12 text-white"
			/>
		</section>
		<h4 class="text-base my-5">
			Not a member yet?<a
				href="/signup"
				class="text-lg text-main font-medium mx-1 hover:text-main-bg">Sign up</a
			>
		</h4>
	</form>
	<h1 class="md:mr-56 md:p-4 hidden">there will be a pic here</h1>
</main>
