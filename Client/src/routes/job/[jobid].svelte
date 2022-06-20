<script context="module">
	export const load = async ({ fetch, params }) => {
		const jobId = params.jobId;
		const res = await fetch(`/api/job/${jobId}`);
		let job;
		if (res.status == 200) {
			job = await res.json();
		}
		return { props: { job } };
	};
</script>

<script>
	export let job;
	console.log(job);
</script>
