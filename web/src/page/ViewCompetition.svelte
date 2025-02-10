<script lang="ts">
    import { Card, CardBody, CardTitle, CardText, Badge, Container, Button, Table } from "sveltestrap";
    import { getCompetition, type Competition } from "../common/competition";
    import { getTaskIdList, getTask, type Task } from "../common/task";
    import { check_token_or_login } from "../common/util";
    import { push } from "svelte-spa-router";

    export let params: any;

    check_token_or_login();

    const token = sessionStorage.getItem("token") ?? "";
    const competition_id: string = params.id;

    let loading = true;
    let competition: Competition;
    let tasks: Task[] = [];

    async function load() {
        let temp: Competition | null = await getCompetition(competition_id, token);

        if (temp) {
            competition = temp;
        }

        const task_id_list: string[] = await getTaskIdList(competition_id, token);

        for (const task_id of task_id_list) {
            const task: Task | null = await getTask(competition_id, task_id, token);

            if (task == null) {
                continue;
            }

            tasks = [...tasks, task]
        }

        loading = false;
    }

    function deleteTask(competition_id: string, task_id: string) {
        alert(`Deleting task: ${competition_id}/${task_id}}`);
    }

    load();
</script>

{#if loading}
<p>Loading competition...</p>
{:else}
<Container class="mt-5">
    <Card>
        <CardBody>
            <CardTitle>{competition.title}</CardTitle>
            <CardText>
                <strong>Description:</strong>
                {competition.description}
            </CardText>

            <CardText>
                <strong>Language:</strong>
                {competition.language}
            </CardText>

            <CardText>
                <strong>Public:</strong>
                {#if competition.public}
                <Badge color="success">Yes</Badge>
                {:else}
                <Badge color="danger">No</Badge>
                {/if}
            </CardText>
            <CardText>
                <strong>Overall Time:</strong>
                {#if competition.use_overall_time}
                <Badge color="info">Enabled</Badge>
                {:else}
                <Badge color="secondary">Disabled</Badge>
                {/if}
            </CardText>
            <CardText>
                <strong>Execution Time:</strong>
                {#if competition.use_execution_time}
                <Badge color="info">Enabled</Badge>
                {:else}
                <Badge color="secondary">Disabled</Badge>
                {/if}
            </CardText>

            <Button color="primary" on:click={() => push(`/competitions/${competition.id}/tasks/create`)}>Create Task</Button>
        </CardBody>
    </Card>

    <Table striped bordered hover>
        <thead>
            <tr>
                <th>#</th>
                <th>Name</th>
                <th>Actions</th>
            </tr>
        </thead>
        <tbody>
            {#each tasks as task, index (task.id)}
                <tr>
                    <td>№{index + 1}</td>
                    <td>{task.name}</td>
                    <td>
                        <Button color="success" size="sm" on:click={() => push(`/competitions/${competition_id}/tasks/${task.id}/view`)}>View</Button>
                        <Button color="primary" size="sm" on:click={() => push(`/competitions/${competition_id}/tasks/${task.id}/edit`)}>Edit</Button>
                        <Button color="danger" size="sm" class="ms-2" on:click={() => deleteTask(competition_id, task.id)}>Delete</Button>
                    </td>
                </tr>
            {/each}
        </tbody>
    </Table>
</Container>
{/if}