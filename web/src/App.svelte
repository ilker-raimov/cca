<script lang="ts">
	import Router, { push } from "svelte-spa-router";
    import { SvelteToast } from '@zerodevx/svelte-toast'

	import Login from './page/Login.svelte';
	import Dashboard from './page/Dashboard.svelte';
	import Task from './page/Task.svelte';

	const routes = {
		'/login': Login,
		'/dashboard': Dashboard,
		'/task': Task
	}
	const token = localStorage.getItem("token");
	const redirect = token === null ? "/login" : "/dashboard";

	push(redirect);
</script>

<Router {routes}/>

<SvelteToast/>