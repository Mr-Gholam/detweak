<script>
	import { loading, onlineFriends, User, ws, Notification, UnseenMsg } from '../store';
	import '../app.css';
	import Navbar from '../components/navbar.svelte';
	import LeftSidebar from '../components/leftSidebar.svelte';
	import Loading from '../components/loading.svelte';
	import RightSidebar from '../components/rightSidebar.svelte';
	import { onMount } from 'svelte';
	import { navigating, page } from '$app/stores';
	$: $loading = !!$navigating;
	onMount(async () => {
		const res = await fetch('/api/jwt');
		const u = await res.json();
		if (res.status == 200) {
			$User = u;
			var h = window.location.href.split('/');
			const webSokect = new WebSocket('ws' + h[0].replace('http', '') + '//' + h[2] + '/api/ws');
			webSokect.onclose = () => {
				$onlineFriends = [];
				console.log('connection lost');
			};
			webSokect.onopen = () => {
				console.log('connected');
			};
			webSokect.onmessage = (e) => {
				const { Friend } = JSON.parse(e.data);
				const { offlineFriend } = JSON.parse(e.data);
				const { notification } = JSON.parse(e.data);
				const { UnSeenMsg } = JSON.parse(e.data);
				if (UnSeenMsg) {
					$UnseenMsg = $UnseenMsg + 1;
					$UnseenMsg = $UnseenMsg;
				}
				if (notification) {
					if (notification.NewFriendReq || notification.Liked) {
						if ($Notification == null) {
							$Notification = 1;
						} else {
							$Notification++;
						}
					} else {
						$Notification--;
					}
				}
				if (Friend) {
					$onlineFriends.push(Friend);
					$onlineFriends = $onlineFriends;
				}
				if (offlineFriend) {
					const location = $onlineFriends.findIndex(
						(user) => user.Username == offlineFriend.Username
					);
					$onlineFriends.splice(location, 1);
					$onlineFriends = $onlineFriends;
				}
			};
			ws.set(webSokect);
		}
	});
</script>

<Navbar />
{#if $User.username}
	{#if $loading}
		<main
			class="md:flex  xl:w-8/12  lg:w-9/12 md:mx-auto md:items-start  {$loading
				? 'blur-sm w-full'
				: ''}"
		>
			<LeftSidebar />
			<slot />
			<RightSidebar />
		</main>
		<Loading />
	{:else}
		<main class="md:flex  xl:w-8/12  lg:w-9/12 md:mx-auto md:items-start 2xl:justify-between">
			<LeftSidebar />
			<slot />
			<RightSidebar />
		</main>
	{/if}
{:else}
	<main class="{$loading ? 'blur-sm ' : ''}{$page.params.userInput ? 'xl:w-8/12 mx-auto' : ''} ">
		<slot />
	</main>
	<Loading />
{/if}
