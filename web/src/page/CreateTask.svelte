<script lang="ts">
    import { Col, Button, Form, FormGroup, Label, Input, Row } from "sveltestrap";

    type Example = {
        id: number,
        input: string;
        output: string;
    }

    type Task = {
        id: number;
        name: string;
        description: string;
        examples: Example[];
        restrictions: string;
        executionTime: string;
        testFile: File | null;
    };

    let examples: Example[] = [ {id: 0, input: "", output: ""} ]
    let task: Task = {
        id: 0,
        name: "",
        description: "",
        examples: examples,
        restrictions: "",
        executionTime: "",
        testFile: null }
    

    function addExample(): void {
        examples = [...examples, { id: examples.length, input: "", output: "" }];
    }

    function removeExample(): void {
        examples = examples.slice(0, examples.length - 1);
    }
</script>

<div class="text-center mb-4">
    <h2>Examples</h2>
</div>

<Form>
    {#each examples as example}
    <FormGroup>
        <Row class="g-2">
            <Col>
                <Label for="example_{example.id}_input">Example {example.id + 1} Input</Label>
                <Input type="textarea" id="example_{example.id}_input" bind:value={examples[example.id].input} placeholder="Enter input" />
            </Col>

            <Col>
                <Label for="example_{example.id}_output">Example {example.id + 1} Output</Label>
                <Input type="textarea" id="example_{example.id}_output" bind:value={examples[example.id].output} placeholder="Enter output" />
            </Col>
        </Row>
    </FormGroup>
    {/each}
</Form>

<div class="d-flex justify-content-between mt-3">
    <Button color="primary" on:click={addExample}>Add Example</Button>
    {#if examples.length > 1}
    <Button color="danger" on:click={removeExample}>Remove Example</Button>
    {/if}
</div>