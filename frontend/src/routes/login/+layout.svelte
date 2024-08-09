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
    <div style="display: flex; flex-direction: row; justify-content:center; align-items:flex-start">
     
    </div>
  </header>
  <main>
    <slot />
  </main>
</div>

<style>

</style>
