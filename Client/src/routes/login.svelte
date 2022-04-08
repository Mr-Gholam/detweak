<script>
	let email;
	let password;
	let passedEmail = false;

	// check Email
	function checkEmail() {
		const error = document.createElement('p');
		error.classList.add('text-error', 'text-center', 'text-base', 'my-2');
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
	/**
	 * todo password
	 */
	// check password
	function checkPassword() {}
</script>

<main class=" flex justify-center m-auto py-8 my-16  items-center">
	<form
		action="/login"
		method="post"
		class="flex flex-col justify-between items-center  mr-56  p-4 gap-4"
	>
		<section class="w-80">
			<label class="text-base" for="email">Email</label>
			<input
				bind:value={email}
				on:change={checkEmail}
				type="email"
				name="email"
				id="email"
				class="outline-none border-2 border-main border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3"
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
				class="outline-none border-2 border-main border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3"
				placeholder="Password"
			/>
		</section>
		<section class="my-3">
			<h4 class="text-base">
				Forgot your Password? <a
					href="/reset-password"
					class="text-lg text-main  mx-1 hover:text-main-darker">Reset Password</a
				>
			</h4>
		</section>
		<section class="w-80">
			<input
				type="submit"
				value="Login"
				class="cursor-pointer text-lg rounded-lg bg-main py-2 px-20 hover:shadow-xl mx-auto block w-11/12 hover:bg-main-darker "
			/>
		</section>
		<h4 class="text-base my-5">
			Not a member yet?<a
				href="/signup"
				class="text-lg text-main font-medium mx-1 hover:text-main-darker ">Sign up</a
			>
		</h4>
	</form>
	<h1 class="mr-56 p-4">there will be a pic here</h1>
</main>
