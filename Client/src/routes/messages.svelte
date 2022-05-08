<script>
	import { io } from 'socket.io-client';
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	let contacts = [];
	let currentChat = [];
	let currentChatInfo;
	let textInput;
	let showList = true;
	let bouncing = false;
	let loading = true;
	const targetUser = $page.url.search;
	const targetUsername = targetUser.split('=')[1];
	onMount(async () => {
		const response = await fetch('/api/load-chatRooms');
		const data = await response.json();
		contacts = data.contacts;
		if (targetUsername) {
			const foundUser = contacts.find((contact) => contact.username == targetUsername);
			if (!foundUser) {
				const result = await fetch('/api/create-room', {
					method: 'POST',
					headers: {
						'Content-Type': 'application/json'
					},
					body: JSON.stringify({
						targetUsername
					})
				});
				const newChatRoom = await result.json();
				if (result.status == 200) {
					contacts.unshift(newChatRoom);
					contacts = contacts;
				}
			}
		}
	});
	function createSendMessage(input) {
		const middePart = document.getElementById('middlePart');
		const messageParent = document.createElement('div');
		middePart.appendChild(messageParent);
		messageParent.classList.add('w-full');
		const message = document.createElement('h1');
		message.classList.add('send-message');
		message.innerText = input;
		messageParent.appendChild(message);
	}
	function createReceiveMessage(input) {
		const middePart = document.getElementById('middlePart');
		const messageParent = document.createElement('div');
		middePart.appendChild(messageParent);
		messageParent.classList.add('w-full');
		const message = document.createElement('h1');
		message.classList.add('receive-message');
		message.innerText = input;
		messageParent.appendChild(message);
	}
	async function sendMessage() {
		const chatRoomId = currentChatInfo.chatRoomId;
		const targetId = currentChatInfo.targetId;
		if (textInput == '' || textInput == undefined) return;
		const response = await fetch('/api/create-message', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				message: textInput,
				chatRoomId,
				targetId
			})
		});
		if (response.status == 200) {
			createSendMessage(textInput);
			textInput = '';
		}
	}
	function showNameList() {
		const nameList = document.getElementById('nameList');
		if (showList) {
			nameList.classList.remove('hidden');
		}
	}
	function closeNameList() {
		const nameList = document.getElementById('nameList');
		nameList.classList.add('hidden');
	}
	async function changeChat(chatRoomId) {
		const response = await fetch('/api/get-chat', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				chatRoomId
			})
		});
		const data = await response.json();
		currentChatInfo = data.chatInfo;
		currentChat = data.currentChat;
		loading = false;
	}
	function typing() {
		const sendBtn = document.getElementById('sendBtn');
		if (bouncing) {
			sendBtn.classList.remove('right-bounce');
			bouncing = false;
		} else {
			sendBtn.classList.add('right-bounce');
			bouncing = true;
		}
		setInterval(() => {
			sendBtn.classList.remove('right-bounce');
			bouncing = false;
		}, 6250);
	}
</script>

<svelte:head>
	<title>Messages</title>
</svelte:head>
<div
	class="w-96 md:w-10/12 md:my-6  relative md:flex md:border-2 md:border-main-bg md:rounded-md md:mr-4 h-screen lg:h-128"
