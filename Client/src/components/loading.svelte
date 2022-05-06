<script>
	import { loading } from '../store';
</script>

{#if $loading}
	<div class="loading ">
		<div class="w-32 h-32 relative mx-auto top-1/2 ">
			<div class="child-loading" />
			<div class="child-loading" />
			<div class="child-loading" />
			<div class="child-loading" />
		</div>
	</div>
{/if}
