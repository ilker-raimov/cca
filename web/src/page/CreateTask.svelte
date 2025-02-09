<script lang="ts">
    import { Col, Button, Form, FormGroup, Label, Input, Row } from "sveltestrap";

    export let params: any;

    const id: string = params.id;

    type Example = {
        id: number,
        input: string;
        output: string;
    }

    type Test = {
        id: number,
        input: FileList | undefined,
        output: FileList | undefined
    }

    type Task = {
        name: string;
        description: string;
        examples: Example[];
        execution_time: string;
        tests: Test[]
    };

    let examples: Example[] = [ {id: 0, input: "", output: ""} ]
    let tests: Test[] = [ {id: 0, input: undefined, output: undefined} ]
    let task: Task = {
        name: "",
        description: "",
        examples: examples,
        execution_time: "",
        tests: tests
    }

    function addExample(): void {
        examples = [...examples, { id: examples.length, input: "", output: "" }];
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
</script>

<div class="text-center mb-4">
    <h2>Examples</h2>
</div>

<main>
    <Form>
        <Row>
            <Col class="w-25">
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
            <Col class="w-25">
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
                            type="text"
                            id={`exampleInput${example.id}`}
                            bind:value={example.input}
                            placeholder="Enter example input"
                        />
                    </FormGroup>
                </Col>
                <Col sm="5">
                    <FormGroup>
                        <Label for={`exampleOutput${example.id}`}>Example {example.id + 1} Output</Label>
                        <Input
                            type="text"
                            id={`exampleOutput${example.id}`}
                            bind:value={example.output}
                            placeholder="Enter example output"
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
            <Col class="w-25">
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

        <Button color="success" class="mt-4" style="width: 100%">
            Save Task
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