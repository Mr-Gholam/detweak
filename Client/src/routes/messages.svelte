<script>
	// @ts-nocheck

	import { io } from 'socket.io-client';
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import formatDistanceToNow from 'date-fns/formatDistanceToNow/index.js';
	let contacts = [];
	let currentChat = [];
	let currentChatInfo;
	let textInput;
	let picInput;
	let hasPhoto = false;
	let showList = true;
	let bouncing = false;
	let loading = true;
	const targetUser = $page.url.search;
	const targetUsername = targetUser.split('=')[1];

	onMount(async () => {
		const response = await fetch('/api/load-chatRooms');
		const data = await response.json();
		contacts = data;
		if (targetUsername) {
			if (contacts == null) {
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
			} else {
				const foundUser = contacts.find((contact) => contact.Username == targetUsername);
				if (foundUser) {
					const response = await fetch('/api/get-chat', {
						method: 'POST',
						headers: {
							'Content-Type': 'application/json'
						},
						body: JSON.stringify({
							RoomId: foundUser.RoomId
						})
					});
					const data = await response.json();
					currentChatInfo = data;
					currentChat = data.Chat;
					loading = false;
				} else {
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
					console.log(newChatRoom);
					if (result.status == 200) {
						contacts.unshift(newChatRoom);
						contacts = contacts;
					}
				}
			}
		}
	});
	function scrollToBottom() {
		const element = document.getElementById('middlePart');
		element.scrollTop = element.scrollHeight;
	}
	function createSendMessage(input, createdAt) {
		const middePart = document.getElementById('middlePart');
		const messageParent = document.createElement('div');
		const parent = document.createElement('div');
		parent.classList.add('w-full');
		middePart.appendChild(parent);
		messageParent.classList.add('flex', 'items-end', 'bg-inherit', 'send-message');
		const message = document.createElement('h1');
		const Time = document.createElement('p');
		message.classList.add('p-1', 'bg-inherit');
		Time.classList.add('text-xs', 'float-left', 'mx-2', 'bg-inherit');
		const time = formatDistanceToNow(new Date(createdAt), { addSuffix: true });
		message.innerText = input;
		Time.innerText = time;
		messageParent.appendChild(message);
		messageParent.appendChild(Time);
		parent.appendChild(messageParent);
		scrollToBottom();
	}
	function createSendImgMessage(imgUrl, createdAt, input) {
		const middePart = document.getElementById('middlePart');
		const messageParent = document.createElement('div');
		const parent = document.createElement('div');
		parent.classList.add('w-full');
		middePart.appendChild(parent);
		messageParent.classList.add('flex', 'items-end', 'bg-inherit', 'send-message', 'flex-col');
		const img = document.createElement('img');
		const message = document.createElement('h1');
		const Time = document.createElement('p');
		img.setAttribute('src', `/api/images/${imgUrl}`);
		message.classList.add('p-1', 'bg-inherit');
		Time.classList.add('text-xs', 'float-left', 'mx-2', 'bg-inherit');
		const time = formatDistanceToNow(new Date(createdAt), { addSuffix: true });
		Time.innerText = time;
		message.innerText = input;
		messageParent.appendChild(img);
		messageParent.appendChild(message);
		messageParent.appendChild(Time);
		parent.appendChild(messageParent);
		scrollToBottom;
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
	function togglePic() {
		if (!hasPhoto) {
			hasPhoto = true;
		} else {
			hasPhoto = false;
		}
	}
	async function sendMessage() {
		const RoomId = currentChatInfo.RoomId;
		const TargetId = currentChatInfo.TargetId;
		const formData = new FormData();
		formData.append('image', picInput.files[0]);
		if (hasPhoto) {
			const image = await fetch(`/api/create-message-img/${TargetId}`, {
				method: 'POST',
				body: formData
			});
			const imgData = await image.json();
			togglePic();

			if (textInput == '' || textInput == undefined) {
				createSendImgMessage(imgData.ImgUrl, imgData.CreatedAt, '');
				scrollToBottom();
				return;
			}
			const response = await fetch('/api/update-message', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					message: textInput,
					messageId: imgData.MessageId
				})
			});
			if (response.ok) {
				createSendImgMessage(imgData.ImgUrl, imgData.CreatedAt, textInput);
				scrollToBottom();
				textInput = '';
			}
		}
		if (textInput == '' || textInput == undefined) return;
		const response = await fetch('/api/create-message', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				Message: textInput,
				RoomId,
				TargetId
			})
		});
		const data = await response.json();
		if (response.status == 200) {
			createSendMessage(textInput, data.CreatedAt);
			scrollToBottom();
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
				RoomId: chatRoomId
			})
		});
		const data = await response.json();
		currentChatInfo = data;
		currentChat = data.Chat;
		setTimeout(() => {
			scrollToBottom();
		}, 5);
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
			{#if contacts}
				{#each contacts as contact}
					<section
						class="  py-2 px-3  md:my-2 items-center  w-full hover:opacity-90 cursor-pointer "
						on:click={changeChat(contact.RoomId)}
						id={contact.Username}
					>
						<!--name and username-->
						<section class="flex  items-center  ">
							<!--profile img-->
							{#if contact.ImgUrl}
								<!-- svelte-ignore a11y-img-redundant-alt -->
								<img
									class="h-12 w-12 object-cover rounded-full  "
									src="/api/images/{contact.ImgUrl}"
									alt="Current profile photo"
								/>
							{:else}
								<div
									class="h-12 w-12 rounded-full hover:opacity-90 bg-main-bg flex items-center justify-center border-2 border-border"
								>
									<i class="fa-solid fa-user text-slate-400 text-2xl" />
								</div>
							{/if}

							<!-- Name and username-->
							<div class="mx-2 w-9/12 ">
								<h4 class=" font-semibold text-sm text-text  ">
									{contact.Firstname}
									{contact.Lastname}
								</h4>
								<h5 class=" text-xs  text-text   ">@{contact.Username}</h5>
							</div>
						</section>
					</section>
				{/each}
			{/if}
		</section>
	</section>
	<!--Main Chat -->
	<section class="md:w-9/12 flex flex-col w-full border-2 border-border rounded ">
		<!--top part-->
		<section class="flex items-center w-full border-border border-b-2 py-2 px-3  ">
			<!--Mobile back-->
			<button class="mr-8 md:hidden" on:click={showNameList}
				><i class="fas fa-chevron-left" /></button
			>
			{#if currentChatInfo}
				<!--profile img-->
				<a href="/profile/{currentChatInfo.Username}" class=" w-fit">
					{#if currentChatInfo.ImgUrl}
						<!-- svelte-ignore a11y-img-redundant-alt -->
						<img
							class="h-12 w-12 object-cover rounded-full  "
							src="/api/images/{currentChatInfo.ImgUrl}"
							alt="Current profile photo"
						/>
					{:else}
						<div
							class="h-12 w-12 rounded-full hover:opacity-90 bg-main-bg flex items-center justify-center border-2 border-border"
						>
							<i class="fa-solid fa-user text-slate-400 text-2xl" />
						</div>
					{/if}
				</a>
				<h4 class=" font-semibold text-sm text-text mx-4 flex-1">
					{currentChatInfo.Firstname}
					{currentChatInfo.Lastname}
				</h4>
				<div class="flex items-center">
					<i
						class="fa-solid fa-ellipsis-vertical px-2.5 text-base cursor-pointer hover:text-main text-text"
					/>
				</div>
			{:else}
				<div class="flex items-center w-full loading-pulse">
					<div
						class="h-12 w-12 rounded-full hover:opacity-90  flex items-center justify-center border-2 border-text"
					>
						<i class="fa-solid fa-user text-text text-2xl" />
					</div>

					<div class="  text-sm  mx-4 flex-1 h-2 flex items-center">
						<div class="h-full w-16 bg-text  rounded mx-1 " />
						<div class="h-full w-16 bg-text  rounded mx-1 " />
					</div>
				</div>
			{/if}
		</section>
		<!--Middle part-->
		<section class="flex-1 m-2 overflow-y-auto overflow-x-hidden flex flex-col" id="middlePart">
			{#if !loading}
				{#if currentChat}
					{#each currentChat as chat}
						<div class="w-full">
							{#if chat.Receive}
								<div class="flex items-end">
									{#if currentChatInfo.ImgUrl}
										<!-- svelte-ignore a11y-img-redundant-alt -->
										<img
											class="h-8 w-8 object-cover rounded-full  "
											src="/api/images/{currentChatInfo.ImgUrl}"
											alt="Current profile photo"
										/>
									{:else}
										<div
											class="h-8 w-8 rounded-full hover:opacity-90 bg-main-bg flex items-center justify-center border-2 border-border"
										>
											<i class="fa-solid fa-user text-slate-400 text-base bg-transparent" />
										</div>
									{/if}
									<div
										class="{chat.Receive
											? 'receive-message'
											: 'send-message'} flex items-end bg-inherit {chat.ImgUrl ? 'flex-col' : ''}"
									>
										{#if chat.ImgUrl}
											<img
												src="/api/images/{chat.ImgUrl}"
												alt=""
												class="w-full h-fit  md:mx-auto  object-cover rounded-sm"
											/>
										{/if}
										<h1 class="p-1 bg-inherit">
											{chat.Message}
										</h1>
										<p class="text-xs float-right mx-2 bg-inherit">
											{formatDistanceToNow(new Date(chat.CreatedAt), { addSuffix: true })}
										</p>
									</div>
								</div>
							{:else}
								<div
									class="{chat.Receive
										? 'receive-message'
										: 'send-message'} flex items-end {chat.ImgUrl ? 'flex-col' : ''}  "
								>
									{#if chat.ImgUrl}
										<img
											src="/api/images/{chat.ImgUrl}"
											alt=""
											class="w-full h-fit  md:mx-auto  object-cover rounded-sm"
										/>
									{/if}
									<h1 class="p-1 bg-inherit">
										{chat.Message}
									</h1>
									<p class="text-xs float-left mx-2 bg-inherit">
										{formatDistanceToNow(new Date(chat.CreatedAt), { addSuffix: true })}
									</p>
								</div>
							{/if}
						</div>
					{/each}
				{/if}
			{:else}
				<section class="w-full loading-pulse">
					<div class=" h-24 loading-msg bg-main-bg w-5/12 " />
					<div class=" h-24 loading-msg bg-main  w-5/12 float-right " />
				</section>
			{/if}
		</section>
		<!--bottom part-->
		<section class="border-t-2 border-border">
			<form
				action=""
				method="post"
				class="flex m-2 gap-2 items-center  rounded-full p-1"
				on:submit|preventDefault={sendMessage}
			>
				<input
					on:focus={typing}
					bind:value={textInput}
					type="text"
					class="flex-1 p-2   focus:outline-none  rounded-full foucs:border-main text-text {loading
						? 'loading-pulse'
						: ''}"
					placeholder="Say Hello"
				/>
				<label
					for="picture"
					class="text-2xl  hover:cursor-pointer hover:text-main mx-2 flex items-center text-text {hasPhoto
						? 'text-main'
						: 'text-text'} "><i class="fa-solid fa-paperclip" /></label
				>
				<input
					type="file"
					name="image"
					id="picture"
					on:change={togglePic}
					class="hidden"
					bind:this={picInput}
				/>
				<button
					class="font-semibold rounded-full border-2 border-main bg-inherit 	justify-center w-11 h-11 mx-2 flex items-center {loading
						? 'loading-pulse'
						: ''}"
					id="sendBtn"><i class="fas fa-chevron-right text-main bg-inherit" /></button
				>
			</form>
		</section>
	</section>
</div>
