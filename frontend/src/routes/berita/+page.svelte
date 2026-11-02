<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { Article, Category } from '$lib/types';
  import { Search, Calendar, ChevronRight, BookOpen, Loader2 } from 'lucide-svelte';

  let articles = $state<Article[]>([]);
  let categories = $state<Category[]>([]);
  let selectedCategory = $state('');
  let searchQuery = $state('');
  let loading = $state(true);
  let total = $state(0);
  let page = $state(1);

  async function loadData() {
    loading = true;
    try {
      const [artRes, catRes] = await Promise.all([
        api.getArticles({ category: selectedCategory, q: searchQuery, page }),
        api.getCategories('berita'),
      ]);
      articles = artRes.data;
      total = artRes.total;
      categories = catRes;
    } catch (e) {
      console.error(e);
    } finally {
      loading = false;
    }
  }

  function handleFilter(slug: string) {
    selectedCategory = slug;
    page = 1;
    loadData();
  }

  function handleSearchSubmit(e: SubmitEvent) {
    e.preventDefault();
    page = 1;
    loadData();
  }

  onMount(() => {
    loadData();
  });
</script>

<svelte:head>
  <title>Berita & Artikel | Pondok Pesantren Al-Hamid</title>
</svelte:head>

<div class="bg-gradient-to-r from-emerald-950 via-emerald-900 to-emerald-800 text-white py-16 px-4 sm:px-8">
  <div class="max-w-6xl mx-auto space-y-4">
    <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-emerald-500/20 text-emerald-300 text-xs font-semibold border border-emerald-500/30">
      Literasi & Informasi
    </div>
    <h1 class="font-serif-luxury font-bold text-3xl sm:text-5xl text-white">
      Kabar, Kajian & Pengumuman
    </h1>
    <p class="text-emerald-100/80 text-base max-w-2xl">
      Eksplorasi wawasan keislaman, mutiara kitab kuning, pengumuman resmi, dan catatan prestasi santri Al-Hamid.
    </p>
  </div>
</div>

<div class="max-w-6xl mx-auto px-4 sm:px-8 py-12 space-y-8">
  <!-- Search & Filter Controls -->
  <div class="flex flex-col md:flex-row md:items-center justify-between gap-4">
    <!-- Category Pills -->
    <div class="flex flex-wrap items-center gap-2">
      <button
        onclick={() => handleFilter('')}
        class="px-4 py-2 rounded-xl text-xs font-semibold transition {!selectedCategory
          ? 'bg-emerald-800 text-white shadow-md'
          : 'bg-white text-slate-600 hover:bg-slate-100 border border-slate-200'}"
      >
        Semua Kategori
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

    <!-- Search Box -->
    <form onsubmit={handleSearchSubmit} class="relative w-full md:w-72">
      <input
        type="text"
        placeholder="Cari artikel..."
        bind:value={searchQuery}
        class="w-full pl-9 pr-4 py-2 rounded-xl bg-white border border-slate-200 text-xs font-medium text-slate-800 placeholder-slate-400 focus:outline-none focus:ring-2 focus:ring-emerald-600/30"
      />
      <Search class="w-4 h-4 text-slate-400 absolute left-3 top-2.5" />
    </form>
  </div>

  <!-- Articles Grid -->
  {#if loading}
    <div class="py-24 text-center text-slate-400 flex items-center justify-center gap-2">
      <Loader2 class="w-5 h-5 animate-spin text-emerald-600" />
      <span>Memuat artikel...</span>
    </div>
  {:else if articles.length === 0}
    <div class="py-20 text-center bg-white rounded-3xl border border-slate-100 p-8 space-y-2">
      <BookOpen class="w-10 h-10 text-slate-300 mx-auto" />
      <h3 class="font-bold text-slate-700">Tidak ada artikel ditemukan</h3>
      <p class="text-xs text-slate-400">Coba gunakan kata kunci pencarian lain atau pilih kategori berbeda.</p>
    </div>
  {:else}
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
      {#each articles as article}
        <article class="bg-white rounded-3xl border border-slate-100 shadow-sm hover:shadow-xl transition-all duration-300 overflow-hidden flex flex-col group">
          <a href="/berita/{article.slug}" class="relative h-52 overflow-hidden bg-slate-900 block">
            <img
              src={article.thumbnail_url}
              alt={article.title}
              class="w-full h-full object-cover group-hover:scale-105 transition duration-500"
            />
            <div class="absolute top-3 left-3">
              <span class="px-2.5 py-1 rounded-full bg-emerald-800/90 text-white text-[11px] font-semibold backdrop-blur-sm">
                {article.category?.name || 'Kajian'}
              </span>
            </div>
          </a>

          <div class="p-6 flex-1 flex flex-col justify-between space-y-3">
            <div>
              <div class="text-xs text-slate-400 mb-2 flex items-center gap-2">
                <Calendar class="w-3.5 h-3.5" />
                <span>{new Date(article.published_at).toLocaleDateString('id-ID', { day: 'numeric', month: 'long', year: 'numeric' })}</span>
                <span>&bull;</span>
                <span>{article.author}</span>
              </div>
              <h2 class="font-serif-luxury font-bold text-lg text-slate-900 group-hover:text-emerald-700 transition line-clamp-2">
                <a href="/berita/{article.slug}">{article.title}</a>
              </h2>
              <p class="text-slate-600 text-xs sm:text-sm mt-2 line-clamp-2 leading-relaxed">
                {article.excerpt}
              </p>
            </div>

            <div class="pt-3 border-t border-slate-100 flex items-center justify-between">
              <span class="text-xs text-slate-400">{article.view_count} pembaca</span>
              <a
                href="/berita/{article.slug}"
                class="inline-flex items-center gap-1 text-xs font-bold text-emerald-700 hover:text-emerald-800"
              >
                <span>Baca Lengkap</span>
                <ChevronRight class="w-3.5 h-3.5" />
              </a>
            </div>
          </div>
        </article>
      {/each}
    </div>
  {/if}
</div>

