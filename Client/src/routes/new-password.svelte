<script>
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import { loading, User } from '../store';
	onMount(async () => {
		if ($User.username) {
			goto('/dashboard');
		}
	});
</script>

<main class=" flex justify-center m-auto py-8  items-center h-80">
	<form
		action="/new-password"
		method="post"
		class="flex flex-col justify-between items-center gap-4"
	>
		<section class="w-80">
			<label class="text-base" for="password">Password</label>
			<input
				type="password"
				name="password"
				id="password"
				class="outline-none border-2 border-main border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3"
				placeholder="Password"
			/>
		</section>
		<section class="w-80">
			<label class="text-base" for="confirmPassword">Confirm Password</label>
			<input
				type="password"
				name="confirmPassword"
				id="confirmPassword"
				class="outline-none border-2 border-main border-solid rounded-lg px-1 mx-auto block w-11/12 p-2 my-3"
				placeholder="Confirm Password"
			/>
		</section>
		<section>
			<input type="hidden" name="token" value="<%= token %>" />
			<input type="hidden" name="userId" value="<%= userId %>" />
			<input
				type="submit"
				value="Reset Password"
				class="cursor-pointer text-lg rounded-lg bg-main py-2 px-20 hover:shadow-xl hover:bg-main-darker"
			/>
		</section>
	</form>
</main>
