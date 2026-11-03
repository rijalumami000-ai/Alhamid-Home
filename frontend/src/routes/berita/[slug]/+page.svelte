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
    <div class="animate-spin w-8 h-8 border-4 border-teal-600 border-t-transparent rounded-full mx-auto mb-4"></div>
    Memuat artikel...
  </div>
{:else if !article}
  <div class="py-32 text-center max-w-md mx-auto px-4 space-y-4">
    <h1 class="text-2xl font-display font-black text-slate-800">Artikel Tidak Ditemukan</h1>
    <p class="text-sm text-slate-500">Artikel yang Anda cari mungkin telah dihapus atau URL tidak valid.</p>
    <a href="/berita" class="inline-flex px-6 py-2.5 rounded-full bg-slate-900 text-white font-bold text-xs uppercase tracking-wider">
      Kembali ke Berita
    </a>
  </div>
{:else}
  <article class="max-w-4xl mx-auto px-4 sm:px-8 py-16 space-y-8">
    <!-- Breadcrumb & Back -->
    <div class="flex items-center justify-between text-xs text-slate-500">
      <a href="/berita" class="inline-flex items-center gap-1.5 hover:text-teal-700 transition font-display font-bold">
        <ArrowLeft class="w-4 h-4" /> Kembali ke Warta
      </a>
      <span class="px-3 py-1 rounded-full bg-teal-50 text-teal-800 border border-teal-200/60 font-black text-[11px] uppercase tracking-wider">
        {article.category?.name || 'Kajian'}
      </span>
    </div>

    <!-- Article Title -->
    <h1 class="font-display font-black text-3xl sm:text-5xl lg:text-6xl text-slate-950 leading-[1.12] tracking-tight">
      {article.title}
    </h1>

    <!-- Meta Information -->
    <div class="flex flex-wrap items-center gap-4 text-xs text-slate-500 pb-6 border-b border-slate-200/80">
      <div class="flex items-center gap-1.5 font-bold text-slate-900">
        <User class="w-4 h-4 text-teal-600" />
        <span>{article.author}</span>
      </div>
      <div class="flex items-center gap-1.5 font-medium">
        <Calendar class="w-4 h-4 text-teal-600" />
        <span>{new Date(article.published_at).toLocaleDateString('id-ID', { day: 'numeric', month: 'long', year: 'numeric' })}</span>
      </div>
      <div class="flex items-center gap-1.5 font-medium">
        <Eye class="w-4 h-4 text-teal-600" />
        <span>{article.view_count} kali dibaca</span>
      </div>
    </div>

    <!-- Hero Thumbnail Image -->
    {#if article.thumbnail_url}
      <div class="rounded-3xl overflow-hidden shadow-2xl border border-slate-200/80 bg-slate-950 aspect-[16/9]">
        <img
          src={article.thumbnail_url}
          alt={article.title}
          class="w-full h-full object-cover"
        />
      </div>
    {/if}

    <!-- Excerpt Card -->
    {#if article.excerpt}
      <div class="p-6 sm:p-8 rounded-3xl bg-slate-100/80 border-l-4 border-teal-500 text-slate-700 font-editorial italic text-lg sm:text-xl leading-relaxed">
        "{article.excerpt}"
      </div>
    {/if}

    <!-- Article Content -->
    <div class="text-slate-700 text-base sm:text-lg leading-relaxed space-y-4 prose prose-teal max-w-none font-normal">
      {@html article.content}
    </div>

    <!-- Share Actions -->
    <div class="pt-8 border-t border-slate-200/80 flex flex-wrap items-center justify-between gap-4">
      <div class="text-xs font-bold text-slate-500 flex items-center gap-2">
        <Share2 class="w-4 h-4 text-teal-600" />
        <span>Bagikan Artikel Ini:</span>
      </div>
      <div class="flex items-center gap-2">
        <a
          href="https://wa.me/?text={encodeURIComponent(article.title + ' ' + (typeof window !== 'undefined' ? window.location.href : ''))}"
          target="_blank"
          rel="noopener noreferrer"
          class="px-5 py-2.5 rounded-full bg-teal-600 hover:bg-teal-500 text-white text-xs font-display font-bold tracking-wide transition shadow-sm"
        >
          WhatsApp
        </a>
        <button
          onclick={copyUrl}
          class="px-5 py-2.5 rounded-full bg-slate-100 hover:bg-slate-200 text-slate-700 text-xs font-display font-bold tracking-wide transition border border-slate-200 flex items-center gap-1.5"
        >
          {#if copied}
            <Check class="w-3.5 h-3.5 text-teal-600" />
            <span>Tersalin!</span>
          {:else}
            <span>Salin Link</span>
          {/if}
        </button>
      </div>
    </div>

    <!-- Related Articles -->
    {#if related && related.length > 0}
      <div class="pt-16 border-t border-slate-200/80 space-y-6">
        <h2 class="font-display font-black text-2xl text-slate-950 tracking-tight">
          Artikel Terkait Lainnya
        </h2>
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
          {#each related as rel}
            <a
              href="/berita/{rel.slug}"
              class="p-5 rounded-2xl bg-white border border-slate-200/80 shadow-sm hover:shadow-xl hover:border-teal-500/40 transition block group"
            >
              <div class="text-[11px] font-bold text-teal-700 uppercase tracking-wider mb-1">
                {rel.category?.name || 'Kajian'}
              </div>
              <h3 class="font-display font-bold text-base text-slate-900 group-hover:text-teal-700 transition leading-snug">
                {rel.title}
              </h3>
            </a>
          {/each}
        </div>
      </div>
    {/if}
  </article>
{/if}
