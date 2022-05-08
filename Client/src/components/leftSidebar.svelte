<script>
	import { User } from '../store';
	let username;
	User.subscribe((value) => (username = value.username));
	import { page } from '$app/stores';
</script>

<div class=" h-fit p-4 my-2  sticky top-0 hidden md:block  w-52 inline  ">
	<section class=" flex justify-evenly flex-col text-lg p-2 font-semibold w-full ">
		<a
			href="/dashboard"
			class="hover:text-main {$page.url.pathname === '/dashboard' ? 'text-main' : ''} pl-2 py-2"
		>
			<i class="fa-solid fa-house mr-1" /> Home</a
		>
		<a
			href="/profile/{username}"
			class="hover:text-main  pl-2 py-2 {$page.url.pathname === `/profile/${username}`
				? 'text-main'
				: ''}"
		>
			<i class="fa-solid fa-house-user" /> My Profile</a
		>
		<a
			href="/notification"
			class="hover:text-main   pl-2 py-2 {$page.url.pathname === '/notification'
				? 'text-main'
				: ''}"><i class="fa-solid fa-bell mr-1" /> Notification</a
		>
		<a
			href="/messages"
			class="hover:text-main pl-2  py-2{$page.url.pathname === '/messages' ? 'text-main' : ''}"
			><i class="fa-solid fa-comments" /> Messages</a
		>
		<a
			href="/liked-posts"
			class="hover:text-main pl-2  py-2{$page.url.pathname === '/liked-posts' ? 'text-main' : ''}"
		>
			<i class="fa-solid fa-thumbs-up mr-1" /> Liked Posts</a
		>
		<a
			href="/setting"
			class="hover:text-main pl-2 py-2 {$page.url.pathname === '/setting' ? 'text-main' : ''}"
			><i class="fa-solid fa-gears mr-1" /> Setting</a
		>
	</section>
</div>
