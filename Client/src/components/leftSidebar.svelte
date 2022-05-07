<script>
	import { User } from '../store';
	let username;
	User.subscribe((value) => (username = value.username));
	import { page } from '$app/stores';
</script>

<div class=" h-fit p-4 my-2  sticky top-0 hidden md:block  w-fit inline  ">
	<section class=" flex justify-evenly flex-col text-lg gap-2 p-2 font-semibold w-fit ">
		<a
			href="/dashboard"
			class="hover:text-main {$page.url.pathname === '/dashboard' ? 'text-main' : ''}">Dashboard</a
		>
		<a
			href="/profile/{username}"
			class="hover:text-main {$page.url.pathname === `/profile/${username}` ? 'text-main' : ''}"
			>My Profile</a
		>
		<a
			href="/notification"
			class="hover:text-main {$page.url.pathname === '/notification' ? 'text-main' : ''}"
			>Notification</a
		>
		<a
			href="/messages"
			class="hover:text-main {$page.url.pathname === '/messages' ? 'text-main' : ''}">Messages</a
		>
		<a
			href="/liked-posts"
			class="hover:text-main {$page.url.pathname === '/liked-posts' ? 'text-main' : ''}"
		>
			Liked Posts</a
		>
		<a
			href="/setting"
			class="hover:text-main {$page.url.pathname === '/setting' ? 'text-main' : ''}">Setting</a
		>
	</section>
</div>
