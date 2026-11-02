<script lang="ts">
  import '../app.css';
  import { page } from '$app/state';
  import Navbar from '$lib/components/Navbar.svelte';
  import Footer from '$lib/components/Footer.svelte';

  let { children } = $props();

  let isAdmin = $derived(page.url.pathname.startsWith('/admin'));
</script>

<div class="min-h-screen flex flex-col font-sans-modern bg-slate-50 text-slate-900 selection:bg-emerald-600 selection:text-white">
  {#if !isAdmin}
    <Navbar />
  {/if}

  <main class="flex-1">
    {@render children()}
  </main>

  {#if !isAdmin}
    <Footer />
  {/if}
</div>
