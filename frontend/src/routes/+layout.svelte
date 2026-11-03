<script lang="ts">
  import '../app.css';
  import { onMount } from 'svelte';
  import { page } from '$app/state';
  import Navbar from '$lib/components/Navbar.svelte';
  import Footer from '$lib/components/Footer.svelte';
  import { applyFontPreset, getSavedFontPresetId } from '$lib/fonts';
  import { initSettings, siteSettings } from '$lib/stores/settings';
  import { api } from '$lib/api';

  let { children } = $props();

  let isAdmin = $derived(page.url.pathname.startsWith('/admin'));

  onMount(async () => {
    // 1. Apply from local cache immediately
    const cached = getSavedFontPresetId();
    applyFontPreset(cached);

    // 2. Fetch server settings globally & synchronize site identity
    await initSettings();
    if ($siteSettings.site_font_preset && $siteSettings.site_font_preset !== cached) {
      applyFontPreset($siteSettings.site_font_preset);
    }

    // 3. Listen for font updates across tabs or admin panel
    window.addEventListener('storage', (e) => {
      if (e.key === 'alhamid_font_preset' && e.newValue) {
        applyFontPreset(e.newValue);
      }
    });
  });
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
