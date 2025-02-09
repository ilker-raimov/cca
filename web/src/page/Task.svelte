<script lang="ts">
    import { Container, Row, Col, Button, Card, CardBody, CardTitle } from 'sveltestrap';
    import * as ace from 'brace';
    import 'brace/mode/java.js';
    import 'brace/theme/chrome.js';
    import AceEditor from 'svelte-ace/src/AceEditor.svelte';

    let code = `function hello() {\n  console.log("Hello, World!");\n}`;

    function saveCode() {
      console.log("Code saved:", code);
    }
  
    function compileCode() {
      console.log("Compiling code...");
    }
  
    function checkCode() {
      console.log("Checking code...");
    }
  
    function submitCode() {
      console.log("Submitting code...");
    }
  </script>
  <!-- REWORK NOT TO USE CARD -->
  <Container fluid class="vh-100">
    <Row class="h-100">
        <Col md="5" class="border-right p-4 overflow-auto">
            <Card>
                <CardBody>
                    <CardTitle><h3>Task Description</h3></CardTitle>
                    <p><strong>Problem:</strong> Write a function that prints "Hello, World!"</p>
                    <p><strong>Constraints:</strong></p>
                    <ul>
                    <li>The function must be named <code>hello</code>.</li>
                    <li>It should print exactly "Hello, World!"</li>
                    <li>Use JavaScript</li>
                    </ul>
                </CardBody>
            </Card>
        </Col>

        <Col md="7" class="d-flex flex-column p-4">
            <Card class="flex-grow-1">
                <CardBody>
                    <CardTitle><h3>Code Editor</h3></CardTitle>
                    <AceEditor lang="java" theme="chrome" height=600px bind:value={code}/>
                </CardBody>
            </Card>

            <div class="d-flex justify-content-between mt-3">
                <Button color="secondary" on:click={saveCode}>Save</Button>
                <Button color="info" on:click={compileCode}>Compile</Button>
                <Button color="warning" on:click={checkCode}>Check</Button>
                <Button color="success" on:click={submitCode}>Submit</Button>
            </div>
        </Col>
    </Row>
</Container>
  