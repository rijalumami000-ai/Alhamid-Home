<script lang="ts">
  import { onMount } from 'svelte';
  import { api } from '$lib/api';
  import type { Article, Category } from '$lib/types';
  import { Search, Calendar, ChevronRight, BookOpen, Loader2, Sparkles } from 'lucide-svelte';

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

<!-- Header -->
<div class="bg-[#090D16] text-white py-20 px-4 sm:px-8 relative overflow-hidden border-b border-slate-800/80">
  <div class="absolute inset-0 bg-grid-dark opacity-30 pointer-events-none"></div>
  <div class="absolute -top-24 left-1/2 -translate-x-1/2 w-[600px] h-[300px] bg-teal-500/15 blur-3xl pointer-events-none rounded-full"></div>

  <div class="max-w-6xl mx-auto space-y-4 relative z-10">
    <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-teal-500/15 text-teal-300 text-xs font-bold border border-teal-500/30">
      <Sparkles class="w-3.5 h-3.5 text-teal-400" />
      <span>Literasi & Warta Pesantren</span>
    </div>
    <h1 class="font-display font-black text-3xl sm:text-5xl lg:text-6xl text-white tracking-tight">
      Kabar, Kajian & Opini Santri
    </h1>
    <p class="text-slate-300 text-base sm:text-lg max-w-2xl font-normal leading-relaxed">
      Kumpulan artikel keilmuan, rilis kegiatan resmi, kabar prestasi, dan dinamika keislaman di lingkungan Pondok Pesantren Al-Hamid.
    </p>
  </div>
</div>

<div class="max-w-6xl mx-auto px-4 sm:px-8 py-20 space-y-10">
  <!-- Search and Categories Toolbar -->
  <div class="flex flex-col md:flex-row items-center justify-between gap-4">
    <div class="flex flex-wrap items-center gap-2 w-full md:w-auto">
      <button
        onclick={() => handleFilter('')}
        class="px-5 py-2.5 rounded-full text-xs font-display font-bold transition {selectedCategory === ''
          ? 'bg-slate-900 text-white shadow-md'
          : 'bg-white text-slate-600 hover:bg-slate-100 border border-slate-200'}"
      >
        Semua Kategori
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

    <form onsubmit={handleSearchSubmit} class="relative w-full md:w-72">
      <input
        type="text"
        placeholder="Cari artikel..."
        bind:value={searchQuery}
        class="w-full pl-10 pr-4 py-2.5 rounded-full bg-white border border-slate-200 text-xs font-medium focus:outline-none focus:ring-2 focus:ring-teal-500 shadow-sm"
      />
      <Search class="w-4 h-4 text-slate-400 absolute left-3.5 top-3" />
    </form>
  </div>

  <!-- Articles List -->
  {#if loading}
    <div class="py-24 text-center text-slate-400 flex items-center justify-center gap-2">
      <Loader2 class="w-6 h-6 animate-spin text-teal-600" />
      <span>Memuat warta dan artikel...</span>
    </div>
  {:else if articles.length === 0}
    <div class="py-24 text-center text-slate-500 bg-white rounded-3xl border border-slate-200/80 p-8">
      Tidak ada artikel yang sesuai dengan pencarian Anda.
    </div>
  {:else}
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
      {#each articles as article}
        <article class="bg-white rounded-3xl border border-slate-200/80 shadow-sm hover:shadow-2xl hover:-translate-y-1 transition-all duration-300 overflow-hidden flex flex-col group">
          <a href="/berita/{article.slug}" class="relative h-52 overflow-hidden bg-slate-950 block">
            <img
              src={article.thumbnail_url}
              alt={article.title}
              class="w-full h-full object-cover group-hover:scale-105 transition duration-500"
            />
            <div class="absolute inset-0 bg-gradient-to-t from-slate-950/70 via-transparent to-transparent"></div>
            <div class="absolute top-3 left-3">
              <span class="px-3 py-1 rounded-full bg-slate-900/90 text-teal-300 text-[11px] font-bold backdrop-blur-md border border-white/10">
                {article.category?.name || 'Kajian'}
              </span>
            </div>
          </a>

          <div class="p-6 flex-1 flex flex-col justify-between space-y-4">
            <div>
              <div class="text-xs text-slate-400 mb-2 flex items-center gap-2 font-medium">
                <Calendar class="w-3.5 h-3.5 text-teal-600" />
                <span>{new Date(article.published_at).toLocaleDateString('id-ID', { day: 'numeric', month: 'long', year: 'numeric' })}</span>
                <span>&bull;</span>
                <span>{article.author}</span>
              </div>
              <h2 class="font-display font-bold text-lg text-slate-950 group-hover:text-teal-700 transition line-clamp-2 leading-snug">
                <a href="/berita/{article.slug}">{article.title}</a>
              </h2>
              <p class="text-slate-600 text-xs sm:text-sm mt-2 line-clamp-2 leading-relaxed">
                {article.excerpt}
              </p>
            </div>

            <div class="pt-4 border-t border-slate-100 flex items-center justify-between">
              <span class="text-xs text-slate-400 font-medium">{article.view_count} pembaca</span>
              <a
                href="/berita/{article.slug}"
                class="inline-flex items-center gap-1 text-xs font-display font-bold text-teal-700 hover:text-teal-800"
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
