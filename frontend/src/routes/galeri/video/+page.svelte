<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { MediaVideo, Category } from '$lib/types';
  import { Play, X, Video, Sparkles } from 'lucide-svelte';

  let videos = $state<MediaVideo[]>([]);
  let categories = $state<Category[]>([]);
  let selectedCategory = $state('');
  let activeVideo = $state<MediaVideo | null>(null);

  async function loadData() {
    try {
      const [vidRes, catRes] = await Promise.all([
        api.getVideos(selectedCategory),
        api.getCategories('galeri_video'),
      ]);
      videos = vidRes;
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
  <title>Galeri Video | Pondok Pesantren Al-Hamid</title>
</svelte:head>

<!-- Video Player Modal -->
{#if activeVideo}
  <div
    class="fixed inset-0 z-50 bg-[#090D16]/95 backdrop-blur-xl flex items-center justify-center p-4"
    role="dialog"
    aria-modal="true"
    tabindex="-1"
    onkeydown={(e) => e.key === 'Escape' && (activeVideo = null)}
  >
    <button
      class="absolute inset-0 w-full h-full cursor-default"
      onclick={() => (activeVideo = null)}
      aria-label="Tutup pemutar video"
    ></button>

    <button
      class="absolute top-6 right-6 z-10 p-3 rounded-full bg-white/10 hover:bg-white/20 text-white transition"
      onclick={() => (activeVideo = null)}
      aria-label="Tutup video"
    >
      <X class="w-6 h-6" />
    </button>

    <div class="relative z-10 max-w-4xl w-full aspect-video bg-black rounded-3xl overflow-hidden shadow-2xl border border-white/10">
      <iframe
        src="https://www.youtube.com/embed/{activeVideo.youtube_id}?autoplay=1"
        title={activeVideo.title}
        class="w-full h-full border-0"
        allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
        allowfullscreen
      ></iframe>
    </div>
  </div>
{/if}

<!-- Header -->
<div class="bg-[#090D16] text-white py-20 px-4 sm:px-8 relative overflow-hidden border-b border-slate-800/80">
  <div class="absolute inset-0 bg-grid-dark opacity-30 pointer-events-none"></div>
  <div class="absolute -top-24 left-1/2 -translate-x-1/2 w-[600px] h-[300px] bg-rose-500/15 blur-3xl pointer-events-none rounded-full"></div>

  <div class="max-w-6xl mx-auto space-y-4 relative z-10">
    <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-rose-500/15 text-rose-300 text-xs font-bold border border-rose-500/30">
      <Sparkles class="w-3.5 h-3.5 text-rose-400" />
      <span>Arsip Audio Visual</span>
    </div>
    <h1 class="font-display font-black text-3xl sm:text-5xl lg:text-6xl text-white tracking-tight">
      Galeri Video & Dokumentasi
    </h1>
    <p class="text-slate-300 text-base sm:text-lg max-w-2xl font-normal leading-relaxed">
      Tayangan kajian kitab kuning, profil lembaga, kreasi seni hadroh santri, dan rekaman momentum istimewa pesantren.
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
      Semua Kategori Video
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

  <!-- Video Grid -->
  <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
    {#each videos as video}
      <button
        class="bg-white rounded-3xl border border-slate-200/80 shadow-sm hover:shadow-2xl hover:-translate-y-1 transition-all duration-300 overflow-hidden group cursor-pointer text-left flex flex-col"
        onclick={() => (activeVideo = video)}
      >
        <div class="relative aspect-video overflow-hidden bg-slate-950 w-full">
          <img
            src={video.thumbnail_url || `https://img.youtube.com/vi/${video.youtube_id}/hqdefault.jpg`}
            alt={video.title}
            class="w-full h-full object-cover group-hover:scale-105 transition duration-500"
          />
          <div class="absolute inset-0 bg-slate-950/40 group-hover:bg-slate-950/20 transition flex items-center justify-center">
            <div class="w-14 h-14 rounded-full bg-rose-600 text-white flex items-center justify-center shadow-xl group-hover:scale-110 transition duration-300">
              <Play class="w-6 h-6 fill-white ml-0.5" />
            </div>
          </div>
        </div>
        <div class="p-6 space-y-2 flex-1">
          <div class="text-[10px] font-black text-rose-600 uppercase tracking-wider">{video.category?.name || 'Video'}</div>
          <h2 class="font-display font-bold text-base text-slate-950 group-hover:text-rose-600 transition leading-snug line-clamp-2">
            {video.title}
          </h2>
          {#if video.description}
            <p class="text-xs text-slate-500 line-clamp-2 leading-relaxed">{video.description}</p>
          {/if}
        </div>
      </button>
    {/each}
  </div>
</div>
