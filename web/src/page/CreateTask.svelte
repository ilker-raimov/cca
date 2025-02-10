<script lang="ts">
    import { Col, Button, Form, FormGroup, Label, Input, Row } from "sveltestrap";
    import Editor from "./Editor.svelte";
    import { check_or_warning, error, warning, success } from "../common/toast";
    import { check_token_or_login } from "../common/util";
    import { push } from "svelte-spa-router";

    export let params: any;

    check_token_or_login();

    const competition_id: string = params.id;

    type Test = {
        id: number,
        input: FileList | undefined,
        output: FileList | undefined
    }

    type Task = {
        name: string;
        description: string;
        execution_time: number;
        setup_code: string;
        user_code: string;
    };

    let examples: Test[] = [ {id: 0, input: undefined, output: undefined} ]
    let tests: Test[] = [ {id: 0, input: undefined, output: undefined} ]
    let task: Task = {
        name: "",
        description: "",
        execution_time: 2,
        setup_code: "",
        user_code: ""
    }

    function addExample(): void {
        examples = [...examples, { id: examples.length, input: undefined, output: undefined }];
    }

    function removeExample(): void {
        examples = examples.slice(0, examples.length - 1);
    }

    function addTest(): void {
        tests = [...tests, { id: tests.length, input: undefined, output: undefined }];
    }

    function removeTest(): void {
        tests = tests.slice(0, tests.length - 1);
    }

    const handleSetupCodeUpdate = (code: string) => {
        task.setup_code = code;
    };

    const handleUserCodeUpdate = (code: string) => {
        task.user_code = code;
    };

    function checkInput(): boolean {
        let check_name: boolean = check_or_warning(task.name, "Name is required!");
        let check_description: boolean = check_or_warning(task.description, "Description is required!");
        let check_examples: boolean = check_or_warning(examples.length !== 0, "Examples are required!");
        let check_tests: boolean = check_or_warning(tests.length !== 0, "Tests are required!");
        let check_setup_code: boolean = check_or_warning(task.setup_code, "Setup code is required!");
        let check_user_code: boolean = check_or_warning(task.user_code, "User code is required!");
        let check_execution_time: boolean = check_or_warning(task.execution_time > 0, "Execution time must be a non-negative integer!");

        if (check_name || check_description || check_examples || check_tests || check_setup_code || check_user_code || check_execution_time) {
            return false
        }

        return true;
    }

    async function createTask() {
        let is_input_ok: boolean = checkInput();

        if (!is_input_ok) {
            return;
        }

        const token: string = sessionStorage.getItem("token") ?? "";

        try {
            const response: Response = await fetch(`/api/competitions/${competition_id}/tasks`, {
                method: "POST",
                headers: { "Authorization": token },
                body: JSON.stringify(task)
            });

            if (!response.ok) {
                const data: string = await response.text();

                warning(data);

                return;
            }

            success("Successfully created task!");

            push(`/competitions/${competition_id}/view`)
        } catch(err: any) {
            error(err);
        }
    }
</script>

<div class="text-center mb-4">
    <h2>Examples</h2>
</div>

<main>
    <Form>
        <Row>
            <Col>
                <FormGroup>
                    <Label for="taskName">Task Name</Label>
                    <Input
                        type="text"
                        id="taskName"
                        bind:value={task.name}
                        placeholder="Enter task name"
                    />
                </FormGroup>
            </Col>
        </Row>

        <Row>
            <Col>
                <FormGroup>
                    <Label for="taskDescription">Description</Label>
                    <Input
                        type="textarea"
                        id="taskDescription"
                        bind:value={task.description}
                        placeholder="Enter task description"
                    />
                </FormGroup>
            </Col>
        </Row>

        <h5>Examples</h5>
        {#each examples as example, index}
            <Row>
                <Col sm="5">
                    <FormGroup>
                        <Label for={`exampleInput${example.id}`}>Example {example.id + 1} Input</Label>
                        <Input
                            type="file"
                            id={`exampleInput${example.id}`}
                            accept=".txt"
                            bind:files={example.input}
                        />
                    </FormGroup>
                </Col>
                <Col sm="5">
                    <FormGroup>
                        <Label for={`exampleOutput${example.id}`}>Example {example.id + 1} Output</Label>
                        <Input
                            type="file"
                            id={`exampleOutput${example.id}`}
                            accept=".txt"
                            bind:files={example.output}
                        />
                    </FormGroup>
                </Col>
                <Col sm="2">
                    <Button
                        color="danger"
                        on:click={() => removeExample()}
                        style="margin-top: 32px; width: 100%"
                    >
                        Remove Example
                    </Button>
                </Col>
            </Row>
        {/each}

        <Button color="primary" on:click={addExample}>
            Add Example
        </Button>

        <h5 class="mt-4">Tests</h5>
        {#each tests as test, index}
            <Row>
                <Col sm="5">
                    <FormGroup>
                        <Label for={`testInput${test.id}`}>Test {test.id + 1} Input</Label>
                        <Input
                            type="file"
                            id={`testInput${test.id}`}
                            accept=".txt"
                            bind:files={test.input}
                        />
                    </FormGroup>
                </Col>
                <Col sm="5">
                    <FormGroup>
                        <Label for={`testOutput${test.id}`}>Test {test.id + 1} Output</Label>
                        <Input
                            type="file"
                            id={`testOutput${test.id}`}
                            accept=".txt"
                            bind:files={test.output}
                        />
                    </FormGroup>
                </Col>
                <Col sm="2">
                    <Button
                        color="danger"
                        on:click={() => removeTest()}
                        style="margin-top: 32px; width: 100%"
                    >
                        Remove Test
                    </Button>
                </Col>
            </Row>
        {/each}

        <Button color="primary" on:click={addTest}>
            Add Test
        </Button>

        <Row class="mt-2">
            <Col>
                <FormGroup>
                    <Label for="executionTime">Execution Time</Label>
                    <Input
                        type="number"
                        id="executionTime"
                        bind:value={task.execution_time}
                        placeholder="Enter task execution time"
                    />
                </FormGroup>
            </Col>
        </Row>

        <Row class="mt-2">
            <Col>
                <FormGroup>
                    <Label for="setup_code">Setup Code</Label>
                    <Editor id="editor_setup" onCodeChange={handleSetupCodeUpdate}></Editor>
                </FormGroup>
            </Col>

            <Col>
                <FormGroup>
                    <Label for="user_code">User Code</Label>
                    <Editor id="editor_user" onCodeChange={handleUserCodeUpdate}></Editor>
                </FormGroup>
            </Col>
        </Row>

        <Button color="success" class="mt-4" style="width: 100%" on:click={createTask}>
            Create Task
        </Button>
    </Form>
</main>

<style>
    main {
        padding: 20px;
    }
    .mt-4 {
        margin-top: 20px;
    }
</style>