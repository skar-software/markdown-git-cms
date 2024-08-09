<script lang="ts">
  import "../../style.css";
  let err = "";

  async function handleSubmit(event: SubmitEvent) {
    const form = event.target as HTMLFormElement;
    const data = new FormData(form);

    let url = "/api/set/?";
    let tknVal = data.get("tkn")?.toString();
    if (tknVal === undefined) {
      tknVal = "";
    }
    const response = await fetch(url + new URLSearchParams({ tkn: tknVal }).toString(), {
      method: "Get",
    });
    if (response.status !== 200) {
      err = await response.text();
    } else {
      window.location.href = "/repos";
    }
  }
</script>

<main>
  <!-- <h1>Set up your gihub token!</h1> -->
  <form
    class="login-form"
    on:submit|preventDefault={handleSubmit}
  >
    <label>
      <span>Enter your Github token</span>
      <input
        class="file-input"
        type="password"
        name="tkn"
      />
    </label>
    <button
      class="login-button"
      type="submit">Submit</button
    >
    {#if err !== ""}
      <span style="color: #ff3e00;">{err}</span>
    {/if}
    <a href="https://skar-software.github.io/pubdeskmd" target="_blank">Help</a>
  </form>
</main>

<style>
  main {
    position: relative;
    display: flex;
    justify-content: center;
    align-items: center;
    height: 50vh;
  }
  a{
    font-size: 1.5rem;
    margin-left: 15px;
    font-weight: 600;
    text-decoration: underline;
    text-decoration-color: red;
    color: #415a77;
  }
</style>