>
	<!--Name list -->
	<section
		class="md:w-3/12 w-full border-r-2 border-main-bg  overflow-y-auto hidden md:block absolute top-0 left-0 z-10  h-full md:bg-transparent md:z-0 md:static"
		id="nameList"
	>
		<div
			class="absolute  top-0 left-0 -z-10 w-screen h-screen md:hidden"
			on:click={closeNameList}
		/>
		<!--Every profile-->
		<section class="w-1/2 bg-main-bg  h-full relative  md:w-full md:h-fit md:bg-transparent">
			<i class="fa-solid fa-x absolute top-0 right-0 m-2 md:hidden z-10" on:click={closeNameList} />
			{#each contacts as contact}
				<section
					class="  py-2 px-3  md:my-2 items-center  w-full hover:opacity-90 cursor-pointer "
					on:click={changeChat(contact.chatRoomId)}
					id={contact.username}
				>
					<!--name and username-->
					<section class="flex  items-center  ">
						<!--profile img-->
						{#if contact.profileImg}
							<!-- svelte-ignore a11y-img-redundant-alt -->
							<img
								class="h-12 w-12 object-cover rounded-full  "
								src="/api/{contact.profileImg}"
								alt="Current profile photo"
							/>
						{:else}
							<div
								class="h-12 w-12 rounded-full hover:opacity-90 bg-main-bg flex items-center justify-center"
							>
								<i class="fa-solid fa-user text-slate-400 text-2xl" />
							</div>
						{/if}

						<!-- Name and username-->
						<div class="mx-2 w-9/12 ">
							<h4 class=" font-semibold text-sm text-gray-900  ">
								{contact.firstName}
								{contact.lastName}
							</h4>
							<h5 class=" text-xs  text-gray-900   ">@{contact.username}</h5>
						</div>
					</section>
				</section>
			{/each}
		</section>
	</section>
	<!--Main Chat -->
	<section class="md:w-9/12 flex flex-col w-full ">
		<!--top part-->
		<section class="flex items-center w-full border-b-main-bg border-2 py-2 px-3  ">
			<!--Mobile back-->
			<button class="mr-8 md:hidden" on:click={showNameList}
				><i class="fas fa-chevron-left" /></button
			>
			{#if currentChatInfo}
				<!--profile img-->
				<a href="/profile/{currentChatInfo.username}" class=" w-fit">
					{#if currentChatInfo.profileImg}
						<!-- svelte-ignore a11y-img-redundant-alt -->
						<img
							class="h-12 w-12 object-cover rounded-full  "
							src="/api/{currentChatInfo.profileImg}"
							alt="Current profile photo"
						/>
					{:else}
						<div
							class="h-12 w-12 rounded-full hover:opacity-90 bg-main-bg flex items-center justify-center"
						>
							<i class="fa-solid fa-user text-slate-400 text-2xl" />
						</div>
					{/if}
				</a>
				<h4 class=" font-semibold text-sm text-gray-900 mx-4 flex-1">
					{currentChatInfo.firstName}
					{currentChatInfo.lastName}
				</h4>
				<h5 class="text-xs text-orange mx-2 cursor-default">{currentChatInfo.onlineTime}</h5>
			{:else}
				<div class="flex items-center w-full loading-pulse">
					<div
						class="h-12 w-12 rounded-full hover:opacity-90 bg-main-bg flex items-center justify-center"
					>
						<i class="fa-solid fa-user text-slate-400 text-2xl" />
					</div>

					<div class="  text-sm  mx-4 flex-1 h-2 flex items-center">
						<div class="h-full w-16 bg-main-bg  rounded mx-1 " />
						<div class="h-full w-16 bg-main-bg  rounded mx-1 " />
					</div>
					<div class=" mx-2 cursor-default h-2  flex items-center">
						<div class="h-full w-16 bg-main-bg  rounded " />
					</div>
				</div>
			{/if}
		</section>
		<!--Middle part-->
		<section class="flex-1 m-2 overflow-y-auto overflow-x-hidden flex flex-col" id="middlePart">
			{#if !loading}
				{#each currentChat as chat}
					<section class="w-full">
						<h1 class={chat.receive ? 'receive-message' : 'send-message'}>
							{chat.message}
						</h1>
					</section>
				{/each}
			{:else}
				<section class="w-full loading-pulse">
					<div class=" h-24 loading-msg bg-main-bg w-5/12 " />
					<div class=" h-24 loading-msg bg-main  w-5/12 float-right " />
				</section>
			{/if}
		</section>
		<!--bottom part-->
		<section class="">
			<form
				action=""
				method="post"
				class="flex m-2 gap-2 items-center"
				on:submit|preventDefault={sendMessage}
			>
				<input
					on:focus={typing}
					bind:value={textInput}
					type="text"
					class="flex-1 p-2  border-2 border-main rounded-full focus:outline-none foucs:border-main {loading
						? 'loading-pulse'
						: ''}"
					placeholder="Say Hello"
				/>
				<button
					class="font-semibold rounded-full bg-main py-2  px-4 w-11 h-11 {loading
						? 'loading-pulse'
						: ''}"
					id="sendBtn"><i class="fas fa-chevron-right" /></button
				>
			</form>
		</section>
	</section>
</div>
