<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { MediaPhoto, Category } from '$lib/types';
  import { Camera, X, ZoomIn } from 'lucide-svelte';

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
    class="fixed inset-0 z-50 bg-slate-950/90 backdrop-blur-md flex items-center justify-center p-4"
    role="button"
    tabindex="0"
    onclick={() => (lightboxPhoto = null)}
    onkeydown={(e) => e.key === 'Escape' && (lightboxPhoto = null)}
  >
    <button
      class="absolute top-6 right-6 p-2 rounded-full bg-white/10 hover:bg-white/20 text-white transition"
      onclick={() => (lightboxPhoto = null)}
      aria-label="Tutup foto"
    >
      <X class="w-6 h-6" />
    </button>
    <div
      class="max-w-4xl w-full max-h-[85vh] flex flex-col items-center"
      role="document"
      tabindex="-1"
      onclick={(e) => e.stopPropagation()}
      onkeydown={(e) => e.stopPropagation()}
    >
      <img
        src={lightboxPhoto.image_url}
        alt={lightboxPhoto.title}
        class="max-h-[70vh] object-contain rounded-2xl shadow-2xl"
      />
      <div class="mt-4 text-center text-white space-y-1">
        <h3 class="font-serif-luxury font-bold text-xl">{lightboxPhoto.title}</h3>
        {#if lightboxPhoto.caption}
          <p class="text-xs text-slate-300 max-w-lg">{lightboxPhoto.caption}</p>
        {/if}
      </div>
    </div>
  </div>
{/if}

<div class="bg-gradient-to-r from-emerald-950 via-emerald-900 to-emerald-800 text-white py-16 px-4 sm:px-8">
  <div class="max-w-6xl mx-auto space-y-4">
    <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-emerald-500/20 text-emerald-300 text-xs font-semibold border border-emerald-500/30">
      Dokumentasi Visual
    </div>
    <h1 class="font-serif-luxury font-bold text-3xl sm:text-5xl text-white">
      Galeri Foto Kegiatan & Kehidupan Santri
    </h1>
    <p class="text-emerald-100/80 text-base max-w-2xl">
      Potret dinamika keseharian santri Al-Hamid: dari heningnya majelis sorogan, keceriaan halaqoh, perhelatan wisuda, hingga keindahan arsitektur pesantren.
    </p>
  </div>
</div>

<div class="max-w-6xl mx-auto px-4 sm:px-8 py-12 space-y-8">
  <!-- Categories Filter -->
  <div class="flex flex-wrap items-center gap-2">
    <button
      onclick={() => handleFilter('')}
      class="px-4 py-2 rounded-xl text-xs font-semibold transition {!selectedCategory
        ? 'bg-emerald-800 text-white shadow-md'
        : 'bg-white text-slate-600 hover:bg-slate-100 border border-slate-200'}"
    >
      Semua Album Foto
    </button>
    {#each categories as cat}
      <button
        onclick={() => handleFilter(cat.slug)}
        class="px-4 py-2 rounded-xl text-xs font-semibold transition {selectedCategory === cat.slug
          ? 'bg-emerald-800 text-white shadow-md'
          : 'bg-white text-slate-600 hover:bg-slate-100 border border-slate-200'}"
      >
        {cat.name}
      </button>
    {/each}
  </div>

  <!-- Photo Grid -->
  <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
    {#each photos as photo}
      <div
        class="bg-white rounded-2xl border border-slate-100 shadow-sm hover:shadow-xl transition-all duration-300 overflow-hidden group cursor-pointer relative"
        role="button"
        tabindex="0"
        onclick={() => (lightboxPhoto = photo)}
        onkeydown={(e) => e.key === 'Enter' && (lightboxPhoto = photo)}
      >
        <div class="relative aspect-square overflow-hidden bg-slate-900">
          <img
            src={photo.image_url}
            alt={photo.title}
            class="w-full h-full object-cover group-hover:scale-105 transition duration-500"
          />
          <div class="absolute inset-0 bg-slate-950/40 opacity-0 group-hover:opacity-100 transition-opacity flex items-center justify-center text-white">
            <ZoomIn class="w-8 h-8" />
          </div>
        </div>
        <div class="p-4 space-y-1">
          <div class="text-[11px] font-bold text-emerald-700 uppercase">{photo.category?.name}</div>
          <h2 class="font-semibold text-sm text-slate-900 line-clamp-1">{photo.title}</h2>
          {#if photo.caption}
            <p class="text-xs text-slate-500 line-clamp-1">{photo.caption}</p>
          {/if}
        </div>
      </div>
    {/each}
  </div>
</div>

