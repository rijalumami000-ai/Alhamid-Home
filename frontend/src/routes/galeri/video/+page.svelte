<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { MediaVideo, Category } from '$lib/types';
  import { Play, X, Video } from 'lucide-svelte';

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
  <title>Galeri Video YouTube | Pondok Pesantren Al-Hamid</title>
</svelte:head>

<!-- Video Player Modal -->
{#if activeVideo}
  <div
    class="fixed inset-0 z-50 bg-slate-950/90 backdrop-blur-md flex items-center justify-center p-4"
    role="button"
    tabindex="0"
    onclick={() => (activeVideo = null)}
    onkeydown={(e) => e.key === 'Escape' && (activeVideo = null)}
  >
    <button
      class="absolute top-6 right-6 p-2 rounded-full bg-white/10 hover:bg-white/20 text-white transition"
      onclick={() => (activeVideo = null)}
      aria-label="Tutup video"
    >
      <X class="w-6 h-6" />
    </button>
    <div
      class="max-w-4xl w-full aspect-video bg-black rounded-2xl overflow-hidden shadow-2xl"
      role="document"
      tabindex="-1"
      onclick={(e) => e.stopPropagation()}
      onkeydown={(e) => e.stopPropagation()}
    >
      <iframe
        src="https://www.youtube-nocookie.com/embed/{activeVideo.youtube_id}?autoplay=1"
        title={activeVideo.title}
        class="w-full h-full border-0"
        allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
        allowfullscreen
      ></iframe>
    </div>
  </div>
{/if}

<div class="bg-gradient-to-r from-emerald-950 via-emerald-900 to-emerald-800 text-white py-16 px-4 sm:px-8">
  <div class="max-w-6xl mx-auto space-y-4">
    <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-emerald-500/20 text-emerald-300 text-xs font-semibold border border-emerald-500/30">
      Dokumenter & Audio-Visual
    </div>
    <h1 class="font-serif-luxury font-bold text-3xl sm:text-5xl text-white">
      Galeri Video Resmi YouTube
    </h1>
    <p class="text-emerald-100/80 text-base max-w-2xl">
      Saksikan rekaman kajian dewan masyayikh, film profil pesantren, lantunan hadroh santri, dan liputan perhelatan akbar Al-Hamid.
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
      Semua Video
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

  <!-- Video Grid -->
  <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
    {#each videos as video}
      <div
        class="bg-white rounded-3xl border border-slate-100 shadow-sm hover:shadow-xl transition-all duration-300 overflow-hidden group cursor-pointer flex flex-col"
        role="button"
        tabindex="0"
        onclick={() => (activeVideo = video)}
        onkeydown={(e) => e.key === 'Enter' && (activeVideo = video)}
      >
        <div class="relative aspect-video overflow-hidden bg-slate-900">
          <img
            src="https://img.youtube.com/vi/{video.youtube_id}/hqdefault.jpg"
            alt={video.title}
            class="w-full h-full object-cover group-hover:scale-105 transition duration-500"
          />
          <div class="absolute inset-0 bg-slate-950/30 flex items-center justify-center">
            <div class="w-14 h-14 rounded-full bg-red-600/90 text-white flex items-center justify-center shadow-lg group-hover:scale-110 group-hover:bg-red-600 transition">
              <Play class="w-6 h-6 fill-white ml-0.5" />
            </div>
          </div>
          <div class="absolute top-3 left-3">
            <span class="px-2.5 py-1 rounded-full bg-slate-950/80 text-white text-[11px] font-semibold backdrop-blur-sm">
              {video.category?.name || 'Video'}
            </span>
          </div>
        </div>

        <div class="p-6 flex-1 flex flex-col justify-between space-y-2">
          <div>
            <h2 class="font-serif-luxury font-bold text-lg text-slate-900 group-hover:text-emerald-700 transition line-clamp-2">
              {video.title}
            </h2>
            {#if video.description}
              <p class="text-slate-500 text-xs mt-2 line-clamp-2 leading-relaxed">
                {video.description}
              </p>
            {/if}
          </div>

          <div class="pt-3 border-t border-slate-100 text-xs font-semibold text-emerald-700 flex items-center gap-1.5">
            <Play class="w-3.5 h-3.5" /> Putar Video YouTube
          </div>
        </div>
      </div>
    {/each}
  </div>
</div>

