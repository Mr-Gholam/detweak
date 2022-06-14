<script>
	let email;
	let error;
	let errorMsg;
	async function submit() {
		const sample =
			/^(([^<>()[\]\.,;:\s@\"]+(\.[^<>()[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$/i;
		if (email.match(sample)) {
			const res = await fetch('/api/reset-password', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					email
				})
			});
			if (res.status == 400) {
				error = await res.json();
				errorMsg = error.error.message;
			}
			if (res.status == 200) {
				error = undefined;
				errorMsg = undefined;
			}
		}
	}
</script>

<main class="w-form flex justify-center m-auto py-8  items-center h-80">
	<form
		on:submit|preventDefault={submit}
		method="post"
		class="flex flex-col justify-between items-center h-48"
	>
		<section class="w-80">
			<h4 class="text-center text-error {error ? 'block' : 'hidden'}">{errorMsg}</h4>
			<label class="text-base text-text" for="email">Email</label>
			<input
				bind:value={email}
				type="email"
				name="email"
				id="email"
				class="outline-none border-2  border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-5 text-text {error
					? 'border-error'
					: 'border-border'}"
				placeholder="Email"
			/>
		</section>
		<section>
			<input type="submit" value="Reset Password" class="main-btn py-2 px-20 " />
		</section>
	</form>
</main>
