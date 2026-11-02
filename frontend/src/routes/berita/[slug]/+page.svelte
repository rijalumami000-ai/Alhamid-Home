<script lang="ts">
  import { onMount } from 'svelte';
  import { page } from '$app/state';
  import { api } from '$lib/api';
  import type { Article } from '$lib/types';
  import { Calendar, User, Eye, ArrowLeft, Share2, Check, ChevronRight } from 'lucide-svelte';

  let article = $state<Article | null>(null);
  let related = $state<Article[]>([]);
  let loading = $state(true);
  let copied = $state(false);

  async function loadArticle(slug: string) {
    loading = true;
    try {
      const res = await api.getArticleBySlug(slug);
      article = res.article;
      related = res.related;
    } catch (e) {
      console.error(e);
    } finally {
      loading = false;
    }
  }

  $effect(() => {
    const slug = page.params.slug;
    if (slug) {
      loadArticle(slug);
    }
  });

  function copyUrl() {
    if (typeof window !== 'undefined') {
      navigator.clipboard.writeText(window.location.href);
      copied = true;
      setTimeout(() => (copied = false), 2500);
    }
  }
</script>

<svelte:head>
  <title>{article?.title ? `${article.title} | Pesantren Al-Hamid` : 'Artikel Pesantren Al-Hamid'}</title>
</svelte:head>

{#if loading}
  <div class="py-32 text-center text-slate-400">
    <div class="animate-spin w-8 h-8 border-4 border-emerald-600 border-t-transparent rounded-full mx-auto mb-4"></div>
    Memuat artikel...
  </div>
{:else if !article}
  <div class="py-32 text-center max-w-md mx-auto px-4 space-y-4">
    <h1 class="text-2xl font-bold text-slate-800">Artikel Tidak Ditemukan</h1>
    <p class="text-sm text-slate-500">Artikel yang Anda cari mungkin telah dihapus atau URL tidak valid.</p>
    <a href="/berita" class="inline-flex px-6 py-2.5 rounded-xl bg-emerald-800 text-white font-semibold text-sm">
      Kembali ke Berita
    </a>
  </div>
{:else}
  <article class="max-w-4xl mx-auto px-4 sm:px-8 py-12 space-y-8">
    <!-- Breadcrumb & Back -->
    <div class="flex items-center justify-between text-xs text-slate-500">
      <a href="/berita" class="inline-flex items-center gap-1 hover:text-emerald-700 transition font-medium">
        <ArrowLeft class="w-4 h-4" /> Kembali ke Semua Berita
      </a>
      <span class="px-2.5 py-1 rounded-full bg-emerald-100 text-emerald-800 font-semibold">
        {article.category?.name || 'Kajian'}
      </span>
    </div>

    <!-- Article Title -->
    <h1 class="font-serif-luxury font-bold text-3xl sm:text-5xl text-slate-900 leading-[1.2]">
      {article.title}
    </h1>

    <!-- Meta Information -->
    <div class="flex flex-wrap items-center gap-4 text-xs text-slate-500 pb-6 border-b border-slate-200">
      <div class="flex items-center gap-1.5 font-medium text-slate-700">
        <User class="w-4 h-4 text-emerald-600" />
        <span>{article.author}</span>
      </div>
      <div class="flex items-center gap-1.5">
        <Calendar class="w-4 h-4 text-emerald-600" />
        <span>{new Date(article.published_at).toLocaleDateString('id-ID', { day: 'numeric', month: 'long', year: 'numeric' })}</span>
      </div>
      <div class="flex items-center gap-1.5">
        <Eye class="w-4 h-4 text-emerald-600" />
        <span>{article.view_count} kali dibaca</span>
      </div>
    </div>

    <!-- Hero Thumbnail Image -->
    {#if article.thumbnail_url}
      <div class="rounded-3xl overflow-hidden shadow-lg border border-slate-100 bg-slate-900 aspect-[16/9]">
        <img
          src={article.thumbnail_url}
          alt={article.title}
          class="w-full h-full object-cover"
        />
      </div>
    {/if}

    <!-- Excerpt Card -->
    {#if article.excerpt}
      <div class="p-6 rounded-2xl bg-emerald-50/70 border-l-4 border-emerald-600 text-slate-700 font-medium text-base sm:text-lg leading-relaxed italic">
        {article.excerpt}
      </div>
    {/if}

    <!-- Article Content -->
    <div class="text-slate-800 text-base sm:text-lg leading-relaxed space-y-4 prose prose-emerald max-w-none">
      {@html article.content}
    </div>

    <!-- Share Actions -->
    <div class="pt-8 border-t border-slate-200 flex flex-wrap items-center justify-between gap-4">
      <div class="text-xs font-semibold text-slate-500 flex items-center gap-2">
        <Share2 class="w-4 h-4 text-emerald-600" />
        <span>Bagikan Artikel Ini:</span>
      </div>
      <div class="flex items-center gap-2">
        <a
          href="https://wa.me/?text={encodeURIComponent(article.title + ' ' + (typeof window !== 'undefined' ? window.location.href : ''))}"
          target="_blank"
          rel="noopener noreferrer"
          class="px-4 py-2 rounded-xl bg-emerald-800 hover:bg-emerald-700 text-white text-xs font-semibold transition"
        >
          WhatsApp
        </a>
        <button
          onclick={copyUrl}
          class="px-4 py-2 rounded-xl bg-slate-100 hover:bg-slate-200 text-slate-700 text-xs font-semibold transition flex items-center gap-1"
        >
          {#if copied}
            <Check class="w-3.5 h-3.5 text-emerald-600" /> Tautan Disalin
          {:else}
            Salin Link
          {/if}
        </button>
      </div>
    </div>

    <!-- Related Articles -->
    {#if related.length > 0}
      <div class="pt-12 space-y-6">
        <h2 class="font-serif-luxury font-bold text-2xl text-slate-900">
          Artikel Terkait Lainnya
        </h2>
        <div class="grid grid-cols-1 sm:grid-cols-3 gap-6">
          {#each related as r}
            <a
              href="/berita/{r.slug}"
              class="bg-white rounded-2xl border border-slate-100 p-4 shadow-sm hover:shadow-md transition space-y-2 group block"
            >
              <div class="h-32 rounded-xl overflow-hidden bg-slate-100 mb-2">
                <img
                  src={r.thumbnail_url}
                  alt={r.title}
                  class="w-full h-full object-cover group-hover:scale-105 transition"
                />
              </div>
              <h3 class="font-serif-luxury font-bold text-sm text-slate-900 group-hover:text-emerald-700 transition line-clamp-2">
                {r.title}
              </h3>
              <p class="text-xs text-slate-500 line-clamp-2">
                {r.excerpt}
              </p>
            </a>
          {/each}
        </div>
      </div>
    {/if}
  </article>
{/if}
