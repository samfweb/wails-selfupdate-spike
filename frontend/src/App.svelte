<script lang="ts">
  import { onMount } from "svelte";
  import logo from "./assets/images/logo-universal.png";
  import {
    Greet,
    CheckForUpdate,
    GetCurrentVersion,
    DoSelfUpdate,
    TestDb,
  } from "../wailsjs/go/main/App.js";

  let resultText: string = "Please enter your name below 👇";
  let name: string;
  let currentVersion = "";
  let pendingVersion = "";

  function greet(): void {
    Greet(name).then((result) => (resultText = result));
  }

  onMount(async () => {
    currentVersion = await GetCurrentVersion();
  });

  const checkVersion = async () => {
    console.log("Checking for update...");
    const res = await CheckForUpdate();
    if (res !== "") {
      console.log("New version available: " + res);
      pendingVersion = res;
    } else {
      console.log("No new version available");
    }
  };

  const testDb = async () => {
    const res = await TestDb();
    console.log(res);
  };
</script>

<main>
  <img alt="Wails logo" id="logo" src={logo} />
  <div class="result" id="result">Current version: {currentVersion}</div>
  <div class="result" id="result">Available version: {pendingVersion}</div>

  <div class="input-box" id="input">
    <input
      autocomplete="off"
      bind:value={name}
      class="input"
      id="name"
      type="text"
    />
    <button class="btn" on:click={checkVersion}>Check!!!!</button>
    <button class="btn" on:click={DoSelfUpdate}>Update!!!!</button>
    <button class="btn" on:click={testDb}>TestDB!!!!</button>
  </div>
</main>

<style>
  #logo {
    display: block;
    width: 50%;
    height: 50%;
    margin: auto;
    padding: 10% 0 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
  }

  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  .input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }
</style>
