<script context="module">
	export const load = async ({ url, fetch }) => {
		const res = await fetch(`${url.protocol}//${url.hostname}/api/jwt`);
		if (res.status == 401) {
			if (
				url.pathname == '/dashboard' ||
				url.pathname.includes('/friends') ||
				url.pathname.includes('/liked-by') ||
				url.pathname.includes('/post') ||
				url.pathname == '/liked-posts' ||
				url.pathname == '/messages' ||
				url.pathname == '/notifiction' ||
				url.pathname == '/sitting'
			) {
				return { status: 301, redirect: '/login' };
			}
		}
		if (res.status == 200) {
			const u = await res.json();
			User.set(u);
			var h = window.location.href.split('/');
			const webSokect = new WebSocket('ws' + h[0].replace('http', '') + '//' + h[2] + '/api/ws');
			webSokect.onclose = () => {
				onlineFriends.set([]);
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
					UnseenMsg.set(get(UnseenMsg) + 1);
				}
				if (notification) {
					if (notification.NewFriendReq || notification.Liked) {
						if (Notification == null) {
							Notification.set(1);
						} else {
							Notification.set(get(Notification) + 1);
						}
					} else {
						Notification.set(get(Notification) - 1);
					}
				}
				if (Friend) {
					onlineFriends.set([...get(onlineFriends), Friend]);
				}
				if (offlineFriend) {
					const location = get(onlineFriends).findIndex(
						(user) => user.Username == offlineFriend.Username
					);
					get(onlineFriends).splice(location, 1);
				}
			};
			ws.set(webSokect);
			if (
				url.pathname == '/signup' ||
				url.pathname == '/login' ||
				url.pathname == '/new-password' ||
				url.pathname == '/set-resume' ||
				url.pathname == '/set-profile' ||
				url.pathname == '/reset-password' ||
				url.pathname == '/'
			) {
				return { status: 301, redirect: '/dashboard' };
			}
		}
		return {};
	};
</script>

<script>
	import { get } from 'svelte/store';
	import { loading, onlineFriends, User, ws, Notification, UnseenMsg } from '../store';
	import '../app.css';
	import Navbar from '../components/navbar.svelte';
	import LeftSidebar from '../components/leftSidebar.svelte';
	import Loading from '../components/loading.svelte';
	import RightSidebar from '../components/rightSidebar.svelte';
	import { navigating, page } from '$app/stores';
	$: $loading = !!$navigating;
	console.log($User);
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
