<script lang="ts">
	import Router, { push } from "svelte-spa-router";
    import { SvelteToast } from '@zerodevx/svelte-toast'

    import Welcome from './page/Welcome.svelte';
	import Login from './page/Login.svelte';
    import Register from './page/Register.svelte';
	import Dashboard from './page/Dashboard.svelte';
	import Task from './page/Task.svelte';

	const routes = {
        '/welcome': Welcome,
		'/login': Login,
        '/register': Register,
		'/dashboard': Dashboard,
		'/task': Task
	}

	push("/welcome");
</script>

<Router {routes}/>

<SvelteToast/>