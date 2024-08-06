<svelte:options immutable={false} />

<script lang="ts">
  import { onMount } from "svelte";

  interface Rep {
    Name: string;
    Owner: string;
  }
  let page = 0;
  let r: Rep[] = [{ Name: "loading...", Owner: "" }];
  let end = false;
  const url: string = "/api/rep";
  async function getReps(): Promise<Rep[]> {
    try {
      const response = await fetch(url);
      const res = JSON.parse(await response.text()) as Rep[];
      r = res;
      page = 1;
    } catch (error) {
      console.error("Error:", error);
    }
    return r;
  }
  async function loadMore() {
    page++;
    const response = await fetch(url + "?page=" + page);
    const res = JSON.parse(await response.text()) as Rep[];
    r = r.concat(res);
    if (res.length === 0) {
      end = true;
    }
  }
  onMount(async () => {
    await getReps();
  });
</script>

<main>
  <h1>Choose your repository</h1>
  <table>
    <tr><p>---------------------</p></tr>
    {#each r as rep}
      {#if rep.Name !== ""}
        <tr>
          {#if rep.Name === "loading..."}
            <p class="button">{rep.Name}</p>
          {:else}
            <a
              href="/repos/{rep.Owner}/{rep.Name}/"
              class="button">{rep.Name}</a
            >
          {/if}
        </tr>
        <tr><p>---------------------</p></tr>
      {/if}
    {/each}
  </table>
  {#if !end}
    <button on:click={loadMore}>LOAD MORE</button>
  {/if}
</main>

<style>
  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }

  h1 {
    color: #ff3e00;
    text-transform: uppercase;
    font-size: 3em;
    font-weight: 100;
  }

  table {
    color: #ff3e00;
    margin-left: auto;
    margin-right: auto;
    font-size: 1.5em;
    font-weight: 100;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>
