<script lang="ts">
	import Router, { push } from "svelte-spa-router";
    import { SvelteToast } from '@zerodevx/svelte-toast'

    import Welcome from './page/Welcome.svelte';
	import Login from './page/Login.svelte';
    import Register from './page/Register.svelte';
	import Dashboard from './page/Dashboard.svelte';
    import CreateCompetition from './page/CreateCompetition.svelte';
	import Task from './page/Task.svelte';

	const routes = {
        '/': Welcome,
		'/login': Login,
        '/register': Register,
		'/dashboard': Dashboard,
        '/competitions/create': CreateCompetition,
		'/task': Task
	}
</script>

<Router {routes}/>

<SvelteToast/>