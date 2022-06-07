<script>
	import { goto } from '$app/navigation';
	import { loading, User } from '../store';
	let email;
	let password;
	let passedEmail = false;
	let passedPassword = false;
	let uppercaseError = false;
	let lowercaseError = false;
	let numberError = false;
	let lengthError = false;
	let passedUppercase = false;
	let passedLowercase = false;
	let passedNumber = false;
	let passedLength = false;
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
			uppercaseError = true;
			passedUppercase = false;
		} else if (!/[a-z]/.test(password)) {
			el.classList.add('border-error');
			uppercaseError = false;
			passedUppercase = true;
			passedLowercase = false;
			lowercaseError = true;
		} else if (!/[0-9]/.test(password)) {
			el.classList.add('border-error');
			lowercaseError = false;
			passedLowercase = true;
			passedNumber = false;
			numberError = true;
		} else if (password.length <= 7) {
			el.classList.add('border-error');
			numberError = false;
			passedNumber = true;
			passedLength = false;
			lengthError = true;
		} else {
			el.classList.remove('border-error');
			lengthError = false;
			numberError = false;
			lowercaseError = false;
			uppercaseError = false;
			passedLength = true;
			passedUppercase = true;
			passedLowercase = true;
			passedNumber = true;
			passedLength = true;
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
			if (response.ok) {
				$User = data;
				$loading = false;
				goto('/dashboard');
			}
			// handleing invalid email
			if (response.status == 400 && data.error.message == 'Email Not found') {
				const error = document.createElement('p');
				error.classList.add('error');
				const el = document.getElementById('email');
				error.innerHTML = 'Email is not vaild';
				el.classList.add('border-error');
				el.parentNode.insertBefore(error, el.previousElementSibling);
				$loading = false;
			}
			// handleing invalid password
			if (response.status == 400 && data.error.message == 'Incorrect Password') {
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
			<label class="text-base text-text" for="email">Email</label>
			<input
				bind:value={email}
				on:change={checkEmail}
				type="email"
				name="email"
				id="email"
				class="outline-none border-2 border-border rounded-lg px-1 mx-auto block w-11/12 p-2 my-3 text-text"
				placeholder="Email"
			/>
		</section>
		<section class="w-80">
			<label class="text-lg text-text" for="password">Password</label>
			<input
				bind:value={password}
				on:input={checkPassword}
				type="password"
				name="password"
				id="password"
				class="outline-none border-2 border-border text-text rounded-lg px-1 mx-auto block w-11/12 p-2 my-3"
				placeholder="Password"
			/>
			<ul class=" px-4">
				<li
					class="text-xs my-1   {uppercaseError
						? 'text-error'
						: passedUppercase
						? 'text-green'
						: 'text-text'}"
					id="uppercase"
				>
					<p>At least one uppercase letter</p>
				</li>
				<li
					class="text-xs my-1 {lowercaseError
						? 'text-error'
						: passedLowercase
						? 'text-green'
						: 'text-text'}"
					id="lowercase"
				>
					<p>At least one lowercase letter</p>
				</li>
				<li
					class="text-xs my-1 {numberError
						? 'text-error'
						: passedNumber
						? 'text-green'
						: 'text-text'}"
					id="number"
				>
					<p>contain at least one number</p>
				</li>
				<li
					class="text-xs my-1 {lengthError
						? 'text-error'
						: passedLength
						? 'text-green'
						: 'text-text'}"
					id="length"
				>
					<p>should be more than 8 character</p>
				</li>
			</ul>
		</section>
		<section class="my-3">
			<h4 class="text-base text-text">
				Forgot your Password? <a
					href="/reset-password"
					class="text-lg hover:text-main  mx-1 text-white">Reset Password</a
				>
			</h4>
		</section>
		<section class="w-80">
			<input
				type="submit"
				value="Login"
				class="main-btn py-2  px-20  mx-auto block w-11/12  font-semibold border-2"
			/>
		</section>
		<h4 class="text-base my-5 text-text">
			Not a member yet?<a href="/signup" class="text-lg hover:text-main font-medium mx-1 text-white"
				>Sign up</a
			>
		</h4>
	</form>
	<h1 class="md:mr-56 md:p-4 hidden">there will be a pic here</h1>
</main>
