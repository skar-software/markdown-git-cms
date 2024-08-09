<script lang="ts">
  function handleOnClick(route: string, type = "normal") {
    if (type === "normal") {
      window.location.href = route;
    } else if (type === "help") {
      window.open(route, "_blank");
    }
  }
</script>

<div class="app">
  <header>
    <h1>Pubdeskmd</h1>
    <h2>Markdown editor with Github integration - to edit and update files</h2>
  </header>
  <main>
    <slot />
  </main>
</div>

<style>
  h2 {
    color: #415a77;
    text-decoration: underline;
    font-size: 1.5em;
    margin-top: 0;
    font-weight: 600;
    text-align: center;
}
</style>