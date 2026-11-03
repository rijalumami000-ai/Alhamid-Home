<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { MediaPhoto, Category } from '$lib/types';
  import { Camera, X, ZoomIn, Sparkles } from 'lucide-svelte';

  let photos = $state<MediaPhoto[]>([]);
  let categories = $state<Category[]>([]);
  let selectedCategory = $state('');
  let lightboxPhoto = $state<MediaPhoto | null>(null);

  async function loadData() {
    try {
      const [photoRes, catRes] = await Promise.all([
        api.getPhotos(selectedCategory),
        api.getCategories('galeri_foto'),
      ]);
      photos = photoRes;
      categories = catRes;
    } catch (e) {
      console.error(e);
    }
  }

  function handleFilter(slug: string) {
    selectedCategory = slug;
    loadData();
  }

  onMount(() => {
    loadData();
  });
</script>

<svelte:head>
  <title>Galeri Foto | Pondok Pesantren Al-Hamid</title>
</svelte:head>

<!-- Lightbox Modal -->
{#if lightboxPhoto}
  <div
    class="fixed inset-0 z-50 bg-[#090D16]/95 backdrop-blur-xl flex items-center justify-center p-4"
    role="dialog"
    aria-modal="true"
    tabindex="-1"
    onkeydown={(e) => e.key === 'Escape' && (lightboxPhoto = null)}
  >
    <!-- Background overlay click -->
    <button
      class="absolute inset-0 w-full h-full cursor-default"
      onclick={() => (lightboxPhoto = null)}
      aria-label="Tutup preview"
    ></button>

    <button
      class="absolute top-6 right-6 z-10 p-3 rounded-full bg-white/10 hover:bg-white/20 text-white transition"
      onclick={() => (lightboxPhoto = null)}
      aria-label="Tutup foto"
    >
      <X class="w-6 h-6" />
    </button>

    <div class="relative z-10 max-w-4xl w-full max-h-[85vh] flex flex-col items-center">
      <img
        src={lightboxPhoto.image_url}
        alt={lightboxPhoto.title}
        class="max-h-[70vh] object-contain rounded-3xl shadow-2xl border border-white/10"
      />
      <div class="mt-4 text-center text-white space-y-1">
        <h3 class="font-display font-bold text-xl">{lightboxPhoto.title}</h3>
        {#if lightboxPhoto.caption}
          <p class="text-xs text-slate-400 max-w-lg font-normal">{lightboxPhoto.caption}</p>
        {/if}
      </div>
    </div>
  </div>
{/if}

<!-- Header -->
<div class="bg-[#090D16] text-white py-20 px-4 sm:px-8 relative overflow-hidden border-b border-slate-800/80">
  <div class="absolute inset-0 bg-grid-dark opacity-30 pointer-events-none"></div>
  <div class="absolute -top-24 left-1/2 -translate-x-1/2 w-[600px] h-[300px] bg-teal-500/15 blur-3xl pointer-events-none rounded-full"></div>

  <div class="max-w-6xl mx-auto space-y-4 relative z-10">
    <p class="text-xs font-mono uppercase tracking-[0.25em] font-semibold text-teal-400">
      Dokumentasi Visual
    </p>
    <h1 class="font-display font-black text-3xl sm:text-5xl lg:text-6xl text-white tracking-tight">
      Galeri Foto Kehidupan Santri
    </h1>
    <p class="text-slate-300 text-base sm:text-lg max-w-2xl font-normal leading-relaxed">
      Potret dinamika keseharian santri Al-Hamid: dari heningnya majelis sorogan, keceriaan halaqoh, perhelatan wisuda, hingga keindahan arsitektur pesantren.
    </p>
  </div>
</div>

<div class="max-w-6xl mx-auto px-4 sm:px-8 py-20 space-y-10">
  <!-- Categories Filter -->
  <div class="flex flex-wrap items-center gap-2">
    <button
      onclick={() => handleFilter('')}
      class="px-5 py-2.5 rounded-full text-xs font-display font-bold transition {!selectedCategory
        ? 'bg-slate-900 text-white shadow-md'
        : 'bg-white text-slate-600 hover:bg-slate-100 border border-slate-200'}"
    >
      Semua Album Foto
    </button>
    {#each categories as cat}
      <button
        onclick={() => handleFilter(cat.slug)}
        class="px-5 py-2.5 rounded-full text-xs font-display font-bold transition {selectedCategory === cat.slug
          ? 'bg-slate-900 text-white shadow-md'
          : 'bg-white text-slate-600 hover:bg-slate-100 border border-slate-200'}"
      >
        {cat.name}
      </button>
    {/each}
  </div>

  <!-- Photo Grid -->
  <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
    {#each photos as photo}
      <button
        class="bg-white rounded-3xl border border-slate-200/80 shadow-sm hover:shadow-2xl hover:-translate-y-1 transition-all duration-300 overflow-hidden group cursor-pointer text-left flex flex-col"
        onclick={() => (lightboxPhoto = photo)}
      >
        <div class="relative aspect-square overflow-hidden bg-slate-950 w-full">
          <img
            src={photo.image_url}
            alt={photo.title}
            class="w-full h-full object-cover group-hover:scale-105 transition duration-500"
          />
          <div class="absolute inset-0 bg-slate-950/40 opacity-0 group-hover:opacity-100 transition-opacity flex items-center justify-center text-white">
            <ZoomIn class="w-8 h-8 text-teal-300" />
          </div>
        </div>
        <div class="p-5 space-y-1 flex-1">
          <div class="text-[10px] font-black text-teal-700 uppercase tracking-wider">{photo.category?.name}</div>
          <h2 class="font-display font-bold text-sm text-slate-950 line-clamp-1 group-hover:text-teal-700 transition">{photo.title}</h2>
          {#if photo.caption}
            <p class="text-xs text-slate-500 line-clamp-1">{photo.caption}</p>
          {/if}
        </div>
      </button>
    {/each}
  </div>
</div>
