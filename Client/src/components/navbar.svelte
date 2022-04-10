<script>
	let IsloggedIn = false;
</script>

<nav class="flex bg-main-bg h-12 items-center justify-between">
	<a href="/" class="mx-2 text-main">Logo</a>

	<section>
		<input
			type="text"
			placeholder="search"
			class=" w-96 border-2 rounded-md py-0.5  px-2 focus:outline-hidden focus:outline-none  "
		/>
	</section>
	<section class="flex justify-between items-center">
		{#if IsloggedIn}
			<a href="/logout">Logout</a>
		{:else}
			<a
				href="/signup"
				class="rounded-lg border-main border-solid mx-2 border-2 px-2 py-1 font-semibold text-base text-gray-300 hover:bg-gray-800 hover:shadow-xl hover:text-main"
				>Sign up</a
			>
			<a
				href="/login"
				class="rounded-lg border-main border-solid mx-2 border-2 px-2 py-1 font-semibold text-base text-gray-300 hover:bg-gray-800 hover:shadow-xl hover:text-main"
				>Login</a
			>
		{/if}
	</section>
</nav>
