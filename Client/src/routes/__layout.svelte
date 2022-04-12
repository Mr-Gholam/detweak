<script>
	import '../app.css';
	import Navbar from '../components/navbar.svelte';
	import LeftSidebar from '../components/leftSidebar.svelte';
</script>

<Navbar />
<!-- <main class="flex  xl:w-8/12  lg:w-9/12 mx-auto items-start justify-center md:justify-between">
	<LeftSidebar />
	<slot />
</main> -->
<slot />
