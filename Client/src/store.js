import { writable } from 'svelte/store'



export const User = writable({ username: null })

export const loading = writable(false)

export const onlineFriends = writable([])
export const ws = writable(null)
export const Notification = writable(null)