<script context="module">
	export const load = async ({ fetch }) => {
		const res = await fetch('/api/my-jobs', {
			credentials: 'include'
		});
		let jobs = null;
		if (res.status == 200) {
			jobs = await res.json();
		}
		return { props: { jobs } };
	};
</script>

<script>
	export let jobs;
	console.log(jobs);
</script>

<svelte:head>
	<title>My Job</title>
</svelte:head>
