<script lang="ts">
  import { setContext } from "svelte";
  import { title } from "../../stores/titleStore";

  setContext("title", title);
</script>

<div class="app">
  <header>
    {#if $title !== ""}
      <h1 class="title">{$title}</h1>
    {/if}
  </header>
  <main>
    <slot />
  </main>
</div>

<style>
  .title {
    text-align: center;
    border: white;
    width: 96.5%;
  }
</style>
