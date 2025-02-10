<script lang="ts">
	import Router, { push } from "svelte-spa-router";
    import { SvelteToast } from '@zerodevx/svelte-toast'

    import Welcome from './page/Welcome.svelte';
	import Login from './page/Login.svelte';
    import Register from './page/Register.svelte';
	import Dashboard from './page/Dashboard.svelte';
    import CreateCompetition from './page/CreateCompetition.svelte';
	import Task from './page/Task.svelte';
    import ViewCompetition from "./page/ViewCompetition.svelte";
    import CreateTask from "./page/CreateTask.svelte";
    import ViewTask from "./page/ViewTask.svelte";

	const routes = {
        '/': Welcome,
		'/login': Login,
        '/register': Register,
		'/dashboard': Dashboard,
        '/competitions/create': CreateCompetition,
        '/competitions/:id/view': ViewCompetition,
        '/competitions/:id/tasks/create': CreateTask,
        '/competitions/:cid/tasks/:tid/view': ViewTask,
		'/task': Task
	}
</script>

<Router {routes}/>

<SvelteToast/>