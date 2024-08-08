<svelte:options immutable={false} />

<script lang="ts">
  import { onMount } from "svelte";
  import "../../style.css";
  import { getContext } from "svelte";

const titleStore = getContext("title");
// Update the title when this component is mounted
titleStore.set(`Choose your repository`);

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
  <table>
    <tr>
      <th>#</th>
      <th>Name</th>
    </tr>
    {#each r as rep, i}
      {#if rep.Name !== ""}
        <tr>
          {#if rep.Name !== "loading..."}
            <td>{i + 1}</td>
            <td
              ><a
                href="/repos/{rep.Owner}/{rep.Name}/"
                class="button">{rep.Name}</a
              ></td
            >
          {/if}
        </tr>
        <!-- <tr><p>---------------------</p></tr> -->
      {/if}
    {/each}
    <tr style=" margin: 0; padding: 0;">
      <td
        colspan="2"
        style=" margin: 0; padding: 0;"
      >
        {#if !end}
          <button
            style="margin: 0; padding: 5px 0; height: 100%; width: 100%;"
            on:click={loadMore}>LOAD MORE</button
          >
        {/if}
      </td>
    </tr>
  </table>
</main>
