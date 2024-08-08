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
    <h1>Hey, set up your gihub token!</h1>
  </header>
  <main>
    <slot />
  </main>
</div>