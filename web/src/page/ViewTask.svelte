<script lang="ts">
    import { Card, CardBody, CardTitle, CardText, Badge, Button, Container } from 'sveltestrap';
    import { push } from 'svelte-spa-router';
    import { getTask, type Task } from '../common/task';
    import { check_token_or_login } from '../common/util';

    export let params: any;

    check_token_or_login()

    const token: string = sessionStorage.getItem("token") ?? "";
    const competition_id: string = params.cid;
    const task_id: string = params.tid;

    let loading = true;
    let task: Task;

    async function load() {
        let temp: Task | null = await getTask(competition_id, task_id, token)

        if (temp !== null) {
            task = temp;
            loading = false;
        }
    }

    load()
</script>

{#if loading}
<div>Loading task...</div>
{:else}
<Container class="mt-5">
    <Card>
        <CardBody>
            <CardTitle>{task.name}</CardTitle>
  
            <CardText>
                <strong>Description:</strong> {task.description}
            </CardText>
  
            <CardText>
                <strong>Execution Time:</strong> {task.execution_time} seconds
            </CardText>
  
            <CardText>
                <strong>Setup Code:</strong>
                <pre><code>{task.setup_code}</code></pre>
            </CardText>
  
            <CardText>
                <strong>User Code:</strong>
                <pre><code>{task.user_code}</code></pre>
            </CardText>
  
            <Button color="primary" on:click={() => push(`/competitions/${task.competition_id}/tasks/${task.id}/edit`)}>
                Edit Task
            </Button>
        </CardBody>
    </Card>
</Container>
{/if}